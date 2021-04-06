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
 * The version of the OpenAPI document: 1.7.6
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiCallback;
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.ApiResponse;
import org.openapitools.client.Configuration;
import org.openapitools.client.Pair;
import org.openapitools.client.ProgressRequestBody;
import org.openapitools.client.ProgressResponseBody;

import com.google.gson.reflect.TypeToken;

import java.io.IOException;


import java.io.File;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class ArtifactsStoresV1Api {
    private ApiClient localVarApiClient;

    public ArtifactsStoresV1Api() {
        this(Configuration.getDefaultApiClient());
    }

    public ArtifactsStoresV1Api(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    public ApiClient getApiClient() {
        return localVarApiClient;
    }

    public void setApiClient(ApiClient apiClient) {
        this.localVarApiClient = apiClient;
    }

    /**
     * Build call for uploadArtifact
     * @param owner Owner of the namespace (required)
     * @param uuid Unique integer identifier of the entity (required)
     * @param uploadfile The file to upload. (required)
     * @param path File path query params. (optional)
     * @param overwrite File path query params. (optional)
     * @param _callback Callback for upload/download progress
     * @return Call to execute
     * @throws ApiException If fail to serialize the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> A successful response. </td><td>  -  </td></tr>
        <tr><td> 204 </td><td> No content. </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> You don&#39;t have permission to access the resource. </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> Resource does not exist. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call uploadArtifactCall(String owner, String uuid, File uploadfile, String path, Boolean overwrite, final ApiCallback _callback) throws ApiException {
        Object localVarPostBody = null;

        // create path and map variables
        String localVarPath = "/api/v1/catalogs/{owner}/artifacts/{uuid}/upload"
            .replaceAll("\\{" + "owner" + "\\}", localVarApiClient.escapeString(owner.toString()))
            .replaceAll("\\{" + "uuid" + "\\}", localVarApiClient.escapeString(uuid.toString()));

        List<Pair> localVarQueryParams = new ArrayList<Pair>();
        List<Pair> localVarCollectionQueryParams = new ArrayList<Pair>();
        if (path != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("path", path));
        }

        if (overwrite != null) {
            localVarQueryParams.addAll(localVarApiClient.parameterToPair("overwrite", overwrite));
        }

        Map<String, String> localVarHeaderParams = new HashMap<String, String>();
        Map<String, String> localVarCookieParams = new HashMap<String, String>();
        Map<String, Object> localVarFormParams = new HashMap<String, Object>();
        if (uploadfile != null) {
            localVarFormParams.put("uploadfile", uploadfile);
        }

        final String[] localVarAccepts = {
            "application/json"
        };
        final String localVarAccept = localVarApiClient.selectHeaderAccept(localVarAccepts);
        if (localVarAccept != null) {
            localVarHeaderParams.put("Accept", localVarAccept);
        }

        final String[] localVarContentTypes = {
            "multipart/form-data"
        };
        final String localVarContentType = localVarApiClient.selectHeaderContentType(localVarContentTypes);
        localVarHeaderParams.put("Content-Type", localVarContentType);

        String[] localVarAuthNames = new String[] { "ApiKey" };
        return localVarApiClient.buildCall(localVarPath, "POST", localVarQueryParams, localVarCollectionQueryParams, localVarPostBody, localVarHeaderParams, localVarCookieParams, localVarFormParams, localVarAuthNames, _callback);
    }

    @SuppressWarnings("rawtypes")
    private okhttp3.Call uploadArtifactValidateBeforeCall(String owner, String uuid, File uploadfile, String path, Boolean overwrite, final ApiCallback _callback) throws ApiException {
        
        // verify the required parameter 'owner' is set
        if (owner == null) {
            throw new ApiException("Missing the required parameter 'owner' when calling uploadArtifact(Async)");
        }
        
        // verify the required parameter 'uuid' is set
        if (uuid == null) {
            throw new ApiException("Missing the required parameter 'uuid' when calling uploadArtifact(Async)");
        }
        
        // verify the required parameter 'uploadfile' is set
        if (uploadfile == null) {
            throw new ApiException("Missing the required parameter 'uploadfile' when calling uploadArtifact(Async)");
        }
        

        okhttp3.Call localVarCall = uploadArtifactCall(owner, uuid, uploadfile, path, overwrite, _callback);
        return localVarCall;

    }

    /**
     * Upload artifact to a store
     * 
     * @param owner Owner of the namespace (required)
     * @param uuid Unique integer identifier of the entity (required)
     * @param uploadfile The file to upload. (required)
     * @param path File path query params. (optional)
     * @param overwrite File path query params. (optional)
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> A successful response. </td><td>  -  </td></tr>
        <tr><td> 204 </td><td> No content. </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> You don&#39;t have permission to access the resource. </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> Resource does not exist. </td><td>  -  </td></tr>
     </table>
     */
    public void uploadArtifact(String owner, String uuid, File uploadfile, String path, Boolean overwrite) throws ApiException {
        uploadArtifactWithHttpInfo(owner, uuid, uploadfile, path, overwrite);
    }

    /**
     * Upload artifact to a store
     * 
     * @param owner Owner of the namespace (required)
     * @param uuid Unique integer identifier of the entity (required)
     * @param uploadfile The file to upload. (required)
     * @param path File path query params. (optional)
     * @param overwrite File path query params. (optional)
     * @return ApiResponse&lt;Void&gt;
     * @throws ApiException If fail to call the API, e.g. server error or cannot deserialize the response body
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> A successful response. </td><td>  -  </td></tr>
        <tr><td> 204 </td><td> No content. </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> You don&#39;t have permission to access the resource. </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> Resource does not exist. </td><td>  -  </td></tr>
     </table>
     */
    public ApiResponse<Void> uploadArtifactWithHttpInfo(String owner, String uuid, File uploadfile, String path, Boolean overwrite) throws ApiException {
        okhttp3.Call localVarCall = uploadArtifactValidateBeforeCall(owner, uuid, uploadfile, path, overwrite, null);
        return localVarApiClient.execute(localVarCall);
    }

    /**
     * Upload artifact to a store (asynchronously)
     * 
     * @param owner Owner of the namespace (required)
     * @param uuid Unique integer identifier of the entity (required)
     * @param uploadfile The file to upload. (required)
     * @param path File path query params. (optional)
     * @param overwrite File path query params. (optional)
     * @param _callback The callback to be executed when the API call finishes
     * @return The request call
     * @throws ApiException If fail to process the API call, e.g. serializing the request body object
     * @http.response.details
     <table summary="Response Details" border="1">
        <tr><td> Status Code </td><td> Description </td><td> Response Headers </td></tr>
        <tr><td> 200 </td><td> A successful response. </td><td>  -  </td></tr>
        <tr><td> 204 </td><td> No content. </td><td>  -  </td></tr>
        <tr><td> 403 </td><td> You don&#39;t have permission to access the resource. </td><td>  -  </td></tr>
        <tr><td> 404 </td><td> Resource does not exist. </td><td>  -  </td></tr>
     </table>
     */
    public okhttp3.Call uploadArtifactAsync(String owner, String uuid, File uploadfile, String path, Boolean overwrite, final ApiCallback<Void> _callback) throws ApiException {

        okhttp3.Call localVarCall = uploadArtifactValidateBeforeCall(owner, uuid, uploadfile, path, overwrite, _callback);
        localVarApiClient.executeAsync(localVarCall, _callback);
        return localVarCall;
    }
}
