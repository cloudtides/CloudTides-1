// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UploadInstanceFileOKCode is the HTTP code returned for type UploadInstanceFileOK
const UploadInstanceFileOKCode int = 200

/*UploadInstanceFileOK OK

swagger:response uploadInstanceFileOK
*/
type UploadInstanceFileOK struct {

	/*
	  In: Body
	*/
	Payload *UploadInstanceFileOKBody `json:"body,omitempty"`
}

// NewUploadInstanceFileOK creates UploadInstanceFileOK with default headers values
func NewUploadInstanceFileOK() *UploadInstanceFileOK {

	return &UploadInstanceFileOK{}
}

// WithPayload adds the payload to the upload instance file o k response
func (o *UploadInstanceFileOK) WithPayload(payload *UploadInstanceFileOKBody) *UploadInstanceFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload instance file o k response
func (o *UploadInstanceFileOK) SetPayload(payload *UploadInstanceFileOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadInstanceFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
