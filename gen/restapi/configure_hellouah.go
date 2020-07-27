// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/laqiiz/go-swagger-oauth2-security/gen/models"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/laqiiz/go-swagger-oauth2-security/gen/restapi/hellouah"
	"github.com/laqiiz/go-swagger-oauth2-security/gen/restapi/hellouah/example"
)

//go:generate swagger generate server --target ..\..\gen --name Hellouah --spec ..\..\swagger.yml --api-package hellouah --principal interface{}

func configureFlags(api *hellouah.HellouahAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *hellouah.HellouahAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GoogleOauthSecurityAuth = func(token string, scopes []string) (*models.Principal, error) {
		log.Println("KeyCloakOauthSecurityAuth path")

		// This handler is called by the runtime whenever a route needs authentication
		// against the 'OAuthSecurity' scheme.
		// It is passed a token extracted from the Authentication Bearer header, and
		// the list of scopes mentioned by the spec for this route.

		// NOTE: in this simple implementation, we do not check scopes against
		// the signed claims in the JWT token.
		// So whatever the required scope (passed a parameter by the runtime),
		// this will succeed provided we get a valid token.

		log.Println("authenticated", token)

		// authenticated validates a JWT token at userInfoURL
		ok, err := authenticated(token)
		if err != nil {
			return nil, errors.New(401, "error authenticate")
		}
		if !ok {
			return nil, errors.New(401, "invalid token")
		}

		// returns the authenticated principal (here just filled in with its token)
		prin := models.Principal(token)
		return &prin, nil
	}

	api.GetAuthCallbackHandler = hellouah.GetAuthCallbackHandlerFunc(func(params hellouah.GetAuthCallbackParams) middleware.Responder {
		log.Println("callback path")

		// implements the callback operation
		token, err := callback(params.HTTPRequest)
		if err != nil {
			return middleware.NotImplemented("operation .GetAuthCallback error")
		}
		//log.Println("Token", token)

		return middleware.ResponderFunc(func(w http.ResponseWriter, pr runtime.Producer) {
			http.SetCookie(w, &http.Cookie{
				Name:  "Authorization",
				Value: "Bearer " + token, // 行儀が悪いので真似しないねで
				Path:  "/",
			})
			http.Redirect(w, params.HTTPRequest, "/v1/hello", http.StatusFound)
		})
	})

	api.GetLoginHandler = hellouah.GetLoginHandlerFunc(func(params hellouah.GetLoginParams) middleware.Responder {
		log.Println("login path")
		return login(params.HTTPRequest)
	})

	// ⚡⚡ Implement ⚡⚡
	api.ExampleHelloHandler = example.HelloHandlerFunc(func(params example.HelloParams, principal *models.Principal) middleware.Responder {
		log.Println("hello path")
		return example.NewHelloOK().WithPayload(&models.Hello{
			Message: "hello",
		})
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
