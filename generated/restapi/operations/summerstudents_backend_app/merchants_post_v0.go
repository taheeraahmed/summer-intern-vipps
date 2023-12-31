// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MerchantsPostV0HandlerFunc turns a function with the right signature into a merchants post v0 handler
type MerchantsPostV0HandlerFunc func(MerchantsPostV0Params) middleware.Responder

// Handle executing the request and returning a response
func (fn MerchantsPostV0HandlerFunc) Handle(params MerchantsPostV0Params) middleware.Responder {
	return fn(params)
}

// MerchantsPostV0Handler interface for that can handle valid merchants post v0 params
type MerchantsPostV0Handler interface {
	Handle(MerchantsPostV0Params) middleware.Responder
}

// NewMerchantsPostV0 creates a new http.Handler for the merchants post v0 operation
func NewMerchantsPostV0(ctx *middleware.Context, handler MerchantsPostV0Handler) *MerchantsPostV0 {
	return &MerchantsPostV0{Context: ctx, Handler: handler}
}

/*
	MerchantsPostV0 swagger:route POST /v0/merchants summerstudents-backend-app v0-api merchantsPostV0

# Post a merchant

Post a merchant to the database
*/
type MerchantsPostV0 struct {
	Context *middleware.Context
	Handler MerchantsPostV0Handler
}

func (o *MerchantsPostV0) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewMerchantsPostV0Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
