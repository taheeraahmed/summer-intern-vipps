// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vippsas/summerstudents-backend/generated/models"
)

// MerchantsPatchV0OKCode is the HTTP code returned for type MerchantsPatchV0OK
const MerchantsPatchV0OKCode int = 200

/*
MerchantsPatchV0OK Success

swagger:response merchantsPatchV0OK
*/
type MerchantsPatchV0OK struct {

	/*
	  In: Body
	*/
	Payload *models.Merchant `json:"body,omitempty"`
}

// NewMerchantsPatchV0OK creates MerchantsPatchV0OK with default headers values
func NewMerchantsPatchV0OK() *MerchantsPatchV0OK {

	return &MerchantsPatchV0OK{}
}

// WithPayload adds the payload to the merchants patch v0 o k response
func (o *MerchantsPatchV0OK) WithPayload(payload *models.Merchant) *MerchantsPatchV0OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the merchants patch v0 o k response
func (o *MerchantsPatchV0OK) SetPayload(payload *models.Merchant) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MerchantsPatchV0OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
