// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListInstanceFilesOKCode is the HTTP code returned for type ListInstanceFilesOK
const ListInstanceFilesOKCode int = 200

/*ListInstanceFilesOK OK

swagger:response listInstanceFilesOK
*/
type ListInstanceFilesOK struct {

	/*
	  In: Body
	*/
	Payload []*ListInstanceFilesOKBodyItems0 `json:"body,omitempty"`
}

// NewListInstanceFilesOK creates ListInstanceFilesOK with default headers values
func NewListInstanceFilesOK() *ListInstanceFilesOK {

	return &ListInstanceFilesOK{}
}

// WithPayload adds the payload to the list instance files o k response
func (o *ListInstanceFilesOK) WithPayload(payload []*ListInstanceFilesOKBodyItems0) *ListInstanceFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list instance files o k response
func (o *ListInstanceFilesOK) SetPayload(payload []*ListInstanceFilesOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListInstanceFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ListInstanceFilesOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
