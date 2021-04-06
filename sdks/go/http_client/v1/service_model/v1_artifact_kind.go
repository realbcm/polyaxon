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
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1ArtifactKind Artifact kind
//
// - model: Model asset/event
//  - audio: Audio asset/event
//  - video: Vidio asset/event
//  - histogram: Histogram asset/event
//  - image: Image asset/event
//  - tensor: Tensor asset/event
//  - dataframe: Dataframe asset/event
//  - chart: plotly/bokeh/vega chart
//  - csv: Comma separated values
//  - tsv: Tab separated values
//  - psv: Pipe separated values
//  - ssv: Space separated values
//  - metric: Metric asset/event
//  - env: Env file
//  - html: HTML asset/event
//  - text: Text asset/event
//  - file: File asset/lineage
//  - dir: Dir asset/lineage
//  - dockerfile: Dockerfile asset
//  - docker_image: Docker image
//  - data: Data asset/event
//  - coderef: Coderef lineage
//  - table: Table asset/event
//  - tensorboard: Tensorboard lineage
//  - curve: Curve event
//  - analysis: Analysis lineage
//  - iteration: Iteration lineage
//  - markdown: Mardown event
//
// swagger:model v1ArtifactKind
type V1ArtifactKind string

func NewV1ArtifactKind(value V1ArtifactKind) *V1ArtifactKind {
	v := value
	return &v
}

const (

	// V1ArtifactKindModel captures enum value "model"
	V1ArtifactKindModel V1ArtifactKind = "model"

	// V1ArtifactKindAudio captures enum value "audio"
	V1ArtifactKindAudio V1ArtifactKind = "audio"

	// V1ArtifactKindVideo captures enum value "video"
	V1ArtifactKindVideo V1ArtifactKind = "video"

	// V1ArtifactKindHistogram captures enum value "histogram"
	V1ArtifactKindHistogram V1ArtifactKind = "histogram"

	// V1ArtifactKindImage captures enum value "image"
	V1ArtifactKindImage V1ArtifactKind = "image"

	// V1ArtifactKindTensor captures enum value "tensor"
	V1ArtifactKindTensor V1ArtifactKind = "tensor"

	// V1ArtifactKindDataframe captures enum value "dataframe"
	V1ArtifactKindDataframe V1ArtifactKind = "dataframe"

	// V1ArtifactKindChart captures enum value "chart"
	V1ArtifactKindChart V1ArtifactKind = "chart"

	// V1ArtifactKindCsv captures enum value "csv"
	V1ArtifactKindCsv V1ArtifactKind = "csv"

	// V1ArtifactKindTsv captures enum value "tsv"
	V1ArtifactKindTsv V1ArtifactKind = "tsv"

	// V1ArtifactKindPsv captures enum value "psv"
	V1ArtifactKindPsv V1ArtifactKind = "psv"

	// V1ArtifactKindSsv captures enum value "ssv"
	V1ArtifactKindSsv V1ArtifactKind = "ssv"

	// V1ArtifactKindMetric captures enum value "metric"
	V1ArtifactKindMetric V1ArtifactKind = "metric"

	// V1ArtifactKindEnv captures enum value "env"
	V1ArtifactKindEnv V1ArtifactKind = "env"

	// V1ArtifactKindHTML captures enum value "html"
	V1ArtifactKindHTML V1ArtifactKind = "html"

	// V1ArtifactKindText captures enum value "text"
	V1ArtifactKindText V1ArtifactKind = "text"

	// V1ArtifactKindFile captures enum value "file"
	V1ArtifactKindFile V1ArtifactKind = "file"

	// V1ArtifactKindDir captures enum value "dir"
	V1ArtifactKindDir V1ArtifactKind = "dir"

	// V1ArtifactKindDockerfile captures enum value "dockerfile"
	V1ArtifactKindDockerfile V1ArtifactKind = "dockerfile"

	// V1ArtifactKindDockerImage captures enum value "docker_image"
	V1ArtifactKindDockerImage V1ArtifactKind = "docker_image"

	// V1ArtifactKindData captures enum value "data"
	V1ArtifactKindData V1ArtifactKind = "data"

	// V1ArtifactKindCoderef captures enum value "coderef"
	V1ArtifactKindCoderef V1ArtifactKind = "coderef"

	// V1ArtifactKindTable captures enum value "table"
	V1ArtifactKindTable V1ArtifactKind = "table"

	// V1ArtifactKindTensorboard captures enum value "tensorboard"
	V1ArtifactKindTensorboard V1ArtifactKind = "tensorboard"

	// V1ArtifactKindCurve captures enum value "curve"
	V1ArtifactKindCurve V1ArtifactKind = "curve"

	// V1ArtifactKindAnalysis captures enum value "analysis"
	V1ArtifactKindAnalysis V1ArtifactKind = "analysis"

	// V1ArtifactKindIteration captures enum value "iteration"
	V1ArtifactKindIteration V1ArtifactKind = "iteration"

	// V1ArtifactKindMarkdown captures enum value "markdown"
	V1ArtifactKindMarkdown V1ArtifactKind = "markdown"
)

// for schema
var v1ArtifactKindEnum []interface{}

func init() {
	var res []V1ArtifactKind
	if err := json.Unmarshal([]byte(`["model","audio","video","histogram","image","tensor","dataframe","chart","csv","tsv","psv","ssv","metric","env","html","text","file","dir","dockerfile","docker_image","data","coderef","table","tensorboard","curve","analysis","iteration","markdown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1ArtifactKindEnum = append(v1ArtifactKindEnum, v)
	}
}

func (m V1ArtifactKind) validateV1ArtifactKindEnum(path, location string, value V1ArtifactKind) error {
	if err := validate.EnumCase(path, location, value, v1ArtifactKindEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 artifact kind
func (m V1ArtifactKind) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1ArtifactKindEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this v1 artifact kind based on context it is used
func (m V1ArtifactKind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
