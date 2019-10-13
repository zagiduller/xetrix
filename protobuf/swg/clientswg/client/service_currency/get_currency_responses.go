// Code generated by go-swagger; DO NOT EDIT.

package service_currency

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "mxp/mxp-protobuf/swg/clientswg/models"
)

// GetCurrencyReader is a Reader for the GetCurrency structure.
type GetCurrencyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrencyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCurrencyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCurrencyOK creates a GetCurrencyOK with default headers values
func NewGetCurrencyOK() *GetCurrencyOK {
	return &GetCurrencyOK{}
}

/*GetCurrencyOK handles this case with default header values.

A successful response.
*/
type GetCurrencyOK struct {
	Payload *models.OrderResponseCurrencyQuery
}

func (o *GetCurrencyOK) Error() string {
	return fmt.Sprintf("[GET /v1/get_currency][%d] getCurrencyOK  %+v", 200, o.Payload)
}

func (o *GetCurrencyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderResponseCurrencyQuery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}