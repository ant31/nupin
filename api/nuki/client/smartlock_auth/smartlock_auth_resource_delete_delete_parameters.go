// Code generated by go-swagger; DO NOT EDIT.

package smartlock_auth

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

// NewSmartlockAuthResourceDeleteDeleteParams creates a new SmartlockAuthResourceDeleteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlockAuthResourceDeleteDeleteParams() *SmartlockAuthResourceDeleteDeleteParams {
	return &SmartlockAuthResourceDeleteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlockAuthResourceDeleteDeleteParamsWithTimeout creates a new SmartlockAuthResourceDeleteDeleteParams object
// with the ability to set a timeout on a request.
func NewSmartlockAuthResourceDeleteDeleteParamsWithTimeout(timeout time.Duration) *SmartlockAuthResourceDeleteDeleteParams {
	return &SmartlockAuthResourceDeleteDeleteParams{
		timeout: timeout,
	}
}

// NewSmartlockAuthResourceDeleteDeleteParamsWithContext creates a new SmartlockAuthResourceDeleteDeleteParams object
// with the ability to set a context for a request.
func NewSmartlockAuthResourceDeleteDeleteParamsWithContext(ctx context.Context) *SmartlockAuthResourceDeleteDeleteParams {
	return &SmartlockAuthResourceDeleteDeleteParams{
		Context: ctx,
	}
}

// NewSmartlockAuthResourceDeleteDeleteParamsWithHTTPClient creates a new SmartlockAuthResourceDeleteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlockAuthResourceDeleteDeleteParamsWithHTTPClient(client *http.Client) *SmartlockAuthResourceDeleteDeleteParams {
	return &SmartlockAuthResourceDeleteDeleteParams{
		HTTPClient: client,
	}
}

/*
SmartlockAuthResourceDeleteDeleteParams contains all the parameters to send to the API endpoint

	for the smartlock auth resource delete delete operation.

	Typically these are written to a http.Request.
*/
type SmartlockAuthResourceDeleteDeleteParams struct {

	/* ID.

	   The smartlock authorization unique id
	*/
	ID string

	/* SmartlockID.

	   The smartlock id
	*/
	SmartlockID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlock auth resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockAuthResourceDeleteDeleteParams) WithDefaults() *SmartlockAuthResourceDeleteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlock auth resource delete delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockAuthResourceDeleteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) WithTimeout(timeout time.Duration) *SmartlockAuthResourceDeleteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) WithContext(ctx context.Context) *SmartlockAuthResourceDeleteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) WithHTTPClient(client *http.Client) *SmartlockAuthResourceDeleteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) WithID(id string) *SmartlockAuthResourceDeleteDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) SetID(id string) {
	o.ID = id
}

// WithSmartlockID adds the smartlockID to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) WithSmartlockID(smartlockID int64) *SmartlockAuthResourceDeleteDeleteParams {
	o.SetSmartlockID(smartlockID)
	return o
}

// SetSmartlockID adds the smartlockId to the smartlock auth resource delete delete params
func (o *SmartlockAuthResourceDeleteDeleteParams) SetSmartlockID(smartlockID int64) {
	o.SmartlockID = smartlockID
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlockAuthResourceDeleteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param smartlockId
	if err := r.SetPathParam("smartlockId", swag.FormatInt64(o.SmartlockID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}