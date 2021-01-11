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

// Code generated by go-swagger; DO NOT EDIT.

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Spark Spark specification
//
// swagger:model v1Spark
type V1Spark struct {

	// Arguments is a list of arguments to be passed to the application.
	Arguments []string `json:"arguments"`

	// Optional connections section
	Connections []string `json:"connections"`

	// Mode is the deployment mode of the Spark application.
	DeployMode SparkDeployMode `json:"deploy_mode,omitempty"`

	// Optional spark driver definition
	Driver *V1SparkReplica `json:"driver,omitempty"`

	// Optional spark executor definition
	Executor *V1SparkReplica `json:"executor,omitempty"`

	// HadoopConf carries user-specified Hadoop configuration properties as they would use the  the "--conf" option
	// in spark-submit.  The SparkApplication controller automatically adds prefix "spark.hadoop." to Hadoop
	// configuration properties.
	HadoopConf map[string]string `json:"hadoop_conf,omitempty"`

	// HadoopConfigMap carries the name of the ConfigMap containing Hadoop configuration files such as core-site.xml.
	// The controller will add environment variable HADOOP_CONF_DIR to the path where the ConfigMap is mounted to.
	HadoopConfigMap string `json:"hadoop_config_map,omitempty"`

	// Kind of runtime, should be equal to "spark"
	Kind *string `json:"kind,omitempty"`

	// MainFile is the path to a bundled JAR, Python, or R file of the application.
	MainApplicationFile string `json:"main_application_file,omitempty"`

	// MainClass is the fully-qualified main class of the Spark application.
	// This only applies to Java/Scala Spark applications.
	MainClass string `json:"main_class,omitempty"`

	// Spark version is the version of Spark the application uses.
	PythonVersion string `json:"python_version,omitempty"`

	// HadoopConf carries user-specified Hadoop configuration properties as they would use the  the "--conf" option
	// in spark-submit.  The SparkApplication controller automatically adds prefix "spark.hadoop." to Hadoop
	// configuration properties.
	SparkConf map[string]string `json:"spark_conf,omitempty"`

	// SparkConfigMap carries the name of the ConfigMap containing Spark configuration files such as log4j.properties.
	// The controller will add environment variable SPARK_CONF_DIR to the path where the ConfigMap is mounted to.
	SparkConfigMap string `json:"spark_config_map,omitempty"`

	// Spark version is the version of Spark the application uses.
	SparkVersion string `json:"spark_version,omitempty"`

	// Type tells the type of the Spark application.
	Type V1SparkType `json:"type,omitempty"`

	// Volumes is a list of volumes that can be mounted.
	Volumes []V1Volume `json:"volumes"`
}

// Validate validates this v1 spark
func (m *V1Spark) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeployMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Spark) validateDeployMode(formats strfmt.Registry) error {

	if swag.IsZero(m.DeployMode) { // not required
		return nil
	}

	if err := m.DeployMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("deploy_mode")
		}
		return err
	}

	return nil
}

func (m *V1Spark) validateDriver(formats strfmt.Registry) error {

	if swag.IsZero(m.Driver) { // not required
		return nil
	}

	if m.Driver != nil {
		if err := m.Driver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("driver")
			}
			return err
		}
	}

	return nil
}

func (m *V1Spark) validateExecutor(formats strfmt.Registry) error {

	if swag.IsZero(m.Executor) { // not required
		return nil
	}

	if m.Executor != nil {
		if err := m.Executor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("executor")
			}
			return err
		}
	}

	return nil
}

func (m *V1Spark) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Spark) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Spark) UnmarshalBinary(b []byte) error {
	var res V1Spark
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
