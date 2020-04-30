#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
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

    The version of the OpenAPI document: 1.0.84
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1EventModel(object):
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
    openapi_types = {"framework": "str", "path": "str", "spec": "object"}

    attribute_map = {"framework": "framework", "path": "path", "spec": "spec"}

    def __init__(
        self, framework=None, path=None, spec=None, local_vars_configuration=None
    ):  # noqa: E501
        """V1EventModel - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._framework = None
        self._path = None
        self._spec = None
        self.discriminator = None

        if framework is not None:
            self.framework = framework
        if path is not None:
            self.path = path
        if spec is not None:
            self.spec = spec

    @property
    def framework(self):
        """Gets the framework of this V1EventModel.  # noqa: E501


        :return: The framework of this V1EventModel.  # noqa: E501
        :rtype: str
        """
        return self._framework

    @framework.setter
    def framework(self, framework):
        """Sets the framework of this V1EventModel.


        :param framework: The framework of this V1EventModel.  # noqa: E501
        :type: str
        """

        self._framework = framework

    @property
    def path(self):
        """Gets the path of this V1EventModel.  # noqa: E501


        :return: The path of this V1EventModel.  # noqa: E501
        :rtype: str
        """
        return self._path

    @path.setter
    def path(self, path):
        """Sets the path of this V1EventModel.


        :param path: The path of this V1EventModel.  # noqa: E501
        :type: str
        """

        self._path = path

    @property
    def spec(self):
        """Gets the spec of this V1EventModel.  # noqa: E501


        :return: The spec of this V1EventModel.  # noqa: E501
        :rtype: object
        """
        return self._spec

    @spec.setter
    def spec(self, spec):
        """Sets the spec of this V1EventModel.


        :param spec: The spec of this V1EventModel.  # noqa: E501
        :type: object
        """

        self._spec = spec

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(
                    map(lambda x: x.to_dict() if hasattr(x, "to_dict") else x, value)
                )
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(
                    map(
                        lambda item: (item[0], item[1].to_dict())
                        if hasattr(item[1], "to_dict")
                        else item,
                        value.items(),
                    )
                )
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
        if not isinstance(other, V1EventModel):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1EventModel):
            return True

        return self.to_dict() != other.to_dict()
