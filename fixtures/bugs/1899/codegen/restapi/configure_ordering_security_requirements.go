// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/yamlpc"
	"github.com/go-openapi/swag/yamlutils"

	"github.com/go-swagger/go-swagger/fixtures/bugs/1899/codegen/restapi/operations"
	"github.com/go-swagger/go-swagger/fixtures/bugs/1899/codegen/restapi/operations/connect"
)

//go:generate swagger generate server --target ../../codegen --name OrderingSecurityRequirements --spec ../../swagger.yml --principal any

func configureFlags(api *operations.OrderingSecurityRequirementsAPI) {
	// api.CommandLineOptionsGroups = []cmdutils.CommandLineOptionsGroup{ ... }
	_ = api
}

func configureAPI(api *operations.OrderingSecurityRequirementsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...any)
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.YamlConsumer = yamlpc.YAMLConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.ConnectConfigPutHandler = connect.ConfigPutHandlerFunc(func(params connect.ConfigPutParams) middleware.Responder {
		buf, err := io.ReadAll(params.Config)
		if err != nil {
			return connect.NewConfigPutBadRequest()
		}
		log.Printf("received: %q", string(buf))

		doc, err := yamlutils.BytesToYAMLDoc(buf)
		if err != nil {
			log.Printf("bad YAML: %v", err)
			return connect.NewConfigPutBadRequest()
		}

		jazon, err := yamlutils.YAMLToJSON(doc)
		if err != nil {
			log.Printf("bad JSON: %v", err)
			return connect.NewConfigPutBadRequest()
		}

		log.Println(jazon)

		return connect.NewConfigPutNoContent()
	})

	api.PreServerShutdown = func() {}

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

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
