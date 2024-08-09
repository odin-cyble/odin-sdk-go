// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExposedSearchCount exposed search count
//
// swagger:model exposed.SearchCount
type ExposedSearchCount struct {

	// count
	Count int64 `json:"count,omitempty"`
}

// Validate validates this exposed search count
func (m *ExposedSearchCount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exposed search count based on context it is used
func (m *ExposedSearchCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExposedSearchCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExposedSearchCount) UnmarshalBinary(b []byte) error {
	var res ExposedSearchCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
