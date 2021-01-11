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

package model_registry_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// BookmarkModelRegistryReader is a Reader for the BookmarkModelRegistry structure.
type BookmarkModelRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookmarkModelRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookmarkModelRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewBookmarkModelRegistryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewBookmarkModelRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBookmarkModelRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBookmarkModelRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBookmarkModelRegistryOK creates a BookmarkModelRegistryOK with default headers values
func NewBookmarkModelRegistryOK() *BookmarkModelRegistryOK {
	return &BookmarkModelRegistryOK{}
}

/*BookmarkModelRegistryOK handles this case with default header values.

A successful response.
*/
type BookmarkModelRegistryOK struct {
}

func (o *BookmarkModelRegistryOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/{name}/bookmark][%d] bookmarkModelRegistryOK ", 200)
}

func (o *BookmarkModelRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBookmarkModelRegistryNoContent creates a BookmarkModelRegistryNoContent with default headers values
func NewBookmarkModelRegistryNoContent() *BookmarkModelRegistryNoContent {
	return &BookmarkModelRegistryNoContent{}
}

/*BookmarkModelRegistryNoContent handles this case with default header values.

No content.
*/
type BookmarkModelRegistryNoContent struct {
	Payload interface{}
}

func (o *BookmarkModelRegistryNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/{name}/bookmark][%d] bookmarkModelRegistryNoContent  %+v", 204, o.Payload)
}

func (o *BookmarkModelRegistryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkModelRegistryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkModelRegistryForbidden creates a BookmarkModelRegistryForbidden with default headers values
func NewBookmarkModelRegistryForbidden() *BookmarkModelRegistryForbidden {
	return &BookmarkModelRegistryForbidden{}
}

/*BookmarkModelRegistryForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type BookmarkModelRegistryForbidden struct {
	Payload interface{}
}

func (o *BookmarkModelRegistryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/{name}/bookmark][%d] bookmarkModelRegistryForbidden  %+v", 403, o.Payload)
}

func (o *BookmarkModelRegistryForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkModelRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkModelRegistryNotFound creates a BookmarkModelRegistryNotFound with default headers values
func NewBookmarkModelRegistryNotFound() *BookmarkModelRegistryNotFound {
	return &BookmarkModelRegistryNotFound{}
}

/*BookmarkModelRegistryNotFound handles this case with default header values.

Resource does not exist.
*/
type BookmarkModelRegistryNotFound struct {
	Payload interface{}
}

func (o *BookmarkModelRegistryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/{name}/bookmark][%d] bookmarkModelRegistryNotFound  %+v", 404, o.Payload)
}

func (o *BookmarkModelRegistryNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BookmarkModelRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookmarkModelRegistryDefault creates a BookmarkModelRegistryDefault with default headers values
func NewBookmarkModelRegistryDefault(code int) *BookmarkModelRegistryDefault {
	return &BookmarkModelRegistryDefault{
		_statusCode: code,
	}
}

/*BookmarkModelRegistryDefault handles this case with default header values.

An unexpected error response.
*/
type BookmarkModelRegistryDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the bookmark model registry default response
func (o *BookmarkModelRegistryDefault) Code() int {
	return o._statusCode
}

func (o *BookmarkModelRegistryDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry/{name}/bookmark][%d] BookmarkModelRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *BookmarkModelRegistryDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *BookmarkModelRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
