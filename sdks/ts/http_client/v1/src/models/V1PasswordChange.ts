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
 * The version of the OpenAPI document: 1.5.0
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
 * @interface V1PasswordChange
 */
export interface V1PasswordChange {
    /**
     * 
     * @type {string}
     * @memberof V1PasswordChange
     */
    old_password?: string;
    /**
     * 
     * @type {string}
     * @memberof V1PasswordChange
     */
    new_password1?: string;
    /**
     * 
     * @type {string}
     * @memberof V1PasswordChange
     */
    new_password2?: string;
}

export function V1PasswordChangeFromJSON(json: any): V1PasswordChange {
    return V1PasswordChangeFromJSONTyped(json, false);
}

export function V1PasswordChangeFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1PasswordChange {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'old_password': !exists(json, 'old_password') ? undefined : json['old_password'],
        'new_password1': !exists(json, 'new_password1') ? undefined : json['new_password1'],
        'new_password2': !exists(json, 'new_password2') ? undefined : json['new_password2'],
    };
}

export function V1PasswordChangeToJSON(value?: V1PasswordChange | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'old_password': value.old_password,
        'new_password1': value.new_password1,
        'new_password2': value.new_password2,
    };
}


