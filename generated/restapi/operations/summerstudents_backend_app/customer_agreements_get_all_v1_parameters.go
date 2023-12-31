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

// NewCustomerAgreementsGetAllV1Params creates a new CustomerAgreementsGetAllV1Params object
//
// There are no default values defined in the spec.
func NewCustomerAgreementsGetAllV1Params() CustomerAgreementsGetAllV1Params {

	return CustomerAgreementsGetAllV1Params{}
}

// CustomerAgreementsGetAllV1Params contains all the bound params for the customer agreements get all v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters CustomerAgreementsGetAllV1
type CustomerAgreementsGetAllV1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Id of the customer
	  Required: true
	  In: path
	*/
	CustomerID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCustomerAgreementsGetAllV1Params() beforehand.
func (o *CustomerAgreementsGetAllV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rCustomerID, rhkCustomerID, _ := route.Params.GetOK("customerId")
	if err := o.bindCustomerID(rCustomerID, rhkCustomerID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCustomerID binds and validates parameter CustomerID from path.
func (o *CustomerAgreementsGetAllV1Params) bindCustomerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("customerId", "path", "int64", raw)
	}
	o.CustomerID = value

	return nil
}
