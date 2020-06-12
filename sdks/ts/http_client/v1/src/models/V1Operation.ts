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
    V1Cache,
    V1CacheFromJSON,
    V1CacheFromJSONTyped,
    V1CacheToJSON,
    V1Component,
    V1ComponentFromJSON,
    V1ComponentFromJSONTyped,
    V1ComponentToJSON,
    V1Param,
    V1ParamFromJSON,
    V1ParamFromJSONTyped,
    V1ParamToJSON,
    V1Plugins,
    V1PluginsFromJSON,
    V1PluginsFromJSONTyped,
    V1PluginsToJSON,
    V1Termination,
    V1TerminationFromJSON,
    V1TerminationFromJSONTyped,
    V1TerminationToJSON,
    V1TriggerPolicy,
    V1TriggerPolicyFromJSON,
    V1TriggerPolicyFromJSONTyped,
    V1TriggerPolicyToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1Operation
 */
export interface V1Operation {
    /**
     * 
     * @type {number}
     * @memberof V1Operation
     */
    version?: number;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    kind?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    description?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Operation
     */
    tags?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    profile?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    queue?: string;
    /**
     * 
     * @type {V1Cache}
     * @memberof V1Operation
     */
    cache?: V1Cache;
    /**
     * 
     * @type {object}
     * @memberof V1Operation
     */
    schedule?: object;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1Operation
     */
    events?: Array<object>;
    /**
     * 
     * @type {object}
     * @memberof V1Operation
     */
    matrix?: object;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Operation
     */
    dependencies?: Array<string>;
    /**
     * 
     * @type {V1TriggerPolicy}
     * @memberof V1Operation
     */
    trigger?: V1TriggerPolicy;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Operation
     */
    conditions?: Array<string>;
    /**
     * 
     * @type {boolean}
     * @memberof V1Operation
     */
    skip_on_upstream_skip?: boolean;
    /**
     * 
     * @type {V1Termination}
     * @memberof V1Operation
     */
    termination?: V1Termination;
    /**
     * 
     * @type {V1Plugins}
     * @memberof V1Operation
     */
    plugins?: V1Plugins;
    /**
     * 
     * @type {{ [key: string]: V1Param; }}
     * @memberof V1Operation
     */
    params?: { [key: string]: V1Param; };
    /**
     * 
     * @type {object}
     * @memberof V1Operation
     */
    run_patch?: object;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    dag_ref?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    url_ref?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    path_ref?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Operation
     */
    hub_ref?: string;
    /**
     * 
     * @type {V1Component}
     * @memberof V1Operation
     */
    component?: V1Component;
}

export function V1OperationFromJSON(json: any): V1Operation {
    return V1OperationFromJSONTyped(json, false);
}

export function V1OperationFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Operation {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'version': !exists(json, 'version') ? undefined : json['version'],
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'tags': !exists(json, 'tags') ? undefined : json['tags'],
        'profile': !exists(json, 'profile') ? undefined : json['profile'],
        'queue': !exists(json, 'queue') ? undefined : json['queue'],
        'cache': !exists(json, 'cache') ? undefined : V1CacheFromJSON(json['cache']),
        'schedule': !exists(json, 'schedule') ? undefined : json['schedule'],
        'events': !exists(json, 'events') ? undefined : json['events'],
        'matrix': !exists(json, 'matrix') ? undefined : json['matrix'],
        'dependencies': !exists(json, 'dependencies') ? undefined : json['dependencies'],
        'trigger': !exists(json, 'trigger') ? undefined : V1TriggerPolicyFromJSON(json['trigger']),
        'conditions': !exists(json, 'conditions') ? undefined : json['conditions'],
        'skip_on_upstream_skip': !exists(json, 'skip_on_upstream_skip') ? undefined : json['skip_on_upstream_skip'],
        'termination': !exists(json, 'termination') ? undefined : V1TerminationFromJSON(json['termination']),
        'plugins': !exists(json, 'plugins') ? undefined : V1PluginsFromJSON(json['plugins']),
        'params': !exists(json, 'params') ? undefined : (mapValues(json['params'], V1ParamFromJSON)),
        'run_patch': !exists(json, 'run_patch') ? undefined : json['run_patch'],
        'dag_ref': !exists(json, 'dag_ref') ? undefined : json['dag_ref'],
        'url_ref': !exists(json, 'url_ref') ? undefined : json['url_ref'],
        'path_ref': !exists(json, 'path_ref') ? undefined : json['path_ref'],
        'hub_ref': !exists(json, 'hub_ref') ? undefined : json['hub_ref'],
        'component': !exists(json, 'component') ? undefined : V1ComponentFromJSON(json['component']),
    };
}

export function V1OperationToJSON(value?: V1Operation | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'version': value.version,
        'kind': value.kind,
        'name': value.name,
        'description': value.description,
        'tags': value.tags,
        'profile': value.profile,
        'queue': value.queue,
        'cache': V1CacheToJSON(value.cache),
        'schedule': value.schedule,
        'events': value.events,
        'matrix': value.matrix,
        'dependencies': value.dependencies,
        'trigger': V1TriggerPolicyToJSON(value.trigger),
        'conditions': value.conditions,
        'skip_on_upstream_skip': value.skip_on_upstream_skip,
        'termination': V1TerminationToJSON(value.termination),
        'plugins': V1PluginsToJSON(value.plugins),
        'params': value.params === undefined ? undefined : (mapValues(value.params, V1ParamToJSON)),
        'run_patch': value.run_patch,
        'dag_ref': value.dag_ref,
        'url_ref': value.url_ref,
        'path_ref': value.path_ref,
        'hub_ref': value.hub_ref,
        'component': V1ComponentToJSON(value.component),
    };
}


