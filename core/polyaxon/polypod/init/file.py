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

from typing import List, Optional

from polyaxon.auxiliaries import V1PolyaxonInitContainer
from polyaxon.containers.names import (
    INIT_FILE_CONTAINER_PREFIX,
    generate_container_name,
)
from polyaxon.contexts import paths as ctx_paths
from polyaxon.k8s import k8s_schemas
from polyaxon.polypod.common import constants
from polyaxon.polypod.common.containers import patch_container
from polyaxon.polypod.common.env_vars import get_run_instance_env_var
from polyaxon.polypod.common.mounts import (
    get_auth_context_mount,
    get_connections_context_mount,
)
from polyaxon.polypod.common.volumes import get_volume_name
from polyaxon.polypod.specs.contexts import PluginsContextsSpec
from polyaxon.schemas.types import V1FileType
from polyaxon.utils.list_utils import to_list


def get_file_init_container(
    polyaxon_init: V1PolyaxonInitContainer,
    file_args: V1FileType,
    contexts: PluginsContextsSpec,
    run_path: str,
    run_instance: str,
    container: Optional[k8s_schemas.V1Container] = None,
    env: List[k8s_schemas.V1EnvVar] = None,
    mount_path: Optional[str] = None,
) -> k8s_schemas.V1Container:
    env = to_list(env, check_none=True)
    env = env + [get_run_instance_env_var(run_instance)]

    container_name = generate_container_name(INIT_FILE_CONTAINER_PREFIX)
    if not container:
        container = k8s_schemas.V1Container(name=container_name)

    volume_name = (
        get_volume_name(mount_path) if mount_path else constants.VOLUME_MOUNT_ARTIFACTS
    )
    mount_path = mount_path or ctx_paths.CONTEXT_MOUNT_ARTIFACTS
    volume_mounts = [
        get_connections_context_mount(name=volume_name, mount_path=mount_path)
    ]
    if contexts and contexts.auth:
        volume_mounts.append(get_auth_context_mount(read_only=True))

    file_args.filename = file_args.filename or "file"
    return patch_container(
        container=container,
        name=container_name,
        image=polyaxon_init.get_image(),
        image_pull_policy=polyaxon_init.image_pull_policy,
        command=["polyaxon", "initializer", "file"],
        args=[
            "--file-context={}".format(file_args.to_dict(dump=True)),
            "--filepath={}".format(mount_path),
            "--copy-path={}".format(
                ctx_paths.CONTEXT_MOUNT_RUN_OUTPUTS_FORMAT.format(run_path)
            ),
            "--track",
        ],
        env=env,
        volume_mounts=volume_mounts,
        resources=polyaxon_init.get_resources(),
    )
