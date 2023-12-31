// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AgreementReturn agreement return
//
// swagger:model AgreementReturn
type AgreementReturn struct {

	// UUID
	UUID string `json:"UUID,omitempty"`

	// agreement Id
	AgreementID int64 `json:"agreementId,omitempty"`

	// charge Id
	ChargeID string `json:"chargeId,omitempty"`

	// vipps confirmation Url
	VippsConfirmationURL string `json:"vippsConfirmationUrl,omitempty"`
}

// Validate validates this agreement return
func (m *AgreementReturn) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this agreement return based on context it is used
func (m *AgreementReturn) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AgreementReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgreementReturn) UnmarshalBinary(b []byte) error {
	var res AgreementReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
