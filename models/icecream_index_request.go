// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IcecreamIndexRequest icecream index request
// swagger:model IcecreamIndexRequest
type IcecreamIndexRequest struct {

	// allergy info
	AllergyInfo *CusString `json:"allergy_info,omitempty"`

	// created at
	CreatedAt *CusTime `json:"created_at,omitempty"`

	// description
	Description *CusString `json:"description,omitempty"`

	// dietary certifications
	DietaryCertifications *CusString `json:"dietary_certifications,omitempty"`

	// id
	ID *CusInt64 `json:"id,omitempty"`

	// image closed
	ImageClosed *CusString `json:"image_closed,omitempty"`

	// image open
	ImageOpen *CusString `json:"image_open,omitempty"`

	// ingredients
	Ingredients *CusArrayString `json:"ingredients,omitempty"`

	// last updated at
	LastUpdatedAt *CusTime `json:"last_updated_at,omitempty"`

	// name
	Name *CusString `json:"name,omitempty"`

	// product id
	ProductID *CusString `json:"product_id,omitempty"`

	// sourcing values
	SourcingValues *CusArrayString `json:"sourcing_values,omitempty"`

	// story
	Story *CusString `json:"story,omitempty"`
}

// Validate validates this icecream index request
func (m *IcecreamIndexRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllergyInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDietaryCertifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageClosed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageOpen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcingValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IcecreamIndexRequest) validateAllergyInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.AllergyInfo) { // not required
		return nil
	}

	if m.AllergyInfo != nil {
		if err := m.AllergyInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allergy_info")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if m.CreatedAt != nil {
		if err := m.CreatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_at")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if m.Description != nil {
		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateDietaryCertifications(formats strfmt.Registry) error {

	if swag.IsZero(m.DietaryCertifications) { // not required
		return nil
	}

	if m.DietaryCertifications != nil {
		if err := m.DietaryCertifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dietary_certifications")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateImageClosed(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageClosed) { // not required
		return nil
	}

	if m.ImageClosed != nil {
		if err := m.ImageClosed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image_closed")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateImageOpen(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageOpen) { // not required
		return nil
	}

	if m.ImageOpen != nil {
		if err := m.ImageOpen.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image_open")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateIngredients(formats strfmt.Registry) error {

	if swag.IsZero(m.Ingredients) { // not required
		return nil
	}

	if m.Ingredients != nil {
		if err := m.Ingredients.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingredients")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateLastUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedAt) { // not required
		return nil
	}

	if m.LastUpdatedAt != nil {
		if err := m.LastUpdatedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_updated_at")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateProductID(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductID) { // not required
		return nil
	}

	if m.ProductID != nil {
		if err := m.ProductID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product_id")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateSourcingValues(formats strfmt.Registry) error {

	if swag.IsZero(m.SourcingValues) { // not required
		return nil
	}

	if m.SourcingValues != nil {
		if err := m.SourcingValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourcing_values")
			}
			return err
		}
	}

	return nil
}

func (m *IcecreamIndexRequest) validateStory(formats strfmt.Registry) error {

	if swag.IsZero(m.Story) { // not required
		return nil
	}

	if m.Story != nil {
		if err := m.Story.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("story")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IcecreamIndexRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IcecreamIndexRequest) UnmarshalBinary(b []byte) error {
	var res IcecreamIndexRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
