// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vippsas/summerstudents-backend/generated/models"
)

// RecurringAgreementsGetV0OKCode is the HTTP code returned for type RecurringAgreementsGetV0OK
const RecurringAgreementsGetV0OKCode int = 200

/*
RecurringAgreementsGetV0OK Success

swagger:response recurringAgreementsGetV0OK
*/
type RecurringAgreementsGetV0OK struct {

	/*
	  In: Body
	*/
	Payload *models.AgreementReturn `json:"body,omitempty"`
}

// NewRecurringAgreementsGetV0OK creates RecurringAgreementsGetV0OK with default headers values
func NewRecurringAgreementsGetV0OK() *RecurringAgreementsGetV0OK {

	return &RecurringAgreementsGetV0OK{}
}

// WithPayload adds the payload to the recurring agreements get v0 o k response
func (o *RecurringAgreementsGetV0OK) WithPayload(payload *models.AgreementReturn) *RecurringAgreementsGetV0OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recurring agreements get v0 o k response
func (o *RecurringAgreementsGetV0OK) SetPayload(payload *models.AgreementReturn) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RecurringAgreementsGetV0OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RecurringAgreementsGetV0NotFoundCode is the HTTP code returned for type RecurringAgreementsGetV0NotFound
const RecurringAgreementsGetV0NotFoundCode int = 404

/*
RecurringAgreementsGetV0NotFound Not found

swagger:response recurringAgreementsGetV0NotFound
*/
type RecurringAgreementsGetV0NotFound struct {
}

// NewRecurringAgreementsGetV0NotFound creates RecurringAgreementsGetV0NotFound with default headers values
func NewRecurringAgreementsGetV0NotFound() *RecurringAgreementsGetV0NotFound {

	return &RecurringAgreementsGetV0NotFound{}
}

// WriteResponse to the client
func (o *RecurringAgreementsGetV0NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// RecurringAgreementsGetV0InternalServerErrorCode is the HTTP code returned for type RecurringAgreementsGetV0InternalServerError
const RecurringAgreementsGetV0InternalServerErrorCode int = 500

/*
RecurringAgreementsGetV0InternalServerError Internal server error

swagger:response recurringAgreementsGetV0InternalServerError
*/
type RecurringAgreementsGetV0InternalServerError struct {
}

// NewRecurringAgreementsGetV0InternalServerError creates RecurringAgreementsGetV0InternalServerError with default headers values
func NewRecurringAgreementsGetV0InternalServerError() *RecurringAgreementsGetV0InternalServerError {

	return &RecurringAgreementsGetV0InternalServerError{}
}

// WriteResponse to the client
func (o *RecurringAgreementsGetV0InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
