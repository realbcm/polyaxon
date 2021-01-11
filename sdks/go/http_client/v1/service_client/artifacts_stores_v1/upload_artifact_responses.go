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

package artifacts_stores_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UploadArtifactReader is a Reader for the UploadArtifact structure.
type UploadArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUploadArtifactNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUploadArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadArtifactOK creates a UploadArtifactOK with default headers values
func NewUploadArtifactOK() *UploadArtifactOK {
	return &UploadArtifactOK{}
}

/*UploadArtifactOK handles this case with default header values.

A successful response.
*/
type UploadArtifactOK struct {
}

func (o *UploadArtifactOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/catalogs/{owner}/artifacts/{uuid}/upload][%d] uploadArtifactOK ", 200)
}

func (o *UploadArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadArtifactNoContent creates a UploadArtifactNoContent with default headers values
func NewUploadArtifactNoContent() *UploadArtifactNoContent {
	return &UploadArtifactNoContent{}
}

/*UploadArtifactNoContent handles this case with default header values.

No content.
*/
type UploadArtifactNoContent struct {
	Payload interface{}
}

func (o *UploadArtifactNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/catalogs/{owner}/artifacts/{uuid}/upload][%d] uploadArtifactNoContent  %+v", 204, o.Payload)
}

func (o *UploadArtifactNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadArtifactNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadArtifactForbidden creates a UploadArtifactForbidden with default headers values
func NewUploadArtifactForbidden() *UploadArtifactForbidden {
	return &UploadArtifactForbidden{}
}

/*UploadArtifactForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UploadArtifactForbidden struct {
	Payload interface{}
}

func (o *UploadArtifactForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/catalogs/{owner}/artifacts/{uuid}/upload][%d] uploadArtifactForbidden  %+v", 403, o.Payload)
}

func (o *UploadArtifactForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadArtifactNotFound creates a UploadArtifactNotFound with default headers values
func NewUploadArtifactNotFound() *UploadArtifactNotFound {
	return &UploadArtifactNotFound{}
}

/*UploadArtifactNotFound handles this case with default header values.

Resource does not exist.
*/
type UploadArtifactNotFound struct {
	Payload interface{}
}

func (o *UploadArtifactNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/catalogs/{owner}/artifacts/{uuid}/upload][%d] uploadArtifactNotFound  %+v", 404, o.Payload)
}

func (o *UploadArtifactNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UploadArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
