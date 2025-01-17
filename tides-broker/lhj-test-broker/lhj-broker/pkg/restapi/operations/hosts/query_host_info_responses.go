// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"lhj-broker/pkg/models"
)

// QueryHostInfoOKCode is the HTTP code returned for type QueryHostInfoOK
const QueryHostInfoOKCode int = 200

/*QueryHostInfoOK list the hosts' info

swagger:response queryHostInfoOK
*/
type QueryHostInfoOK struct {

	/*
	  In: Body
	*/
	Payload []*QueryHostInfoOKBodyItems0 `json:"body,omitempty"`
}

// NewQueryHostInfoOK creates QueryHostInfoOK with default headers values
func NewQueryHostInfoOK() *QueryHostInfoOK {

	return &QueryHostInfoOK{}
}

// WithPayload adds the payload to the query host info o k response
func (o *QueryHostInfoOK) WithPayload(payload []*QueryHostInfoOKBodyItems0) *QueryHostInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the query host info o k response
func (o *QueryHostInfoOK) SetPayload(payload []*QueryHostInfoOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QueryHostInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*QueryHostInfoOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*QueryHostInfoDefault generic error response

swagger:response queryHostInfoDefault
*/
type QueryHostInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewQueryHostInfoDefault creates QueryHostInfoDefault with default headers values
func NewQueryHostInfoDefault(code int) *QueryHostInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &QueryHostInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the query host info default response
func (o *QueryHostInfoDefault) WithStatusCode(code int) *QueryHostInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the query host info default response
func (o *QueryHostInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the query host info default response
func (o *QueryHostInfoDefault) WithPayload(payload *models.Error) *QueryHostInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the query host info default response
func (o *QueryHostInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *QueryHostInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
