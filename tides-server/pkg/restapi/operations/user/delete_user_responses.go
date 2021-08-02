// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteUserOKCode is the HTTP code returned for type DeleteUserOK
const DeleteUserOKCode int = 200

/*DeleteUserOK deletion success

swagger:response deleteUserOK
*/
type DeleteUserOK struct {

	/*
	  In: Body
	*/
	Payload *DeleteUserOKBody `json:"body,omitempty"`
}

// NewDeleteUserOK creates DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {

	return &DeleteUserOK{}
}

// WithPayload adds the payload to the delete user o k response
func (o *DeleteUserOK) WithPayload(payload *DeleteUserOKBody) *DeleteUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user o k response
func (o *DeleteUserOK) SetPayload(payload *DeleteUserOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserUnauthorizedCode is the HTTP code returned for type DeleteUserUnauthorized
const DeleteUserUnauthorizedCode int = 401

/*DeleteUserUnauthorized Unauthorized

swagger:response deleteUserUnauthorized
*/
type DeleteUserUnauthorized struct {
}

// NewDeleteUserUnauthorized creates DeleteUserUnauthorized with default headers values
func NewDeleteUserUnauthorized() *DeleteUserUnauthorized {

	return &DeleteUserUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DeleteUserForbiddenCode is the HTTP code returned for type DeleteUserForbidden
const DeleteUserForbiddenCode int = 403

/*DeleteUserForbidden Forbidden

swagger:response deleteUserForbidden
*/
type DeleteUserForbidden struct {
}

// NewDeleteUserForbidden creates DeleteUserForbidden with default headers values
func NewDeleteUserForbidden() *DeleteUserForbidden {

	return &DeleteUserForbidden{}
}

// WriteResponse to the client
func (o *DeleteUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeleteUserNotFoundCode is the HTTP code returned for type DeleteUserNotFound
const DeleteUserNotFoundCode int = 404

/*DeleteUserNotFound resource not found

swagger:response deleteUserNotFound
*/
type DeleteUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *DeleteUserNotFoundBody `json:"body,omitempty"`
}

// NewDeleteUserNotFound creates DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {

	return &DeleteUserNotFound{}
}

// WithPayload adds the payload to the delete user not found response
func (o *DeleteUserNotFound) WithPayload(payload *DeleteUserNotFoundBody) *DeleteUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user not found response
func (o *DeleteUserNotFound) SetPayload(payload *DeleteUserNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}