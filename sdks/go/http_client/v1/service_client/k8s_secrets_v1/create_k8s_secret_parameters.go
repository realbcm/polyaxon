// Copyright 2019 Polyaxon, Inc.
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

package k8s_secrets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewCreateK8sSecretParams creates a new CreateK8sSecretParams object
// with the default values initialized.
func NewCreateK8sSecretParams() *CreateK8sSecretParams {
	var ()
	return &CreateK8sSecretParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateK8sSecretParamsWithTimeout creates a new CreateK8sSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateK8sSecretParamsWithTimeout(timeout time.Duration) *CreateK8sSecretParams {
	var ()
	return &CreateK8sSecretParams{

		timeout: timeout,
	}
}

// NewCreateK8sSecretParamsWithContext creates a new CreateK8sSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateK8sSecretParamsWithContext(ctx context.Context) *CreateK8sSecretParams {
	var ()
	return &CreateK8sSecretParams{

		Context: ctx,
	}
}

// NewCreateK8sSecretParamsWithHTTPClient creates a new CreateK8sSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateK8sSecretParamsWithHTTPClient(client *http.Client) *CreateK8sSecretParams {
	var ()
	return &CreateK8sSecretParams{
		HTTPClient: client,
	}
}

/*CreateK8sSecretParams contains all the parameters to send to the API endpoint
for the create k8s secret operation typically these are written to a http.Request
*/
type CreateK8sSecretParams struct {

	/*Body
	  Artifact store body

	*/
	Body *service_model.V1K8sResource
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create k8s secret params
func (o *CreateK8sSecretParams) WithTimeout(timeout time.Duration) *CreateK8sSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create k8s secret params
func (o *CreateK8sSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create k8s secret params
func (o *CreateK8sSecretParams) WithContext(ctx context.Context) *CreateK8sSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create k8s secret params
func (o *CreateK8sSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create k8s secret params
func (o *CreateK8sSecretParams) WithHTTPClient(client *http.Client) *CreateK8sSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create k8s secret params
func (o *CreateK8sSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create k8s secret params
func (o *CreateK8sSecretParams) WithBody(body *service_model.V1K8sResource) *CreateK8sSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create k8s secret params
func (o *CreateK8sSecretParams) SetBody(body *service_model.V1K8sResource) {
	o.Body = body
}

// WithOwner adds the owner to the create k8s secret params
func (o *CreateK8sSecretParams) WithOwner(owner string) *CreateK8sSecretParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the create k8s secret params
func (o *CreateK8sSecretParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *CreateK8sSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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