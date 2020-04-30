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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.V1Environment;
import org.openapitools.client.model.V1Init;

/**
 * V1KFReplica
 */

public class V1KFReplica {
  public static final String SERIALIZED_NAME_REPLICAS = "replicas";
  @SerializedName(SERIALIZED_NAME_REPLICAS)
  private Integer replicas;

  public static final String SERIALIZED_NAME_ENVIRONMENT = "environment";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT)
  private V1Environment environment;

  public static final String SERIALIZED_NAME_CONNECTIONS = "connections";
  @SerializedName(SERIALIZED_NAME_CONNECTIONS)
  private List<String> connections = null;

  public static final String SERIALIZED_NAME_VOLUMES = "volumes";
  @SerializedName(SERIALIZED_NAME_VOLUMES)
  private List<Object> volumes = null;

  public static final String SERIALIZED_NAME_INIT = "init";
  @SerializedName(SERIALIZED_NAME_INIT)
  private List<V1Init> init = null;

  public static final String SERIALIZED_NAME_SIDECARS = "sidecars";
  @SerializedName(SERIALIZED_NAME_SIDECARS)
  private List<Object> sidecars = null;

  public static final String SERIALIZED_NAME_CONTAINER = "container";
  @SerializedName(SERIALIZED_NAME_CONTAINER)
  private Object container;


  public V1KFReplica replicas(Integer replicas) {
    
    this.replicas = replicas;
    return this;
  }

   /**
   * Get replicas
   * @return replicas
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getReplicas() {
    return replicas;
  }


  public void setReplicas(Integer replicas) {
    this.replicas = replicas;
  }


  public V1KFReplica environment(V1Environment environment) {
    
    this.environment = environment;
    return this;
  }

   /**
   * Get environment
   * @return environment
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1Environment getEnvironment() {
    return environment;
  }


  public void setEnvironment(V1Environment environment) {
    this.environment = environment;
  }


  public V1KFReplica connections(List<String> connections) {
    
    this.connections = connections;
    return this;
  }

  public V1KFReplica addConnectionsItem(String connectionsItem) {
    if (this.connections == null) {
      this.connections = new ArrayList<String>();
    }
    this.connections.add(connectionsItem);
    return this;
  }

   /**
   * Get connections
   * @return connections
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getConnections() {
    return connections;
  }


  public void setConnections(List<String> connections) {
    this.connections = connections;
  }


  public V1KFReplica volumes(List<Object> volumes) {
    
    this.volumes = volumes;
    return this;
  }

  public V1KFReplica addVolumesItem(Object volumesItem) {
    if (this.volumes == null) {
      this.volumes = new ArrayList<Object>();
    }
    this.volumes.add(volumesItem);
    return this;
  }

   /**
   * Get volumes
   * @return volumes
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getVolumes() {
    return volumes;
  }


  public void setVolumes(List<Object> volumes) {
    this.volumes = volumes;
  }


  public V1KFReplica init(List<V1Init> init) {
    
    this.init = init;
    return this;
  }

  public V1KFReplica addInitItem(V1Init initItem) {
    if (this.init == null) {
      this.init = new ArrayList<V1Init>();
    }
    this.init.add(initItem);
    return this;
  }

   /**
   * Get init
   * @return init
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<V1Init> getInit() {
    return init;
  }


  public void setInit(List<V1Init> init) {
    this.init = init;
  }


  public V1KFReplica sidecars(List<Object> sidecars) {
    
    this.sidecars = sidecars;
    return this;
  }

  public V1KFReplica addSidecarsItem(Object sidecarsItem) {
    if (this.sidecars == null) {
      this.sidecars = new ArrayList<Object>();
    }
    this.sidecars.add(sidecarsItem);
    return this;
  }

   /**
   * Get sidecars
   * @return sidecars
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getSidecars() {
    return sidecars;
  }


  public void setSidecars(List<Object> sidecars) {
    this.sidecars = sidecars;
  }


  public V1KFReplica container(Object container) {
    
    this.container = container;
    return this;
  }

   /**
   * Get container
   * @return container
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getContainer() {
    return container;
  }


  public void setContainer(Object container) {
    this.container = container;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1KFReplica v1KFReplica = (V1KFReplica) o;
    return Objects.equals(this.replicas, v1KFReplica.replicas) &&
        Objects.equals(this.environment, v1KFReplica.environment) &&
        Objects.equals(this.connections, v1KFReplica.connections) &&
        Objects.equals(this.volumes, v1KFReplica.volumes) &&
        Objects.equals(this.init, v1KFReplica.init) &&
        Objects.equals(this.sidecars, v1KFReplica.sidecars) &&
        Objects.equals(this.container, v1KFReplica.container);
  }

  @Override
  public int hashCode() {
    return Objects.hash(replicas, environment, connections, volumes, init, sidecars, container);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1KFReplica {\n");
    sb.append("    replicas: ").append(toIndentedString(replicas)).append("\n");
    sb.append("    environment: ").append(toIndentedString(environment)).append("\n");
    sb.append("    connections: ").append(toIndentedString(connections)).append("\n");
    sb.append("    volumes: ").append(toIndentedString(volumes)).append("\n");
    sb.append("    init: ").append(toIndentedString(init)).append("\n");
    sb.append("    sidecars: ").append(toIndentedString(sidecars)).append("\n");
    sb.append("    container: ").append(toIndentedString(container)).append("\n");
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

