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

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.5.1
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
 * @interface V1ModelRegistrySettings
 */
export interface V1ModelRegistrySettings {
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ModelRegistrySettings
     */
    teams?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1ModelRegistrySettings
     */
    projects?: Array<string>;
}

export function V1ModelRegistrySettingsFromJSON(json: any): V1ModelRegistrySettings {
    return V1ModelRegistrySettingsFromJSONTyped(json, false);
}

export function V1ModelRegistrySettingsFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ModelRegistrySettings {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'teams': !exists(json, 'teams') ? undefined : json['teams'],
        'projects': !exists(json, 'projects') ? undefined : json['projects'],
    };
}

export function V1ModelRegistrySettingsToJSON(value?: V1ModelRegistrySettings | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'teams': value.teams,
        'projects': value.projects,
    };
}


