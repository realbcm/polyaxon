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
 * The version of the OpenAPI document: 1.5.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1Compatibility;
import org.openapitools.client.model.V1Installation;
import org.openapitools.client.model.V1LogHandler;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for VersionsV1Api
 */
@Ignore
public class VersionsV1ApiTest {

    private final VersionsV1Api api = new VersionsV1Api();

    
    /**
     * Get compatibility versions
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getCompatibilityTest() throws ApiException {
        String uuid = null;
        String version = null;
        String service = null;
        V1Compatibility response = api.getCompatibility(uuid, version, service);

        // TODO: test validations
    }
    
    /**
     * Get installation versions
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getInstallationTest() throws ApiException {
        Boolean auth = null;
        V1Installation response = api.getInstallation(auth);

        // TODO: test validations
    }
    
    /**
     * Get log handler versions
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getLogHandlerTest() throws ApiException {
        V1LogHandler response = api.getLogHandler();

        // TODO: test validations
    }
    
}
