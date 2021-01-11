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
)

// NewDeleteTeamMemberParams creates a new DeleteTeamMemberParams object
// with the default values initialized.
func NewDeleteTeamMemberParams() *DeleteTeamMemberParams {
	var ()
	return &DeleteTeamMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamMemberParamsWithTimeout creates a new DeleteTeamMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTeamMemberParamsWithTimeout(timeout time.Duration) *DeleteTeamMemberParams {
	var ()
	return &DeleteTeamMemberParams{

		timeout: timeout,
	}
}

// NewDeleteTeamMemberParamsWithContext creates a new DeleteTeamMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTeamMemberParamsWithContext(ctx context.Context) *DeleteTeamMemberParams {
	var ()
	return &DeleteTeamMemberParams{

		Context: ctx,
	}
}

// NewDeleteTeamMemberParamsWithHTTPClient creates a new DeleteTeamMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTeamMemberParamsWithHTTPClient(client *http.Client) *DeleteTeamMemberParams {
	var ()
	return &DeleteTeamMemberParams{
		HTTPClient: client,
	}
}

/*DeleteTeamMemberParams contains all the parameters to send to the API endpoint
for the delete team member operation typically these are written to a http.Request
*/
type DeleteTeamMemberParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Team
	  Team under namesapce

	*/
	Team string
	/*User
	  Member under team

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete team member params
func (o *DeleteTeamMemberParams) WithTimeout(timeout time.Duration) *DeleteTeamMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete team member params
func (o *DeleteTeamMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete team member params
func (o *DeleteTeamMemberParams) WithContext(ctx context.Context) *DeleteTeamMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete team member params
func (o *DeleteTeamMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete team member params
func (o *DeleteTeamMemberParams) WithHTTPClient(client *http.Client) *DeleteTeamMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete team member params
func (o *DeleteTeamMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the delete team member params
func (o *DeleteTeamMemberParams) WithOwner(owner string) *DeleteTeamMemberParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete team member params
func (o *DeleteTeamMemberParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithTeam adds the team to the delete team member params
func (o *DeleteTeamMemberParams) WithTeam(team string) *DeleteTeamMemberParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the delete team member params
func (o *DeleteTeamMemberParams) SetTeam(team string) {
	o.Team = team
}

// WithUser adds the user to the delete team member params
func (o *DeleteTeamMemberParams) WithUser(user string) *DeleteTeamMemberParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the delete team member params
func (o *DeleteTeamMemberParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
