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

package agents_v1

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

// NewGetAgentTokenParams creates a new GetAgentTokenParams object
// with the default values initialized.
func NewGetAgentTokenParams() *GetAgentTokenParams {
	var ()
	return &GetAgentTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAgentTokenParamsWithTimeout creates a new GetAgentTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAgentTokenParamsWithTimeout(timeout time.Duration) *GetAgentTokenParams {
	var ()
	return &GetAgentTokenParams{

		timeout: timeout,
	}
}

// NewGetAgentTokenParamsWithContext creates a new GetAgentTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAgentTokenParamsWithContext(ctx context.Context) *GetAgentTokenParams {
	var ()
	return &GetAgentTokenParams{

		Context: ctx,
	}
}

// NewGetAgentTokenParamsWithHTTPClient creates a new GetAgentTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAgentTokenParamsWithHTTPClient(client *http.Client) *GetAgentTokenParams {
	var ()
	return &GetAgentTokenParams{
		HTTPClient: client,
	}
}

/*GetAgentTokenParams contains all the parameters to send to the API endpoint
for the get agent token operation typically these are written to a http.Request
*/
type GetAgentTokenParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get agent token params
func (o *GetAgentTokenParams) WithTimeout(timeout time.Duration) *GetAgentTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get agent token params
func (o *GetAgentTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get agent token params
func (o *GetAgentTokenParams) WithContext(ctx context.Context) *GetAgentTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get agent token params
func (o *GetAgentTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get agent token params
func (o *GetAgentTokenParams) WithHTTPClient(client *http.Client) *GetAgentTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get agent token params
func (o *GetAgentTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get agent token params
func (o *GetAgentTokenParams) WithOwner(owner string) *GetAgentTokenParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get agent token params
func (o *GetAgentTokenParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get agent token params
func (o *GetAgentTokenParams) WithUUID(uuid string) *GetAgentTokenParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get agent token params
func (o *GetAgentTokenParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAgentTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
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
