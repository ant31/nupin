// Code generated by go-swagger; DO NOT EDIT.

package smartlock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SmartlockSyncResourcePostPostReader is a Reader for the SmartlockSyncResourcePostPost structure.
type SmartlockSyncResourcePostPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SmartlockSyncResourcePostPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSmartlockSyncResourcePostPostNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSmartlockSyncResourcePostPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSmartlockSyncResourcePostPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /smartlock/{smartlockId}/sync] SmartlockSyncResource_post_post", response, response.Code())
	}
}

// NewSmartlockSyncResourcePostPostNoContent creates a SmartlockSyncResourcePostPostNoContent with default headers values
func NewSmartlockSyncResourcePostPostNoContent() *SmartlockSyncResourcePostPostNoContent {
	return &SmartlockSyncResourcePostPostNoContent{}
}

/*
SmartlockSyncResourcePostPostNoContent describes a response with status code 204, with default header values.

Ok
*/
type SmartlockSyncResourcePostPostNoContent struct {
}

// IsSuccess returns true when this smartlock sync resource post post no content response has a 2xx status code
func (o *SmartlockSyncResourcePostPostNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this smartlock sync resource post post no content response has a 3xx status code
func (o *SmartlockSyncResourcePostPostNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock sync resource post post no content response has a 4xx status code
func (o *SmartlockSyncResourcePostPostNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this smartlock sync resource post post no content response has a 5xx status code
func (o *SmartlockSyncResourcePostPostNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock sync resource post post no content response a status code equal to that given
func (o *SmartlockSyncResourcePostPostNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the smartlock sync resource post post no content response
func (o *SmartlockSyncResourcePostPostNoContent) Code() int {
	return 204
}

func (o *SmartlockSyncResourcePostPostNoContent) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/sync][%d] smartlockSyncResourcePostPostNoContent ", 204)
}

func (o *SmartlockSyncResourcePostPostNoContent) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/sync][%d] smartlockSyncResourcePostPostNoContent ", 204)
}

func (o *SmartlockSyncResourcePostPostNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockSyncResourcePostPostBadRequest creates a SmartlockSyncResourcePostPostBadRequest with default headers values
func NewSmartlockSyncResourcePostPostBadRequest() *SmartlockSyncResourcePostPostBadRequest {
	return &SmartlockSyncResourcePostPostBadRequest{}
}

/*
SmartlockSyncResourcePostPostBadRequest describes a response with status code 400, with default header values.

Bad Parameter
*/
type SmartlockSyncResourcePostPostBadRequest struct {
}

// IsSuccess returns true when this smartlock sync resource post post bad request response has a 2xx status code
func (o *SmartlockSyncResourcePostPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock sync resource post post bad request response has a 3xx status code
func (o *SmartlockSyncResourcePostPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock sync resource post post bad request response has a 4xx status code
func (o *SmartlockSyncResourcePostPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock sync resource post post bad request response has a 5xx status code
func (o *SmartlockSyncResourcePostPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock sync resource post post bad request response a status code equal to that given
func (o *SmartlockSyncResourcePostPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the smartlock sync resource post post bad request response
func (o *SmartlockSyncResourcePostPostBadRequest) Code() int {
	return 400
}

func (o *SmartlockSyncResourcePostPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/sync][%d] smartlockSyncResourcePostPostBadRequest ", 400)
}

func (o *SmartlockSyncResourcePostPostBadRequest) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/sync][%d] smartlockSyncResourcePostPostBadRequest ", 400)
}

func (o *SmartlockSyncResourcePostPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSmartlockSyncResourcePostPostUnauthorized creates a SmartlockSyncResourcePostPostUnauthorized with default headers values
func NewSmartlockSyncResourcePostPostUnauthorized() *SmartlockSyncResourcePostPostUnauthorized {
	return &SmartlockSyncResourcePostPostUnauthorized{}
}

/*
SmartlockSyncResourcePostPostUnauthorized describes a response with status code 401, with default header values.

Not authorized
*/
type SmartlockSyncResourcePostPostUnauthorized struct {
}

// IsSuccess returns true when this smartlock sync resource post post unauthorized response has a 2xx status code
func (o *SmartlockSyncResourcePostPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this smartlock sync resource post post unauthorized response has a 3xx status code
func (o *SmartlockSyncResourcePostPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this smartlock sync resource post post unauthorized response has a 4xx status code
func (o *SmartlockSyncResourcePostPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this smartlock sync resource post post unauthorized response has a 5xx status code
func (o *SmartlockSyncResourcePostPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this smartlock sync resource post post unauthorized response a status code equal to that given
func (o *SmartlockSyncResourcePostPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the smartlock sync resource post post unauthorized response
func (o *SmartlockSyncResourcePostPostUnauthorized) Code() int {
	return 401
}

func (o *SmartlockSyncResourcePostPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/sync][%d] smartlockSyncResourcePostPostUnauthorized ", 401)
}

func (o *SmartlockSyncResourcePostPostUnauthorized) String() string {
	return fmt.Sprintf("[POST /smartlock/{smartlockId}/sync][%d] smartlockSyncResourcePostPostUnauthorized ", 401)
}

func (o *SmartlockSyncResourcePostPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}