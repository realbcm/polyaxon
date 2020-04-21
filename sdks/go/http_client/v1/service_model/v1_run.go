// Copyright 2018-2020 Polyaxon, Inc.
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
	"github.com/go-openapi/validate"
)

// V1Run Run specification
//
// swagger:model v1Run
type V1Run struct {

	// Optional if this entity was bookmarked
	Bookmarked bool `json:"bookmarked,omitempty"`

	// Optional content of the entity's spec
	Content string `json:"content,omitempty"`

	// Optional time when the entity was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Optional if the entity has been deleted
	Deleted bool `json:"deleted,omitempty"`

	// Optional description
	Description string `json:"description,omitempty"`

	// Optional last time the entity was started
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finished_at,omitempty"`

	// Optional hub reference versioned component
	Hub string `json:"hub,omitempty"`

	// Optional inputs of this entity
	Inputs interface{} `json:"inputs,omitempty"`

	// Optional is helper run
	IsHelper bool `json:"is_helper,omitempty"`

	// Optional flag to tell if this entity is managed by the platform
	IsManaged string `json:"is_managed,omitempty"`

	// Optional kind to tell the kind of this run
	Kind V1RunKind `json:"kind,omitempty"`

	// Optional run meta info
	MetaInfo interface{} `json:"meta_info,omitempty"`

	// Optional meta kind to tell the nature of this run
	MetaKind V1RunKind `json:"meta_kind,omitempty"`

	// Optional name
	Name string `json:"name,omitempty"`

	// Optional original run meta information
	Original *V1Cloning `json:"original,omitempty"`

	// Optional outputs of this entity
	Outputs interface{} `json:"outputs,omitempty"`

	// Required name of owner of this entity
	Owner string `json:"owner,omitempty"`

	// Optional pipeline run meta information
	Pipeline *V1Pipeline `json:"pipeline,omitempty"`

	// Required project name
	Project string `json:"project,omitempty"`

	// Current user's role in this (org/teams)/project/runs
	Role string `json:"role,omitempty"`

	// Optional run time of the entity
	RunTime int32 `json:"run_time,omitempty"`

	// Optional settings
	Settings *V1RunSettings `json:"settings,omitempty"`

	// Optional last time the entity was started
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// Optional latest status of this entity
	Status V1Statuses `json:"status,omitempty"`

	// The status conditions timeline
	StatusConditions []*V1StatusCondition `json:"status_conditions"`

	// Optional tags of this entity
	Tags []string `json:"tags"`

	// Optional last time the entity was updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// Required name of user started this entity
	User string `json:"user,omitempty"`

	// UUID
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 run
func (m *V1Run) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipeline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Run) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateFinishedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", m.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(m.Kind) { // not required
		return nil
	}

	if err := m.Kind.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("kind")
		}
		return err
	}

	return nil
}

func (m *V1Run) validateMetaKind(formats strfmt.Registry) error {

	if swag.IsZero(m.MetaKind) { // not required
		return nil
	}

	if err := m.MetaKind.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("meta_kind")
		}
		return err
	}

	return nil
}

func (m *V1Run) validateOriginal(formats strfmt.Registry) error {

	if swag.IsZero(m.Original) { // not required
		return nil
	}

	if m.Original != nil {
		if err := m.Original.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validatePipeline(formats strfmt.Registry) error {

	if swag.IsZero(m.Pipeline) { // not required
		return nil
	}

	if m.Pipeline != nil {
		if err := m.Pipeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pipeline")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *V1Run) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Run) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *V1Run) validateStatusConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusConditions); i++ {
		if swag.IsZero(m.StatusConditions[i]) { // not required
			continue
		}

		if m.StatusConditions[i] != nil {
			if err := m.StatusConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1Run) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Run) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Run) UnmarshalBinary(b []byte) error {
	var res V1Run
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
