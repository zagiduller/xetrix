// Code generated by go-swagger; DO NOT EDIT.

package service_commission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new service commission API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service commission API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Calc calc API
*/
func (a *Client) Calc(params *CalcParams) (*CalcOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCalcParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Calc",
		Method:             "POST",
		PathPattern:        "/_v1/calc_commission",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CalcReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CalcOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}