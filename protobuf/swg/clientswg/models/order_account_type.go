// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// OrderAccountType enum structs
// swagger:model orderAccountType
type OrderAccountType string

const (

	// OrderAccountTypeINACTIVE captures enum value "INACTIVE"
	OrderAccountTypeINACTIVE OrderAccountType = "INACTIVE"

	// OrderAccountTypeINTERNAL captures enum value "INTERNAL"
	OrderAccountTypeINTERNAL OrderAccountType = "INTERNAL"

	// OrderAccountTypeEXTERNAL captures enum value "EXTERNAL"
	OrderAccountTypeEXTERNAL OrderAccountType = "EXTERNAL"

	// OrderAccountTypeSYSTEM captures enum value "SYSTEM"
	OrderAccountTypeSYSTEM OrderAccountType = "SYSTEM"
)

// for schema
var orderAccountTypeEnum []interface{}

func init() {
	var res []OrderAccountType
	if err := json.Unmarshal([]byte(`["INACTIVE","INTERNAL","EXTERNAL","SYSTEM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		orderAccountTypeEnum = append(orderAccountTypeEnum, v)
	}
}

func (m OrderAccountType) validateOrderAccountTypeEnum(path, location string, value OrderAccountType) error {
	if err := validate.Enum(path, location, value, orderAccountTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this order account type
func (m OrderAccountType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrderAccountTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
