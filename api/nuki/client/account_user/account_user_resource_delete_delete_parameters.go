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

// NewAccountUserResourceDeleteDeleteParams creates a new AccountUserResourceDeleteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountUserResourceDeleteDeleteParams() *AccountUserResourceDeleteDeleteParams {
	return &AccountUserResourceDeleteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountUserResourceDeleteDeleteParamsWithTimeout creates a new AccountUserResourceDeleteDeleteParams object
// with the ability to set a timeout on a request.
func NewAccountUserResourceDeleteDeleteParamsWithTimeout(timeout time.Duration) *AccountUserResourceDeleteDeleteParams {
	return &AccountUserResourceDeleteDeleteParams{
		timeout: timeout,
	}
}

// NewAccountUserResourceDeleteDeleteParamsWithContext creates a new AccountUserResourceDeleteDeleteParams object
// with the ability to set a context for a request.
func NewAccountUserResourceDeleteDeleteParamsWithContext(ctx context.Context) *AccountUserResourceDeleteDeleteParams {
	return &AccountUserResourceDeleteDeleteParams{
		Context: ctx,
	}
}

// NewAccountUserResourceDeleteDeleteParamsWithHTTPClient creates a new AccountUserResourceDeleteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountUserResourceDeleteDeleteParamsWithHTTPClient(client *http.Client) *AccountUserResourceDeleteDeleteParams {
	return &AccountUserResourceDeleteDeleteParams{
		HTTPClient: client,
	}
}

/*
AccountUserResourceDeleteDeleteParams contains all the parameters to send to the API endpoint

	for the account user resource delete delete operation.

	Typically these are written to a http.Request.
*/
type AccountUserResourceDeleteDeleteParams struct {

	/* AccountUserID.

	   The account user id
	*/
	AccountUserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account user resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountUserResourceDeleteDeleteParams) WithDefaults() *AccountUserResourceDeleteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account user resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountUserResourceDeleteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account user resource delete delete params
func (o *AccountUserResourceDeleteDeleteParams) WithTimeout(timeout time.Duration) *AccountUserResourceDeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account user resource delete delete params
func (o *AccountUserResourceDeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account user resource delete delete params
func (o *AccountUserResourceDeleteDeleteParams) WithContext(ctx context.Context) *AccountUserResourceDeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account user resource delete delete params
func (o *AccountUserResourceDeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account user resource delete delete params
func (o *AccountUserResourceDeleteDeleteParams) WithHTTPClient(client *http.Client) *AccountUserResourceDeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account user resource delete delete params
func (o *AccountUserResourceDeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountUserID adds the accountUserID to the account user resource delete delete params
func (o *AccountUserResourceDeleteDeleteParams) WithAccountUserID(accountUserID int64) *AccountUserResourceDeleteDeleteParams {
	o.SetAccountUserID(accountUserID)
	return o
}

// SetAccountUserID adds the accountUserId to the account user resource delete delete params
func (o *AccountUserResourceDeleteDeleteParams) SetAccountUserID(accountUserID int64) {
	o.AccountUserID = accountUserID
}

// WriteToRequest writes these params to a swagger request
func (o *AccountUserResourceDeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountUserId
	if err := r.SetPathParam("accountUserId", swag.FormatInt64(o.AccountUserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}