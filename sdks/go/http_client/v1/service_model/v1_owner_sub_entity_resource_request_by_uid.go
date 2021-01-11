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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1OwnerSubEntityResourceRequestByUID Request data to get/delete sub-entity
//
// swagger:model v1OwnerSubEntityResourceRequestByUid
type V1OwnerSubEntityResourceRequestByUID struct {

	// Entity: project name, hub name, registry name, ...
	Entity string `json:"entity,omitempty"`

	// Owner of the namespace
	Owner string `json:"owner,omitempty"`

	// Uuid identifier of the sub-entity
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this v1 owner sub entity resource request by Uid
func (m *V1OwnerSubEntityResourceRequestByUID) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1OwnerSubEntityResourceRequestByUID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OwnerSubEntityResourceRequestByUID) UnmarshalBinary(b []byte) error {
	var res V1OwnerSubEntityResourceRequestByUID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
