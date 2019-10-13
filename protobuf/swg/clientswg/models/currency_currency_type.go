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

// CurrencyCurrencyType currency currency type
// swagger:model Currency_currencyType
type CurrencyCurrencyType string

const (

	// CurrencyCurrencyTypeCRYPTOCURRENCY captures enum value "CRYPTO_CURRENCY"
	CurrencyCurrencyTypeCRYPTOCURRENCY CurrencyCurrencyType = "CRYPTO_CURRENCY"

	// CurrencyCurrencyTypeFIATCURRENCY captures enum value "FIAT_CURRENCY"
	CurrencyCurrencyTypeFIATCURRENCY CurrencyCurrencyType = "FIAT_CURRENCY"
)

// for schema
var currencyCurrencyTypeEnum []interface{}

func init() {
	var res []CurrencyCurrencyType
	if err := json.Unmarshal([]byte(`["CRYPTO_CURRENCY","FIAT_CURRENCY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		currencyCurrencyTypeEnum = append(currencyCurrencyTypeEnum, v)
	}
}

func (m CurrencyCurrencyType) validateCurrencyCurrencyTypeEnum(path, location string, value CurrencyCurrencyType) error {
	if err := validate.Enum(path, location, value, currencyCurrencyTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this currency currency type
func (m CurrencyCurrencyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCurrencyCurrencyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}