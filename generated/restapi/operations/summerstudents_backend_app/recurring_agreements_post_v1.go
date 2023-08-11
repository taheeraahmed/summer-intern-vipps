// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RecurringAgreementsPostV1HandlerFunc turns a function with the right signature into a recurring agreements post v1 handler
type RecurringAgreementsPostV1HandlerFunc func(RecurringAgreementsPostV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn RecurringAgreementsPostV1HandlerFunc) Handle(params RecurringAgreementsPostV1Params) middleware.Responder {
	return fn(params)
}

// RecurringAgreementsPostV1Handler interface for that can handle valid recurring agreements post v1 params
type RecurringAgreementsPostV1Handler interface {
	Handle(RecurringAgreementsPostV1Params) middleware.Responder
}

// NewRecurringAgreementsPostV1 creates a new http.Handler for the recurring agreements post v1 operation
func NewRecurringAgreementsPostV1(ctx *middleware.Context, handler RecurringAgreementsPostV1Handler) *RecurringAgreementsPostV1 {
	return &RecurringAgreementsPostV1{Context: ctx, Handler: handler}
}

/*
	RecurringAgreementsPostV1 swagger:route POST /v1/recurring-agreements summerstudents-backend-app Recurrizz v1 endpoints recurringAgreementsPostV1

# Posting recurring agreement

Try to created a recurring agreement in the recurring api
*/
type RecurringAgreementsPostV1 struct {
	Context *middleware.Context
	Handler RecurringAgreementsPostV1Handler
}

func (o *RecurringAgreementsPostV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRecurringAgreementsPostV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}