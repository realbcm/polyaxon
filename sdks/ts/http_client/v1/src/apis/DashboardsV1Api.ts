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
 * The version of the OpenAPI document: 1.0.84
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
    V1Dashboard,
    V1DashboardFromJSON,
    V1DashboardToJSON,
    V1ListDashboardsResponse,
    V1ListDashboardsResponseFromJSON,
    V1ListDashboardsResponseToJSON,
} from '../models';

export interface CreateDashboardRequest {
    owner: string;
    body: V1Dashboard;
}

export interface DeleteDashboardRequest {
    owner: string;
    uuid: string;
}

export interface GetDashboardRequest {
    owner: string;
    uuid: string;
}

export interface ListDashboardNamesRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
}

export interface ListDashboardsRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
}

export interface PatchDashboardRequest {
    owner: string;
    dashboardUuid: string;
    body: V1Dashboard;
}

export interface UpdateDashboardRequest {
    owner: string;
    dashboardUuid: string;
    body: V1Dashboard;
}

/**
 * Polyaxon sdk
 */
export class DashboardsV1Api extends runtime.BaseAPI {

    /**
     */
    async createDashboardRaw(requestParameters: CreateDashboardRequest): Promise<runtime.ApiResponse<V1Dashboard>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createDashboard.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createDashboard.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1DashboardToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1DashboardFromJSON(jsonValue));
    }

    /**
     */
    async createDashboard(requestParameters: CreateDashboardRequest): Promise<V1Dashboard> {
        const response = await this.createDashboardRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async deleteDashboardRaw(requestParameters: DeleteDashboardRequest): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteDashboard.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling deleteDashboard.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteDashboard(requestParameters: DeleteDashboardRequest): Promise<void> {
        await this.deleteDashboardRaw(requestParameters);
    }

    /**
     */
    async getDashboardRaw(requestParameters: GetDashboardRequest): Promise<runtime.ApiResponse<V1Dashboard>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getDashboard.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling getDashboard.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1DashboardFromJSON(jsonValue));
    }

    /**
     */
    async getDashboard(requestParameters: GetDashboardRequest): Promise<V1Dashboard> {
        const response = await this.getDashboardRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listDashboardNamesRaw(requestParameters: ListDashboardNamesRequest): Promise<runtime.ApiResponse<V1ListDashboardsResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listDashboardNames.');
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
            path: `/api/v1/orgs/{owner}/dashboards/names`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListDashboardsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listDashboardNames(requestParameters: ListDashboardNamesRequest): Promise<V1ListDashboardsResponse> {
        const response = await this.listDashboardNamesRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async listDashboardsRaw(requestParameters: ListDashboardsRequest): Promise<runtime.ApiResponse<V1ListDashboardsResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listDashboards.');
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
            path: `/api/v1/orgs/{owner}/dashboards`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListDashboardsResponseFromJSON(jsonValue));
    }

    /**
     */
    async listDashboards(requestParameters: ListDashboardsRequest): Promise<V1ListDashboardsResponse> {
        const response = await this.listDashboardsRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async patchDashboardRaw(requestParameters: PatchDashboardRequest): Promise<runtime.ApiResponse<V1Dashboard>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchDashboard.');
        }

        if (requestParameters.dashboardUuid === null || requestParameters.dashboardUuid === undefined) {
            throw new runtime.RequiredError('dashboardUuid','Required parameter requestParameters.dashboardUuid was null or undefined when calling patchDashboard.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchDashboard.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards/{dashboard.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"dashboard.uuid"}}`, encodeURIComponent(String(requestParameters.dashboardUuid))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1DashboardToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1DashboardFromJSON(jsonValue));
    }

    /**
     */
    async patchDashboard(requestParameters: PatchDashboardRequest): Promise<V1Dashboard> {
        const response = await this.patchDashboardRaw(requestParameters);
        return await response.value();
    }

    /**
     */
    async updateDashboardRaw(requestParameters: UpdateDashboardRequest): Promise<runtime.ApiResponse<V1Dashboard>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateDashboard.');
        }

        if (requestParameters.dashboardUuid === null || requestParameters.dashboardUuid === undefined) {
            throw new runtime.RequiredError('dashboardUuid','Required parameter requestParameters.dashboardUuid was null or undefined when calling updateDashboard.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateDashboard.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/dashboards/{dashboard.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"dashboard.uuid"}}`, encodeURIComponent(String(requestParameters.dashboardUuid))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1DashboardToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1DashboardFromJSON(jsonValue));
    }

    /**
     */
    async updateDashboard(requestParameters: UpdateDashboardRequest): Promise<V1Dashboard> {
        const response = await this.updateDashboardRaw(requestParameters);
        return await response.value();
    }

}
