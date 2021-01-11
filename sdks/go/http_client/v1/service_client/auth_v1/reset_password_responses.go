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

package auth_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ResetPasswordReader is a Reader for the ResetPassword structure.
type ResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewResetPasswordNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewResetPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResetPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResetPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResetPasswordOK creates a ResetPasswordOK with default headers values
func NewResetPasswordOK() *ResetPasswordOK {
	return &ResetPasswordOK{}
}

/*ResetPasswordOK handles this case with default header values.

A successful response.
*/
type ResetPasswordOK struct {
}

func (o *ResetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password][%d] resetPasswordOK ", 200)
}

func (o *ResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetPasswordNoContent creates a ResetPasswordNoContent with default headers values
func NewResetPasswordNoContent() *ResetPasswordNoContent {
	return &ResetPasswordNoContent{}
}

/*ResetPasswordNoContent handles this case with default header values.

No content.
*/
type ResetPasswordNoContent struct {
	Payload interface{}
}

func (o *ResetPasswordNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password][%d] resetPasswordNoContent  %+v", 204, o.Payload)
}

func (o *ResetPasswordNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ResetPasswordNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordForbidden creates a ResetPasswordForbidden with default headers values
func NewResetPasswordForbidden() *ResetPasswordForbidden {
	return &ResetPasswordForbidden{}
}

/*ResetPasswordForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ResetPasswordForbidden struct {
	Payload interface{}
}

func (o *ResetPasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password][%d] resetPasswordForbidden  %+v", 403, o.Payload)
}

func (o *ResetPasswordForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ResetPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordNotFound creates a ResetPasswordNotFound with default headers values
func NewResetPasswordNotFound() *ResetPasswordNotFound {
	return &ResetPasswordNotFound{}
}

/*ResetPasswordNotFound handles this case with default header values.

Resource does not exist.
*/
type ResetPasswordNotFound struct {
	Payload interface{}
}

func (o *ResetPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password][%d] resetPasswordNotFound  %+v", 404, o.Payload)
}

func (o *ResetPasswordNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ResetPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetPasswordDefault creates a ResetPasswordDefault with default headers values
func NewResetPasswordDefault(code int) *ResetPasswordDefault {
	return &ResetPasswordDefault{
		_statusCode: code,
	}
}

/*ResetPasswordDefault handles this case with default header values.

An unexpected error response.
*/
type ResetPasswordDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the reset password default response
func (o *ResetPasswordDefault) Code() int {
	return o._statusCode
}

func (o *ResetPasswordDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/auth/reset-password][%d] ResetPassword default  %+v", o._statusCode, o.Payload)
}

func (o *ResetPasswordDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ResetPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
