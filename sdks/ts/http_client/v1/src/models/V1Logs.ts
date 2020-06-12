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
 * The version of the OpenAPI document: 1.0.97
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    V1Log,
    V1LogFromJSON,
    V1LogFromJSONTyped,
    V1LogToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1Logs
 */
export interface V1Logs {
    /**
     * 
     * @type {Array<V1Log>}
     * @memberof V1Logs
     */
    logs?: Array<V1Log>;
    /**
     * 
     * @type {Date}
     * @memberof V1Logs
     */
    last_time?: Date;
    /**
     * 
     * @type {string}
     * @memberof V1Logs
     */
    last_file?: string;
}

export function V1LogsFromJSON(json: any): V1Logs {
    return V1LogsFromJSONTyped(json, false);
}

export function V1LogsFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Logs {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'logs': !exists(json, 'logs') ? undefined : ((json['logs'] as Array<any>).map(V1LogFromJSON)),
        'last_time': !exists(json, 'last_time') ? undefined : (new Date(json['last_time'])),
        'last_file': !exists(json, 'last_file') ? undefined : json['last_file'],
    };
}

export function V1LogsToJSON(value?: V1Logs | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'logs': value.logs === undefined ? undefined : ((value.logs as Array<any>).map(V1LogToJSON)),
        'last_time': value.last_time === undefined ? undefined : (value.last_time.toISOString()),
        'last_file': value.last_file,
    };
}


