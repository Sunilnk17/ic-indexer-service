// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GenericError Icecream Generic Error
// swagger:model GenericError
type GenericError struct {

	// code
	Code int64 `json:"code,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// request Id
	RequestID string `json:"request_id,omitempty"`
}

// Validate validates this generic error
func (m *GenericError) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenericError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericError) UnmarshalBinary(b []byte) error {
	var res GenericError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
