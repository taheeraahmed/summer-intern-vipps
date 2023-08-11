// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RecurringAgreementsGetV1HandlerFunc turns a function with the right signature into a recurring agreements get v1 handler
type RecurringAgreementsGetV1HandlerFunc func(RecurringAgreementsGetV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn RecurringAgreementsGetV1HandlerFunc) Handle(params RecurringAgreementsGetV1Params) middleware.Responder {
	return fn(params)
}

// RecurringAgreementsGetV1Handler interface for that can handle valid recurring agreements get v1 params
type RecurringAgreementsGetV1Handler interface {
	Handle(RecurringAgreementsGetV1Params) middleware.Responder
}

// NewRecurringAgreementsGetV1 creates a new http.Handler for the recurring agreements get v1 operation
func NewRecurringAgreementsGetV1(ctx *middleware.Context, handler RecurringAgreementsGetV1Handler) *RecurringAgreementsGetV1 {
	return &RecurringAgreementsGetV1{Context: ctx, Handler: handler}
}

/*
	RecurringAgreementsGetV1 swagger:route GET /v1/recurring-agreements/{agreementId} summerstudents-backend-app Recurrizz v1 endpoints recurringAgreementsGetV1

# Get recurring agreement

Given a recurrent agreement id, get the agreement
*/
type RecurringAgreementsGetV1 struct {
	Context *middleware.Context
	Handler RecurringAgreementsGetV1Handler
}

func (o *RecurringAgreementsGetV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRecurringAgreementsGetV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
