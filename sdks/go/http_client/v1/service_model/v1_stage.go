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

// V1Stage Stage specification
//
// swagger:model v1Stage
type V1Stage struct {

	// The current stage
	Stage V1Stages `json:"stage,omitempty"`

	// The stage conditions timeline
	StageConditions []*V1StageCondition `json:"stage_conditions"`

	// The uuid of the stage
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 stage
func (m *V1Stage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Stage) validateStage(formats strfmt.Registry) error {

	if swag.IsZero(m.Stage) { // not required
		return nil
	}

	if err := m.Stage.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("stage")
		}
		return err
	}

	return nil
}

func (m *V1Stage) validateStageConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.StageConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.StageConditions); i++ {
		if swag.IsZero(m.StageConditions[i]) { // not required
			continue
		}

		if m.StageConditions[i] != nil {
			if err := m.StageConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stage_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Stage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Stage) UnmarshalBinary(b []byte) error {
	var res V1Stage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
