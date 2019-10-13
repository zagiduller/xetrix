// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OrderAccount order account
// swagger:model orderAccount
type OrderAccount struct {

	// address
	Address string `json:"Address,omitempty"`

	// @inject_tag: bson:"_id,omitempty"
	// Format: byte
	BsonID strfmt.Base64 `json:"bsonId,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// currency
	Currency *OrderCurrency `json:"currency,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// owner Id
	OwnerID string `json:"ownerId,omitempty"`

	// status
	Status OrderAccountStatus `json:"status,omitempty"`

	// type
	Type OrderAccountType `json:"type,omitempty"`
}

// Validate validates this order account
func (m *OrderAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBsonID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderAccount) validateBsonID(formats strfmt.Registry) error {

	if swag.IsZero(m.BsonID) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

func (m *OrderAccount) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if m.Currency != nil {
		if err := m.Currency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currency")
			}
			return err
		}
	}

	return nil
}

func (m *OrderAccount) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *OrderAccount) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderAccount) UnmarshalBinary(b []byte) error {
	var res OrderAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
