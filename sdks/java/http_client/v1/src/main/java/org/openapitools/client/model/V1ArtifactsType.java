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
 * The version of the OpenAPI document: 1.0.97
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
import java.util.ArrayList;
import java.util.List;

/**
 * V1ArtifactsType
 */

public class V1ArtifactsType {
  public static final String SERIALIZED_NAME_CONNECTION = "connection";
  @SerializedName(SERIALIZED_NAME_CONNECTION)
  private String connection;

  public static final String SERIALIZED_NAME_FILES = "files";
  @SerializedName(SERIALIZED_NAME_FILES)
  private List<String> files = null;

  public static final String SERIALIZED_NAME_DIRS = "dirs";
  @SerializedName(SERIALIZED_NAME_DIRS)
  private List<String> dirs = null;

  public static final String SERIALIZED_NAME_INIT = "init";
  @SerializedName(SERIALIZED_NAME_INIT)
  private Boolean init;

  public static final String SERIALIZED_NAME_WORKERS = "workers";
  @SerializedName(SERIALIZED_NAME_WORKERS)
  private Integer workers;


  public V1ArtifactsType connection(String connection) {
    
    this.connection = connection;
    return this;
  }

   /**
   * Get connection
   * @return connection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getConnection() {
    return connection;
  }


  public void setConnection(String connection) {
    this.connection = connection;
  }


  public V1ArtifactsType files(List<String> files) {
    
    this.files = files;
    return this;
  }

  public V1ArtifactsType addFilesItem(String filesItem) {
    if (this.files == null) {
      this.files = new ArrayList<String>();
    }
    this.files.add(filesItem);
    return this;
  }

   /**
   * Get files
   * @return files
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getFiles() {
    return files;
  }


  public void setFiles(List<String> files) {
    this.files = files;
  }


  public V1ArtifactsType dirs(List<String> dirs) {
    
    this.dirs = dirs;
    return this;
  }

  public V1ArtifactsType addDirsItem(String dirsItem) {
    if (this.dirs == null) {
      this.dirs = new ArrayList<String>();
    }
    this.dirs.add(dirsItem);
    return this;
  }

   /**
   * Get dirs
   * @return dirs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getDirs() {
    return dirs;
  }


  public void setDirs(List<String> dirs) {
    this.dirs = dirs;
  }


  public V1ArtifactsType init(Boolean init) {
    
    this.init = init;
    return this;
  }

   /**
   * Get init
   * @return init
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getInit() {
    return init;
  }


  public void setInit(Boolean init) {
    this.init = init;
  }


  public V1ArtifactsType workers(Integer workers) {
    
    this.workers = workers;
    return this;
  }

   /**
   * Get workers
   * @return workers
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getWorkers() {
    return workers;
  }


  public void setWorkers(Integer workers) {
    this.workers = workers;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1ArtifactsType v1ArtifactsType = (V1ArtifactsType) o;
    return Objects.equals(this.connection, v1ArtifactsType.connection) &&
        Objects.equals(this.files, v1ArtifactsType.files) &&
        Objects.equals(this.dirs, v1ArtifactsType.dirs) &&
        Objects.equals(this.init, v1ArtifactsType.init) &&
        Objects.equals(this.workers, v1ArtifactsType.workers);
  }

  @Override
  public int hashCode() {
    return Objects.hash(connection, files, dirs, init, workers);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1ArtifactsType {\n");
    sb.append("    connection: ").append(toIndentedString(connection)).append("\n");
    sb.append("    files: ").append(toIndentedString(files)).append("\n");
    sb.append("    dirs: ").append(toIndentedString(dirs)).append("\n");
    sb.append("    init: ").append(toIndentedString(init)).append("\n");
    sb.append("    workers: ").append(toIndentedString(workers)).append("\n");
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

