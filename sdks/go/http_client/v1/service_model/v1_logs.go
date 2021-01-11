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
	"github.com/go-openapi/validate"
)

// V1Logs Logs
//
// swagger:model v1Logs
type V1Logs struct {

	// Log files
	Files []string `json:"files"`

	// Last log file
	LastFile string `json:"last_file,omitempty"`

	// Last log time
	// Format: date-time
	LastTime strfmt.DateTime `json:"last_time,omitempty"`

	// Log lines
	Logs []*V1Log `json:"logs"`
}

// Validate validates this v1 logs
func (m *V1Logs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Logs) validateLastTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastTime) { // not required
		return nil
	}

	if err := validate.FormatOf("last_time", "body", "date-time", m.LastTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1Logs) validateLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Logs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Logs) UnmarshalBinary(b []byte) error {
	var res V1Logs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
