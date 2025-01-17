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

// UpdateApplicationInstanceHandlerFunc turns a function with the right signature into a update application instance handler
type UpdateApplicationInstanceHandlerFunc func(UpdateApplicationInstanceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateApplicationInstanceHandlerFunc) Handle(params UpdateApplicationInstanceParams) middleware.Responder {
	return fn(params)
}

// UpdateApplicationInstanceHandler interface for that can handle valid update application instance params
type UpdateApplicationInstanceHandler interface {
	Handle(UpdateApplicationInstanceParams) middleware.Responder
}

// NewUpdateApplicationInstance creates a new http.Handler for the update application instance operation
func NewUpdateApplicationInstance(ctx *middleware.Context, handler UpdateApplicationInstanceHandler) *UpdateApplicationInstance {
	return &UpdateApplicationInstance{Context: ctx, Handler: handler}
}

/* UpdateApplicationInstance swagger:route PUT /application/instance application updateApplicationInstance

update application instance

*/
type UpdateApplicationInstance struct {
	Context *middleware.Context
	Handler UpdateApplicationInstanceHandler
}

func (o *UpdateApplicationInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateApplicationInstanceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateApplicationInstanceBody update application instance body
//
// swagger:model UpdateApplicationInstanceBody
type UpdateApplicationInstanceBody struct {

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

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this update application instance body
func (o *UpdateApplicationInstanceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update application instance body based on context it is used
func (o *UpdateApplicationInstanceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateApplicationInstanceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateApplicationInstanceBody) UnmarshalBinary(b []byte) error {
	var res UpdateApplicationInstanceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateApplicationInstanceForbiddenBody update application instance forbidden body
//
// swagger:model UpdateApplicationInstanceForbiddenBody
type UpdateApplicationInstanceForbiddenBody struct {

	// msg
	Msg string `json:"msg,omitempty"`
}

// Validate validates this update application instance forbidden body
func (o *UpdateApplicationInstanceForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update application instance forbidden body based on context it is used
func (o *UpdateApplicationInstanceForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateApplicationInstanceForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateApplicationInstanceForbiddenBody) UnmarshalBinary(b []byte) error {
	var res UpdateApplicationInstanceForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateApplicationInstanceOKBody update application instance o k body
//
// swagger:model UpdateApplicationInstanceOKBody
type UpdateApplicationInstanceOKBody struct {

	// link
	Link string `json:"link,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this update application instance o k body
func (o *UpdateApplicationInstanceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update application instance o k body based on context it is used
func (o *UpdateApplicationInstanceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateApplicationInstanceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateApplicationInstanceOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateApplicationInstanceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
