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

/*
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


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1ComponentHub;
import org.openapitools.client.model.V1ComponentHubSettings;
import org.openapitools.client.model.V1ComponentVersion;
import org.openapitools.client.model.V1EntityStageBodyRequest;
import org.openapitools.client.model.V1ListComponentHubsResponse;
import org.openapitools.client.model.V1ListComponentVersionsResponse;
import org.openapitools.client.model.V1Stage;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for ComponentHubV1Api
 */
@Ignore
public class ComponentHubV1ApiTest {

    private final ComponentHubV1Api api = new ComponentHubV1Api();

    
    /**
     * Archive hub component
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void archiveComponentHubTest() throws ApiException {
        String owner = null;
        String name = null;
        api.archiveComponentHub(owner, name);

        // TODO: test validations
    }
    
    /**
     * Bookmark component hub
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void bookmarkComponentHubTest() throws ApiException {
        String owner = null;
        String name = null;
        api.bookmarkComponentHub(owner, name);

        // TODO: test validations
    }
    
    /**
     * Create hub component
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createComponentHubTest() throws ApiException {
        String owner = null;
        V1ComponentHub body = null;
        V1ComponentHub response = api.createComponentHub(owner, body);

        // TODO: test validations
    }
    
    /**
     * Create component version
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createComponentVersionTest() throws ApiException {
        String owner = null;
        String component = null;
        V1ComponentVersion body = null;
        V1ComponentVersion response = api.createComponentVersion(owner, component, body);

        // TODO: test validations
    }
    
    /**
     * Create new component version stage
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createComponentVersionStageTest() throws ApiException {
        String owner = null;
        String entity = null;
        String name = null;
        V1EntityStageBodyRequest body = null;
        V1Stage response = api.createComponentVersionStage(owner, entity, name, body);

        // TODO: test validations
    }
    
    /**
     * Delete hub component
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteComponentHubTest() throws ApiException {
        String owner = null;
        String name = null;
        api.deleteComponentHub(owner, name);

        // TODO: test validations
    }
    
    /**
     * Delete component version
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteComponentVersionTest() throws ApiException {
        String owner = null;
        String entity = null;
        String name = null;
        api.deleteComponentVersion(owner, entity, name);

        // TODO: test validations
    }
    
    /**
     * Get hub component
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getComponentHubTest() throws ApiException {
        String owner = null;
        String name = null;
        V1ComponentHub response = api.getComponentHub(owner, name);

        // TODO: test validations
    }
    
    /**
     * Get hub component settings
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getComponentHubSettingsTest() throws ApiException {
        String owner = null;
        String name = null;
        V1ComponentHubSettings response = api.getComponentHubSettings(owner, name);

        // TODO: test validations
    }
    
    /**
     * Get component version
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getComponentVersionTest() throws ApiException {
        String owner = null;
        String entity = null;
        String name = null;
        V1ComponentVersion response = api.getComponentVersion(owner, entity, name);

        // TODO: test validations
    }
    
    /**
     * Get component version stages
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getComponentVersionStagesTest() throws ApiException {
        String owner = null;
        String entity = null;
        String name = null;
        V1Stage response = api.getComponentVersionStages(owner, entity, name);

        // TODO: test validations
    }
    
    /**
     * List hub component names
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listComponentHubNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListComponentHubsResponse response = api.listComponentHubNames(owner, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * List hub components
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listComponentHubsTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListComponentHubsResponse response = api.listComponentHubs(owner, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * List component version names
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listComponentVersionNamesTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListComponentVersionsResponse response = api.listComponentVersionNames(owner, name, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * List component versions
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listComponentVersionsTest() throws ApiException {
        String owner = null;
        String name = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListComponentVersionsResponse response = api.listComponentVersions(owner, name, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * Patch hub component
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchComponentHubTest() throws ApiException {
        String owner = null;
        String componentName = null;
        V1ComponentHub body = null;
        V1ComponentHub response = api.patchComponentHub(owner, componentName, body);

        // TODO: test validations
    }
    
    /**
     * Patch hub component settings
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchComponentHubSettingsTest() throws ApiException {
        String owner = null;
        String component = null;
        V1ComponentHubSettings body = null;
        V1ComponentHubSettings response = api.patchComponentHubSettings(owner, component, body);

        // TODO: test validations
    }
    
    /**
     * Patch component version
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchComponentVersionTest() throws ApiException {
        String owner = null;
        String component = null;
        String versionName = null;
        V1ComponentVersion body = null;
        V1ComponentVersion response = api.patchComponentVersion(owner, component, versionName, body);

        // TODO: test validations
    }
    
    /**
     * Restore hub component
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void restoreComponentHubTest() throws ApiException {
        String owner = null;
        String name = null;
        api.restoreComponentHub(owner, name);

        // TODO: test validations
    }
    
    /**
     * Unbookmark component hub
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void unbookmarkComponentHubTest() throws ApiException {
        String owner = null;
        String name = null;
        api.unbookmarkComponentHub(owner, name);

        // TODO: test validations
    }
    
    /**
     * Update hub component
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateComponentHubTest() throws ApiException {
        String owner = null;
        String componentName = null;
        V1ComponentHub body = null;
        V1ComponentHub response = api.updateComponentHub(owner, componentName, body);

        // TODO: test validations
    }
    
    /**
     * Update hub component settings
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateComponentHubSettingsTest() throws ApiException {
        String owner = null;
        String component = null;
        V1ComponentHubSettings body = null;
        V1ComponentHubSettings response = api.updateComponentHubSettings(owner, component, body);

        // TODO: test validations
    }
    
    /**
     * Update component version
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateComponentVersionTest() throws ApiException {
        String owner = null;
        String component = null;
        String versionName = null;
        V1ComponentVersion body = null;
        V1ComponentVersion response = api.updateComponentVersion(owner, component, versionName, body);

        // TODO: test validations
    }
    
}
