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

import { exists, mapValues } from '../runtime';
import {
    V1Environment,
    V1EnvironmentFromJSON,
    V1EnvironmentFromJSONTyped,
    V1EnvironmentToJSON,
    V1Init,
    V1InitFromJSON,
    V1InitFromJSONTyped,
    V1InitToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1Dask
 */
export interface V1Dask {
    /**
     * 
     * @type {string}
     * @memberof V1Dask
     */
    kind?: string;
    /**
     * 
     * @type {number}
     * @memberof V1Dask
     */
    scale?: number;
    /**
     * 
     * @type {number}
     * @memberof V1Dask
     */
    adapt_min?: number;
    /**
     * 
     * @type {number}
     * @memberof V1Dask
     */
    adapt_max?: number;
    /**
     * 
     * @type {string}
     * @memberof V1Dask
     */
    adapt_interval?: string;
    /**
     * 
     * @type {V1Environment}
     * @memberof V1Dask
     */
    environment?: V1Environment;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Dask
     */
    connections?: Array<string>;
    /**
     * Volumes is a list of volumes that can be mounted.
     * @type {Array<object>}
     * @memberof V1Dask
     */
    volumes?: Array<object>;
    /**
     * 
     * @type {Array<V1Init>}
     * @memberof V1Dask
     */
    init?: Array<V1Init>;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1Dask
     */
    sidecars?: Array<object>;
    /**
     * 
     * @type {object}
     * @memberof V1Dask
     */
    container?: object;
}

export function V1DaskFromJSON(json: any): V1Dask {
    return V1DaskFromJSONTyped(json, false);
}

export function V1DaskFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Dask {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'scale': !exists(json, 'scale') ? undefined : json['scale'],
        'adapt_min': !exists(json, 'adapt_min') ? undefined : json['adapt_min'],
        'adapt_max': !exists(json, 'adapt_max') ? undefined : json['adapt_max'],
        'adapt_interval': !exists(json, 'adapt_interval') ? undefined : json['adapt_interval'],
        'environment': !exists(json, 'environment') ? undefined : V1EnvironmentFromJSON(json['environment']),
        'connections': !exists(json, 'connections') ? undefined : json['connections'],
        'volumes': !exists(json, 'volumes') ? undefined : json['volumes'],
        'init': !exists(json, 'init') ? undefined : ((json['init'] as Array<any>).map(V1InitFromJSON)),
        'sidecars': !exists(json, 'sidecars') ? undefined : json['sidecars'],
        'container': !exists(json, 'container') ? undefined : json['container'],
    };
}

export function V1DaskToJSON(value?: V1Dask | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': value.kind,
        'scale': value.scale,
        'adapt_min': value.adapt_min,
        'adapt_max': value.adapt_max,
        'adapt_interval': value.adapt_interval,
        'environment': V1EnvironmentToJSON(value.environment),
        'connections': value.connections,
        'volumes': value.volumes,
        'init': value.init === undefined ? undefined : ((value.init as Array<any>).map(V1InitToJSON)),
        'sidecars': value.sidecars,
        'container': value.container,
    };
}


