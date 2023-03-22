// Code generated by go-swagger; DO NOT EDIT.

package hosts_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddOneHostCreatedCode is the HTTP code returned for type AddOneHostCreated
const AddOneHostCreatedCode int = 201

/*AddOneHostCreated local machine registered successfully

swagger:response addOneHostCreated
*/
type AddOneHostCreated struct {

	/*
	  In: Body
	*/
	Payload *AddOneHostCreatedBody `json:"body,omitempty"`
}

// NewAddOneHostCreated creates AddOneHostCreated with default headers values
func NewAddOneHostCreated() *AddOneHostCreated {

	return &AddOneHostCreated{}
}

// WithPayload adds the payload to the add one host created response
func (o *AddOneHostCreated) WithPayload(payload *AddOneHostCreatedBody) *AddOneHostCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add one host created response
func (o *AddOneHostCreated) SetPayload(payload *AddOneHostCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddOneHostCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddOneHostBadRequestCode is the HTTP code returned for type AddOneHostBadRequest
const AddOneHostBadRequestCode int = 400

/*AddOneHostBadRequest bad request

swagger:response addOneHostBadRequest
*/
type AddOneHostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *AddOneHostBadRequestBody `json:"body,omitempty"`
}

// NewAddOneHostBadRequest creates AddOneHostBadRequest with default headers values
func NewAddOneHostBadRequest() *AddOneHostBadRequest {

	return &AddOneHostBadRequest{}
}

// WithPayload adds the payload to the add one host bad request response
func (o *AddOneHostBadRequest) WithPayload(payload *AddOneHostBadRequestBody) *AddOneHostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add one host bad request response
func (o *AddOneHostBadRequest) SetPayload(payload *AddOneHostBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddOneHostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddOneHostUnauthorizedCode is the HTTP code returned for type AddOneHostUnauthorized
const AddOneHostUnauthorizedCode int = 401

/*AddOneHostUnauthorized Unauthorized

swagger:response addOneHostUnauthorized
*/
type AddOneHostUnauthorized struct {
}

// NewAddOneHostUnauthorized creates AddOneHostUnauthorized with default headers values
func NewAddOneHostUnauthorized() *AddOneHostUnauthorized {

	return &AddOneHostUnauthorized{}
}

// WriteResponse to the client
func (o *AddOneHostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}