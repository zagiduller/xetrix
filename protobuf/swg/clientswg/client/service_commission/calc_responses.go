// Code generated by go-swagger; DO NOT EDIT.

package service_commission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "mxp/mxp-protobuf/swg/clientswg/models"
)

// CalcReader is a Reader for the Calc structure.
type CalcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CalcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCalcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCalcOK creates a CalcOK with default headers values
func NewCalcOK() *CalcOK {
	return &CalcOK{}
}

/*CalcOK handles this case with default header values.

A successful response.
*/
type CalcOK struct {
	Payload *models.OrderCommission
}

func (o *CalcOK) Error() string {
	return fmt.Sprintf("[POST /_v1/calc_commission][%d] calcOK  %+v", 200, o.Payload)
}

func (o *CalcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderCommission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
