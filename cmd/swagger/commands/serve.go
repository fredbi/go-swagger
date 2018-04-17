package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/gorilla/handlers"
	"github.com/toqueteos/webbrowser"
)

const (
	// defaultSwaggerUIURL is the default URL contacted to server SwaggerUI style UI. This may be overriden with --ui-url arg.
	defaultSwaggerUIURL = "http://petstore.swagger.io/"
)

// ServeCmd to serve a swagger spec with docs ui
type ServeCmd struct {
	BasePath string   `long:"base-path" description:"the base path to serve the spec and UI at" default:"/"`
	Flavor   string   `short:"F" long:"flavor" description:"the flavor of docs, can be swagger or redoc" default:"redoc" choice:"redoc" choice:"swagger"`
	UIURL    string   `long:"ui-url" description:"override the url where to fetch the UI rendering resources. Default is to fetch these resources online according to the UI flavor choice"`
	UILocal  string   `long:"ui-local" description:"when ui-url is a local URL, the file to be served for this URL"`
	DocURL   string   `long:"doc-url" description:"override the route at which UI is served" default:"docs"`
	SpecJSON string   `long:"spec-json" description:"override the document to fetch the spec" default:"swagger.json"`
	NoOpen   bool     `long:"no-open" description:"when present won't open the the browser to show the url"`
	NoUI     bool     `long:"no-ui" description:"when present, only the swagger spec will be served"`
	Port     int      `long:"port" short:"p" description:"the port to serve this site" env:"PORT"`
	Host     string   `long:"host" description:"the interface to serve this site" default:"0.0.0.0" env:"HOST"`
	NoRoot   bool     `long:"no-root" description:"always add basePath from spec to route when serving spec. Enabled by default when serving multiple specs. Spec basePath is added to the BasePath arg"`
	NoCors   bool     `long:"no-cors" description:"disable CORS on spec serve route"`
	Titles   []string `long:"title" description:"override the title tag of the rendered doc page. Default is to use the API title in spec, or if empty: 'API documentation'"`
	// TODO
	// IndexTemplate
	// Style
	// TLS
	// YAML spec
	basePaths      map[string]bool
	host           string
	port           int
	defaultOpenURL string
}

// Execute the serve command
func (s *ServeCmd) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("specify the spec(s) to serve as argument to the serve command")
	}

	listener, errL := net.Listen("tcp4", net.JoinHostPort(s.Host, strconv.Itoa(s.Port)))
	if errL != nil {
		return errL
	}

	if err := s.resolveHostAndPort(listener); err != nil {
		return err
	}

	mux := http.NewServeMux()

	s.basePaths = map[string]bool{}

	for i, arg := range args {
		var specPath string

		specDoc, err := loads.Spec(arg)
		if err != nil {
			return fmt.Errorf("on spec %s, could not load spec: %v", arg, err)
		}

		if len(args) > 1 || s.NoRoot {
			// when serving multiple specs, add the spec's base path to construct the route, e.g. /{basePath}/{specBasePath}
			specPath = path.Join(s.BasePath, specDoc.BasePath())
		} else {
			specPath = s.BasePath
		}

		// skips undiscriminated routes
		if !s.isRouteUnique(specPath, i) {
			log.Printf("on spec %s, cannot serve %s: path conflicts with another API to serve. Skipping.", arg, specPath)
			continue
		}

		if !s.NoUI {
			// serve UI
			if err := s.addHandlerUI(mux, specPath, i, specDoc); err != nil {
				return fmt.Errorf("on spec %s, could not handle UI: %v", arg, err)
			}
		}

		// serve spec
		if err := s.addHandlerSpec(mux, specPath, specDoc); err != nil {
			return fmt.Errorf("on spec %s, could not handle spec: %v", arg, err)
		}
	}

	// serve local UI if UIURL is local
	if s.UIURL != "" {
		if err := s.addHandlerLocalUI(mux); err != nil {
			return fmt.Errorf("could not handle UI URL: %v", err)
		}
	}

	errFuture := make(chan error)

	go func() {
		docServer := new(http.Server)
		docServer.SetKeepAlivesEnabled(true)
		docServer.Handler = mux
		errFuture <- docServer.Serve(listener)
	}()

	if !s.NoOpen && !s.NoUI {
		err := webbrowser.Open(s.defaultOpenURL)
		if err != nil {
			return err
		}
	}
	return <-errFuture
}

func (s *ServeCmd) resolveHostAndPort(listener net.Listener) error {
	// resolveHostAndPort retrieves host and port allocated by listener
	// (if no port is specified, a random one is chosen)
	var err error
	s.host, s.port, err = swag.SplitHostPort(listener.Addr().String())
	if err != nil {
		return err
	}
	if s.host == "0.0.0.0" {
		// TODO: windows setup 127.0.0.1
		s.host = "localhost"
	}
	log.Printf("listening on %s:%d", s.host, s.port)
	return nil
}

func (s *ServeCmd) addHandlerUI(mux *http.ServeMux, specPath string, specIndex int, specDoc *loads.Document) error {
	// addHandlerUI injects the UI serving middleware and route
	handler := http.NotFoundHandler()

	// retrieve page titles from spec or from command line args
	if specIndex > len(s.Titles)-1 {
		if specDoc.Spec().Info != nil && specDoc.Spec().Info.Title != "" {
			// use spec title as page title
			s.Titles = append(s.Titles, specDoc.Spec().Info.Title)
		} else {
			// default title
			s.Titles = append(s.Titles, "API documentation")
		}
	}

	// serve UI
	if s.Flavor == "redoc" {
		// redoc-style UI
		opts := middleware.RedocOpts{
			BasePath: s.BasePath,
			SpecURL:  path.Join(specPath, s.SpecJSON),
			Path:     path.Join(specPath, s.DocURL),
			Title:    s.Titles[specIndex],
		}

		if s.UIURL != "" {
			opts.RedocURL = s.UIURL
			// default RedocURL is baked in go-openapi/runtime/middleware
		}

		handler = middleware.Redoc(opts, handler)

		visit := path.Join(specPath, s.DocURL)
		mux.Handle(visit, handler)
		log.Println("serving UI docs at", visit)
	} else if s.UIURL != "" || s.Flavor == "swagger" {
		var visit string
		// swaggerUI-style UI
		// TODO: test
		if s.UIURL == "" {
			// external UI resource
			visit = defaultSwaggerUIURL
		} else {
			visit = s.UIURL
		}
		u, err := url.Parse(visit)
		if err != nil {
			return err
		}
		q := u.Query()
		// TODO: https setup
		q.Add("url", fmt.Sprintf("http://%s:%d%s", s.host, s.port, path.Join(specPath, s.SpecJSON)))
		u.RawQuery = q.Encode()
		visit = u.String()
		log.Println("serving UI docs at ", visit)
	}
	return nil
}

func (s *ServeCmd) addHandlerSpec(mux *http.ServeMux, specPath string, specDoc *loads.Document) error {
	// addHandlerSpec adds a handler and route to serve the spec

	// the raw spec
	b, err := json.MarshalIndent(specDoc.Spec(), "", "  ")
	if err != nil {
		return err
	}

	handler := http.NotFoundHandler()

	// TODO: serve as YAML
	visit := path.Join(specPath, s.SpecJSON)
	if s.NoCors {
		log.Printf("serving spec with CORS disabled at %s", visit)
		mux.Handle(visit, middleware.Spec(specPath, b, handler))
	} else {
		// enable CORS to serve Spec
		log.Printf("serving spec with CORS enabled at %s", visit)
		mux.Handle(visit, handlers.CORS()(middleware.Spec(specPath, b, handler)))
	}
	return nil
}

func (s *ServeCmd) addHandlerLocalUI(mux *http.ServeMux) error {
	// addHandlerLocalUI serves UI resources from local files
	u, err := url.Parse(s.UIURL)
	if err != nil {
		return err
	}
	if u.Hostname() == s.host && u.Port() == strconv.Itoa(s.port) && s.UILocal != "" {
		// TODO: or no hostname or not URI
		// UIURL points to this server: serve local file or dir as UI
		log.Printf("serving local UI at %s from files in %s", s.UIURL, s.UILocal)
		mux.Handle(u.Path, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			http.ServeFile(rw, r, s.UILocal)
		}))
	}
	return nil
}

func (s *ServeCmd) isRouteUnique(specPath string, specIndex int) bool {
	// isRouteUnique skips duplicates when specs cannot be discriminated by their base paths
	if _, exists := s.basePaths[specPath]; exists {
		if specIndex > len(s.Titles)-1 {
			s.Titles = append(s.Titles, "")
		}
		return false
	}
	s.basePaths[specPath] = true
	return true
}
