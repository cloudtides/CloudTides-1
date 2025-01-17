// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewDownInstanceFileParams creates a new DownInstanceFileParams object
//
// There are no default values defined in the spec.
func NewDownInstanceFileParams() DownInstanceFileParams {

	return DownInstanceFileParams{}
}

// DownInstanceFileParams contains all the bound params for the down instance file operation
// typically these are obtained from a http.Request
//
// swagger:parameters downInstanceFile
type DownInstanceFileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Name string
	/*
	  Required: true
	  In: path
	*/
	Token string
	/*
	  Required: true
	  In: path
	*/
	UID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDownInstanceFileParams() beforehand.
func (o *DownInstanceFileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rToken, rhkToken, _ := route.Params.GetOK("token")
	if err := o.bindToken(rToken, rhkToken, route.Formats); err != nil {
		res = append(res, err)
	}

	rUID, rhkUID, _ := route.Params.GetOK("uid")
	if err := o.bindUID(rUID, rhkUID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from path.
func (o *DownInstanceFileParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Name = raw

	return nil
}

// bindToken binds and validates parameter Token from path.
func (o *DownInstanceFileParams) bindToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Token = raw

	return nil
}

// bindUID binds and validates parameter UID from path.
func (o *DownInstanceFileParams) bindUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.UID = raw

	return nil
}
