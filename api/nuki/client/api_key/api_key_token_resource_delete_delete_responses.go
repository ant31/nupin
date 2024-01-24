// Code generated by go-swagger; DO NOT EDIT.

package api_key

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// APIKeyTokenResourceDeleteDeleteReader is a Reader for the APIKeyTokenResourceDeleteDelete structure.
type APIKeyTokenResourceDeleteDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIKeyTokenResourceDeleteDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAPIKeyTokenResourceDeleteDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAPIKeyTokenResourceDeleteDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /api/key/{apiKeyId}/token/{id}] ApiKeyTokenResource_delete_delete", response, response.Code())
	}
}

// NewAPIKeyTokenResourceDeleteDeleteNoContent creates a APIKeyTokenResourceDeleteDeleteNoContent with default headers values
func NewAPIKeyTokenResourceDeleteDeleteNoContent() *APIKeyTokenResourceDeleteDeleteNoContent {
	return &APIKeyTokenResourceDeleteDeleteNoContent{}
}

/*
APIKeyTokenResourceDeleteDeleteNoContent describes a response with status code 204, with default header values.

Ok
*/
type APIKeyTokenResourceDeleteDeleteNoContent struct {
}

// IsSuccess returns true when this api key token resource delete delete no content response has a 2xx status code
func (o *APIKeyTokenResourceDeleteDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this api key token resource delete delete no content response has a 3xx status code
func (o *APIKeyTokenResourceDeleteDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api key token resource delete delete no content response has a 4xx status code
func (o *APIKeyTokenResourceDeleteDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this api key token resource delete delete no content response has a 5xx status code
func (o *APIKeyTokenResourceDeleteDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this api key token resource delete delete no content response a status code equal to that given
func (o *APIKeyTokenResourceDeleteDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the api key token resource delete delete no content response
func (o *APIKeyTokenResourceDeleteDeleteNoContent) Code() int {
	return 204
}

func (o *APIKeyTokenResourceDeleteDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/key/{apiKeyId}/token/{id}][%d] apiKeyTokenResourceDeleteDeleteNoContent ", 204)
}

func (o *APIKeyTokenResourceDeleteDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/key/{apiKeyId}/token/{id}][%d] apiKeyTokenResourceDeleteDeleteNoContent ", 204)
}

func (o *APIKeyTokenResourceDeleteDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAPIKeyTokenResourceDeleteDeleteUnauthorized creates a APIKeyTokenResourceDeleteDeleteUnauthorized with default headers values
func NewAPIKeyTokenResourceDeleteDeleteUnauthorized() *APIKeyTokenResourceDeleteDeleteUnauthorized {
	return &APIKeyTokenResourceDeleteDeleteUnauthorized{}
}

/*
APIKeyTokenResourceDeleteDeleteUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type APIKeyTokenResourceDeleteDeleteUnauthorized struct {
}

// IsSuccess returns true when this api key token resource delete delete unauthorized response has a 2xx status code
func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this api key token resource delete delete unauthorized response has a 3xx status code
func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this api key token resource delete delete unauthorized response has a 4xx status code
func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this api key token resource delete delete unauthorized response has a 5xx status code
func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this api key token resource delete delete unauthorized response a status code equal to that given
func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the api key token resource delete delete unauthorized response
func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) Code() int {
	return 401
}

func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/key/{apiKeyId}/token/{id}][%d] apiKeyTokenResourceDeleteDeleteUnauthorized ", 401)
}

func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/key/{apiKeyId}/token/{id}][%d] apiKeyTokenResourceDeleteDeleteUnauthorized ", 401)
}

func (o *APIKeyTokenResourceDeleteDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}