// Code generated by go-swagger; DO NOT EDIT.

package service_participant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "mxp/mxp-protobuf/swg/clientswg/models"
)

// GetParticipantReader is a Reader for the GetParticipant structure.
type GetParticipantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParticipantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParticipantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParticipantOK creates a GetParticipantOK with default headers values
func NewGetParticipantOK() *GetParticipantOK {
	return &GetParticipantOK{}
}

/*GetParticipantOK handles this case with default header values.

A successful response.
*/
type GetParticipantOK struct {
	Payload *models.OrderResponseParticipant
}

func (o *GetParticipantOK) Error() string {
	return fmt.Sprintf("[GET /_v1/get_participant][%d] getParticipantOK  %+v", 200, o.Payload)
}

func (o *GetParticipantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderResponseParticipant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}