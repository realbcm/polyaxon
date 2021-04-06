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

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    The version of the OpenAPI document: 1.7.6
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1Environment(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'labels': 'dict(str, str)',
        'annotations': 'dict(str, str)',
        'node_selector': 'dict(str, str)',
        'affinity': 'V1Affinity',
        'tolerations': 'list[V1Toleration]',
        'node_name': 'str',
        'service_account_name': 'str',
        'host_aliases': 'list[V1HostAlias]',
        'security_context': 'V1PodSecurityContext',
        'image_pull_secrets': 'list[str]',
        'host_network': 'bool',
        'host_pid': 'str',
        'dns_policy': 'str',
        'dns_config': 'V1PodDNSConfig',
        'scheduler_name': 'str',
        'priority_class_name': 'str',
        'priority': 'int',
        'restart_policy': 'str'
    }

    attribute_map = {
        'labels': 'labels',
        'annotations': 'annotations',
        'node_selector': 'nodeSelector',
        'affinity': 'affinity',
        'tolerations': 'tolerations',
        'node_name': 'nodeName',
        'service_account_name': 'serviceAccountName',
        'host_aliases': 'hostAliases',
        'security_context': 'securityContext',
        'image_pull_secrets': 'imagePullSecrets',
        'host_network': 'hostNetwork',
        'host_pid': 'hostPID',
        'dns_policy': 'dnsPolicy',
        'dns_config': 'dnsConfig',
        'scheduler_name': 'schedulerName',
        'priority_class_name': 'priorityClassName',
        'priority': 'priority',
        'restart_policy': 'restartPolicy'
    }

    def __init__(self, labels=None, annotations=None, node_selector=None, affinity=None, tolerations=None, node_name=None, service_account_name=None, host_aliases=None, security_context=None, image_pull_secrets=None, host_network=None, host_pid=None, dns_policy=None, dns_config=None, scheduler_name=None, priority_class_name=None, priority=None, restart_policy=None, local_vars_configuration=None):  # noqa: E501
        """V1Environment - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._labels = None
        self._annotations = None
        self._node_selector = None
        self._affinity = None
        self._tolerations = None
        self._node_name = None
        self._service_account_name = None
        self._host_aliases = None
        self._security_context = None
        self._image_pull_secrets = None
        self._host_network = None
        self._host_pid = None
        self._dns_policy = None
        self._dns_config = None
        self._scheduler_name = None
        self._priority_class_name = None
        self._priority = None
        self._restart_policy = None
        self.discriminator = None

        if labels is not None:
            self.labels = labels
        if annotations is not None:
            self.annotations = annotations
        if node_selector is not None:
            self.node_selector = node_selector
        if affinity is not None:
            self.affinity = affinity
        if tolerations is not None:
            self.tolerations = tolerations
        if node_name is not None:
            self.node_name = node_name
        if service_account_name is not None:
            self.service_account_name = service_account_name
        if host_aliases is not None:
            self.host_aliases = host_aliases
        if security_context is not None:
            self.security_context = security_context
        if image_pull_secrets is not None:
            self.image_pull_secrets = image_pull_secrets
        if host_network is not None:
            self.host_network = host_network
        if host_pid is not None:
            self.host_pid = host_pid
        if dns_policy is not None:
            self.dns_policy = dns_policy
        if dns_config is not None:
            self.dns_config = dns_config
        if scheduler_name is not None:
            self.scheduler_name = scheduler_name
        if priority_class_name is not None:
            self.priority_class_name = priority_class_name
        if priority is not None:
            self.priority = priority
        if restart_policy is not None:
            self.restart_policy = restart_policy

    @property
    def labels(self):
        """Gets the labels of this V1Environment.  # noqa: E501


        :return: The labels of this V1Environment.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._labels

    @labels.setter
    def labels(self, labels):
        """Sets the labels of this V1Environment.


        :param labels: The labels of this V1Environment.  # noqa: E501
        :type: dict(str, str)
        """

        self._labels = labels

    @property
    def annotations(self):
        """Gets the annotations of this V1Environment.  # noqa: E501


        :return: The annotations of this V1Environment.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._annotations

    @annotations.setter
    def annotations(self, annotations):
        """Sets the annotations of this V1Environment.


        :param annotations: The annotations of this V1Environment.  # noqa: E501
        :type: dict(str, str)
        """

        self._annotations = annotations

    @property
    def node_selector(self):
        """Gets the node_selector of this V1Environment.  # noqa: E501


        :return: The node_selector of this V1Environment.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._node_selector

    @node_selector.setter
    def node_selector(self, node_selector):
        """Sets the node_selector of this V1Environment.


        :param node_selector: The node_selector of this V1Environment.  # noqa: E501
        :type: dict(str, str)
        """

        self._node_selector = node_selector

    @property
    def affinity(self):
        """Gets the affinity of this V1Environment.  # noqa: E501


        :return: The affinity of this V1Environment.  # noqa: E501
        :rtype: V1Affinity
        """
        return self._affinity

    @affinity.setter
    def affinity(self, affinity):
        """Sets the affinity of this V1Environment.


        :param affinity: The affinity of this V1Environment.  # noqa: E501
        :type: V1Affinity
        """

        self._affinity = affinity

    @property
    def tolerations(self):
        """Gets the tolerations of this V1Environment.  # noqa: E501

        Optional Tolerations to apply.  # noqa: E501

        :return: The tolerations of this V1Environment.  # noqa: E501
        :rtype: list[V1Toleration]
        """
        return self._tolerations

    @tolerations.setter
    def tolerations(self, tolerations):
        """Sets the tolerations of this V1Environment.

        Optional Tolerations to apply.  # noqa: E501

        :param tolerations: The tolerations of this V1Environment.  # noqa: E501
        :type: list[V1Toleration]
        """

        self._tolerations = tolerations

    @property
    def node_name(self):
        """Gets the node_name of this V1Environment.  # noqa: E501

        Optional NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.  # noqa: E501

        :return: The node_name of this V1Environment.  # noqa: E501
        :rtype: str
        """
        return self._node_name

    @node_name.setter
    def node_name(self, node_name):
        """Sets the node_name of this V1Environment.

        Optional NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.  # noqa: E501

        :param node_name: The node_name of this V1Environment.  # noqa: E501
        :type: str
        """

        self._node_name = node_name

    @property
    def service_account_name(self):
        """Gets the service_account_name of this V1Environment.  # noqa: E501


        :return: The service_account_name of this V1Environment.  # noqa: E501
        :rtype: str
        """
        return self._service_account_name

    @service_account_name.setter
    def service_account_name(self, service_account_name):
        """Sets the service_account_name of this V1Environment.


        :param service_account_name: The service_account_name of this V1Environment.  # noqa: E501
        :type: str
        """

        self._service_account_name = service_account_name

    @property
    def host_aliases(self):
        """Gets the host_aliases of this V1Environment.  # noqa: E501

        Optional HostAliases is an optional list of hosts and IPs that will be injected into the pod spec.  # noqa: E501

        :return: The host_aliases of this V1Environment.  # noqa: E501
        :rtype: list[V1HostAlias]
        """
        return self._host_aliases

    @host_aliases.setter
    def host_aliases(self, host_aliases):
        """Sets the host_aliases of this V1Environment.

        Optional HostAliases is an optional list of hosts and IPs that will be injected into the pod spec.  # noqa: E501

        :param host_aliases: The host_aliases of this V1Environment.  # noqa: E501
        :type: list[V1HostAlias]
        """

        self._host_aliases = host_aliases

    @property
    def security_context(self):
        """Gets the security_context of this V1Environment.  # noqa: E501


        :return: The security_context of this V1Environment.  # noqa: E501
        :rtype: V1PodSecurityContext
        """
        return self._security_context

    @security_context.setter
    def security_context(self, security_context):
        """Sets the security_context of this V1Environment.


        :param security_context: The security_context of this V1Environment.  # noqa: E501
        :type: V1PodSecurityContext
        """

        self._security_context = security_context

    @property
    def image_pull_secrets(self):
        """Gets the image_pull_secrets of this V1Environment.  # noqa: E501


        :return: The image_pull_secrets of this V1Environment.  # noqa: E501
        :rtype: list[str]
        """
        return self._image_pull_secrets

    @image_pull_secrets.setter
    def image_pull_secrets(self, image_pull_secrets):
        """Sets the image_pull_secrets of this V1Environment.


        :param image_pull_secrets: The image_pull_secrets of this V1Environment.  # noqa: E501
        :type: list[str]
        """

        self._image_pull_secrets = image_pull_secrets

    @property
    def host_network(self):
        """Gets the host_network of this V1Environment.  # noqa: E501

        Host networking requested for this workflow pod. Default to false.  # noqa: E501

        :return: The host_network of this V1Environment.  # noqa: E501
        :rtype: bool
        """
        return self._host_network

    @host_network.setter
    def host_network(self, host_network):
        """Sets the host_network of this V1Environment.

        Host networking requested for this workflow pod. Default to false.  # noqa: E501

        :param host_network: The host_network of this V1Environment.  # noqa: E501
        :type: bool
        """

        self._host_network = host_network

    @property
    def host_pid(self):
        """Gets the host_pid of this V1Environment.  # noqa: E501

        Use the host's pid namespace. Default to false.  # noqa: E501

        :return: The host_pid of this V1Environment.  # noqa: E501
        :rtype: str
        """
        return self._host_pid

    @host_pid.setter
    def host_pid(self, host_pid):
        """Sets the host_pid of this V1Environment.

        Use the host's pid namespace. Default to false.  # noqa: E501

        :param host_pid: The host_pid of this V1Environment.  # noqa: E501
        :type: str
        """

        self._host_pid = host_pid

    @property
    def dns_policy(self):
        """Gets the dns_policy of this V1Environment.  # noqa: E501

        Set DNS policy for the pod. Defaults to \"ClusterFirst\". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.  # noqa: E501

        :return: The dns_policy of this V1Environment.  # noqa: E501
        :rtype: str
        """
        return self._dns_policy

    @dns_policy.setter
    def dns_policy(self, dns_policy):
        """Sets the dns_policy of this V1Environment.

        Set DNS policy for the pod. Defaults to \"ClusterFirst\". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.  # noqa: E501

        :param dns_policy: The dns_policy of this V1Environment.  # noqa: E501
        :type: str
        """

        self._dns_policy = dns_policy

    @property
    def dns_config(self):
        """Gets the dns_config of this V1Environment.  # noqa: E501


        :return: The dns_config of this V1Environment.  # noqa: E501
        :rtype: V1PodDNSConfig
        """
        return self._dns_config

    @dns_config.setter
    def dns_config(self, dns_config):
        """Sets the dns_config of this V1Environment.


        :param dns_config: The dns_config of this V1Environment.  # noqa: E501
        :type: V1PodDNSConfig
        """

        self._dns_config = dns_config

    @property
    def scheduler_name(self):
        """Gets the scheduler_name of this V1Environment.  # noqa: E501


        :return: The scheduler_name of this V1Environment.  # noqa: E501
        :rtype: str
        """
        return self._scheduler_name

    @scheduler_name.setter
    def scheduler_name(self, scheduler_name):
        """Sets the scheduler_name of this V1Environment.


        :param scheduler_name: The scheduler_name of this V1Environment.  # noqa: E501
        :type: str
        """

        self._scheduler_name = scheduler_name

    @property
    def priority_class_name(self):
        """Gets the priority_class_name of this V1Environment.  # noqa: E501

        If specified, indicates the pod's priority. \"system-node-critical\" and \"system-cluster-critical\" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.  # noqa: E501

        :return: The priority_class_name of this V1Environment.  # noqa: E501
        :rtype: str
        """
        return self._priority_class_name

    @priority_class_name.setter
    def priority_class_name(self, priority_class_name):
        """Sets the priority_class_name of this V1Environment.

        If specified, indicates the pod's priority. \"system-node-critical\" and \"system-cluster-critical\" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.  # noqa: E501

        :param priority_class_name: The priority_class_name of this V1Environment.  # noqa: E501
        :type: str
        """

        self._priority_class_name = priority_class_name

    @property
    def priority(self):
        """Gets the priority of this V1Environment.  # noqa: E501

        The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.  # noqa: E501

        :return: The priority of this V1Environment.  # noqa: E501
        :rtype: int
        """
        return self._priority

    @priority.setter
    def priority(self, priority):
        """Sets the priority of this V1Environment.

        The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.  # noqa: E501

        :param priority: The priority of this V1Environment.  # noqa: E501
        :type: int
        """

        self._priority = priority

    @property
    def restart_policy(self):
        """Gets the restart_policy of this V1Environment.  # noqa: E501


        :return: The restart_policy of this V1Environment.  # noqa: E501
        :rtype: str
        """
        return self._restart_policy

    @restart_policy.setter
    def restart_policy(self, restart_policy):
        """Sets the restart_policy of this V1Environment.


        :param restart_policy: The restart_policy of this V1Environment.  # noqa: E501
        :type: str
        """

        self._restart_policy = restart_policy

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1Environment):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1Environment):
            return True

        return self.to_dict() != other.to_dict()
