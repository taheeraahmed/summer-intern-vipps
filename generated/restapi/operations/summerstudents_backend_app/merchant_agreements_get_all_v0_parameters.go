// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewMerchantAgreementsGetAllV0Params creates a new MerchantAgreementsGetAllV0Params object
//
// There are no default values defined in the spec.
func NewMerchantAgreementsGetAllV0Params() MerchantAgreementsGetAllV0Params {

	return MerchantAgreementsGetAllV0Params{}
}

// MerchantAgreementsGetAllV0Params contains all the bound params for the merchant agreements get all v0 operation
// typically these are obtained from a http.Request
//
// swagger:parameters MerchantAgreementsGetAllV0
type MerchantAgreementsGetAllV0Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Vippsummer of merchant
	  Required: true
	  In: path
	*/
	Vippsnummer int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewMerchantAgreementsGetAllV0Params() beforehand.
func (o *MerchantAgreementsGetAllV0Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rVippsnummer, rhkVippsnummer, _ := route.Params.GetOK("vippsnummer")
	if err := o.bindVippsnummer(rVippsnummer, rhkVippsnummer, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindVippsnummer binds and validates parameter Vippsnummer from path.
func (o *MerchantAgreementsGetAllV0Params) bindVippsnummer(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("vippsnummer", "path", "int64", raw)
	}
	o.Vippsnummer = value

	return nil
}
