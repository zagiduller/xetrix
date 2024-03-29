// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OrderResponseTx order response tx
// swagger:model orderResponse_Tx
type OrderResponseTx struct {

	// query status
	QueryStatus OrderQueryStatus `json:"QueryStatus,omitempty"`

	// confirmed
	Confirmed bool `json:"confirmed,omitempty"`

	// created
	Created bool `json:"created,omitempty"`

	// items
	Items []*OrderTx `json:"items"`

	// items count
	ItemsCount int64 `json:"itemsCount,omitempty"`

	// object
	Object *OrderTx `json:"object,omitempty"`
}

// Validate validates this order response tx
func (m *OrderResponseTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueryStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderResponseTx) validateQueryStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.QueryStatus) { // not required
		return nil
	}

	if err := m.QueryStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("QueryStatus")
		}
		return err
	}

	return nil
}

func (m *OrderResponseTx) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrderResponseTx) validateObject(formats strfmt.Registry) error {

	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrderResponseTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderResponseTx) UnmarshalBinary(b []byte) error {
	var res OrderResponseTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
