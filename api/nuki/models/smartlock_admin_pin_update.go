// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SmartlockAdminPinUpdate smartlock admin pin update
//
// swagger:model SmartlockAdminPinUpdate
type SmartlockAdminPinUpdate struct {

	// The admin pin
	// Required: true
	// Maximum: 9999
	// Minimum: 0
	AdminPin *int32 `json:"adminPin"`
}

// Validate validates this smartlock admin pin update
func (m *SmartlockAdminPinUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminPin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartlockAdminPinUpdate) validateAdminPin(formats strfmt.Registry) error {

	if err := validate.Required("adminPin", "body", m.AdminPin); err != nil {
		return err
	}

	if err := validate.MinimumInt("adminPin", "body", int64(*m.AdminPin), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("adminPin", "body", int64(*m.AdminPin), 9999, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this smartlock admin pin update based on context it is used
func (m *SmartlockAdminPinUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SmartlockAdminPinUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmartlockAdminPinUpdate) UnmarshalBinary(b []byte) error {
	var res SmartlockAdminPinUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}