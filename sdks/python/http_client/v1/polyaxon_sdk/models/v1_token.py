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

    The version of the OpenAPI document: 1.5.0
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1Token(object):
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
        'uuid': 'str',
        'key': 'str',
        'name': 'str',
        'scopes': 'list[str]',
        'services': 'list[str]',
        'started_at': 'datetime',
        'expires_at': 'datetime',
        'created_at': 'datetime',
        'updated_at': 'datetime',
        'expiration': 'int'
    }

    attribute_map = {
        'uuid': 'uuid',
        'key': 'key',
        'name': 'name',
        'scopes': 'scopes',
        'services': 'services',
        'started_at': 'started_at',
        'expires_at': 'expires_at',
        'created_at': 'created_at',
        'updated_at': 'updated_at',
        'expiration': 'expiration'
    }

    def __init__(self, uuid=None, key=None, name=None, scopes=None, services=None, started_at=None, expires_at=None, created_at=None, updated_at=None, expiration=None, local_vars_configuration=None):  # noqa: E501
        """V1Token - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._uuid = None
        self._key = None
        self._name = None
        self._scopes = None
        self._services = None
        self._started_at = None
        self._expires_at = None
        self._created_at = None
        self._updated_at = None
        self._expiration = None
        self.discriminator = None

        if uuid is not None:
            self.uuid = uuid
        if key is not None:
            self.key = key
        if name is not None:
            self.name = name
        if scopes is not None:
            self.scopes = scopes
        if services is not None:
            self.services = services
        if started_at is not None:
            self.started_at = started_at
        if expires_at is not None:
            self.expires_at = expires_at
        if created_at is not None:
            self.created_at = created_at
        if updated_at is not None:
            self.updated_at = updated_at
        if expiration is not None:
            self.expiration = expiration

    @property
    def uuid(self):
        """Gets the uuid of this V1Token.  # noqa: E501


        :return: The uuid of this V1Token.  # noqa: E501
        :rtype: str
        """
        return self._uuid

    @uuid.setter
    def uuid(self, uuid):
        """Sets the uuid of this V1Token.


        :param uuid: The uuid of this V1Token.  # noqa: E501
        :type: str
        """

        self._uuid = uuid

    @property
    def key(self):
        """Gets the key of this V1Token.  # noqa: E501


        :return: The key of this V1Token.  # noqa: E501
        :rtype: str
        """
        return self._key

    @key.setter
    def key(self, key):
        """Sets the key of this V1Token.


        :param key: The key of this V1Token.  # noqa: E501
        :type: str
        """

        self._key = key

    @property
    def name(self):
        """Gets the name of this V1Token.  # noqa: E501


        :return: The name of this V1Token.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this V1Token.


        :param name: The name of this V1Token.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def scopes(self):
        """Gets the scopes of this V1Token.  # noqa: E501


        :return: The scopes of this V1Token.  # noqa: E501
        :rtype: list[str]
        """
        return self._scopes

    @scopes.setter
    def scopes(self, scopes):
        """Sets the scopes of this V1Token.


        :param scopes: The scopes of this V1Token.  # noqa: E501
        :type: list[str]
        """

        self._scopes = scopes

    @property
    def services(self):
        """Gets the services of this V1Token.  # noqa: E501


        :return: The services of this V1Token.  # noqa: E501
        :rtype: list[str]
        """
        return self._services

    @services.setter
    def services(self, services):
        """Sets the services of this V1Token.


        :param services: The services of this V1Token.  # noqa: E501
        :type: list[str]
        """

        self._services = services

    @property
    def started_at(self):
        """Gets the started_at of this V1Token.  # noqa: E501


        :return: The started_at of this V1Token.  # noqa: E501
        :rtype: datetime
        """
        return self._started_at

    @started_at.setter
    def started_at(self, started_at):
        """Sets the started_at of this V1Token.


        :param started_at: The started_at of this V1Token.  # noqa: E501
        :type: datetime
        """

        self._started_at = started_at

    @property
    def expires_at(self):
        """Gets the expires_at of this V1Token.  # noqa: E501


        :return: The expires_at of this V1Token.  # noqa: E501
        :rtype: datetime
        """
        return self._expires_at

    @expires_at.setter
    def expires_at(self, expires_at):
        """Sets the expires_at of this V1Token.


        :param expires_at: The expires_at of this V1Token.  # noqa: E501
        :type: datetime
        """

        self._expires_at = expires_at

    @property
    def created_at(self):
        """Gets the created_at of this V1Token.  # noqa: E501


        :return: The created_at of this V1Token.  # noqa: E501
        :rtype: datetime
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at):
        """Sets the created_at of this V1Token.


        :param created_at: The created_at of this V1Token.  # noqa: E501
        :type: datetime
        """

        self._created_at = created_at

    @property
    def updated_at(self):
        """Gets the updated_at of this V1Token.  # noqa: E501


        :return: The updated_at of this V1Token.  # noqa: E501
        :rtype: datetime
        """
        return self._updated_at

    @updated_at.setter
    def updated_at(self, updated_at):
        """Sets the updated_at of this V1Token.


        :param updated_at: The updated_at of this V1Token.  # noqa: E501
        :type: datetime
        """

        self._updated_at = updated_at

    @property
    def expiration(self):
        """Gets the expiration of this V1Token.  # noqa: E501


        :return: The expiration of this V1Token.  # noqa: E501
        :rtype: int
        """
        return self._expiration

    @expiration.setter
    def expiration(self, expiration):
        """Sets the expiration of this V1Token.


        :param expiration: The expiration of this V1Token.  # noqa: E501
        :type: int
        """

        self._expiration = expiration

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
        if not isinstance(other, V1Token):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1Token):
            return True

        return self.to_dict() != other.to_dict()
