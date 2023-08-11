// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/vippsas/summerstudents-backend/internal/app"
	v0_handlers "github.com/vippsas/summerstudents-backend/internal/handlers/v0"
	v1_handlers "github.com/vippsas/summerstudents-backend/internal/handlers/v1"

	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/vippsas/summerstudents-backend/internal/handlers"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/vippsas/summerstudents-backend/generated/restapi/operations"
)

//go:generate swagger generate server --target ../../generated --name SummerstudentsBackend --spec ../../manifests/k8s/base/openApiSwagger/summerstudents-backend-apim.yaml --principal interface{}

func configureFlags(api *operations.SummerstudentsBackendAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SummerstudentsBackendAPI) http.Handler {
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

	ap := app.New()

	// v0 handlers
	api.SummerstudentsBackendAppMerchantsGetV0Handler = v0_handlers.MerchantsGetHandler(ap)
	api.SummerstudentsBackendAppMerchantsGetAllV0Handler = v0_handlers.MerchantsGetAllHandler(ap)
	api.SummerstudentsBackendAppMerchantsPatchV0Handler = v0_handlers.MerchantsPatchHandler(ap)
	api.SummerstudentsBackendAppMerchantsPostV0Handler = v0_handlers.MerchantsPostHandler(ap)

	api.SummerstudentsBackendAppRecurringAgreementsGetV0Handler = v0_handlers.RecurringAgreementsGetHandler(ap)
	api.SummerstudentsBackendAppRecurringAgreementsPatchV0Handler = v0_handlers.RecurringAgreementsPatchHandler(ap)
	api.SummerstudentsBackendAppRecurringAgreementsPostV0Handler = v0_handlers.RecurringAgreementsPostHandler(ap)

	api.SummerstudentsBackendAppCustomerAgreementsGetAllV0Handler = v0_handlers.CustomerAgreementsGetAllHandler(ap)
	api.SummerstudentsBackendAppMerchantAgreementsGetAllV0Handler = v0_handlers.MerchantAgreementsGetAllHandler(ap)
	// v1 handlers
	api.SummerstudentsBackendAppMerchantsGetV1Handler = v1_handlers.MerchantsGetHandler(ap)
	api.SummerstudentsBackendAppMerchantsGetAllV1Handler = v1_handlers.MerchantsGetAllHandler(ap)
	api.SummerstudentsBackendAppMerchantsPatchV1Handler = v1_handlers.MerchantsPatchHandler(ap)
	api.SummerstudentsBackendAppMerchantsPostV1Handler = v1_handlers.MerchantsPostHandler(ap)

	api.SummerstudentsBackendAppRecurringAgreementsPostV1Handler = v1_handlers.RecurringAgreementsPostHandler(ap)
	api.SummerstudentsBackendAppRecurringAgreementsPatchV1Handler = v1_handlers.RecurringAgreementsPatchHandler(ap)
	api.SummerstudentsBackendAppRecurringAgreementsGetV1Handler = v1_handlers.RecurringAgreementsGetHandler(ap)

	api.SummerstudentsBackendAppCustomerAgreementsGetAllV1Handler = v1_handlers.CustomerAgreementsGetAllHandler(ap)
	api.SummerstudentsBackendAppMerchantAgreementsGetAllV1Handler = v1_handlers.MerchantAgreementsGetAllHandler(ap)

	// Use localhost if running in dev environment, otherwise use the service name
	var host string
	if ap.Config.IsLocal() {
		host = "localhost"
	} else {
		host = "0.0.0.0"
	}

	healthSrv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, 8081),
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}
	go serveHealth(healthSrv, ap.Logger)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	})

	handler := c.Handler(api.Serve(setupMiddlewares))

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	//return setupGlobalMiddleware(api.Serve(setupMiddlewares))
	return handler
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
	return handler
}

func serveHealth(srv *http.Server, log logrus.FieldLogger) {
	healthHandler := http.NewServeMux()
	healthHandler.HandleFunc("/health", handlers.HealthHandler)
	srv.Handler = healthHandler
	log.WithFields(logrus.Fields{"event": "starting.health.server"}).Infof("Starting health server on http://%s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.WithError(err).WithField("event", "error.health.server.start").
			Fatal("Unable to start health http server")
	}
}
