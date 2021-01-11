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

package model_registry_v1

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

// NewRestoreModelRegistryParams creates a new RestoreModelRegistryParams object
// with the default values initialized.
func NewRestoreModelRegistryParams() *RestoreModelRegistryParams {
	var ()
	return &RestoreModelRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreModelRegistryParamsWithTimeout creates a new RestoreModelRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRestoreModelRegistryParamsWithTimeout(timeout time.Duration) *RestoreModelRegistryParams {
	var ()
	return &RestoreModelRegistryParams{

		timeout: timeout,
	}
}

// NewRestoreModelRegistryParamsWithContext creates a new RestoreModelRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewRestoreModelRegistryParamsWithContext(ctx context.Context) *RestoreModelRegistryParams {
	var ()
	return &RestoreModelRegistryParams{

		Context: ctx,
	}
}

// NewRestoreModelRegistryParamsWithHTTPClient creates a new RestoreModelRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRestoreModelRegistryParamsWithHTTPClient(client *http.Client) *RestoreModelRegistryParams {
	var ()
	return &RestoreModelRegistryParams{
		HTTPClient: client,
	}
}

/*RestoreModelRegistryParams contains all the parameters to send to the API endpoint
for the restore model registry operation typically these are written to a http.Request
*/
type RestoreModelRegistryParams struct {

	/*Name
	  Component under namesapce

	*/
	Name string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the restore model registry params
func (o *RestoreModelRegistryParams) WithTimeout(timeout time.Duration) *RestoreModelRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore model registry params
func (o *RestoreModelRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore model registry params
func (o *RestoreModelRegistryParams) WithContext(ctx context.Context) *RestoreModelRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore model registry params
func (o *RestoreModelRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore model registry params
func (o *RestoreModelRegistryParams) WithHTTPClient(client *http.Client) *RestoreModelRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore model registry params
func (o *RestoreModelRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the restore model registry params
func (o *RestoreModelRegistryParams) WithName(name string) *RestoreModelRegistryParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the restore model registry params
func (o *RestoreModelRegistryParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the restore model registry params
func (o *RestoreModelRegistryParams) WithOwner(owner string) *RestoreModelRegistryParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the restore model registry params
func (o *RestoreModelRegistryParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreModelRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
