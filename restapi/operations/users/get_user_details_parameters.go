package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserDetailsParams creates a new GetUserDetailsParams object
// with the default values initialized.
func NewGetUserDetailsParams() GetUserDetailsParams {
	var (
		limitDefault = int64(20)
		sinceDefault = int64(0)
	)
	return GetUserDetailsParams{
		Limit: &limitDefault,

		Since: &sinceDefault,
	}
}

// GetUserDetailsParams contains all the bound params for the get user details operation
// typically these are obtained from a http.Request
//
// swagger:parameters getUserDetails
type GetUserDetailsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Limit
	  In: query
	  Default: 20
	*/
	Limit *int64
	/*Offset
	  In: query
	  Default: 0
	*/
	Since *int64
	/*User ID
	  Required: true
	  In: path
	*/
	UserID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetUserDetailsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qSince, qhkSince, _ := qs.GetOK("since")
	if err := o.bindSince(qSince, qhkSince, route.Formats); err != nil {
		res = append(res, err)
	}

	rUserID, rhkUserID, _ := route.Params.GetOK("user_id")
	if err := o.bindUserID(rUserID, rhkUserID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserDetailsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var limitDefault int64 = int64(20)
		o.Limit = &limitDefault
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

func (o *GetUserDetailsParams) bindSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		var sinceDefault int64 = int64(0)
		o.Since = &sinceDefault
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("since", "query", "int64", raw)
	}
	o.Since = &value

	return nil
}

func (o *GetUserDetailsParams) bindUserID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("user_id", "path", "int64", raw)
	}
	o.UserID = value

	return nil
}
