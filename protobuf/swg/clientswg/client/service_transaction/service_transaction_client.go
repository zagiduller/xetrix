// Code generated by go-swagger; DO NOT EDIT.

package service_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new service transaction API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service transaction API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetTx get tx API
*/
func (a *Client) GetTx(params *GetTxParams) (*GetTxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTx",
		Method:             "GET",
		PathPattern:        "/_v1/get_tx/{txId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTxOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
