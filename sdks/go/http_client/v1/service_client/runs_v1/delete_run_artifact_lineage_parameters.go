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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteRunArtifactLineageParams creates a new DeleteRunArtifactLineageParams object
// with the default values initialized.
func NewDeleteRunArtifactLineageParams() *DeleteRunArtifactLineageParams {
	var ()
	return &DeleteRunArtifactLineageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRunArtifactLineageParamsWithTimeout creates a new DeleteRunArtifactLineageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRunArtifactLineageParamsWithTimeout(timeout time.Duration) *DeleteRunArtifactLineageParams {
	var ()
	return &DeleteRunArtifactLineageParams{

		timeout: timeout,
	}
}

// NewDeleteRunArtifactLineageParamsWithContext creates a new DeleteRunArtifactLineageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRunArtifactLineageParamsWithContext(ctx context.Context) *DeleteRunArtifactLineageParams {
	var ()
	return &DeleteRunArtifactLineageParams{

		Context: ctx,
	}
}

// NewDeleteRunArtifactLineageParamsWithHTTPClient creates a new DeleteRunArtifactLineageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRunArtifactLineageParamsWithHTTPClient(client *http.Client) *DeleteRunArtifactLineageParams {
	var ()
	return &DeleteRunArtifactLineageParams{
		HTTPClient: client,
	}
}

/*DeleteRunArtifactLineageParams contains all the parameters to send to the API endpoint
for the delete run artifact lineage operation typically these are written to a http.Request
*/
type DeleteRunArtifactLineageParams struct {

	/*Name
	  Artifact name

	*/
	Name string
	/*Namespace
	  namespace.

	*/
	Namespace *string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the run will be assigned

	*/
	Project string
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) WithTimeout(timeout time.Duration) *DeleteRunArtifactLineageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) WithContext(ctx context.Context) *DeleteRunArtifactLineageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) WithHTTPClient(client *http.Client) *DeleteRunArtifactLineageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) WithName(name string) *DeleteRunArtifactLineageParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) WithNamespace(namespace *string) *DeleteRunArtifactLineageParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) WithOwner(owner string) *DeleteRunArtifactLineageParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) WithProject(project string) *DeleteRunArtifactLineageParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) SetProject(project string) {
	o.Project = project
}

// WithUUID adds the uuid to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) WithUUID(uuid string) *DeleteRunArtifactLineageParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete run artifact lineage params
func (o *DeleteRunArtifactLineageParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRunArtifactLineageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string
		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {
			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}

	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
