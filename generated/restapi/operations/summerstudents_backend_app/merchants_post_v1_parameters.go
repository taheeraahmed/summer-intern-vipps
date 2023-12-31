// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/vippsas/summerstudents-backend/generated/models"
)

// NewMerchantsPostV1Params creates a new MerchantsPostV1Params object
//
// There are no default values defined in the spec.
func NewMerchantsPostV1Params() MerchantsPostV1Params {

	return MerchantsPostV1Params{}
}

// MerchantsPostV1Params contains all the bound params for the merchants post v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters MerchantsPostV1
type MerchantsPostV1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The merchant to be posted
	  In: body
	*/
	Merchant *models.Merchant
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewMerchantsPostV1Params() beforehand.
func (o *MerchantsPostV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Merchant
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("merchant", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Merchant = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
