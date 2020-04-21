// Copyright 2018-2020 Polyaxon, Inc.
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

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.0.81
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    RuntimeError,
    RuntimeErrorFromJSON,
    RuntimeErrorToJSON,
    V1ListOrganizationMembersResponse,
    V1ListOrganizationMembersResponseFromJSON,
    V1ListOrganizationMembersResponseToJSON,
    V1ListOrganizationsResponse,
    V1ListOrganizationsResponseFromJSON,
    V1ListOrganizationsResponseToJSON,
    V1Organization,
    V1OrganizationFromJSON,
    V1OrganizationToJSON,
    V1OrganizationMember,
    V1OrganizationMemberFromJSON,
    V1OrganizationMemberToJSON,
} from '../models';

export interface CreateOrganizationRequest {
    body: V1Organization;
}

export interface CreateOrganizationMemberRequest {
    owner: string;
    body: V1OrganizationMember;
}

export interface DeleteOrganizationRequest {
    owner: string;
}

export interface DeleteOrganizationMemberRequest {
    owner: string;
    user: string;
}

export interface GetOrganizationRequest {
    owner: string;
}

export interface GetOrganizationMemberRequest {
    owner: string;
    user: string;
}

export interface ListOrganizationMembersRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
}

export interface PatchOrganizationRequest {
    owner: string;
    body: V1Organization;
}

export interface PatchOrganizationMemberRequest {
    owner: string;
    memberUser: string;
    body: V1OrganizationMember;
}

export interface UpdateOrganizationRequest {
    owner: string;
    body: V1Organization;
}

export interface UpdateOrganizationMemberRequest {
    owner: string;
    memberUser: string;
    body: V1OrganizationMember;
}

/**
 * Polyaxon's typescript client
 */
export class OrganizationsV1Api extends runtime.BaseAPI {

    /**
     */
    async createOrganizationRaw(requestParameters: CreateOrganizationRequest): Promise<runtime.ApiResponse<V1Organization>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createOrganization.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/create`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1OrganizationToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1OrganizationFromJSON(jsonValue));
    }

    /**
     */
    async createOrganization(requestParameters: CreateOrganizationRequest): Promise<V1Organization> {
        const response = await this.createOrganizationRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async createOrganizationMemberRaw(requestParameters: CreateOrganizationMemberRequest): Promise<runtime.ApiResponse<V1OrganizationMember>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createOrganizationMember.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createOrganizationMember.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/members`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1OrganizationMemberToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1OrganizationMemberFromJSON(jsonValue));
    }

    /**
     */
    async createOrganizationMember(requestParameters: CreateOrganizationMemberRequest): Promise<V1OrganizationMember> {
        const response = await this.createOrganizationMemberRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async deleteOrganizationRaw(requestParameters: DeleteOrganizationRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteOrganization.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async deleteOrganization(requestParameters: DeleteOrganizationRequest): Promise<object> {
        const response = await this.deleteOrganizationRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async deleteOrganizationMemberRaw(requestParameters: DeleteOrganizationMemberRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteOrganizationMember.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling deleteOrganizationMember.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/members/{user}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"user"}}`, encodeURIComponent(String(requestParameters.user))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     */
    async deleteOrganizationMember(requestParameters: DeleteOrganizationMemberRequest): Promise<object> {
        const response = await this.deleteOrganizationMemberRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getOrganizationRaw(requestParameters: GetOrganizationRequest): Promise<runtime.ApiResponse<V1Organization>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getOrganization.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1OrganizationFromJSON(jsonValue));
    }

    /**
     */
    async getOrganization(requestParameters: GetOrganizationRequest): Promise<V1Organization> {
        const response = await this.getOrganizationRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async getOrganizationMemberRaw(requestParameters: GetOrganizationMemberRequest): Promise<runtime.ApiResponse<V1OrganizationMember>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getOrganizationMember.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling getOrganizationMember.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/members/{user}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"user"}}`, encodeURIComponent(String(requestParameters.user))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1OrganizationMemberFromJSON(jsonValue));
    }

    /**
     */
    async getOrganizationMember(requestParameters: GetOrganizationMemberRequest): Promise<V1OrganizationMember> {
        const response = await this.getOrganizationMemberRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listOrganizationMembersRaw(requestParameters: ListOrganizationMembersRequest): Promise<runtime.ApiResponse<V1ListOrganizationMembersResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listOrganizationMembers.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.sort !== undefined) {
            queryParameters['sort'] = requestParameters.sort;
        }

        if (requestParameters.query !== undefined) {
            queryParameters['query'] = requestParameters.query;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/members`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListOrganizationMembersResponseFromJSON(jsonValue));
    }

    /**
     */
    async listOrganizationMembers(requestParameters: ListOrganizationMembersRequest): Promise<V1ListOrganizationMembersResponse> {
        const response = await this.listOrganizationMembersRaw(requestParameters);
        return await response.value();
    }

    /**
     * Get versions
     */
    async listOrganizationNamesRaw(): Promise<runtime.ApiResponse<V1ListOrganizationsResponse>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/names`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListOrganizationsResponseFromJSON(jsonValue));
    }

    /**
     * Get versions
     */
    async listOrganizationNames(): Promise<V1ListOrganizationsResponse> {
        const response = await this.listOrganizationNamesRaw();
        return await response.value();
    }

    /**
     * Get log handler
     */
    async listOrganizationsRaw(): Promise<runtime.ApiResponse<V1ListOrganizationsResponse>> {
        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/list`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListOrganizationsResponseFromJSON(jsonValue));
    }

    /**
     * Get log handler
     */
    async listOrganizations(): Promise<V1ListOrganizationsResponse> {
        const response = await this.listOrganizationsRaw();
        return await response.value();
    }

    /**
     */
    async patchOrganizationRaw(requestParameters: PatchOrganizationRequest): Promise<runtime.ApiResponse<V1Organization>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchOrganization.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchOrganization.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1OrganizationToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1OrganizationFromJSON(jsonValue));
    }

    /**
     */
    async patchOrganization(requestParameters: PatchOrganizationRequest): Promise<V1Organization> {
        const response = await this.patchOrganizationRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async patchOrganizationMemberRaw(requestParameters: PatchOrganizationMemberRequest): Promise<runtime.ApiResponse<V1OrganizationMember>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchOrganizationMember.');
        }

        if (requestParameters.memberUser === null || requestParameters.memberUser === undefined) {
            throw new runtime.RequiredError('memberUser','Required parameter requestParameters.memberUser was null or undefined when calling patchOrganizationMember.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchOrganizationMember.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/members/{member.user}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"member.user"}}`, encodeURIComponent(String(requestParameters.memberUser))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1OrganizationMemberToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1OrganizationMemberFromJSON(jsonValue));
    }

    /**
     */
    async patchOrganizationMember(requestParameters: PatchOrganizationMemberRequest): Promise<V1OrganizationMember> {
        const response = await this.patchOrganizationMemberRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async updateOrganizationRaw(requestParameters: UpdateOrganizationRequest): Promise<runtime.ApiResponse<V1Organization>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateOrganization.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateOrganization.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1OrganizationToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1OrganizationFromJSON(jsonValue));
    }

    /**
     */
    async updateOrganization(requestParameters: UpdateOrganizationRequest): Promise<V1Organization> {
        const response = await this.updateOrganizationRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async updateOrganizationMemberRaw(requestParameters: UpdateOrganizationMemberRequest): Promise<runtime.ApiResponse<V1OrganizationMember>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateOrganizationMember.');
        }

        if (requestParameters.memberUser === null || requestParameters.memberUser === undefined) {
            throw new runtime.RequiredError('memberUser','Required parameter requestParameters.memberUser was null or undefined when calling updateOrganizationMember.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateOrganizationMember.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/members/{member.user}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"member.user"}}`, encodeURIComponent(String(requestParameters.memberUser))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1OrganizationMemberToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1OrganizationMemberFromJSON(jsonValue));
    }

    /**
     */
    async updateOrganizationMember(requestParameters: UpdateOrganizationMemberRequest): Promise<V1OrganizationMember> {
        const response = await this.updateOrganizationMemberRaw(requestParameters);
        return await response.value();
    }

}
