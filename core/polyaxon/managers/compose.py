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

from polyaxon.managers.base import BaseConfigManager


class ComposeConfigManager(BaseConfigManager):
    """Manages access cli configuration .compose file."""

    VISIBILITY = BaseConfigManager.VISIBILITY_GLOBAL
    CONFIG_FILE_NAME = ".compose/.env"
    FREQUENCY = 3

    @classmethod
    def get_config_filepath(cls, create=True):
        path = super().get_config_filepath(create=create)
        values = path.split("/")[:-1]
        cls._create_dir("/".join(values))
        return path
