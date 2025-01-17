// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateApplicationInstanceOKCode is the HTTP code returned for type UpdateApplicationInstanceOK
const UpdateApplicationInstanceOKCode int = 200

/*UpdateApplicationInstanceOK OK

swagger:response updateApplicationInstanceOK
*/
type UpdateApplicationInstanceOK struct {

	/*
	  In: Body
	*/
	Payload *UpdateApplicationInstanceOKBody `json:"body,omitempty"`
}

// NewUpdateApplicationInstanceOK creates UpdateApplicationInstanceOK with default headers values
func NewUpdateApplicationInstanceOK() *UpdateApplicationInstanceOK {

	return &UpdateApplicationInstanceOK{}
}

// WithPayload adds the payload to the update application instance o k response
func (o *UpdateApplicationInstanceOK) WithPayload(payload *UpdateApplicationInstanceOKBody) *UpdateApplicationInstanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update application instance o k response
func (o *UpdateApplicationInstanceOK) SetPayload(payload *UpdateApplicationInstanceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateApplicationInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateApplicationInstanceUnauthorizedCode is the HTTP code returned for type UpdateApplicationInstanceUnauthorized
const UpdateApplicationInstanceUnauthorizedCode int = 401

/*UpdateApplicationInstanceUnauthorized Unauthorized

swagger:response updateApplicationInstanceUnauthorized
*/
type UpdateApplicationInstanceUnauthorized struct {
}

// NewUpdateApplicationInstanceUnauthorized creates UpdateApplicationInstanceUnauthorized with default headers values
func NewUpdateApplicationInstanceUnauthorized() *UpdateApplicationInstanceUnauthorized {

	return &UpdateApplicationInstanceUnauthorized{}
}

// WriteResponse to the client
func (o *UpdateApplicationInstanceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// UpdateApplicationInstanceForbiddenCode is the HTTP code returned for type UpdateApplicationInstanceForbidden
const UpdateApplicationInstanceForbiddenCode int = 403

/*UpdateApplicationInstanceForbidden Forbidden

swagger:response updateApplicationInstanceForbidden
*/
type UpdateApplicationInstanceForbidden struct {

	/*
	  In: Body
	*/
	Payload *UpdateApplicationInstanceForbiddenBody `json:"body,omitempty"`
}

// NewUpdateApplicationInstanceForbidden creates UpdateApplicationInstanceForbidden with default headers values
func NewUpdateApplicationInstanceForbidden() *UpdateApplicationInstanceForbidden {

	return &UpdateApplicationInstanceForbidden{}
}

// WithPayload adds the payload to the update application instance forbidden response
func (o *UpdateApplicationInstanceForbidden) WithPayload(payload *UpdateApplicationInstanceForbiddenBody) *UpdateApplicationInstanceForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update application instance forbidden response
func (o *UpdateApplicationInstanceForbidden) SetPayload(payload *UpdateApplicationInstanceForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateApplicationInstanceForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
