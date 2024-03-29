// Code generated by go-swagger; DO NOT EDIT.

package service_participant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "mxp/mxp-protobuf/swg/clientswg/models"
)

// NewCreateParticipantParams creates a new CreateParticipantParams object
// with the default values initialized.
func NewCreateParticipantParams() *CreateParticipantParams {
	var ()
	return &CreateParticipantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateParticipantParamsWithTimeout creates a new CreateParticipantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateParticipantParamsWithTimeout(timeout time.Duration) *CreateParticipantParams {
	var ()
	return &CreateParticipantParams{

		timeout: timeout,
	}
}

// NewCreateParticipantParamsWithContext creates a new CreateParticipantParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateParticipantParamsWithContext(ctx context.Context) *CreateParticipantParams {
	var ()
	return &CreateParticipantParams{

		Context: ctx,
	}
}

// NewCreateParticipantParamsWithHTTPClient creates a new CreateParticipantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateParticipantParamsWithHTTPClient(client *http.Client) *CreateParticipantParams {
	var ()
	return &CreateParticipantParams{
		HTTPClient: client,
	}
}

/*CreateParticipantParams contains all the parameters to send to the API endpoint
for the create participant operation typically these are written to a http.Request
*/
type CreateParticipantParams struct {

	/*Body*/
	Body *models.OrderParticipant

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create participant params
func (o *CreateParticipantParams) WithTimeout(timeout time.Duration) *CreateParticipantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create participant params
func (o *CreateParticipantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create participant params
func (o *CreateParticipantParams) WithContext(ctx context.Context) *CreateParticipantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create participant params
func (o *CreateParticipantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create participant params
func (o *CreateParticipantParams) WithHTTPClient(client *http.Client) *CreateParticipantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create participant params
func (o *CreateParticipantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create participant params
func (o *CreateParticipantParams) WithBody(body *models.OrderParticipant) *CreateParticipantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create participant params
func (o *CreateParticipantParams) SetBody(body *models.OrderParticipant) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateParticipantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
