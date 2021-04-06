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


class V1OrganizationMember(object):
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
        'user': 'str',
        'user_email': 'str',
        'role': 'str',
        'created_at': 'datetime',
        'updated_at': 'datetime'
    }

    attribute_map = {
        'user': 'user',
        'user_email': 'user_email',
        'role': 'role',
        'created_at': 'created_at',
        'updated_at': 'updated_at'
    }

    def __init__(self, user=None, user_email=None, role=None, created_at=None, updated_at=None, local_vars_configuration=None):  # noqa: E501
        """V1OrganizationMember - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._user = None
        self._user_email = None
        self._role = None
        self._created_at = None
        self._updated_at = None
        self.discriminator = None

        if user is not None:
            self.user = user
        if user_email is not None:
            self.user_email = user_email
        if role is not None:
            self.role = role
        if created_at is not None:
            self.created_at = created_at
        if updated_at is not None:
            self.updated_at = updated_at

    @property
    def user(self):
        """Gets the user of this V1OrganizationMember.  # noqa: E501


        :return: The user of this V1OrganizationMember.  # noqa: E501
        :rtype: str
        """
        return self._user

    @user.setter
    def user(self, user):
        """Sets the user of this V1OrganizationMember.


        :param user: The user of this V1OrganizationMember.  # noqa: E501
        :type: str
        """

        self._user = user

    @property
    def user_email(self):
        """Gets the user_email of this V1OrganizationMember.  # noqa: E501


        :return: The user_email of this V1OrganizationMember.  # noqa: E501
        :rtype: str
        """
        return self._user_email

    @user_email.setter
    def user_email(self, user_email):
        """Sets the user_email of this V1OrganizationMember.


        :param user_email: The user_email of this V1OrganizationMember.  # noqa: E501
        :type: str
        """

        self._user_email = user_email

    @property
    def role(self):
        """Gets the role of this V1OrganizationMember.  # noqa: E501


        :return: The role of this V1OrganizationMember.  # noqa: E501
        :rtype: str
        """
        return self._role

    @role.setter
    def role(self, role):
        """Sets the role of this V1OrganizationMember.


        :param role: The role of this V1OrganizationMember.  # noqa: E501
        :type: str
        """

        self._role = role

    @property
    def created_at(self):
        """Gets the created_at of this V1OrganizationMember.  # noqa: E501


        :return: The created_at of this V1OrganizationMember.  # noqa: E501
        :rtype: datetime
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at):
        """Sets the created_at of this V1OrganizationMember.


        :param created_at: The created_at of this V1OrganizationMember.  # noqa: E501
        :type: datetime
        """

        self._created_at = created_at

    @property
    def updated_at(self):
        """Gets the updated_at of this V1OrganizationMember.  # noqa: E501


        :return: The updated_at of this V1OrganizationMember.  # noqa: E501
        :rtype: datetime
        """
        return self._updated_at

    @updated_at.setter
    def updated_at(self, updated_at):
        """Sets the updated_at of this V1OrganizationMember.


        :param updated_at: The updated_at of this V1OrganizationMember.  # noqa: E501
        :type: datetime
        """

        self._updated_at = updated_at

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
        if not isinstance(other, V1OrganizationMember):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1OrganizationMember):
            return True

        return self.to_dict() != other.to_dict()
