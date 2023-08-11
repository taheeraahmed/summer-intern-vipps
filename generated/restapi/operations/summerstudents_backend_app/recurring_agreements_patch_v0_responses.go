// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vippsas/summerstudents-backend/generated/models"
)

// RecurringAgreementsPatchV0OKCode is the HTTP code returned for type RecurringAgreementsPatchV0OK
const RecurringAgreementsPatchV0OKCode int = 200

/*
RecurringAgreementsPatchV0OK Success

swagger:response recurringAgreementsPatchV0OK
*/
type RecurringAgreementsPatchV0OK struct {

	/*
	  In: Body
	*/
	Payload *models.Agreement `json:"body,omitempty"`
}

// NewRecurringAgreementsPatchV0OK creates RecurringAgreementsPatchV0OK with default headers values
func NewRecurringAgreementsPatchV0OK() *RecurringAgreementsPatchV0OK {

	return &RecurringAgreementsPatchV0OK{}
}

// WithPayload adds the payload to the recurring agreements patch v0 o k response
func (o *RecurringAgreementsPatchV0OK) WithPayload(payload *models.Agreement) *RecurringAgreementsPatchV0OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recurring agreements patch v0 o k response
func (o *RecurringAgreementsPatchV0OK) SetPayload(payload *models.Agreement) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RecurringAgreementsPatchV0OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}