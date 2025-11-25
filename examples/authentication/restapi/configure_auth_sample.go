// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/go-swagger/go-swagger/examples/authentication/models"
	"github.com/go-swagger/go-swagger/examples/authentication/restapi/operations"
	"github.com/go-swagger/go-swagger/examples/authentication/restapi/operations/customers"
)

//go:generate swagger generate server --target ../../authentication --name AuthSample --spec ../swagger.yml --principal models.Principal

func configureFlags(api *operations.AuthSampleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
	_ = api
}

func configureAPI(api *operations.AuthSampleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-token" header is set
	api.KeyAuth = func(token string) (*models.Principal, error) {
		if token == "abcdefuvwxyz" {
			prin := models.Principal(token)
			return &prin, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", token)
		return nil, errors.New(401, "incorrect api key auth")
	}

	api.CustomersCreateHandler = customers.CreateHandlerFunc(func(params customers.CreateParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation customers.Create has not yet been implemented")
	})
	api.CustomersGetIDHandler = customers.GetIDHandlerFunc(func(params customers.GetIDParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation customers.GetID has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
	_ = tlsConfig
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(server *http.Server, scheme, addr string) {
	_ = server
	_ = scheme
	_ = addr
}

type cachedContext struct {
	Data string
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("setting context after routing")
		ctx := r.Context()
		v1 := ctx.Value("globalCtx")
		log.Printf("retrieved global context: %v", v1)
		setupCtx := context.WithValue(r.Context(), "setupCtx", "my-setup-data")

		*r = *r.WithContext(setupCtx)

		handler.ServeHTTP(w, r)

		v2 := ctx.Value("cachedCtx")
		cached, ok := v2.(*cachedContext)
		if !ok {
			log.Printf("did not find cachedCtx")
			return
		}

		cached.Data = "updated cache"
		cachedCtx := context.WithValue(r.Context(), "cachedCtx", cached)

		log.Printf("setting context after serving")
		afterCtx := context.WithValue(cachedCtx, "afterCtx", "my-after-data")

		*r = *r.WithContext(afterCtx)
	})
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// global context is available down the entire stack
		globalCtx := context.WithValue(r.Context(), "globalCtx", "my-global-data")
		cachedCtx := context.WithValue(globalCtx, "cachedCtx", &cachedContext{Data: "cached"})
		log.Printf("setting context in globalMiddleware")
		//*r = *r.WithContext(cachedCtx)

		handler.ServeHTTP(w, r.WithContext(cachedCtx))

		log.Println("final stack")
		ctx := r.Context()

		v1 := ctx.Value("globalCtx")
		v2 := ctx.Value("setupCtx")
		v3 := ctx.Value("afterCtx")
		v4 := ctx.Value("finalCtx")
		v5 := ctx.Value("cachedCtx")

		log.Println(v1)
		log.Println(v2)
		log.Println(v3)
		log.Println(v4)
		log.Println(v5)
	})
}
