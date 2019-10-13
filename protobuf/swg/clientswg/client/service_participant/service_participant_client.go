// Code generated by go-swagger; DO NOT EDIT.

package service_participant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new service participant API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service participant API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateParticipant create participant API
*/
func (a *Client) CreateParticipant(params *CreateParticipantParams) (*CreateParticipantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParticipantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateParticipant",
		Method:             "POST",
		PathPattern:        "/v1/create_participant",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateParticipantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateParticipantOK), nil

}

/*
GetInfo get info API
*/
func (a *Client) GetInfo(params *GetInfoParams) (*GetInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInfo",
		Method:             "GET",
		PathPattern:        "/_v1/get_info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInfoOK), nil

}

/*
GetParticipant get participant API
*/
func (a *Client) GetParticipant(params *GetParticipantParams) (*GetParticipantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParticipantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetParticipant",
		Method:             "GET",
		PathPattern:        "/_v1/get_participant",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetParticipantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetParticipantOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
