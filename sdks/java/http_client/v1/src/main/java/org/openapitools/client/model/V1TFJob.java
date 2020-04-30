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
import org.openapitools.client.model.V1CleanPodPolicy;
import org.openapitools.client.model.V1KFReplica;

/**
 * V1TFJob
 */

public class V1TFJob {
  public static final String SERIALIZED_NAME_KIND = "kind";
  @SerializedName(SERIALIZED_NAME_KIND)
  private String kind = "tfjob";

  public static final String SERIALIZED_NAME_CLEAN_POD_POLICY = "cleanPodPolicy";
  @SerializedName(SERIALIZED_NAME_CLEAN_POD_POLICY)
  private V1CleanPodPolicy cleanPodPolicy = V1CleanPodPolicy.ALL;

  public static final String SERIALIZED_NAME_CHIEF = "chief";
  @SerializedName(SERIALIZED_NAME_CHIEF)
  private V1KFReplica chief;

  public static final String SERIALIZED_NAME_PS = "ps";
  @SerializedName(SERIALIZED_NAME_PS)
  private V1KFReplica ps;

  public static final String SERIALIZED_NAME_WORKER = "worker";
  @SerializedName(SERIALIZED_NAME_WORKER)
  private V1KFReplica worker;

  public static final String SERIALIZED_NAME_EVALUATOR = "evaluator";
  @SerializedName(SERIALIZED_NAME_EVALUATOR)
  private V1KFReplica evaluator;


  public V1TFJob kind(String kind) {
    
    this.kind = kind;
    return this;
  }

   /**
   * Get kind
   * @return kind
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getKind() {
    return kind;
  }


  public void setKind(String kind) {
    this.kind = kind;
  }


  public V1TFJob cleanPodPolicy(V1CleanPodPolicy cleanPodPolicy) {
    
    this.cleanPodPolicy = cleanPodPolicy;
    return this;
  }

   /**
   * Get cleanPodPolicy
   * @return cleanPodPolicy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1CleanPodPolicy getCleanPodPolicy() {
    return cleanPodPolicy;
  }


  public void setCleanPodPolicy(V1CleanPodPolicy cleanPodPolicy) {
    this.cleanPodPolicy = cleanPodPolicy;
  }


  public V1TFJob chief(V1KFReplica chief) {
    
    this.chief = chief;
    return this;
  }

   /**
   * Get chief
   * @return chief
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1KFReplica getChief() {
    return chief;
  }


  public void setChief(V1KFReplica chief) {
    this.chief = chief;
  }


  public V1TFJob ps(V1KFReplica ps) {
    
    this.ps = ps;
    return this;
  }

   /**
   * Get ps
   * @return ps
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1KFReplica getPs() {
    return ps;
  }


  public void setPs(V1KFReplica ps) {
    this.ps = ps;
  }


  public V1TFJob worker(V1KFReplica worker) {
    
    this.worker = worker;
    return this;
  }

   /**
   * Get worker
   * @return worker
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1KFReplica getWorker() {
    return worker;
  }


  public void setWorker(V1KFReplica worker) {
    this.worker = worker;
  }


  public V1TFJob evaluator(V1KFReplica evaluator) {
    
    this.evaluator = evaluator;
    return this;
  }

   /**
   * Get evaluator
   * @return evaluator
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1KFReplica getEvaluator() {
    return evaluator;
  }


  public void setEvaluator(V1KFReplica evaluator) {
    this.evaluator = evaluator;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1TFJob v1TFJob = (V1TFJob) o;
    return Objects.equals(this.kind, v1TFJob.kind) &&
        Objects.equals(this.cleanPodPolicy, v1TFJob.cleanPodPolicy) &&
        Objects.equals(this.chief, v1TFJob.chief) &&
        Objects.equals(this.ps, v1TFJob.ps) &&
        Objects.equals(this.worker, v1TFJob.worker) &&
        Objects.equals(this.evaluator, v1TFJob.evaluator);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, cleanPodPolicy, chief, ps, worker, evaluator);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1TFJob {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    cleanPodPolicy: ").append(toIndentedString(cleanPodPolicy)).append("\n");
    sb.append("    chief: ").append(toIndentedString(chief)).append("\n");
    sb.append("    ps: ").append(toIndentedString(ps)).append("\n");
    sb.append("    worker: ").append(toIndentedString(worker)).append("\n");
    sb.append("    evaluator: ").append(toIndentedString(evaluator)).append("\n");
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

