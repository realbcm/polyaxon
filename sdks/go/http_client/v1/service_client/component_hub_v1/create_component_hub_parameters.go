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

package component_hub_v1

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

// NewCreateComponentHubParams creates a new CreateComponentHubParams object
// with the default values initialized.
func NewCreateComponentHubParams() *CreateComponentHubParams {
	var ()
	return &CreateComponentHubParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateComponentHubParamsWithTimeout creates a new CreateComponentHubParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateComponentHubParamsWithTimeout(timeout time.Duration) *CreateComponentHubParams {
	var ()
	return &CreateComponentHubParams{

		timeout: timeout,
	}
}

// NewCreateComponentHubParamsWithContext creates a new CreateComponentHubParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateComponentHubParamsWithContext(ctx context.Context) *CreateComponentHubParams {
	var ()
	return &CreateComponentHubParams{

		Context: ctx,
	}
}

// NewCreateComponentHubParamsWithHTTPClient creates a new CreateComponentHubParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateComponentHubParamsWithHTTPClient(client *http.Client) *CreateComponentHubParams {
	var ()
	return &CreateComponentHubParams{
		HTTPClient: client,
	}
}

/*CreateComponentHubParams contains all the parameters to send to the API endpoint
for the create component hub operation typically these are written to a http.Request
*/
type CreateComponentHubParams struct {

	/*Body
	  Component body

	*/
	Body *service_model.V1ComponentHub
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create component hub params
func (o *CreateComponentHubParams) WithTimeout(timeout time.Duration) *CreateComponentHubParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create component hub params
func (o *CreateComponentHubParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create component hub params
func (o *CreateComponentHubParams) WithContext(ctx context.Context) *CreateComponentHubParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create component hub params
func (o *CreateComponentHubParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create component hub params
func (o *CreateComponentHubParams) WithHTTPClient(client *http.Client) *CreateComponentHubParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create component hub params
func (o *CreateComponentHubParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create component hub params
func (o *CreateComponentHubParams) WithBody(body *service_model.V1ComponentHub) *CreateComponentHubParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create component hub params
func (o *CreateComponentHubParams) SetBody(body *service_model.V1ComponentHub) {
	o.Body = body
}

// WithOwner adds the owner to the create component hub params
func (o *CreateComponentHubParams) WithOwner(owner string) *CreateComponentHubParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create component hub params
func (o *CreateComponentHubParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateComponentHubParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
