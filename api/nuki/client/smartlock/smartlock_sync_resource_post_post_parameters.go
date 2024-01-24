// Code generated by go-swagger; DO NOT EDIT.

package smartlock

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
)

// NewSmartlockSyncResourcePostPostParams creates a new SmartlockSyncResourcePostPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSmartlockSyncResourcePostPostParams() *SmartlockSyncResourcePostPostParams {
	return &SmartlockSyncResourcePostPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSmartlockSyncResourcePostPostParamsWithTimeout creates a new SmartlockSyncResourcePostPostParams object
// with the ability to set a timeout on a request.
func NewSmartlockSyncResourcePostPostParamsWithTimeout(timeout time.Duration) *SmartlockSyncResourcePostPostParams {
	return &SmartlockSyncResourcePostPostParams{
		timeout: timeout,
	}
}

// NewSmartlockSyncResourcePostPostParamsWithContext creates a new SmartlockSyncResourcePostPostParams object
// with the ability to set a context for a request.
func NewSmartlockSyncResourcePostPostParamsWithContext(ctx context.Context) *SmartlockSyncResourcePostPostParams {
	return &SmartlockSyncResourcePostPostParams{
		Context: ctx,
	}
}

// NewSmartlockSyncResourcePostPostParamsWithHTTPClient creates a new SmartlockSyncResourcePostPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewSmartlockSyncResourcePostPostParamsWithHTTPClient(client *http.Client) *SmartlockSyncResourcePostPostParams {
	return &SmartlockSyncResourcePostPostParams{
		HTTPClient: client,
	}
}

/*
SmartlockSyncResourcePostPostParams contains all the parameters to send to the API endpoint

	for the smartlock sync resource post post operation.

	Typically these are written to a http.Request.
*/
type SmartlockSyncResourcePostPostParams struct {

	/* SmartlockID.

	   The smartlock id
	*/
	SmartlockID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the smartlock sync resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockSyncResourcePostPostParams) WithDefaults() *SmartlockSyncResourcePostPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the smartlock sync resource post post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SmartlockSyncResourcePostPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the smartlock sync resource post post params
func (o *SmartlockSyncResourcePostPostParams) WithTimeout(timeout time.Duration) *SmartlockSyncResourcePostPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the smartlock sync resource post post params
func (o *SmartlockSyncResourcePostPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the smartlock sync resource post post params
func (o *SmartlockSyncResourcePostPostParams) WithContext(ctx context.Context) *SmartlockSyncResourcePostPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the smartlock sync resource post post params
func (o *SmartlockSyncResourcePostPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the smartlock sync resource post post params
func (o *SmartlockSyncResourcePostPostParams) WithHTTPClient(client *http.Client) *SmartlockSyncResourcePostPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the smartlock sync resource post post params
func (o *SmartlockSyncResourcePostPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSmartlockID adds the smartlockID to the smartlock sync resource post post params
func (o *SmartlockSyncResourcePostPostParams) WithSmartlockID(smartlockID string) *SmartlockSyncResourcePostPostParams {
	o.SetSmartlockID(smartlockID)
	return o
}

// SetSmartlockID adds the smartlockId to the smartlock sync resource post post params
func (o *SmartlockSyncResourcePostPostParams) SetSmartlockID(smartlockID string) {
	o.SmartlockID = smartlockID
}

// WriteToRequest writes these params to a swagger request
func (o *SmartlockSyncResourcePostPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param smartlockId
	if err := r.SetPathParam("smartlockId", o.SmartlockID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}