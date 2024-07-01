// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	alternate "github.com/ianchen0119/go-swagger/examples/external-types/fred"
)

// MyCustomArrayNullable This generate an array type in models, based on the external type.
// Notice the impact of the nullable hint (equivalent to x-nullable at the type level), to produce a slice of pointers.
//
// []*alternate.MyAlternateType
//
// swagger:model MyCustomArrayNullable
type MyCustomArrayNullable []*alternate.MyAlternateType

// Validate validates this my custom array nullable
func (m MyCustomArrayNullable) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this my custom array nullable based on context it is used
func (m MyCustomArrayNullable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
