# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

import os

from kubernetes import client

from polyaxon_schemas.exceptions import PolyaxonConfigurationError
from polyaxon_schemas.k8s.templates import constants
from polyaxon_schemas.settings import RunTypes


def get_cluster_env_var(project, experiment, task_type):
    cluster_name = constants.CONFIG_MAP_CLUSTER_NAME.format(project=project, experiment=experiment)
    config_map_key_ref = client.V1ConfigMapKeySelector(name=cluster_name, key=task_type)
    value = client.V1EnvVarSource(config_map_key_ref=config_map_key_ref)
    key_name = constants.CONFIG_MAP_CLUSTER_KEY_NAME.format(project=project,
                                                            experiment=experiment,
                                                            task_type=task_type)
    return client.V1EnvVar(name=key_name, value_from=value)


def get_gpu_resources(gpu_limits=0, gpu_requests=0):
    limits = constants.GPU_RESOURCES.format(gpu_limits) if gpu_limits > 0 else ''
    requests = constants.GPU_RESOURCES.format(gpu_requests) if gpu_requests > 0 else ''
    return client.V1ResourceRequirements(limits=limits, requests=requests)


def get_gpu_volume_mounts():
    return [
        client.V1VolumeMount(name='bin', mount_path='/usr/local/nvidia/bin'),
        client.V1VolumeMount(name='lib', mount_path='/usr/local/nvidia/lib'),
    ]


def get_gpu_volumes():
    return [
        client.V1Volume(name='bin',
                        host_path=client.V1HostPathVolumeSource(path='/usr/local/nvidia/bin')),
        client.V1Volume(name='lib',
                        host_path=client.V1HostPathVolumeSource(path='/usr/local/nvidia/lib')),
    ]


def get_volume_mount(volume):
    return client.V1VolumeMount(name=volume, mount_path=os.path.join('/', volume))


def get_volume(volume, run_type):
    host_path = None
    nfs = None
    path = os.path.join('/', volume)
    if run_type == RunTypes.MINIKUBE:
        host_path = client.V1HostPathVolumeSource(path=path)
    elif run_type == RunTypes.KUBERNETES:
        nfs = client.V1NFSVolumeSource(path=path)
    else:
        raise PolyaxonConfigurationError('Run type `{}` is not allowed.'.format(run_type))
    return client.V1Volume(name=volume, host_path=host_path, nfs=nfs)


def get_project_pod_spec(project,
                         name,
                         volume_mounts,
                         volumes,
                         args=None,
                         ports=None,
                         gpu_limits=0,
                         gpu_requests=0,
                         env_vars=None,
                         restart_policy='OnFailure'):
    """Pod spec to be used to create pods for project side: tensorboard, notebooks."""
    volume_mounts = volume_mounts or []
    volumes = volumes or []

    volume_mounts += get_gpu_volume_mounts()
    volumes += get_gpu_volumes()

    ports = [client.V1ContainerPort(container_port=port) for port in ports]

    container_name = constants.POD_CONTAINER_PROJECT_NAME.format(project=project, name=name)
    containers = [client.V1Container(name=container_name,
                                     image=constants.DOCKER_IMAGE,
                                     args=args,
                                     ports=ports,
                                     env=env_vars,
                                     resources=get_gpu_resources(gpu_limits, gpu_requests),
                                     volume_mounts=volume_mounts)]
    return client.V1PodSpec(restart_policy=restart_policy, containers=containers, volumes=volumes)


def get_task_pod_spec(project,
                      experiment,
                      task_type,
                      task_id,
                      volume_mounts,
                      volumes,
                      env_vars=None,
                      args=None,
                      ports=None,
                      gpu_limits=0,
                      gpu_requests=0,
                      restart_policy='OnFailure'):
    """Pod spec to be used to create pods for tasks: master, worker, ps."""
    # env_vars = [
    #     get_cluster_env_var(project=project, experiment=experiment, task_type=task_type)
    # ]

    volume_mounts = volume_mounts or []
    volumes = volumes or []

    volume_mounts += get_gpu_volume_mounts()
    volumes += get_gpu_volumes()

    ports = [client.V1ContainerPort(container_port=port) for port in ports]

    container_name = constants.POD_CONTAINER_TASK_NAME.format(project=project,
                                                              experiment=experiment,
                                                              task_type=task_type,
                                                              task_id=task_id)
    containers = [client.V1Container(name=container_name,
                                     image=constants.DOCKER_IMAGE,
                                     args=args,
                                     ports=ports,
                                     env=env_vars,
                                     resources=get_gpu_resources(gpu_limits, gpu_requests),
                                     volume_mounts=volume_mounts)]
    return client.V1PodSpec(restart_policy=restart_policy, containers=containers, volumes=volumes)


def get_labels(project, experiment, task_type, task_id, task_name):
    return {'project': project,
            'experiment': experiment,
            'task_type': task_type,
            'task_id': task_id,
            'task': task_name}


def get_pod(project, experiment, task_type, task_id, volume_mounts, volumes, ports):
    task_name = constants.TASK_NAME.format(project=project,
                                           experiment=experiment,
                                           task_type=task_type,
                                           task_id=task_id)
    labels = get_labels(project, experiment, task_type, task_id, task_name)
    metadata = client.V1ObjectMeta(name=task_name, labels=labels)

    spec = get_task_pod_spec(project=project,
                             experiment=experiment,
                             task_type=task_type,
                             task_id=task_id,
                             volume_mounts=volume_mounts,
                             volumes=volumes,
                             ports=ports)
    return client.V1Pod(api_version=constants.K8S_API_VERSION_V1,
                        kind=constants.K8S_POD_KIND,
                        metadata=metadata,
                        spec=spec)
