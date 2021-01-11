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

package component_hub_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetComponentVersionStagesReader is a Reader for the GetComponentVersionStages structure.
type GetComponentVersionStagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComponentVersionStagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComponentVersionStagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetComponentVersionStagesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetComponentVersionStagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetComponentVersionStagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetComponentVersionStagesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetComponentVersionStagesOK creates a GetComponentVersionStagesOK with default headers values
func NewGetComponentVersionStagesOK() *GetComponentVersionStagesOK {
	return &GetComponentVersionStagesOK{}
}

/*GetComponentVersionStagesOK handles this case with default header values.

A successful response.
*/
type GetComponentVersionStagesOK struct {
	Payload *service_model.V1Stage
}

func (o *GetComponentVersionStagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/hub/{entity}/versions/{name}/stages][%d] getComponentVersionStagesOK  %+v", 200, o.Payload)
}

func (o *GetComponentVersionStagesOK) GetPayload() *service_model.V1Stage {
	return o.Payload
}

func (o *GetComponentVersionStagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Stage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentVersionStagesNoContent creates a GetComponentVersionStagesNoContent with default headers values
func NewGetComponentVersionStagesNoContent() *GetComponentVersionStagesNoContent {
	return &GetComponentVersionStagesNoContent{}
}

/*GetComponentVersionStagesNoContent handles this case with default header values.

No content.
*/
type GetComponentVersionStagesNoContent struct {
	Payload interface{}
}

func (o *GetComponentVersionStagesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/hub/{entity}/versions/{name}/stages][%d] getComponentVersionStagesNoContent  %+v", 204, o.Payload)
}

func (o *GetComponentVersionStagesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetComponentVersionStagesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentVersionStagesForbidden creates a GetComponentVersionStagesForbidden with default headers values
func NewGetComponentVersionStagesForbidden() *GetComponentVersionStagesForbidden {
	return &GetComponentVersionStagesForbidden{}
}

/*GetComponentVersionStagesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetComponentVersionStagesForbidden struct {
	Payload interface{}
}

func (o *GetComponentVersionStagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/hub/{entity}/versions/{name}/stages][%d] getComponentVersionStagesForbidden  %+v", 403, o.Payload)
}

func (o *GetComponentVersionStagesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetComponentVersionStagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentVersionStagesNotFound creates a GetComponentVersionStagesNotFound with default headers values
func NewGetComponentVersionStagesNotFound() *GetComponentVersionStagesNotFound {
	return &GetComponentVersionStagesNotFound{}
}

/*GetComponentVersionStagesNotFound handles this case with default header values.

Resource does not exist.
*/
type GetComponentVersionStagesNotFound struct {
	Payload interface{}
}

func (o *GetComponentVersionStagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/hub/{entity}/versions/{name}/stages][%d] getComponentVersionStagesNotFound  %+v", 404, o.Payload)
}

func (o *GetComponentVersionStagesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetComponentVersionStagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComponentVersionStagesDefault creates a GetComponentVersionStagesDefault with default headers values
func NewGetComponentVersionStagesDefault(code int) *GetComponentVersionStagesDefault {
	return &GetComponentVersionStagesDefault{
		_statusCode: code,
	}
}

/*GetComponentVersionStagesDefault handles this case with default header values.

An unexpected error response.
*/
type GetComponentVersionStagesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get component version stages default response
func (o *GetComponentVersionStagesDefault) Code() int {
	return o._statusCode
}

func (o *GetComponentVersionStagesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/hub/{entity}/versions/{name}/stages][%d] GetComponentVersionStages default  %+v", o._statusCode, o.Payload)
}

func (o *GetComponentVersionStagesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetComponentVersionStagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
