// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new subscription API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscription API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SubscriptionResourceGetGet(params *SubscriptionResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionResourceGetGetOK, error)

	SubscriptionsResourceGetGet(params *SubscriptionsResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionsResourceGetGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SubscriptionResourceGetGet gets a subscription
*/
func (a *Client) SubscriptionResourceGetGet(params *SubscriptionResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SubscriptionResource_get_get",
		Method:             "GET",
		PathPattern:        "/subscription/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SubscriptionResourceGetGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SubscriptionResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubscriptionsResourceGetGet gets a list of subscriptions
*/
func (a *Client) SubscriptionsResourceGetGet(params *SubscriptionsResourceGetGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionsResourceGetGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionsResourceGetGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SubscriptionsResource_get_get",
		Method:             "GET",
		PathPattern:        "/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SubscriptionsResourceGetGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionsResourceGetGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SubscriptionsResource_get_get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}