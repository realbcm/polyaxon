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

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// OrganizationPlanReader is a Reader for the OrganizationPlan structure.
type OrganizationPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewOrganizationPlanNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewOrganizationPlanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOrganizationPlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationPlanOK creates a OrganizationPlanOK with default headers values
func NewOrganizationPlanOK() *OrganizationPlanOK {
	return &OrganizationPlanOK{}
}

/*OrganizationPlanOK handles this case with default header values.

A successful response.
*/
type OrganizationPlanOK struct {
	Payload *service_model.V1Organization
}

func (o *OrganizationPlanOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/plan][%d] organizationPlanOK  %+v", 200, o.Payload)
}

func (o *OrganizationPlanOK) GetPayload() *service_model.V1Organization {
	return o.Payload
}

func (o *OrganizationPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationPlanNoContent creates a OrganizationPlanNoContent with default headers values
func NewOrganizationPlanNoContent() *OrganizationPlanNoContent {
	return &OrganizationPlanNoContent{}
}

/*OrganizationPlanNoContent handles this case with default header values.

No content.
*/
type OrganizationPlanNoContent struct {
	Payload interface{}
}

func (o *OrganizationPlanNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/plan][%d] organizationPlanNoContent  %+v", 204, o.Payload)
}

func (o *OrganizationPlanNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationPlanNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationPlanForbidden creates a OrganizationPlanForbidden with default headers values
func NewOrganizationPlanForbidden() *OrganizationPlanForbidden {
	return &OrganizationPlanForbidden{}
}

/*OrganizationPlanForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type OrganizationPlanForbidden struct {
	Payload interface{}
}

func (o *OrganizationPlanForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/plan][%d] organizationPlanForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationPlanForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationPlanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationPlanNotFound creates a OrganizationPlanNotFound with default headers values
func NewOrganizationPlanNotFound() *OrganizationPlanNotFound {
	return &OrganizationPlanNotFound{}
}

/*OrganizationPlanNotFound handles this case with default header values.

Resource does not exist.
*/
type OrganizationPlanNotFound struct {
	Payload interface{}
}

func (o *OrganizationPlanNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/plan][%d] organizationPlanNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationPlanNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *OrganizationPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationPlanDefault creates a OrganizationPlanDefault with default headers values
func NewOrganizationPlanDefault(code int) *OrganizationPlanDefault {
	return &OrganizationPlanDefault{
		_statusCode: code,
	}
}

/*OrganizationPlanDefault handles this case with default header values.

An unexpected error response.
*/
type OrganizationPlanDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the organization plan default response
func (o *OrganizationPlanDefault) Code() int {
	return o._statusCode
}

func (o *OrganizationPlanDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/plan][%d] OrganizationPlan default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationPlanDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *OrganizationPlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
