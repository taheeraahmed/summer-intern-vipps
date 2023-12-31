// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vippsas/summerstudents-backend/generated/models"
)

// RecurringAgreementsPostV1OKCode is the HTTP code returned for type RecurringAgreementsPostV1OK
const RecurringAgreementsPostV1OKCode int = 200

/*
RecurringAgreementsPostV1OK Success

swagger:response recurringAgreementsPostV1OK
*/
type RecurringAgreementsPostV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.AgreementReturn `json:"body,omitempty"`
}

// NewRecurringAgreementsPostV1OK creates RecurringAgreementsPostV1OK with default headers values
func NewRecurringAgreementsPostV1OK() *RecurringAgreementsPostV1OK {

	return &RecurringAgreementsPostV1OK{}
}

// WithPayload adds the payload to the recurring agreements post v1 o k response
func (o *RecurringAgreementsPostV1OK) WithPayload(payload *models.AgreementReturn) *RecurringAgreementsPostV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the recurring agreements post v1 o k response
func (o *RecurringAgreementsPostV1OK) SetPayload(payload *models.AgreementReturn) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RecurringAgreementsPostV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RecurringAgreementsPostV1NotFoundCode is the HTTP code returned for type RecurringAgreementsPostV1NotFound
const RecurringAgreementsPostV1NotFoundCode int = 404

/*
RecurringAgreementsPostV1NotFound Not found

swagger:response recurringAgreementsPostV1NotFound
*/
type RecurringAgreementsPostV1NotFound struct {
}

// NewRecurringAgreementsPostV1NotFound creates RecurringAgreementsPostV1NotFound with default headers values
func NewRecurringAgreementsPostV1NotFound() *RecurringAgreementsPostV1NotFound {

	return &RecurringAgreementsPostV1NotFound{}
}

// WriteResponse to the client
func (o *RecurringAgreementsPostV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// RecurringAgreementsPostV1InternalServerErrorCode is the HTTP code returned for type RecurringAgreementsPostV1InternalServerError
const RecurringAgreementsPostV1InternalServerErrorCode int = 500

/*
RecurringAgreementsPostV1InternalServerError Internal server error

swagger:response recurringAgreementsPostV1InternalServerError
*/
type RecurringAgreementsPostV1InternalServerError struct {
}

// NewRecurringAgreementsPostV1InternalServerError creates RecurringAgreementsPostV1InternalServerError with default headers values
func NewRecurringAgreementsPostV1InternalServerError() *RecurringAgreementsPostV1InternalServerError {

	return &RecurringAgreementsPostV1InternalServerError{}
}

// WriteResponse to the client
func (o *RecurringAgreementsPostV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
