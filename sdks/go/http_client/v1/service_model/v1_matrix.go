// Copyright 2018-2021 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Matrix All Schemas message
//
// swagger:model v1Matrix
type V1Matrix struct {

	// bayes
	Bayes *V1Bayes `json:"bayes,omitempty"`

	// grid
	Grid *V1GridSearch `json:"grid,omitempty"`

	// hyperband
	Hyperband *V1Hyperband `json:"hyperband,omitempty"`

	// hyperopt
	Hyperopt *V1Hyperopt `json:"hyperopt,omitempty"`

	// iterative
	Iterative *V1Iterative `json:"iterative,omitempty"`

	// mapping
	Mapping *V1Mapping `json:"mapping,omitempty"`

	// random
	Random *V1RandomSearch `json:"random,omitempty"`
}

// Validate validates this v1 matrix
func (m *V1Matrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBayes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHyperband(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHyperopt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIterative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRandom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Matrix) validateBayes(formats strfmt.Registry) error {

	if swag.IsZero(m.Bayes) { // not required
		return nil
	}

	if m.Bayes != nil {
		if err := m.Bayes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bayes")
			}
			return err
		}
	}

	return nil
}

func (m *V1Matrix) validateGrid(formats strfmt.Registry) error {

	if swag.IsZero(m.Grid) { // not required
		return nil
	}

	if m.Grid != nil {
		if err := m.Grid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("grid")
			}
			return err
		}
	}

	return nil
}

func (m *V1Matrix) validateHyperband(formats strfmt.Registry) error {

	if swag.IsZero(m.Hyperband) { // not required
		return nil
	}

	if m.Hyperband != nil {
		if err := m.Hyperband.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hyperband")
			}
			return err
		}
	}

	return nil
}

func (m *V1Matrix) validateHyperopt(formats strfmt.Registry) error {

	if swag.IsZero(m.Hyperopt) { // not required
		return nil
	}

	if m.Hyperopt != nil {
		if err := m.Hyperopt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hyperopt")
			}
			return err
		}
	}

	return nil
}

func (m *V1Matrix) validateIterative(formats strfmt.Registry) error {

	if swag.IsZero(m.Iterative) { // not required
		return nil
	}

	if m.Iterative != nil {
		if err := m.Iterative.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iterative")
			}
			return err
		}
	}

	return nil
}

func (m *V1Matrix) validateMapping(formats strfmt.Registry) error {

	if swag.IsZero(m.Mapping) { // not required
		return nil
	}

	if m.Mapping != nil {
		if err := m.Mapping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mapping")
			}
			return err
		}
	}

	return nil
}

func (m *V1Matrix) validateRandom(formats strfmt.Registry) error {

	if swag.IsZero(m.Random) { // not required
		return nil
	}

	if m.Random != nil {
		if err := m.Random.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("random")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Matrix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Matrix) UnmarshalBinary(b []byte) error {
	var res V1Matrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
