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

package presets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListPresetsReader is a Reader for the ListPresets structure.
type ListPresetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPresetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPresetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListPresetsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListPresetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListPresetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListPresetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPresetsOK creates a ListPresetsOK with default headers values
func NewListPresetsOK() *ListPresetsOK {
	return &ListPresetsOK{}
}

/*ListPresetsOK handles this case with default header values.

A successful response.
*/
type ListPresetsOK struct {
	Payload *service_model.V1ListPresetsResponse
}

func (o *ListPresetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets][%d] listPresetsOK  %+v", 200, o.Payload)
}

func (o *ListPresetsOK) GetPayload() *service_model.V1ListPresetsResponse {
	return o.Payload
}

func (o *ListPresetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListPresetsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPresetsNoContent creates a ListPresetsNoContent with default headers values
func NewListPresetsNoContent() *ListPresetsNoContent {
	return &ListPresetsNoContent{}
}

/*ListPresetsNoContent handles this case with default header values.

No content.
*/
type ListPresetsNoContent struct {
	Payload interface{}
}

func (o *ListPresetsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets][%d] listPresetsNoContent  %+v", 204, o.Payload)
}

func (o *ListPresetsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListPresetsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPresetsForbidden creates a ListPresetsForbidden with default headers values
func NewListPresetsForbidden() *ListPresetsForbidden {
	return &ListPresetsForbidden{}
}

/*ListPresetsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListPresetsForbidden struct {
	Payload interface{}
}

func (o *ListPresetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets][%d] listPresetsForbidden  %+v", 403, o.Payload)
}

func (o *ListPresetsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListPresetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPresetsNotFound creates a ListPresetsNotFound with default headers values
func NewListPresetsNotFound() *ListPresetsNotFound {
	return &ListPresetsNotFound{}
}

/*ListPresetsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListPresetsNotFound struct {
	Payload interface{}
}

func (o *ListPresetsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets][%d] listPresetsNotFound  %+v", 404, o.Payload)
}

func (o *ListPresetsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListPresetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPresetsDefault creates a ListPresetsDefault with default headers values
func NewListPresetsDefault(code int) *ListPresetsDefault {
	return &ListPresetsDefault{
		_statusCode: code,
	}
}

/*ListPresetsDefault handles this case with default header values.

An unexpected error response.
*/
type ListPresetsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list presets default response
func (o *ListPresetsDefault) Code() int {
	return o._statusCode
}

func (o *ListPresetsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets][%d] ListPresets default  %+v", o._statusCode, o.Payload)
}

func (o *ListPresetsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListPresetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
