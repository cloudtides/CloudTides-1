// Code generated by go-swagger; DO NOT EDIT.

package hosts_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetHostInfoHandlerFunc turns a function with the right signature into a get host info handler
type GetHostInfoHandlerFunc func(GetHostInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHostInfoHandlerFunc) Handle(params GetHostInfoParams) middleware.Responder {
	return fn(params)
}

// GetHostInfoHandler interface for that can handle valid get host info params
type GetHostInfoHandler interface {
	Handle(GetHostInfoParams) middleware.Responder
}

// NewGetHostInfo creates a new http.Handler for the get host info operation
func NewGetHostInfo(ctx *middleware.Context, handler GetHostInfoHandler) *GetHostInfo {
	return &GetHostInfo{Context: ctx, Handler: handler}
}

/* GetHostInfo swagger:route GET /hostInfo hosts_api getHostInfo

GetHostInfo get host info API

*/
type GetHostInfo struct {
	Context *middleware.Context
	Handler GetHostInfoHandler
}

func (o *GetHostInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHostInfoParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetHostInfoBadRequestBody get host info bad request body
//
// swagger:model GetHostInfoBadRequestBody
type GetHostInfoBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get host info bad request body
func (o *GetHostInfoBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get host info bad request body based on context it is used
func (o *GetHostInfoBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHostInfoBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHostInfoBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetHostInfoBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetHostInfoOKBody get host info o k body
//
// swagger:model GetHostInfoOKBody
type GetHostInfoOKBody struct {

	// current RAM
	CurrentRAM float32 `json:"currentRAM,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// ip
	// Min Length: 1
	IP string `json:"ip,omitempty"`

	// percent CPU
	PercentCPU float32 `json:"percentCPU,omitempty"`

	// percent RAM
	PercentRAM float32 `json:"percentRAM,omitempty"`

	// port
	Port string `json:"port,omitempty"`

	// total RAM
	TotalRAM float32 `json:"totalRAM,omitempty"`
}

// Validate validates this get host info o k body
func (o *GetHostInfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHostInfoOKBody) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(o.IP) { // not required
		return nil
	}

	if err := validate.MinLength("getHostInfoOK"+"."+"ip", "body", o.IP, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get host info o k body based on context it is used
func (o *GetHostInfoOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHostInfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHostInfoOKBody) UnmarshalBinary(b []byte) error {
	var res GetHostInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
