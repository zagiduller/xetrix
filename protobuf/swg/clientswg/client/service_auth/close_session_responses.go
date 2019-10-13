// Code generated by go-swagger; DO NOT EDIT.

package service_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "mxp/mxp-protobuf/swg/clientswg/models"
)

// CloseSessionReader is a Reader for the CloseSession structure.
type CloseSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloseSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCloseSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCloseSessionOK creates a CloseSessionOK with default headers values
func NewCloseSessionOK() *CloseSessionOK {
	return &CloseSessionOK{}
}

/*CloseSessionOK handles this case with default header values.

A successful response.
*/
type CloseSessionOK struct {
	Payload *models.OrderBool
}

func (o *CloseSessionOK) Error() string {
	return fmt.Sprintf("[GET /_v1/close_session][%d] closeSessionOK  %+v", 200, o.Payload)
}

func (o *CloseSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderBool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
