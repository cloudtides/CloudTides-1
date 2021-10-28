// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateApplicationInstanceHandlerFunc turns a function with the right signature into a create application instance handler
type CreateApplicationInstanceHandlerFunc func(CreateApplicationInstanceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateApplicationInstanceHandlerFunc) Handle(params CreateApplicationInstanceParams) middleware.Responder {
	return fn(params)
}

// CreateApplicationInstanceHandler interface for that can handle valid create application instance params
type CreateApplicationInstanceHandler interface {
	Handle(CreateApplicationInstanceParams) middleware.Responder
}

// NewCreateApplicationInstance creates a new http.Handler for the create application instance operation
func NewCreateApplicationInstance(ctx *middleware.Context, handler CreateApplicationInstanceHandler) *CreateApplicationInstance {
	return &CreateApplicationInstance{Context: ctx, Handler: handler}
}

/* CreateApplicationInstance swagger:route POST /application/instance application createApplicationInstance

create application instance

*/
type CreateApplicationInstance struct {
	Context *middleware.Context
	Handler CreateApplicationInstanceHandler
}

func (o *CreateApplicationInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateApplicationInstanceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// CreateApplicationInstanceBody create application instance body
//
// swagger:model CreateApplicationInstanceBody
type CreateApplicationInstanceBody struct {

	// app type
	AppType string `json:"appType,omitempty"`

	// instance name
	InstanceName string `json:"instanceName,omitempty"`

	// port
	Port string `json:"port,omitempty"`

	// ssh host
	SSHHost string `json:"sshHost,omitempty"`

	// ssh password
	SSHPassword string `json:"sshPassword,omitempty"`

	// ssh port
	SSHPort int64 `json:"sshPort,omitempty"`

	// ssh user
	SSHUser string `json:"sshUser,omitempty"`
}

// Validate validates this create application instance body
func (o *CreateApplicationInstanceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create application instance body based on context it is used
func (o *CreateApplicationInstanceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateApplicationInstanceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateApplicationInstanceBody) UnmarshalBinary(b []byte) error {
	var res CreateApplicationInstanceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateApplicationInstanceForbiddenBody create application instance forbidden body
//
// swagger:model CreateApplicationInstanceForbiddenBody
type CreateApplicationInstanceForbiddenBody struct {

	// msg
	Msg string `json:"msg,omitempty"`
}

// Validate validates this create application instance forbidden body
func (o *CreateApplicationInstanceForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create application instance forbidden body based on context it is used
func (o *CreateApplicationInstanceForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateApplicationInstanceForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateApplicationInstanceForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CreateApplicationInstanceForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// CreateApplicationInstanceOKBody create application instance o k body
//
// swagger:model CreateApplicationInstanceOKBody
type CreateApplicationInstanceOKBody struct {

	// link
	Link string `json:"link,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this create application instance o k body
func (o *CreateApplicationInstanceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create application instance o k body based on context it is used
func (o *CreateApplicationInstanceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateApplicationInstanceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateApplicationInstanceOKBody) UnmarshalBinary(b []byte) error {
	var res CreateApplicationInstanceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
