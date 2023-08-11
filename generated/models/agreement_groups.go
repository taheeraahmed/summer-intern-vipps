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

// AgreementGroups agreement groups
//
// swagger:model AgreementGroups
type AgreementGroups struct {

	// active
	Active []*AgreementDetails `json:"active"`

	// paused
	Paused []*AgreementDetails `json:"paused"`

	// stopped
	Stopped []*AgreementDetails `json:"stopped"`
}

// Validate validates this agreement groups
func (m *AgreementGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaused(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopped(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgreementGroups) validateActive(formats strfmt.Registry) error {
	if swag.IsZero(m.Active) { // not required
		return nil
	}

	for i := 0; i < len(m.Active); i++ {
		if swag.IsZero(m.Active[i]) { // not required
			continue
		}

		if m.Active[i] != nil {
			if err := m.Active[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("active" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgreementGroups) validatePaused(formats strfmt.Registry) error {
	if swag.IsZero(m.Paused) { // not required
		return nil
	}

	for i := 0; i < len(m.Paused); i++ {
		if swag.IsZero(m.Paused[i]) { // not required
			continue
		}

		if m.Paused[i] != nil {
			if err := m.Paused[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paused" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("paused" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgreementGroups) validateStopped(formats strfmt.Registry) error {
	if swag.IsZero(m.Stopped) { // not required
		return nil
	}

	for i := 0; i < len(m.Stopped); i++ {
		if swag.IsZero(m.Stopped[i]) { // not required
			continue
		}

		if m.Stopped[i] != nil {
			if err := m.Stopped[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stopped" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stopped" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this agreement groups based on the context it is used
func (m *AgreementGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaused(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopped(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgreementGroups) contextValidateActive(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Active); i++ {

		if m.Active[i] != nil {

			if swag.IsZero(m.Active[i]) { // not required
				return nil
			}

			if err := m.Active[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("active" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgreementGroups) contextValidatePaused(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Paused); i++ {

		if m.Paused[i] != nil {

			if swag.IsZero(m.Paused[i]) { // not required
				return nil
			}

			if err := m.Paused[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paused" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("paused" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgreementGroups) contextValidateStopped(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stopped); i++ {

		if m.Stopped[i] != nil {

			if swag.IsZero(m.Stopped[i]) { // not required
				return nil
			}

			if err := m.Stopped[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stopped" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stopped" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgreementGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgreementGroups) UnmarshalBinary(b []byte) error {
	var res AgreementGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}