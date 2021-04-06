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

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.7.6
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1Cache from './V1Cache';
import V1Hook from './V1Hook';
import V1IO from './V1IO';
import V1Plugins from './V1Plugins';
import V1Template from './V1Template';
import V1Termination from './V1Termination';

/**
 * The V1Component model module.
 * @module model/V1Component
 * @version 1.7.6
 */
class V1Component {
    /**
     * Constructs a new <code>V1Component</code>.
     * @alias module:model/V1Component
     */
    constructor() { 
        
        V1Component.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Component</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Component} obj Optional instance to populate.
     * @return {module:model/V1Component} The populated <code>V1Component</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Component();

            if (data.hasOwnProperty('version')) {
                obj['version'] = ApiClient.convertToType(data['version'], 'Number');
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], ['String']);
            }
            if (data.hasOwnProperty('presets')) {
                obj['presets'] = ApiClient.convertToType(data['presets'], ['String']);
            }
            if (data.hasOwnProperty('queue')) {
                obj['queue'] = ApiClient.convertToType(data['queue'], 'String');
            }
            if (data.hasOwnProperty('cache')) {
                obj['cache'] = V1Cache.constructFromObject(data['cache']);
            }
            if (data.hasOwnProperty('termination')) {
                obj['termination'] = V1Termination.constructFromObject(data['termination']);
            }
            if (data.hasOwnProperty('plugins')) {
                obj['plugins'] = V1Plugins.constructFromObject(data['plugins']);
            }
            if (data.hasOwnProperty('hooks')) {
                obj['hooks'] = ApiClient.convertToType(data['hooks'], [V1Hook]);
            }
            if (data.hasOwnProperty('inputs')) {
                obj['inputs'] = ApiClient.convertToType(data['inputs'], [V1IO]);
            }
            if (data.hasOwnProperty('outputs')) {
                obj['outputs'] = ApiClient.convertToType(data['outputs'], [V1IO]);
            }
            if (data.hasOwnProperty('run')) {
                obj['run'] = ApiClient.convertToType(data['run'], Object);
            }
            if (data.hasOwnProperty('template')) {
                obj['template'] = V1Template.constructFromObject(data['template']);
            }
            if (data.hasOwnProperty('isApproved')) {
                obj['isApproved'] = ApiClient.convertToType(data['isApproved'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} version
 */
V1Component.prototype['version'] = undefined;

/**
 * @member {String} kind
 */
V1Component.prototype['kind'] = undefined;

/**
 * @member {String} name
 */
V1Component.prototype['name'] = undefined;

/**
 * @member {String} description
 */
V1Component.prototype['description'] = undefined;

/**
 * @member {Array.<String>} tags
 */
V1Component.prototype['tags'] = undefined;

/**
 * @member {Array.<String>} presets
 */
V1Component.prototype['presets'] = undefined;

/**
 * @member {String} queue
 */
V1Component.prototype['queue'] = undefined;

/**
 * @member {module:model/V1Cache} cache
 */
V1Component.prototype['cache'] = undefined;

/**
 * @member {module:model/V1Termination} termination
 */
V1Component.prototype['termination'] = undefined;

/**
 * @member {module:model/V1Plugins} plugins
 */
V1Component.prototype['plugins'] = undefined;

/**
 * @member {Array.<module:model/V1Hook>} hooks
 */
V1Component.prototype['hooks'] = undefined;

/**
 * @member {Array.<module:model/V1IO>} inputs
 */
V1Component.prototype['inputs'] = undefined;

/**
 * @member {Array.<module:model/V1IO>} outputs
 */
V1Component.prototype['outputs'] = undefined;

/**
 * @member {Object} run
 */
V1Component.prototype['run'] = undefined;

/**
 * @member {module:model/V1Template} template
 */
V1Component.prototype['template'] = undefined;

/**
 * @member {Boolean} isApproved
 */
V1Component.prototype['isApproved'] = undefined;






export default V1Component;

