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
/**
 * 
 * @export
 * @interface V1ArtifactsType
 */
export interface V1ArtifactsType {
    /**
     * 
     * @type {string}
     * @memberof V1ArtifactsType
     */
    connection?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ArtifactsType
     */
    files?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ArtifactsType
     */
    dirs?: Array<string>;
    /**
     * 
     * @type {boolean}
     * @memberof V1ArtifactsType
     */
    init?: boolean;
    /**
     * 
     * @type {number}
     * @memberof V1ArtifactsType
     */
    workers?: number;
}

export function V1ArtifactsTypeFromJSON(json: any): V1ArtifactsType {
    return V1ArtifactsTypeFromJSONTyped(json, false);
}

export function V1ArtifactsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ArtifactsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'connection': !exists(json, 'connection') ? undefined : json['connection'],
        'files': !exists(json, 'files') ? undefined : json['files'],
        'dirs': !exists(json, 'dirs') ? undefined : json['dirs'],
        'init': !exists(json, 'init') ? undefined : json['init'],
        'workers': !exists(json, 'workers') ? undefined : json['workers'],
    };
}

export function V1ArtifactsTypeToJSON(value?: V1ArtifactsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'connection': value.connection,
        'files': value.files,
        'dirs': value.dirs,
        'init': value.init,
        'workers': value.workers,
    };
}


