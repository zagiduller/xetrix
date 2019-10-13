// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OrderFrontMetaData order front meta data
// swagger:model orderFrontMetaData
type OrderFrontMetaData struct {

	// user name
	UserName string `json:"userName,omitempty"`

	// user price input
	UserPriceInput string `json:"userPriceInput,omitempty"`
}

// Validate validates this order front meta data
func (m *OrderFrontMetaData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrderFrontMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrderFrontMetaData) UnmarshalBinary(b []byte) error {
	var res OrderFrontMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}