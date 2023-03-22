// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// InstanceActionStatueOKCode is the HTTP code returned for type InstanceActionStatueOK
const InstanceActionStatueOKCode int = 200

/*InstanceActionStatueOK OK

swagger:response instanceActionStatueOK
*/
type InstanceActionStatueOK struct {
}

// NewInstanceActionStatueOK creates InstanceActionStatueOK with default headers values
func NewInstanceActionStatueOK() *InstanceActionStatueOK {

	return &InstanceActionStatueOK{}
}

// WriteResponse to the client
func (o *InstanceActionStatueOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// InstanceActionStatueUnauthorizedCode is the HTTP code returned for type InstanceActionStatueUnauthorized
const InstanceActionStatueUnauthorizedCode int = 401

/*InstanceActionStatueUnauthorized Unauthorized

swagger:response instanceActionStatueUnauthorized
*/
type InstanceActionStatueUnauthorized struct {
}

// NewInstanceActionStatueUnauthorized creates InstanceActionStatueUnauthorized with default headers values
func NewInstanceActionStatueUnauthorized() *InstanceActionStatueUnauthorized {

	return &InstanceActionStatueUnauthorized{}
}

// WriteResponse to the client
func (o *InstanceActionStatueUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// InstanceActionStatueForbiddenCode is the HTTP code returned for type InstanceActionStatueForbidden
const InstanceActionStatueForbiddenCode int = 403

/*InstanceActionStatueForbidden Forbidden

swagger:response instanceActionStatueForbidden
*/
type InstanceActionStatueForbidden struct {

	/*
	  In: Body
	*/
	Payload *InstanceActionStatueForbiddenBody `json:"body,omitempty"`
}

// NewInstanceActionStatueForbidden creates InstanceActionStatueForbidden with default headers values
func NewInstanceActionStatueForbidden() *InstanceActionStatueForbidden {

	return &InstanceActionStatueForbidden{}
}

// WithPayload adds the payload to the instance action statue forbidden response
func (o *InstanceActionStatueForbidden) WithPayload(payload *InstanceActionStatueForbiddenBody) *InstanceActionStatueForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the instance action statue forbidden response
func (o *InstanceActionStatueForbidden) SetPayload(payload *InstanceActionStatueForbiddenBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InstanceActionStatueForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}