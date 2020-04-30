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

/*
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


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * V1Version
 */

public class V1Version {
  public static final String SERIALIZED_NAME_MIN_VERSION = "min_version";
  @SerializedName(SERIALIZED_NAME_MIN_VERSION)
  private String minVersion;

  public static final String SERIALIZED_NAME_LATEST_VERSION = "latest_version";
  @SerializedName(SERIALIZED_NAME_LATEST_VERSION)
  private String latestVersion;


  public V1Version minVersion(String minVersion) {
    
    this.minVersion = minVersion;
    return this;
  }

   /**
   * Get minVersion
   * @return minVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMinVersion() {
    return minVersion;
  }


  public void setMinVersion(String minVersion) {
    this.minVersion = minVersion;
  }


  public V1Version latestVersion(String latestVersion) {
    
    this.latestVersion = latestVersion;
    return this;
  }

   /**
   * Get latestVersion
   * @return latestVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLatestVersion() {
    return latestVersion;
  }


  public void setLatestVersion(String latestVersion) {
    this.latestVersion = latestVersion;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Version v1Version = (V1Version) o;
    return Objects.equals(this.minVersion, v1Version.minVersion) &&
        Objects.equals(this.latestVersion, v1Version.latestVersion);
  }

  @Override
  public int hashCode() {
    return Objects.hash(minVersion, latestVersion);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Version {\n");
    sb.append("    minVersion: ").append(toIndentedString(minVersion)).append("\n");
    sb.append("    latestVersion: ").append(toIndentedString(latestVersion)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

