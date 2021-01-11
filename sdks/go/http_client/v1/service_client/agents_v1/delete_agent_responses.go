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

package agents_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteAgentReader is a Reader for the DeleteAgent structure.
type DeleteAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteAgentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteAgentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAgentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteAgentOK creates a DeleteAgentOK with default headers values
func NewDeleteAgentOK() *DeleteAgentOK {
	return &DeleteAgentOK{}
}

/*DeleteAgentOK handles this case with default header values.

A successful response.
*/
type DeleteAgentOK struct {
}

func (o *DeleteAgentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentOK ", 200)
}

func (o *DeleteAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAgentNoContent creates a DeleteAgentNoContent with default headers values
func NewDeleteAgentNoContent() *DeleteAgentNoContent {
	return &DeleteAgentNoContent{}
}

/*DeleteAgentNoContent handles this case with default header values.

No content.
*/
type DeleteAgentNoContent struct {
	Payload interface{}
}

func (o *DeleteAgentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentNoContent  %+v", 204, o.Payload)
}

func (o *DeleteAgentNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAgentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentForbidden creates a DeleteAgentForbidden with default headers values
func NewDeleteAgentForbidden() *DeleteAgentForbidden {
	return &DeleteAgentForbidden{}
}

/*DeleteAgentForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteAgentForbidden struct {
	Payload interface{}
}

func (o *DeleteAgentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAgentForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAgentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentNotFound creates a DeleteAgentNotFound with default headers values
func NewDeleteAgentNotFound() *DeleteAgentNotFound {
	return &DeleteAgentNotFound{}
}

/*DeleteAgentNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteAgentNotFound struct {
	Payload interface{}
}

func (o *DeleteAgentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] deleteAgentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAgentNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteAgentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAgentDefault creates a DeleteAgentDefault with default headers values
func NewDeleteAgentDefault(code int) *DeleteAgentDefault {
	return &DeleteAgentDefault{
		_statusCode: code,
	}
}

/*DeleteAgentDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteAgentDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete agent default response
func (o *DeleteAgentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAgentDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/agents/{uuid}][%d] DeleteAgent default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteAgentDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
