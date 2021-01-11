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

// DeleteOrganizationInvitationReader is a Reader for the DeleteOrganizationInvitation structure.
type DeleteOrganizationInvitationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationInvitationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationInvitationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteOrganizationInvitationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteOrganizationInvitationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationInvitationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteOrganizationInvitationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrganizationInvitationOK creates a DeleteOrganizationInvitationOK with default headers values
func NewDeleteOrganizationInvitationOK() *DeleteOrganizationInvitationOK {
	return &DeleteOrganizationInvitationOK{}
}

/*DeleteOrganizationInvitationOK handles this case with default header values.

A successful response.
*/
type DeleteOrganizationInvitationOK struct {
}

func (o *DeleteOrganizationInvitationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/invitations][%d] deleteOrganizationInvitationOK ", 200)
}

func (o *DeleteOrganizationInvitationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationInvitationNoContent creates a DeleteOrganizationInvitationNoContent with default headers values
func NewDeleteOrganizationInvitationNoContent() *DeleteOrganizationInvitationNoContent {
	return &DeleteOrganizationInvitationNoContent{}
}

/*DeleteOrganizationInvitationNoContent handles this case with default header values.

No content.
*/
type DeleteOrganizationInvitationNoContent struct {
	Payload interface{}
}

func (o *DeleteOrganizationInvitationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/invitations][%d] deleteOrganizationInvitationNoContent  %+v", 204, o.Payload)
}

func (o *DeleteOrganizationInvitationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationInvitationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationInvitationForbidden creates a DeleteOrganizationInvitationForbidden with default headers values
func NewDeleteOrganizationInvitationForbidden() *DeleteOrganizationInvitationForbidden {
	return &DeleteOrganizationInvitationForbidden{}
}

/*DeleteOrganizationInvitationForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteOrganizationInvitationForbidden struct {
	Payload interface{}
}

func (o *DeleteOrganizationInvitationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/invitations][%d] deleteOrganizationInvitationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganizationInvitationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationInvitationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationInvitationNotFound creates a DeleteOrganizationInvitationNotFound with default headers values
func NewDeleteOrganizationInvitationNotFound() *DeleteOrganizationInvitationNotFound {
	return &DeleteOrganizationInvitationNotFound{}
}

/*DeleteOrganizationInvitationNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteOrganizationInvitationNotFound struct {
	Payload interface{}
}

func (o *DeleteOrganizationInvitationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/invitations][%d] deleteOrganizationInvitationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationInvitationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationInvitationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationInvitationDefault creates a DeleteOrganizationInvitationDefault with default headers values
func NewDeleteOrganizationInvitationDefault(code int) *DeleteOrganizationInvitationDefault {
	return &DeleteOrganizationInvitationDefault{
		_statusCode: code,
	}
}

/*DeleteOrganizationInvitationDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteOrganizationInvitationDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete organization invitation default response
func (o *DeleteOrganizationInvitationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrganizationInvitationDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}/invitations][%d] DeleteOrganizationInvitation default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationInvitationDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteOrganizationInvitationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
