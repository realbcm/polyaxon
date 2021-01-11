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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewPatchOrganizationSettingsParams creates a new PatchOrganizationSettingsParams object
// with the default values initialized.
func NewPatchOrganizationSettingsParams() *PatchOrganizationSettingsParams {
	var ()
	return &PatchOrganizationSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchOrganizationSettingsParamsWithTimeout creates a new PatchOrganizationSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchOrganizationSettingsParamsWithTimeout(timeout time.Duration) *PatchOrganizationSettingsParams {
	var ()
	return &PatchOrganizationSettingsParams{

		timeout: timeout,
	}
}

// NewPatchOrganizationSettingsParamsWithContext creates a new PatchOrganizationSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchOrganizationSettingsParamsWithContext(ctx context.Context) *PatchOrganizationSettingsParams {
	var ()
	return &PatchOrganizationSettingsParams{

		Context: ctx,
	}
}

// NewPatchOrganizationSettingsParamsWithHTTPClient creates a new PatchOrganizationSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchOrganizationSettingsParamsWithHTTPClient(client *http.Client) *PatchOrganizationSettingsParams {
	var ()
	return &PatchOrganizationSettingsParams{
		HTTPClient: client,
	}
}

/*PatchOrganizationSettingsParams contains all the parameters to send to the API endpoint
for the patch organization settings operation typically these are written to a http.Request
*/
type PatchOrganizationSettingsParams struct {

	/*Body
	  Organization body

	*/
	Body *service_model.V1Organization
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch organization settings params
func (o *PatchOrganizationSettingsParams) WithTimeout(timeout time.Duration) *PatchOrganizationSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch organization settings params
func (o *PatchOrganizationSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch organization settings params
func (o *PatchOrganizationSettingsParams) WithContext(ctx context.Context) *PatchOrganizationSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch organization settings params
func (o *PatchOrganizationSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch organization settings params
func (o *PatchOrganizationSettingsParams) WithHTTPClient(client *http.Client) *PatchOrganizationSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch organization settings params
func (o *PatchOrganizationSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch organization settings params
func (o *PatchOrganizationSettingsParams) WithBody(body *service_model.V1Organization) *PatchOrganizationSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch organization settings params
func (o *PatchOrganizationSettingsParams) SetBody(body *service_model.V1Organization) {
	o.Body = body
}

// WithOwner adds the owner to the patch organization settings params
func (o *PatchOrganizationSettingsParams) WithOwner(owner string) *PatchOrganizationSettingsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch organization settings params
func (o *PatchOrganizationSettingsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *PatchOrganizationSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
