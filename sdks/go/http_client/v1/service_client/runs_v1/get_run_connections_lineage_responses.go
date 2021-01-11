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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetRunConnectionsLineageReader is a Reader for the GetRunConnectionsLineage structure.
type GetRunConnectionsLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunConnectionsLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunConnectionsLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetRunConnectionsLineageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunConnectionsLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunConnectionsLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetRunConnectionsLineageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunConnectionsLineageOK creates a GetRunConnectionsLineageOK with default headers values
func NewGetRunConnectionsLineageOK() *GetRunConnectionsLineageOK {
	return &GetRunConnectionsLineageOK{}
}

/*GetRunConnectionsLineageOK handles this case with default header values.

A successful response.
*/
type GetRunConnectionsLineageOK struct {
	Payload *service_model.V1ListRunConnectionsResponse
}

func (o *GetRunConnectionsLineageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/connections][%d] getRunConnectionsLineageOK  %+v", 200, o.Payload)
}

func (o *GetRunConnectionsLineageOK) GetPayload() *service_model.V1ListRunConnectionsResponse {
	return o.Payload
}

func (o *GetRunConnectionsLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListRunConnectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunConnectionsLineageNoContent creates a GetRunConnectionsLineageNoContent with default headers values
func NewGetRunConnectionsLineageNoContent() *GetRunConnectionsLineageNoContent {
	return &GetRunConnectionsLineageNoContent{}
}

/*GetRunConnectionsLineageNoContent handles this case with default header values.

No content.
*/
type GetRunConnectionsLineageNoContent struct {
	Payload interface{}
}

func (o *GetRunConnectionsLineageNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/connections][%d] getRunConnectionsLineageNoContent  %+v", 204, o.Payload)
}

func (o *GetRunConnectionsLineageNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunConnectionsLineageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunConnectionsLineageForbidden creates a GetRunConnectionsLineageForbidden with default headers values
func NewGetRunConnectionsLineageForbidden() *GetRunConnectionsLineageForbidden {
	return &GetRunConnectionsLineageForbidden{}
}

/*GetRunConnectionsLineageForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetRunConnectionsLineageForbidden struct {
	Payload interface{}
}

func (o *GetRunConnectionsLineageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/connections][%d] getRunConnectionsLineageForbidden  %+v", 403, o.Payload)
}

func (o *GetRunConnectionsLineageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunConnectionsLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunConnectionsLineageNotFound creates a GetRunConnectionsLineageNotFound with default headers values
func NewGetRunConnectionsLineageNotFound() *GetRunConnectionsLineageNotFound {
	return &GetRunConnectionsLineageNotFound{}
}

/*GetRunConnectionsLineageNotFound handles this case with default header values.

Resource does not exist.
*/
type GetRunConnectionsLineageNotFound struct {
	Payload interface{}
}

func (o *GetRunConnectionsLineageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/connections][%d] getRunConnectionsLineageNotFound  %+v", 404, o.Payload)
}

func (o *GetRunConnectionsLineageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunConnectionsLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunConnectionsLineageDefault creates a GetRunConnectionsLineageDefault with default headers values
func NewGetRunConnectionsLineageDefault(code int) *GetRunConnectionsLineageDefault {
	return &GetRunConnectionsLineageDefault{
		_statusCode: code,
	}
}

/*GetRunConnectionsLineageDefault handles this case with default header values.

An unexpected error response.
*/
type GetRunConnectionsLineageDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get run connections lineage default response
func (o *GetRunConnectionsLineageDefault) Code() int {
	return o._statusCode
}

func (o *GetRunConnectionsLineageDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/runs/{uuid}/lineage/connections][%d] GetRunConnectionsLineage default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunConnectionsLineageDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetRunConnectionsLineageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
