// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AgreementBody agreement body
//
// swagger:model AgreementBody
type AgreementBody struct {

	// agreement Url
	AgreementURL string `json:"agreementUrl,omitempty"`

	// amount
	Amount int64 `json:"amount,omitempty"`

	// customer Id
	CustomerID int64 `json:"customerId,omitempty"`

	// interval count
	IntervalCount int64 `json:"intervalCount,omitempty"`

	// interval unit
	// Enum: [WEEK MONTH YEAR]
	IntervalUnit string `json:"intervalUnit,omitempty"`

	// status
	// Enum: [ACTIVE PAUSED STOPPED]
	Status string `json:"status,omitempty"`

	// vippsnummer
	Vippsnummer int64 `json:"vippsnummer,omitempty"`
}

// Validate validates this agreement body
func (m *AgreementBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntervalUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var agreementBodyTypeIntervalUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WEEK","MONTH","YEAR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		agreementBodyTypeIntervalUnitPropEnum = append(agreementBodyTypeIntervalUnitPropEnum, v)
	}
}

const (

	// AgreementBodyIntervalUnitWEEK captures enum value "WEEK"
	AgreementBodyIntervalUnitWEEK string = "WEEK"

	// AgreementBodyIntervalUnitMONTH captures enum value "MONTH"
	AgreementBodyIntervalUnitMONTH string = "MONTH"

	// AgreementBodyIntervalUnitYEAR captures enum value "YEAR"
	AgreementBodyIntervalUnitYEAR string = "YEAR"
)

// prop value enum
func (m *AgreementBody) validateIntervalUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, agreementBodyTypeIntervalUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AgreementBody) validateIntervalUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.IntervalUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateIntervalUnitEnum("intervalUnit", "body", m.IntervalUnit); err != nil {
		return err
	}

	return nil
}

var agreementBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","PAUSED","STOPPED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		agreementBodyTypeStatusPropEnum = append(agreementBodyTypeStatusPropEnum, v)
	}
}

const (

	// AgreementBodyStatusACTIVE captures enum value "ACTIVE"
	AgreementBodyStatusACTIVE string = "ACTIVE"

	// AgreementBodyStatusPAUSED captures enum value "PAUSED"
	AgreementBodyStatusPAUSED string = "PAUSED"

	// AgreementBodyStatusSTOPPED captures enum value "STOPPED"
	AgreementBodyStatusSTOPPED string = "STOPPED"
)

// prop value enum
func (m *AgreementBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, agreementBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AgreementBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this agreement body based on context it is used
func (m *AgreementBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AgreementBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgreementBody) UnmarshalBinary(b []byte) error {
	var res AgreementBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}