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
    V1Agent,
    V1AgentFromJSON,
    V1AgentToJSON,
    V1ListQueuesResponse,
    V1ListQueuesResponseFromJSON,
    V1ListQueuesResponseToJSON,
    V1Queue,
    V1QueueFromJSON,
    V1QueueToJSON,
} from '../models';

export interface CreateQueueRequest {
    owner: string;
    agent: string;
    body: V1Queue;
}

export interface DeleteQueueRequest {
    owner: string;
    agent: string;
    uuid: string;
}

export interface GetQueueRequest {
    owner: string;
    agent: string;
    uuid: string;
}

export interface ListOrganizationQueueNamesRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
}

export interface ListOrganizationQueuesRequest {
    owner: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
}

export interface ListQueueNamesRequest {
    owner: string;
    agent: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
}

export interface ListQueuesRequest {
    owner: string;
    agent: string;
    offset?: number;
    limit?: number;
    sort?: string;
    query?: string;
}

export interface PatchQueueRequest {
    owner: string;
    queueAgent: string;
    queueUuid: string;
    body: V1Queue;
}

export interface UpdateQueueRequest {
    owner: string;
    queueAgent: string;
    queueUuid: string;
    body: V1Queue;
}

/**
 * Polyaxon sdk
 */
export class QueuesV1Api extends runtime.BaseAPI {

    /**
     * Update agent
     */
    async createQueueRaw(requestParameters: CreateQueueRequest): Promise<runtime.ApiResponse<V1Agent>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling createQueue.');
        }

        if (requestParameters.agent === null || requestParameters.agent === undefined) {
            throw new runtime.RequiredError('agent','Required parameter requestParameters.agent was null or undefined when calling createQueue.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createQueue.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/agents/{agent}/queues`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"agent"}}`, encodeURIComponent(String(requestParameters.agent))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: V1QueueToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1AgentFromJSON(jsonValue));
    }

    /**
     * Update agent
     */
    async createQueue(requestParameters: CreateQueueRequest): Promise<V1Agent> {
        const response = await this.createQueueRaw(requestParameters);
        return await response.value();
    }

    /**
     * Sync agent
     */
    async deleteQueueRaw(requestParameters: DeleteQueueRequest): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling deleteQueue.');
        }

        if (requestParameters.agent === null || requestParameters.agent === undefined) {
            throw new runtime.RequiredError('agent','Required parameter requestParameters.agent was null or undefined when calling deleteQueue.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling deleteQueue.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/agents/{agent}/queues/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"agent"}}`, encodeURIComponent(String(requestParameters.agent))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Sync agent
     */
    async deleteQueue(requestParameters: DeleteQueueRequest): Promise<void> {
        await this.deleteQueueRaw(requestParameters);
    }

    /**
     * Patch agent
     */
    async getQueueRaw(requestParameters: GetQueueRequest): Promise<runtime.ApiResponse<V1Queue>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling getQueue.');
        }

        if (requestParameters.agent === null || requestParameters.agent === undefined) {
            throw new runtime.RequiredError('agent','Required parameter requestParameters.agent was null or undefined when calling getQueue.');
        }

        if (requestParameters.uuid === null || requestParameters.uuid === undefined) {
            throw new runtime.RequiredError('uuid','Required parameter requestParameters.uuid was null or undefined when calling getQueue.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/agents/{agent}/queues/{uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"agent"}}`, encodeURIComponent(String(requestParameters.agent))).replace(`{${"uuid"}}`, encodeURIComponent(String(requestParameters.uuid))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1QueueFromJSON(jsonValue));
    }

    /**
     * Patch agent
     */
    async getQueue(requestParameters: GetQueueRequest): Promise<V1Queue> {
        const response = await this.getQueueRaw(requestParameters);
        return await response.value();
    }

    /**
     * List agents names
     */
    async listOrganizationQueueNamesRaw(requestParameters: ListOrganizationQueueNamesRequest): Promise<runtime.ApiResponse<V1ListQueuesResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listOrganizationQueueNames.');
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
            path: `/api/v1/orgs/{owner}/queues/names`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListQueuesResponseFromJSON(jsonValue));
    }

    /**
     * List agents names
     */
    async listOrganizationQueueNames(requestParameters: ListOrganizationQueueNamesRequest): Promise<V1ListQueuesResponse> {
        const response = await this.listOrganizationQueueNamesRaw(requestParameters);
        return await response.value();
    }

    /**
     * List agents
     */
    async listOrganizationQueuesRaw(requestParameters: ListOrganizationQueuesRequest): Promise<runtime.ApiResponse<V1ListQueuesResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listOrganizationQueues.');
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
            path: `/api/v1/orgs/{owner}/queues`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListQueuesResponseFromJSON(jsonValue));
    }

    /**
     * List agents
     */
    async listOrganizationQueues(requestParameters: ListOrganizationQueuesRequest): Promise<V1ListQueuesResponse> {
        const response = await this.listOrganizationQueuesRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create agent
     */
    async listQueueNamesRaw(requestParameters: ListQueueNamesRequest): Promise<runtime.ApiResponse<V1ListQueuesResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listQueueNames.');
        }

        if (requestParameters.agent === null || requestParameters.agent === undefined) {
            throw new runtime.RequiredError('agent','Required parameter requestParameters.agent was null or undefined when calling listQueueNames.');
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
            path: `/api/v1/orgs/{owner}/agents/{agent}/queues/names`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"agent"}}`, encodeURIComponent(String(requestParameters.agent))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListQueuesResponseFromJSON(jsonValue));
    }

    /**
     * Create agent
     */
    async listQueueNames(requestParameters: ListQueueNamesRequest): Promise<V1ListQueuesResponse> {
        const response = await this.listQueueNamesRaw(requestParameters);
        return await response.value();
    }

    /**
     * Get agent
     */
    async listQueuesRaw(requestParameters: ListQueuesRequest): Promise<runtime.ApiResponse<V1ListQueuesResponse>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling listQueues.');
        }

        if (requestParameters.agent === null || requestParameters.agent === undefined) {
            throw new runtime.RequiredError('agent','Required parameter requestParameters.agent was null or undefined when calling listQueues.');
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
            path: `/api/v1/orgs/{owner}/agents/{agent}/queues`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"agent"}}`, encodeURIComponent(String(requestParameters.agent))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1ListQueuesResponseFromJSON(jsonValue));
    }

    /**
     * Get agent
     */
    async listQueues(requestParameters: ListQueuesRequest): Promise<V1ListQueuesResponse> {
        const response = await this.listQueuesRaw(requestParameters);
        return await response.value();
    }

    /**
     * Get State (queues/runs)
     */
    async patchQueueRaw(requestParameters: PatchQueueRequest): Promise<runtime.ApiResponse<V1Queue>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling patchQueue.');
        }

        if (requestParameters.queueAgent === null || requestParameters.queueAgent === undefined) {
            throw new runtime.RequiredError('queueAgent','Required parameter requestParameters.queueAgent was null or undefined when calling patchQueue.');
        }

        if (requestParameters.queueUuid === null || requestParameters.queueUuid === undefined) {
            throw new runtime.RequiredError('queueUuid','Required parameter requestParameters.queueUuid was null or undefined when calling patchQueue.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling patchQueue.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/agents/{queue.agent}/queues/{queue.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"queue.agent"}}`, encodeURIComponent(String(requestParameters.queueAgent))).replace(`{${"queue.uuid"}}`, encodeURIComponent(String(requestParameters.queueUuid))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: V1QueueToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1QueueFromJSON(jsonValue));
    }

    /**
     * Get State (queues/runs)
     */
    async patchQueue(requestParameters: PatchQueueRequest): Promise<V1Queue> {
        const response = await this.patchQueueRaw(requestParameters);
        return await response.value();
    }

    /**
     * Delete agent
     */
    async updateQueueRaw(requestParameters: UpdateQueueRequest): Promise<runtime.ApiResponse<V1Queue>> {
        if (requestParameters.owner === null || requestParameters.owner === undefined) {
            throw new runtime.RequiredError('owner','Required parameter requestParameters.owner was null or undefined when calling updateQueue.');
        }

        if (requestParameters.queueAgent === null || requestParameters.queueAgent === undefined) {
            throw new runtime.RequiredError('queueAgent','Required parameter requestParameters.queueAgent was null or undefined when calling updateQueue.');
        }

        if (requestParameters.queueUuid === null || requestParameters.queueUuid === undefined) {
            throw new runtime.RequiredError('queueUuid','Required parameter requestParameters.queueUuid was null or undefined when calling updateQueue.');
        }

        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling updateQueue.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // ApiKey authentication
        }

        const response = await this.request({
            path: `/api/v1/orgs/{owner}/agents/{queue.agent}/queues/{queue.uuid}`.replace(`{${"owner"}}`, encodeURIComponent(String(requestParameters.owner))).replace(`{${"queue.agent"}}`, encodeURIComponent(String(requestParameters.queueAgent))).replace(`{${"queue.uuid"}}`, encodeURIComponent(String(requestParameters.queueUuid))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: V1QueueToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => V1QueueFromJSON(jsonValue));
    }

    /**
     * Delete agent
     */
    async updateQueue(requestParameters: UpdateQueueRequest): Promise<V1Queue> {
        const response = await this.updateQueueRaw(requestParameters);
        return await response.value();
    }

}
