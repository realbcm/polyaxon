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

// DeleteRunArtifactLineageReader is a Reader for the DeleteRunArtifactLineage structure.
type DeleteRunArtifactLineageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRunArtifactLineageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRunArtifactLineageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteRunArtifactLineageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRunArtifactLineageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRunArtifactLineageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRunArtifactLineageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRunArtifactLineageOK creates a DeleteRunArtifactLineageOK with default headers values
func NewDeleteRunArtifactLineageOK() *DeleteRunArtifactLineageOK {
	return &DeleteRunArtifactLineageOK{}
}

/*DeleteRunArtifactLineageOK handles this case with default header values.

A successful response.
*/
type DeleteRunArtifactLineageOK struct {
}

func (o *DeleteRunArtifactLineageOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] deleteRunArtifactLineageOK ", 200)
}

func (o *DeleteRunArtifactLineageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRunArtifactLineageNoContent creates a DeleteRunArtifactLineageNoContent with default headers values
func NewDeleteRunArtifactLineageNoContent() *DeleteRunArtifactLineageNoContent {
	return &DeleteRunArtifactLineageNoContent{}
}

/*DeleteRunArtifactLineageNoContent handles this case with default header values.

No content.
*/
type DeleteRunArtifactLineageNoContent struct {
	Payload interface{}
}

func (o *DeleteRunArtifactLineageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] deleteRunArtifactLineageNoContent  %+v", 204, o.Payload)
}

func (o *DeleteRunArtifactLineageNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunArtifactLineageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunArtifactLineageForbidden creates a DeleteRunArtifactLineageForbidden with default headers values
func NewDeleteRunArtifactLineageForbidden() *DeleteRunArtifactLineageForbidden {
	return &DeleteRunArtifactLineageForbidden{}
}

/*DeleteRunArtifactLineageForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteRunArtifactLineageForbidden struct {
	Payload interface{}
}

func (o *DeleteRunArtifactLineageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] deleteRunArtifactLineageForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRunArtifactLineageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunArtifactLineageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunArtifactLineageNotFound creates a DeleteRunArtifactLineageNotFound with default headers values
func NewDeleteRunArtifactLineageNotFound() *DeleteRunArtifactLineageNotFound {
	return &DeleteRunArtifactLineageNotFound{}
}

/*DeleteRunArtifactLineageNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteRunArtifactLineageNotFound struct {
	Payload interface{}
}

func (o *DeleteRunArtifactLineageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] deleteRunArtifactLineageNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRunArtifactLineageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteRunArtifactLineageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRunArtifactLineageDefault creates a DeleteRunArtifactLineageDefault with default headers values
func NewDeleteRunArtifactLineageDefault(code int) *DeleteRunArtifactLineageDefault {
	return &DeleteRunArtifactLineageDefault{
		_statusCode: code,
	}
}

/*DeleteRunArtifactLineageDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteRunArtifactLineageDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete run artifact lineage default response
func (o *DeleteRunArtifactLineageDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRunArtifactLineageDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/runs/{uuid}/lineage/artifacts/{name}][%d] DeleteRunArtifactLineage default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRunArtifactLineageDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteRunArtifactLineageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
