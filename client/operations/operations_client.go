// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SaveOrUpdateIcecream saves or update an icecream

SaveOrUpdate an Icecream
*/
func (a *Client) SaveOrUpdateIcecream(params *SaveOrUpdateIcecreamParams) (*SaveOrUpdateIcecreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveOrUpdateIcecreamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SaveOrUpdateIcecream",
		Method:             "PUT",
		PathPattern:        "/icecream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveOrUpdateIcecreamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SaveOrUpdateIcecreamOK), nil

}

/*
GetIcecream gets icecream

Get Icecream
*/
func (a *Client) GetIcecream(params *GetIcecreamParams) (*GetIcecreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIcecreamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getIcecream",
		Method:             "GET",
		PathPattern:        "/icecream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIcecreamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIcecreamOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
