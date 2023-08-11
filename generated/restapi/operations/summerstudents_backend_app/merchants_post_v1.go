// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MerchantsPostV1HandlerFunc turns a function with the right signature into a merchants post v1 handler
type MerchantsPostV1HandlerFunc func(MerchantsPostV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn MerchantsPostV1HandlerFunc) Handle(params MerchantsPostV1Params) middleware.Responder {
	return fn(params)
}

// MerchantsPostV1Handler interface for that can handle valid merchants post v1 params
type MerchantsPostV1Handler interface {
	Handle(MerchantsPostV1Params) middleware.Responder
}

// NewMerchantsPostV1 creates a new http.Handler for the merchants post v1 operation
func NewMerchantsPostV1(ctx *middleware.Context, handler MerchantsPostV1Handler) *MerchantsPostV1 {
	return &MerchantsPostV1{Context: ctx, Handler: handler}
}

/*
	MerchantsPostV1 swagger:route POST /v1/merchants summerstudents-backend-app Recurrizz v1 endpoints merchantsPostV1

# Post a merchant

Post a merchant to the database
*/
type MerchantsPostV1 struct {
	Context *middleware.Context
	Handler MerchantsPostV1Handler
}

func (o *MerchantsPostV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewMerchantsPostV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}