// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OrderQueryOrder order query order
// swagger:model orderQuery_Order
type OrderQueryOrder struct {

	// active
	Active bool `json:"active,omitempty"`

	// amount
	Amount float64 `json:"amount,omitempty"`

	// available
	Available float64 `json:"available,omitempty"`

	// buy currency symbol
	BuyCurrencySymbol string `json:"buyCurrencySymbol,omitempty"`

	// commission
	Commission float64 `json:"commission,omitempty"`

	// front meta data
	FrontMetaData *OrderFrontMetaData `json:"frontMetaData,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// owner Id
	OwnerID string `json:"ownerId,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// receive address
	ReceiveAddress string `json:"receiveAddress,omitempty"`

	// sell currency symbol
	SellCurrencySymbol string `json:"sellCurrencySymbol,omitempty"`

	// sending address
	SendingAddress string `json:"sendingAddress,omitempty"`

	// status
	Status *OrderDealStatus `json:"status,omitempty"`
}

// Validate validates this order query order
func (m *OrderQueryOrder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrontMetaData(formats); err != nil {
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

func (m *OrderQueryOrder) validateFrontMetaData(formats strfmt.Registry) error {

	if swag.IsZero(m.FrontMetaData) { // not required
		return nil
	}

	if m.FrontMetaData != nil {
		if err := m.FrontMetaData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frontMetaData")
			}
			return err
		}
	}

	return nil
}

func (m *OrderQueryOrder) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderQueryOrder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderQueryOrder) UnmarshalBinary(b []byte) error {
	var res OrderQueryOrder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
