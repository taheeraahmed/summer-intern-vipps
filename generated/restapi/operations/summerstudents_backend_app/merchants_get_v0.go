// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MerchantsGetV0HandlerFunc turns a function with the right signature into a merchants get v0 handler
type MerchantsGetV0HandlerFunc func(MerchantsGetV0Params) middleware.Responder

// Handle executing the request and returning a response
func (fn MerchantsGetV0HandlerFunc) Handle(params MerchantsGetV0Params) middleware.Responder {
	return fn(params)
}

// MerchantsGetV0Handler interface for that can handle valid merchants get v0 params
type MerchantsGetV0Handler interface {
	Handle(MerchantsGetV0Params) middleware.Responder
}

// NewMerchantsGetV0 creates a new http.Handler for the merchants get v0 operation
func NewMerchantsGetV0(ctx *middleware.Context, handler MerchantsGetV0Handler) *MerchantsGetV0 {
	return &MerchantsGetV0{Context: ctx, Handler: handler}
}

/*
	MerchantsGetV0 swagger:route GET /v0/merchants/{vippsnummer} summerstudents-backend-app v0-api merchantsGetV0

# Get a merchant given vippsnummer

Given a vippsnumber a merchant will be returned
*/
type MerchantsGetV0 struct {
	Context *middleware.Context
	Handler MerchantsGetV0Handler
}

func (o *MerchantsGetV0) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewMerchantsGetV0Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}