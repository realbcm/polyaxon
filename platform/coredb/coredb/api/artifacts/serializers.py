#!/usr/bin/python
#
# Copyright 2018-2021 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from rest_framework import fields, serializers

from coredb.abstracts.getter import get_lineage_model


class RunArtifactSerializer(serializers.ModelSerializer):
    state = fields.SerializerMethodField()
    name = fields.SerializerMethodField()
    kind = fields.SerializerMethodField()
    path = fields.SerializerMethodField()
    summary = fields.SerializerMethodField()

    class Meta:
        model = get_lineage_model()
        fields = ("name", "kind", "path", "summary", "state", "is_input")

    def get_name(self, obj):
        return obj.artifact.name

    def get_state(self, obj):
        return obj.artifact.state

    def get_kind(self, obj):
        return obj.artifact.kind

    def get_path(self, obj):
        return obj.artifact.path

    def get_summary(self, obj):
        return obj.artifact.summary


class RunArtifactLightSerializer(serializers.ModelSerializer):
    name = fields.SerializerMethodField()
    kind = fields.SerializerMethodField()

    class Meta:
        model = get_lineage_model()
        fields = ("name", "kind", "is_input")

    def get_name(self, obj):
        return obj.artifact.name

    def get_kind(self, obj):
        return obj.artifact.kind


class RunArtifactNameSerializer(serializers.ModelSerializer):
    name = fields.SerializerMethodField()

    class Meta:
        model = get_lineage_model()
        fields = ("name",)

    def get_name(self, obj):
        return obj.artifact.name
