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

package teams_v1

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

// NewPatchTeamParams creates a new PatchTeamParams object
// with the default values initialized.
func NewPatchTeamParams() *PatchTeamParams {
	var ()
	return &PatchTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchTeamParamsWithTimeout creates a new PatchTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchTeamParamsWithTimeout(timeout time.Duration) *PatchTeamParams {
	var ()
	return &PatchTeamParams{

		timeout: timeout,
	}
}

// NewPatchTeamParamsWithContext creates a new PatchTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchTeamParamsWithContext(ctx context.Context) *PatchTeamParams {
	var ()
	return &PatchTeamParams{

		Context: ctx,
	}
}

// NewPatchTeamParamsWithHTTPClient creates a new PatchTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchTeamParamsWithHTTPClient(client *http.Client) *PatchTeamParams {
	var ()
	return &PatchTeamParams{
		HTTPClient: client,
	}
}

/*PatchTeamParams contains all the parameters to send to the API endpoint
for the patch team operation typically these are written to a http.Request
*/
type PatchTeamParams struct {

	/*Body
	  Team body

	*/
	Body *service_model.V1Team
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*TeamName
	  Name

	*/
	TeamName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch team params
func (o *PatchTeamParams) WithTimeout(timeout time.Duration) *PatchTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch team params
func (o *PatchTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch team params
func (o *PatchTeamParams) WithContext(ctx context.Context) *PatchTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch team params
func (o *PatchTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch team params
func (o *PatchTeamParams) WithHTTPClient(client *http.Client) *PatchTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch team params
func (o *PatchTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch team params
func (o *PatchTeamParams) WithBody(body *service_model.V1Team) *PatchTeamParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch team params
func (o *PatchTeamParams) SetBody(body *service_model.V1Team) {
	o.Body = body
}

// WithOwner adds the owner to the patch team params
func (o *PatchTeamParams) WithOwner(owner string) *PatchTeamParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch team params
func (o *PatchTeamParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithTeamName adds the teamName to the patch team params
func (o *PatchTeamParams) WithTeamName(teamName string) *PatchTeamParams {
	o.SetTeamName(teamName)
	return o
}

// SetTeamName adds the teamName to the patch team params
func (o *PatchTeamParams) SetTeamName(teamName string) {
	o.TeamName = teamName
}

// WriteToRequest writes these params to a swagger request
func (o *PatchTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param team.name
	if err := r.SetPathParam("team.name", o.TeamName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
