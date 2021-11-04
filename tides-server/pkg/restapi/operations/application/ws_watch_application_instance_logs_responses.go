// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WsWatchApplicationInstanceLogsOKCode is the HTTP code returned for type WsWatchApplicationInstanceLogsOK
const WsWatchApplicationInstanceLogsOKCode int = 200

/*WsWatchApplicationInstanceLogsOK OK

swagger:response wsWatchApplicationInstanceLogsOK
*/
type WsWatchApplicationInstanceLogsOK struct {
}

// NewWsWatchApplicationInstanceLogsOK creates WsWatchApplicationInstanceLogsOK with default headers values
func NewWsWatchApplicationInstanceLogsOK() *WsWatchApplicationInstanceLogsOK {

	return &WsWatchApplicationInstanceLogsOK{}
}

// WriteResponse to the client
func (o *WsWatchApplicationInstanceLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// WsWatchApplicationInstanceLogsUnauthorizedCode is the HTTP code returned for type WsWatchApplicationInstanceLogsUnauthorized
const WsWatchApplicationInstanceLogsUnauthorizedCode int = 401

/*WsWatchApplicationInstanceLogsUnauthorized Unauthorized

swagger:response wsWatchApplicationInstanceLogsUnauthorized
*/
type WsWatchApplicationInstanceLogsUnauthorized struct {
}

// NewWsWatchApplicationInstanceLogsUnauthorized creates WsWatchApplicationInstanceLogsUnauthorized with default headers values
func NewWsWatchApplicationInstanceLogsUnauthorized() *WsWatchApplicationInstanceLogsUnauthorized {

	return &WsWatchApplicationInstanceLogsUnauthorized{}
}

// WriteResponse to the client
func (o *WsWatchApplicationInstanceLogsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WsWatchApplicationInstanceLogsForbiddenCode is the HTTP code returned for type WsWatchApplicationInstanceLogsForbidden
const WsWatchApplicationInstanceLogsForbiddenCode int = 403

/*WsWatchApplicationInstanceLogsForbidden Forbidden

swagger:response wsWatchApplicationInstanceLogsForbidden
*/
type WsWatchApplicationInstanceLogsForbidden struct {
}

// NewWsWatchApplicationInstanceLogsForbidden creates WsWatchApplicationInstanceLogsForbidden with default headers values
func NewWsWatchApplicationInstanceLogsForbidden() *WsWatchApplicationInstanceLogsForbidden {

	return &WsWatchApplicationInstanceLogsForbidden{}
}

// WriteResponse to the client
func (o *WsWatchApplicationInstanceLogsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
