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

import { exists, mapValues } from '../runtime';
import {
    V1StatusCondition,
    V1StatusConditionFromJSON,
    V1StatusConditionFromJSONTyped,
    V1StatusConditionToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1EntityStatusBodyRequest
 */
export interface V1EntityStatusBodyRequest {
    /**
     * 
     * @type {string}
     * @memberof V1EntityStatusBodyRequest
     */
    owner?: string;
    /**
     * 
     * @type {string}
     * @memberof V1EntityStatusBodyRequest
     */
    project?: string;
    /**
     * 
     * @type {string}
     * @memberof V1EntityStatusBodyRequest
     */
    uuid?: string;
    /**
     * 
     * @type {V1StatusCondition}
     * @memberof V1EntityStatusBodyRequest
     */
    condition?: V1StatusCondition;
}

export function V1EntityStatusBodyRequestFromJSON(json: any): V1EntityStatusBodyRequest {
    return V1EntityStatusBodyRequestFromJSONTyped(json, false);
}

export function V1EntityStatusBodyRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1EntityStatusBodyRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'owner': !exists(json, 'owner') ? undefined : json['owner'],
        'project': !exists(json, 'project') ? undefined : json['project'],
        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'condition': !exists(json, 'condition') ? undefined : V1StatusConditionFromJSON(json['condition']),
    };
}

export function V1EntityStatusBodyRequestToJSON(value?: V1EntityStatusBodyRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'owner': value.owner,
        'project': value.project,
        'uuid': value.uuid,
        'condition': V1StatusConditionToJSON(value.condition),
    };
}


