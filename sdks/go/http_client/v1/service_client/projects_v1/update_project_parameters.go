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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewUpdateProjectParams creates a new UpdateProjectParams object
// with the default values initialized.
func NewUpdateProjectParams() *UpdateProjectParams {
	var ()
	return &UpdateProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectParamsWithTimeout creates a new UpdateProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProjectParamsWithTimeout(timeout time.Duration) *UpdateProjectParams {
	var ()
	return &UpdateProjectParams{

		timeout: timeout,
	}
}

// NewUpdateProjectParamsWithContext creates a new UpdateProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProjectParamsWithContext(ctx context.Context) *UpdateProjectParams {
	var ()
	return &UpdateProjectParams{

		Context: ctx,
	}
}

// NewUpdateProjectParamsWithHTTPClient creates a new UpdateProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProjectParamsWithHTTPClient(client *http.Client) *UpdateProjectParams {
	var ()
	return &UpdateProjectParams{
		HTTPClient: client,
	}
}

/*UpdateProjectParams contains all the parameters to send to the API endpoint
for the update project operation typically these are written to a http.Request
*/
type UpdateProjectParams struct {

	/*Body
	  Project body

	*/
	Body *service_model.V1Project
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*ProjectName
	  Required name

	*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update project params
func (o *UpdateProjectParams) WithTimeout(timeout time.Duration) *UpdateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project params
func (o *UpdateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project params
func (o *UpdateProjectParams) WithContext(ctx context.Context) *UpdateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project params
func (o *UpdateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project params
func (o *UpdateProjectParams) WithHTTPClient(client *http.Client) *UpdateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project params
func (o *UpdateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update project params
func (o *UpdateProjectParams) WithBody(body *service_model.V1Project) *UpdateProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update project params
func (o *UpdateProjectParams) SetBody(body *service_model.V1Project) {
	o.Body = body
}

// WithOwner adds the owner to the update project params
func (o *UpdateProjectParams) WithOwner(owner string) *UpdateProjectParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the update project params
func (o *UpdateProjectParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProjectName adds the projectName to the update project params
func (o *UpdateProjectParams) WithProjectName(projectName string) *UpdateProjectParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the update project params
func (o *UpdateProjectParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project.name
	if err := r.SetPathParam("project.name", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
