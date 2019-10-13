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

// TxReasonTxReasonStatus tx reason tx reason status
// swagger:model TxReason_TxReasonStatus
type TxReasonTxReasonStatus string

const (

	// TxReasonTxReasonStatusUNREASONTX captures enum value "UNREASON_TX"
	TxReasonTxReasonStatusUNREASONTX TxReasonTxReasonStatus = "UNREASON_TX"

	// TxReasonTxReasonStatusFUNDTX captures enum value "FUND_TX"
	TxReasonTxReasonStatusFUNDTX TxReasonTxReasonStatus = "FUND_TX"

	// TxReasonTxReasonStatusWITHDRAWTX captures enum value "WITHDRAW_TX"
	TxReasonTxReasonStatusWITHDRAWTX TxReasonTxReasonStatus = "WITHDRAW_TX"

	// TxReasonTxReasonStatusSELLERCONTRACTTX captures enum value "SELLER_CONTRACT_TX"
	TxReasonTxReasonStatusSELLERCONTRACTTX TxReasonTxReasonStatus = "SELLER_CONTRACT_TX"

	// TxReasonTxReasonStatusBUYERCONTRACTTX captures enum value "BUYER_CONTRACT_TX"
	TxReasonTxReasonStatusBUYERCONTRACTTX TxReasonTxReasonStatus = "BUYER_CONTRACT_TX"

	// TxReasonTxReasonStatusCONTRACTCOMMISSIONTX captures enum value "CONTRACT_COMMISSION_TX"
	TxReasonTxReasonStatusCONTRACTCOMMISSIONTX TxReasonTxReasonStatus = "CONTRACT_COMMISSION_TX"
)

// for schema
var txReasonTxReasonStatusEnum []interface{}

func init() {
	var res []TxReasonTxReasonStatus
	if err := json.Unmarshal([]byte(`["UNREASON_TX","FUND_TX","WITHDRAW_TX","SELLER_CONTRACT_TX","BUYER_CONTRACT_TX","CONTRACT_COMMISSION_TX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		txReasonTxReasonStatusEnum = append(txReasonTxReasonStatusEnum, v)
	}
}

func (m TxReasonTxReasonStatus) validateTxReasonTxReasonStatusEnum(path, location string, value TxReasonTxReasonStatus) error {
	if err := validate.Enum(path, location, value, txReasonTxReasonStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tx reason tx reason status
func (m TxReasonTxReasonStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTxReasonTxReasonStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
