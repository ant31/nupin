// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Application application
//
// swagger:model Application
type Application struct {

	// application
	Application *Application `json:"application,omitempty"`

	// author
	Author string `json:"author,omitempty"`

	// connector service
	ConnectorService *ConnectorService `json:"connectorService,omitempty"`

	// conneg service
	ConnegService *ConnegService `json:"connegService,omitempty"`

	// context
	Context *Context `json:"context,omitempty"`

	// converter service
	ConverterService *ConverterService `json:"converterService,omitempty"`

	// debugging
	Debugging bool `json:"debugging,omitempty"`

	// decoder service
	DecoderService *DecoderService `json:"decoderService,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// encoder service
	EncoderService *EncoderService `json:"encoderService,omitempty"`

	// inbound root
	InboundRoot *Restlet `json:"inboundRoot,omitempty"`

	// logger
	Logger *Logger `json:"logger,omitempty"`

	// metadata service
	MetadataService *MetadataService `json:"metadataService,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// outbound root
	OutboundRoot *Restlet `json:"outboundRoot,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// range service
	RangeService *RangeService `json:"rangeService,omitempty"`

	// roles
	Roles []*Role `json:"roles"`

	// services
	Services []*Service `json:"services"`

	// started
	Started bool `json:"started,omitempty"`

	// status service
	StatusService *StatusService `json:"statusService,omitempty"`

	// stopped
	Stopped bool `json:"stopped,omitempty"`

	// task service
	TaskService *TaskService `json:"taskService,omitempty"`

	// tunnel service
	TunnelService *TunnelService `json:"tunnelService,omitempty"`
}

// Validate validates this application
func (m *Application) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectorService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnegService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConverterService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecoderService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncoderService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInboundRoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogger(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadataService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutboundRoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRangeService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunnelService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Application) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateConnectorService(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectorService) { // not required
		return nil
	}

	if m.ConnectorService != nil {
		if err := m.ConnectorService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectorService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectorService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateConnegService(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnegService) { // not required
		return nil
	}

	if m.ConnegService != nil {
		if err := m.ConnegService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connegService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connegService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateContext(formats strfmt.Registry) error {
	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateConverterService(formats strfmt.Registry) error {
	if swag.IsZero(m.ConverterService) { // not required
		return nil
	}

	if m.ConverterService != nil {
		if err := m.ConverterService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("converterService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("converterService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateDecoderService(formats strfmt.Registry) error {
	if swag.IsZero(m.DecoderService) { // not required
		return nil
	}

	if m.DecoderService != nil {
		if err := m.DecoderService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decoderService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("decoderService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateEncoderService(formats strfmt.Registry) error {
	if swag.IsZero(m.EncoderService) { // not required
		return nil
	}

	if m.EncoderService != nil {
		if err := m.EncoderService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encoderService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encoderService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateInboundRoot(formats strfmt.Registry) error {
	if swag.IsZero(m.InboundRoot) { // not required
		return nil
	}

	if m.InboundRoot != nil {
		if err := m.InboundRoot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inboundRoot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inboundRoot")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateLogger(formats strfmt.Registry) error {
	if swag.IsZero(m.Logger) { // not required
		return nil
	}

	if m.Logger != nil {
		if err := m.Logger.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logger")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateMetadataService(formats strfmt.Registry) error {
	if swag.IsZero(m.MetadataService) { // not required
		return nil
	}

	if m.MetadataService != nil {
		if err := m.MetadataService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadataService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadataService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateOutboundRoot(formats strfmt.Registry) error {
	if swag.IsZero(m.OutboundRoot) { // not required
		return nil
	}

	if m.OutboundRoot != nil {
		if err := m.OutboundRoot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundRoot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outboundRoot")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateRangeService(formats strfmt.Registry) error {
	if swag.IsZero(m.RangeService) { // not required
		return nil
	}

	if m.RangeService != nil {
		if err := m.RangeService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rangeService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rangeService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) validateStatusService(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusService) { // not required
		return nil
	}

	if m.StatusService != nil {
		if err := m.StatusService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateTaskService(formats strfmt.Registry) error {
	if swag.IsZero(m.TaskService) { // not required
		return nil
	}

	if m.TaskService != nil {
		if err := m.TaskService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taskService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) validateTunnelService(formats strfmt.Registry) error {
	if swag.IsZero(m.TunnelService) { // not required
		return nil
	}

	if m.TunnelService != nil {
		if err := m.TunnelService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tunnelService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tunnelService")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application based on the context it is used
func (m *Application) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectorService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnegService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConverterService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDecoderService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncoderService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInboundRoot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogger(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadataService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutboundRoot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRangeService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaskService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTunnelService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Application) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {

		if swag.IsZero(m.Application) { // not required
			return nil
		}

		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateConnectorService(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectorService != nil {

		if swag.IsZero(m.ConnectorService) { // not required
			return nil
		}

		if err := m.ConnectorService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectorService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectorService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateConnegService(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnegService != nil {

		if swag.IsZero(m.ConnegService) { // not required
			return nil
		}

		if err := m.ConnegService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connegService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connegService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateContext(ctx context.Context, formats strfmt.Registry) error {

	if m.Context != nil {

		if swag.IsZero(m.Context) { // not required
			return nil
		}

		if err := m.Context.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateConverterService(ctx context.Context, formats strfmt.Registry) error {

	if m.ConverterService != nil {

		if swag.IsZero(m.ConverterService) { // not required
			return nil
		}

		if err := m.ConverterService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("converterService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("converterService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateDecoderService(ctx context.Context, formats strfmt.Registry) error {

	if m.DecoderService != nil {

		if swag.IsZero(m.DecoderService) { // not required
			return nil
		}

		if err := m.DecoderService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decoderService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("decoderService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateEncoderService(ctx context.Context, formats strfmt.Registry) error {

	if m.EncoderService != nil {

		if swag.IsZero(m.EncoderService) { // not required
			return nil
		}

		if err := m.EncoderService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encoderService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encoderService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateInboundRoot(ctx context.Context, formats strfmt.Registry) error {

	if m.InboundRoot != nil {

		if swag.IsZero(m.InboundRoot) { // not required
			return nil
		}

		if err := m.InboundRoot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inboundRoot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("inboundRoot")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateLogger(ctx context.Context, formats strfmt.Registry) error {

	if m.Logger != nil {

		if swag.IsZero(m.Logger) { // not required
			return nil
		}

		if err := m.Logger.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logger")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logger")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateMetadataService(ctx context.Context, formats strfmt.Registry) error {

	if m.MetadataService != nil {

		if swag.IsZero(m.MetadataService) { // not required
			return nil
		}

		if err := m.MetadataService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadataService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadataService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateOutboundRoot(ctx context.Context, formats strfmt.Registry) error {

	if m.OutboundRoot != nil {

		if swag.IsZero(m.OutboundRoot) { // not required
			return nil
		}

		if err := m.OutboundRoot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundRoot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("outboundRoot")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateRangeService(ctx context.Context, formats strfmt.Registry) error {

	if m.RangeService != nil {

		if swag.IsZero(m.RangeService) { // not required
			return nil
		}

		if err := m.RangeService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rangeService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rangeService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {

			if swag.IsZero(m.Roles[i]) { // not required
				return nil
			}

			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Services); i++ {

		if m.Services[i] != nil {

			if swag.IsZero(m.Services[i]) { // not required
				return nil
			}

			if err := m.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Application) contextValidateStatusService(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusService != nil {

		if swag.IsZero(m.StatusService) { // not required
			return nil
		}

		if err := m.StatusService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateTaskService(ctx context.Context, formats strfmt.Registry) error {

	if m.TaskService != nil {

		if swag.IsZero(m.TaskService) { // not required
			return nil
		}

		if err := m.TaskService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taskService")
			}
			return err
		}
	}

	return nil
}

func (m *Application) contextValidateTunnelService(ctx context.Context, formats strfmt.Registry) error {

	if m.TunnelService != nil {

		if swag.IsZero(m.TunnelService) { // not required
			return nil
		}

		if err := m.TunnelService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tunnelService")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tunnelService")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Application) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Application) UnmarshalBinary(b []byte) error {
	var res Application
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}