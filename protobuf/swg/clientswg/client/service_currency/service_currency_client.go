// Code generated by go-swagger; DO NOT EDIT.

package service_currency

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new service currency API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service currency API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCurrency get currency API
*/
func (a *Client) GetCurrency(params *GetCurrencyParams) (*GetCurrencyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrencyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCurrency",
		Method:             "GET",
		PathPattern:        "/v1/get_currency",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCurrencyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCurrencyOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
