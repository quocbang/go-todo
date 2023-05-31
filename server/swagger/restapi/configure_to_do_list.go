// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"quocbang/golang-to-do-list/impl"
	"quocbang/golang-to-do-list/impl/utils"
	"quocbang/golang-to-do-list/swagger/restapi/operations"
	"quocbang/golang-to-do-list/swagger/restapi/operations/to_do_list"
)

//go:generate swagger generate server --target ..\..\swagger --name ToDoList --spec ..\..\..\swagger.yml --principal interface{}

func configureFlags(api *operations.ToDoListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ToDoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()

	db, err := impl.NewDB("test")
	if err != nil {
		log.Fatalf("failed to register database")
	}

	impl.RegisterHandlers(api, db)
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ToDoListCreateToDoListHandler == nil {
		api.ToDoListCreateToDoListHandler = to_do_list.CreateToDoListHandlerFunc(func(params to_do_list.CreateToDoListParams) middleware.Responder {
			return middleware.NotImplemented("operation to_do_list.CreateToDoList has not yet been implemented")
		})
	}
	if api.ToDoListGetAllToDoListHandler == nil {
		api.ToDoListGetAllToDoListHandler = to_do_list.GetAllToDoListHandlerFunc(func(params to_do_list.GetAllToDoListParams) middleware.Responder {
			return middleware.NotImplemented("operation to_do_list.GetAllToDoList has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return utils.Logging(handler)

}
