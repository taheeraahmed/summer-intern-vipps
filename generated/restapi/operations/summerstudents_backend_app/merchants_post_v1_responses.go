// Code generated by go-swagger; DO NOT EDIT.

package summerstudents_backend_app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// MerchantsPostV1CreatedCode is the HTTP code returned for type MerchantsPostV1Created
const MerchantsPostV1CreatedCode int = 201

/*
MerchantsPostV1Created Success

swagger:response merchantsPostV1Created
*/
type MerchantsPostV1Created struct {
}

// NewMerchantsPostV1Created creates MerchantsPostV1Created with default headers values
func NewMerchantsPostV1Created() *MerchantsPostV1Created {

	return &MerchantsPostV1Created{}
}

// WriteResponse to the client
func (o *MerchantsPostV1Created) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// MerchantsPostV1BadRequestCode is the HTTP code returned for type MerchantsPostV1BadRequest
const MerchantsPostV1BadRequestCode int = 400

/*
MerchantsPostV1BadRequest Bad request

swagger:response merchantsPostV1BadRequest
*/
type MerchantsPostV1BadRequest struct {
}

// NewMerchantsPostV1BadRequest creates MerchantsPostV1BadRequest with default headers values
func NewMerchantsPostV1BadRequest() *MerchantsPostV1BadRequest {

	return &MerchantsPostV1BadRequest{}
}

// WriteResponse to the client
func (o *MerchantsPostV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// MerchantsPostV1ConflictCode is the HTTP code returned for type MerchantsPostV1Conflict
const MerchantsPostV1ConflictCode int = 409

/*
MerchantsPostV1Conflict Conflict

swagger:response merchantsPostV1Conflict
*/
type MerchantsPostV1Conflict struct {
}

// NewMerchantsPostV1Conflict creates MerchantsPostV1Conflict with default headers values
func NewMerchantsPostV1Conflict() *MerchantsPostV1Conflict {

	return &MerchantsPostV1Conflict{}
}

// WriteResponse to the client
func (o *MerchantsPostV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// MerchantsPostV1InternalServerErrorCode is the HTTP code returned for type MerchantsPostV1InternalServerError
const MerchantsPostV1InternalServerErrorCode int = 500

/*
MerchantsPostV1InternalServerError Internal server error

swagger:response merchantsPostV1InternalServerError
*/
type MerchantsPostV1InternalServerError struct {
}

// NewMerchantsPostV1InternalServerError creates MerchantsPostV1InternalServerError with default headers values
func NewMerchantsPostV1InternalServerError() *MerchantsPostV1InternalServerError {

	return &MerchantsPostV1InternalServerError{}
}

// WriteResponse to the client
func (o *MerchantsPostV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
