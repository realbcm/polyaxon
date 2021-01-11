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

// NewStopRunTensorboardParams creates a new StopRunTensorboardParams object
// with the default values initialized.
func NewStopRunTensorboardParams() *StopRunTensorboardParams {
	var ()
	return &StopRunTensorboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopRunTensorboardParamsWithTimeout creates a new StopRunTensorboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopRunTensorboardParamsWithTimeout(timeout time.Duration) *StopRunTensorboardParams {
	var ()
	return &StopRunTensorboardParams{

		timeout: timeout,
	}
}

// NewStopRunTensorboardParamsWithContext creates a new StopRunTensorboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopRunTensorboardParamsWithContext(ctx context.Context) *StopRunTensorboardParams {
	var ()
	return &StopRunTensorboardParams{

		Context: ctx,
	}
}

// NewStopRunTensorboardParamsWithHTTPClient creates a new StopRunTensorboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopRunTensorboardParamsWithHTTPClient(client *http.Client) *StopRunTensorboardParams {
	var ()
	return &StopRunTensorboardParams{
		HTTPClient: client,
	}
}

/*StopRunTensorboardParams contains all the parameters to send to the API endpoint
for the stop run tensorboard operation typically these are written to a http.Request
*/
type StopRunTensorboardParams struct {

	/*Entity
	  Entity: project name, hub name, registry name, ...

	*/
	Entity string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*UUID
	  Uuid identifier of the sub-entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop run tensorboard params
func (o *StopRunTensorboardParams) WithTimeout(timeout time.Duration) *StopRunTensorboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop run tensorboard params
func (o *StopRunTensorboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop run tensorboard params
func (o *StopRunTensorboardParams) WithContext(ctx context.Context) *StopRunTensorboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop run tensorboard params
func (o *StopRunTensorboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop run tensorboard params
func (o *StopRunTensorboardParams) WithHTTPClient(client *http.Client) *StopRunTensorboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop run tensorboard params
func (o *StopRunTensorboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the stop run tensorboard params
func (o *StopRunTensorboardParams) WithEntity(entity string) *StopRunTensorboardParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the stop run tensorboard params
func (o *StopRunTensorboardParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the stop run tensorboard params
func (o *StopRunTensorboardParams) WithOwner(owner string) *StopRunTensorboardParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the stop run tensorboard params
func (o *StopRunTensorboardParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the stop run tensorboard params
func (o *StopRunTensorboardParams) WithUUID(uuid string) *StopRunTensorboardParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the stop run tensorboard params
func (o *StopRunTensorboardParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *StopRunTensorboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entity
	if err := r.SetPathParam("entity", o.Entity); err != nil {
		return err
	}

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
