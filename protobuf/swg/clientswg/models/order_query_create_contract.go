// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OrderQueryCreateContract order query create contract
// swagger:model orderQuery_CreateContract
type OrderQueryCreateContract struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// buyer commission
	BuyerCommission float64 `json:"buyerCommission,omitempty"`

	// buyer Id
	BuyerID string `json:"buyerId,omitempty"`

	// front meta data
	FrontMetaData *OrderFrontMetaData `json:"frontMetaData,omitempty"`

	// order Id
	OrderID string `json:"orderId,omitempty"`

	// receive address
	ReceiveAddress string `json:"receiveAddress,omitempty"`

	// seller commission
	SellerCommission float64 `json:"sellerCommission,omitempty"`

	// sending address
	SendingAddress string `json:"sendingAddress,omitempty"`
}

// Validate validates this order query create contract
func (m *OrderQueryCreateContract) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrontMetaData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderQueryCreateContract) validateFrontMetaData(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *OrderQueryCreateContract) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderQueryCreateContract) UnmarshalBinary(b []byte) error {
	var res OrderQueryCreateContract
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
