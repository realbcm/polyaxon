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

// NewUnbookmarkRunParams creates a new UnbookmarkRunParams object
// with the default values initialized.
func NewUnbookmarkRunParams() *UnbookmarkRunParams {
	var ()
	return &UnbookmarkRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnbookmarkRunParamsWithTimeout creates a new UnbookmarkRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnbookmarkRunParamsWithTimeout(timeout time.Duration) *UnbookmarkRunParams {
	var ()
	return &UnbookmarkRunParams{

		timeout: timeout,
	}
}

// NewUnbookmarkRunParamsWithContext creates a new UnbookmarkRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnbookmarkRunParamsWithContext(ctx context.Context) *UnbookmarkRunParams {
	var ()
	return &UnbookmarkRunParams{

		Context: ctx,
	}
}

// NewUnbookmarkRunParamsWithHTTPClient creates a new UnbookmarkRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnbookmarkRunParamsWithHTTPClient(client *http.Client) *UnbookmarkRunParams {
	var ()
	return &UnbookmarkRunParams{
		HTTPClient: client,
	}
}

/*UnbookmarkRunParams contains all the parameters to send to the API endpoint
for the unbookmark run operation typically these are written to a http.Request
*/
type UnbookmarkRunParams struct {

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

// WithTimeout adds the timeout to the unbookmark run params
func (o *UnbookmarkRunParams) WithTimeout(timeout time.Duration) *UnbookmarkRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unbookmark run params
func (o *UnbookmarkRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unbookmark run params
func (o *UnbookmarkRunParams) WithContext(ctx context.Context) *UnbookmarkRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unbookmark run params
func (o *UnbookmarkRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unbookmark run params
func (o *UnbookmarkRunParams) WithHTTPClient(client *http.Client) *UnbookmarkRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unbookmark run params
func (o *UnbookmarkRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntity adds the entity to the unbookmark run params
func (o *UnbookmarkRunParams) WithEntity(entity string) *UnbookmarkRunParams {
	o.SetEntity(entity)
	return o
}

// SetEntity adds the entity to the unbookmark run params
func (o *UnbookmarkRunParams) SetEntity(entity string) {
	o.Entity = entity
}

// WithOwner adds the owner to the unbookmark run params
func (o *UnbookmarkRunParams) WithOwner(owner string) *UnbookmarkRunParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the unbookmark run params
func (o *UnbookmarkRunParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the unbookmark run params
func (o *UnbookmarkRunParams) WithUUID(uuid string) *UnbookmarkRunParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the unbookmark run params
func (o *UnbookmarkRunParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *UnbookmarkRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
