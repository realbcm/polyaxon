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

package project_searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetProjectSearchReader is a Reader for the GetProjectSearch structure.
type GetProjectSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetProjectSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetProjectSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetProjectSearchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectSearchOK creates a GetProjectSearchOK with default headers values
func NewGetProjectSearchOK() *GetProjectSearchOK {
	return &GetProjectSearchOK{}
}

/*GetProjectSearchOK handles this case with default header values.

A successful response.
*/
type GetProjectSearchOK struct {
	Payload *service_model.V1Search
}

func (o *GetProjectSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchOK  %+v", 200, o.Payload)
}

func (o *GetProjectSearchOK) GetPayload() *service_model.V1Search {
	return o.Payload
}

func (o *GetProjectSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Search)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSearchNoContent creates a GetProjectSearchNoContent with default headers values
func NewGetProjectSearchNoContent() *GetProjectSearchNoContent {
	return &GetProjectSearchNoContent{}
}

/*GetProjectSearchNoContent handles this case with default header values.

No content.
*/
type GetProjectSearchNoContent struct {
	Payload interface{}
}

func (o *GetProjectSearchNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchNoContent  %+v", 204, o.Payload)
}

func (o *GetProjectSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSearchForbidden creates a GetProjectSearchForbidden with default headers values
func NewGetProjectSearchForbidden() *GetProjectSearchForbidden {
	return &GetProjectSearchForbidden{}
}

/*GetProjectSearchForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetProjectSearchForbidden struct {
	Payload interface{}
}

func (o *GetProjectSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSearchNotFound creates a GetProjectSearchNotFound with default headers values
func NewGetProjectSearchNotFound() *GetProjectSearchNotFound {
	return &GetProjectSearchNotFound{}
}

/*GetProjectSearchNotFound handles this case with default header values.

Resource does not exist.
*/
type GetProjectSearchNotFound struct {
	Payload interface{}
}

func (o *GetProjectSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] getProjectSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectSearchDefault creates a GetProjectSearchDefault with default headers values
func NewGetProjectSearchDefault(code int) *GetProjectSearchDefault {
	return &GetProjectSearchDefault{
		_statusCode: code,
	}
}

/*GetProjectSearchDefault handles this case with default header values.

An unexpected error response.
*/
type GetProjectSearchDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get project search default response
func (o *GetProjectSearchDefault) Code() int {
	return o._statusCode
}

func (o *GetProjectSearchDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{entity}/searches/{uuid}][%d] GetProjectSearch default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectSearchDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetProjectSearchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
