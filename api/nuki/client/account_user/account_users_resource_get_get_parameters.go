// Code generated by go-swagger; DO NOT EDIT.

package account_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewAccountUsersResourceGetGetParams creates a new AccountUsersResourceGetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountUsersResourceGetGetParams() *AccountUsersResourceGetGetParams {
	return &AccountUsersResourceGetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountUsersResourceGetGetParamsWithTimeout creates a new AccountUsersResourceGetGetParams object
// with the ability to set a timeout on a request.
func NewAccountUsersResourceGetGetParamsWithTimeout(timeout time.Duration) *AccountUsersResourceGetGetParams {
	return &AccountUsersResourceGetGetParams{
		timeout: timeout,
	}
}

// NewAccountUsersResourceGetGetParamsWithContext creates a new AccountUsersResourceGetGetParams object
// with the ability to set a context for a request.
func NewAccountUsersResourceGetGetParamsWithContext(ctx context.Context) *AccountUsersResourceGetGetParams {
	return &AccountUsersResourceGetGetParams{
		Context: ctx,
	}
}

// NewAccountUsersResourceGetGetParamsWithHTTPClient creates a new AccountUsersResourceGetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountUsersResourceGetGetParamsWithHTTPClient(client *http.Client) *AccountUsersResourceGetGetParams {
	return &AccountUsersResourceGetGetParams{
		HTTPClient: client,
	}
}

/*
AccountUsersResourceGetGetParams contains all the parameters to send to the API endpoint

	for the account users resource get get operation.

	Typically these are written to a http.Request.
*/
type AccountUsersResourceGetGetParams struct {

	/* Email.

	   Filter for email
	*/
	Email *string

	/* Limit.

	   The maximum number of users to return. If the value exceeds the maximum, then the maximum value will be used.
	*/
	Limit *int64

	/* Offset.

	   The offset of the first user in the collection to return
	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account users resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountUsersResourceGetGetParams) WithDefaults() *AccountUsersResourceGetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account users resource get get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountUsersResourceGetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) WithTimeout(timeout time.Duration) *AccountUsersResourceGetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) WithContext(ctx context.Context) *AccountUsersResourceGetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) WithHTTPClient(client *http.Client) *AccountUsersResourceGetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) WithEmail(email *string) *AccountUsersResourceGetGetParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) SetEmail(email *string) {
	o.Email = email
}

// WithLimit adds the limit to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) WithLimit(limit *int64) *AccountUsersResourceGetGetParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) WithOffset(offset *int64) *AccountUsersResourceGetGetParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the account users resource get get params
func (o *AccountUsersResourceGetGetParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *AccountUsersResourceGetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Email != nil {

		// query param email
		var qrEmail string

		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {

			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}