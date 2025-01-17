// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CreateApplicationInstanceOKCode is the HTTP code returned for type CreateApplicationInstanceOK
const CreateApplicationInstanceOKCode int = 200

/*CreateApplicationInstanceOK OK

swagger:response createApplicationInstanceOK
*/
type CreateApplicationInstanceOK struct {

	/*
	  In: Body
	*/
	Payload *CreateApplicationInstanceOKBody `json:"body,omitempty"`
}

// NewCreateApplicationInstanceOK creates CreateApplicationInstanceOK with default headers values
func NewCreateApplicationInstanceOK() *CreateApplicationInstanceOK {

	return &CreateApplicationInstanceOK{}
}

// WithPayload adds the payload to the create application instance o k response
func (o *CreateApplicationInstanceOK) WithPayload(payload *CreateApplicationInstanceOKBody) *CreateApplicationInstanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create application instance o k response
func (o *CreateApplicationInstanceOK) SetPayload(payload *CreateApplicationInstanceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateApplicationInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateApplicationInstanceUnauthorizedCode is the HTTP code returned for type CreateApplicationInstanceUnauthorized
const CreateApplicationInstanceUnauthorizedCode int = 401

/*CreateApplicationInstanceUnauthorized Unauthorized

swagger:response createApplicationInstanceUnauthorized
*/
type CreateApplicationInstanceUnauthorized struct {
}

// NewCreateApplicationInstanceUnauthorized creates CreateApplicationInstanceUnauthorized with default headers values
func NewCreateApplicationInstanceUnauthorized() *CreateApplicationInstanceUnauthorized {

	return &CreateApplicationInstanceUnauthorized{}
}

// WriteResponse to the client
func (o *CreateApplicationInstanceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// CreateApplicationInstanceForbiddenCode is the HTTP code returned for type CreateApplicationInstanceForbidden
const CreateApplicationInstanceForbiddenCode int = 403

/*CreateApplicationInstanceForbidden Forbidden

swagger:response createApplicationInstanceForbidden
*/
type CreateApplicationInstanceForbidden struct {

	/*
	  In: Body
	*/
	Payload *CreateApplicationInstanceForbiddenBody `json:"body,omitempty"`
}

// NewCreateApplicationInstanceForbidden creates CreateApplicationInstanceForbidden with default headers values
func NewCreateApplicationInstanceForbidden() *CreateApplicationInstanceForbidden {

	return &CreateApplicationInstanceForbidden{}
}

// WithPayload adds the payload to the create application instance forbidden response
func (o *CreateApplicationInstanceForbidden) WithPayload(payload *CreateApplicationInstanceForbiddenBody) *CreateApplicationInstanceForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create application instance forbidden response
func (o *CreateApplicationInstanceForbidden) SetPayload(payload *CreateApplicationInstanceForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateApplicationInstanceForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
