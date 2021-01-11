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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Component Component specification
//
// swagger:model v1Component
type V1Component struct {

	// Optional flag to disable cache validation and force run this component
	Cache *V1Cache `json:"cache,omitempty"`

	// Optional component description
	Description string `json:"description,omitempty"`

	// Optional hooks section
	Hooks []*V1Hook `json:"hooks"`

	// Optional inputs definition
	Inputs []*V1IO `json:"inputs"`

	// Optional flag to mark this specification requires approval before running
	IsApproved bool `json:"is_approved,omitempty"`

	// Optional component kind, should be equal to 'operation'
	Kind string `json:"kind,omitempty"`

	// Optional component name, should be a valid fully qualified value: name[:version]
	Name string `json:"name,omitempty"`

	// Optional outputs definition
	Outputs []*V1IO `json:"outputs"`

	// Optional plugins to enable
	Plugins *V1Plugins `json:"plugins,omitempty"`

	// Optional presets to use for running this component
	Presets []string `json:"presets"`

	// Optional queue to use for running this component
	Queue string `json:"queue,omitempty"`

	// Run definition, should be one of: Job/Service/Spark/Flink/Kubeflow/Dask/Dag
	Run interface{} `json:"run,omitempty"`

	// Optional component tags
	Tags []string `json:"tags"`

	// Optional flag to mark this specification as template
	Template *V1Template `json:"template,omitempty"`

	// optional termination section
	Termination *V1Termination `json:"termination,omitempty"`

	// Spec version
	Version float32 `json:"version,omitempty"`
}

// Validate validates this v1 component
func (m *V1Component) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCache(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHooks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Component) validateCache(formats strfmt.Registry) error {

	if swag.IsZero(m.Cache) { // not required
		return nil
	}

	if m.Cache != nil {
		if err := m.Cache.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cache")
			}
			return err
		}
	}

	return nil
}

func (m *V1Component) validateHooks(formats strfmt.Registry) error {

	if swag.IsZero(m.Hooks) { // not required
		return nil
	}

	for i := 0; i < len(m.Hooks); i++ {
		if swag.IsZero(m.Hooks[i]) { // not required
			continue
		}

		if m.Hooks[i] != nil {
			if err := m.Hooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Component) validateInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Inputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Inputs); i++ {
		if swag.IsZero(m.Inputs[i]) { // not required
			continue
		}

		if m.Inputs[i] != nil {
			if err := m.Inputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Component) validateOutputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Outputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Outputs); i++ {
		if swag.IsZero(m.Outputs[i]) { // not required
			continue
		}

		if m.Outputs[i] != nil {
			if err := m.Outputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Component) validatePlugins(formats strfmt.Registry) error {

	if swag.IsZero(m.Plugins) { // not required
		return nil
	}

	if m.Plugins != nil {
		if err := m.Plugins.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plugins")
			}
			return err
		}
	}

	return nil
}

func (m *V1Component) validateTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

func (m *V1Component) validateTermination(formats strfmt.Registry) error {

	if swag.IsZero(m.Termination) { // not required
		return nil
	}

	if m.Termination != nil {
		if err := m.Termination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("termination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Component) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Component) UnmarshalBinary(b []byte) error {
	var res V1Component
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
