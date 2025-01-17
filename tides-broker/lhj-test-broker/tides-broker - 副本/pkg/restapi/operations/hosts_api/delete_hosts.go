// Code generated by go-swagger; DO NOT EDIT.

package hosts_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteHostsHandlerFunc turns a function with the right signature into a delete hosts handler
type DeleteHostsHandlerFunc func(DeleteHostsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteHostsHandlerFunc) Handle(params DeleteHostsParams) middleware.Responder {
	return fn(params)
}

// DeleteHostsHandler interface for that can handle valid delete hosts params
type DeleteHostsHandler interface {
	Handle(DeleteHostsParams) middleware.Responder
}

// NewDeleteHosts creates a new http.Handler for the delete hosts operation
func NewDeleteHosts(ctx *middleware.Context, handler DeleteHostsHandler) *DeleteHosts {
	return &DeleteHosts{Context: ctx, Handler: handler}
}

/* DeleteHosts swagger:route POST /deleteHosts hosts_api deleteHosts

delete a host

*/
type DeleteHosts struct {
	Context *middleware.Context
	Handler DeleteHostsHandler
}

func (o *DeleteHosts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteHostsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteHostsBadRequestBody delete hosts bad request body
//
// swagger:model DeleteHostsBadRequestBody
type DeleteHostsBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete hosts bad request body
func (o *DeleteHostsBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete hosts bad request body based on context it is used
func (o *DeleteHostsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHostsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHostsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteHostsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteHostsBody local machine registration form
//
// swagger:model DeleteHostsBody
type DeleteHostsBody struct {

	// hostname
	Hostname string `json:"hostname,omitempty"`
}

// Validate validates this delete hosts body
func (o *DeleteHostsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete hosts body based on context it is used
func (o *DeleteHostsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHostsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHostsBody) UnmarshalBinary(b []byte) error {
	var res DeleteHostsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteHostsOKBody delete hosts o k body
//
// swagger:model DeleteHostsOKBody
type DeleteHostsOKBody struct {

	// hostname
	Hostname string `json:"hostname,omitempty"`
}

// Validate validates this delete hosts o k body
func (o *DeleteHostsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete hosts o k body based on context it is used
func (o *DeleteHostsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteHostsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteHostsOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteHostsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
