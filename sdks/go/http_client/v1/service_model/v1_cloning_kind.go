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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1CloningKind v1 cloning kind
//
// swagger:model v1CloningKind
type V1CloningKind string

const (

	// V1CloningKindCopy captures enum value "copy"
	V1CloningKindCopy V1CloningKind = "copy"

	// V1CloningKindRestart captures enum value "restart"
	V1CloningKindRestart V1CloningKind = "restart"

	// V1CloningKindCache captures enum value "cache"
	V1CloningKindCache V1CloningKind = "cache"
)

// for schema
var v1CloningKindEnum []interface{}

func init() {
	var res []V1CloningKind
	if err := json.Unmarshal([]byte(`["copy","restart","cache"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CloningKindEnum = append(v1CloningKindEnum, v)
	}
}

func (m V1CloningKind) validateV1CloningKindEnum(path, location string, value V1CloningKind) error {
	if err := validate.EnumCase(path, location, value, v1CloningKindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 cloning kind
func (m V1CloningKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1CloningKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
