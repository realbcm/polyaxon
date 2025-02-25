#!/usr/bin/python
#
# Copyright 2018-2022 Polyaxon, Inc.
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

import pytest

from polyaxon.auxiliaries import get_init_resources
from polyaxon.k8s import k8s_schemas
from polyaxon.utils.test_utils import BaseTestCase


@pytest.mark.polypod_mark
class TestInitUtils(BaseTestCase):
    def test_get_init_resources(self):
        assert get_init_resources() == k8s_schemas.V1ResourceRequirements(
            limits={"cpu": "1", "memory": "500Mi"},
            requests={"cpu": "0.1", "memory": "60Mi"},
        )
