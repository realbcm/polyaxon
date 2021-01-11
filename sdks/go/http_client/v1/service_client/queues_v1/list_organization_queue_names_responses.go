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

package queues_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListOrganizationQueueNamesReader is a Reader for the ListOrganizationQueueNames structure.
type ListOrganizationQueueNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationQueueNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationQueueNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListOrganizationQueueNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListOrganizationQueueNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListOrganizationQueueNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListOrganizationQueueNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOrganizationQueueNamesOK creates a ListOrganizationQueueNamesOK with default headers values
func NewListOrganizationQueueNamesOK() *ListOrganizationQueueNamesOK {
	return &ListOrganizationQueueNamesOK{}
}

/*ListOrganizationQueueNamesOK handles this case with default header values.

A successful response.
*/
type ListOrganizationQueueNamesOK struct {
	Payload *service_model.V1ListQueuesResponse
}

func (o *ListOrganizationQueueNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues/names][%d] listOrganizationQueueNamesOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationQueueNamesOK) GetPayload() *service_model.V1ListQueuesResponse {
	return o.Payload
}

func (o *ListOrganizationQueueNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListQueuesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationQueueNamesNoContent creates a ListOrganizationQueueNamesNoContent with default headers values
func NewListOrganizationQueueNamesNoContent() *ListOrganizationQueueNamesNoContent {
	return &ListOrganizationQueueNamesNoContent{}
}

/*ListOrganizationQueueNamesNoContent handles this case with default header values.

No content.
*/
type ListOrganizationQueueNamesNoContent struct {
	Payload interface{}
}

func (o *ListOrganizationQueueNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues/names][%d] listOrganizationQueueNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListOrganizationQueueNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationQueueNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationQueueNamesForbidden creates a ListOrganizationQueueNamesForbidden with default headers values
func NewListOrganizationQueueNamesForbidden() *ListOrganizationQueueNamesForbidden {
	return &ListOrganizationQueueNamesForbidden{}
}

/*ListOrganizationQueueNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListOrganizationQueueNamesForbidden struct {
	Payload interface{}
}

func (o *ListOrganizationQueueNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues/names][%d] listOrganizationQueueNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListOrganizationQueueNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationQueueNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationQueueNamesNotFound creates a ListOrganizationQueueNamesNotFound with default headers values
func NewListOrganizationQueueNamesNotFound() *ListOrganizationQueueNamesNotFound {
	return &ListOrganizationQueueNamesNotFound{}
}

/*ListOrganizationQueueNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListOrganizationQueueNamesNotFound struct {
	Payload interface{}
}

func (o *ListOrganizationQueueNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues/names][%d] listOrganizationQueueNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListOrganizationQueueNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListOrganizationQueueNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationQueueNamesDefault creates a ListOrganizationQueueNamesDefault with default headers values
func NewListOrganizationQueueNamesDefault(code int) *ListOrganizationQueueNamesDefault {
	return &ListOrganizationQueueNamesDefault{
		_statusCode: code,
	}
}

/*ListOrganizationQueueNamesDefault handles this case with default header values.

An unexpected error response.
*/
type ListOrganizationQueueNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list organization queue names default response
func (o *ListOrganizationQueueNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListOrganizationQueueNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/queues/names][%d] ListOrganizationQueueNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListOrganizationQueueNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListOrganizationQueueNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
