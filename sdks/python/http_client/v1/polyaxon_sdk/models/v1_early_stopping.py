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

    The version of the OpenAPI document: 1.5.1
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1EarlyStopping(object):
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
        'median': 'V1MedianStoppingPolicy',
        'diff': 'V1DiffStoppingPolicy',
        'truncation': 'V1TruncationStoppingPolicy',
        'metric': 'V1MetricEarlyStopping',
        'failure': 'V1FailureEarlyStopping'
    }

    attribute_map = {
        'median': 'median',
        'diff': 'diff',
        'truncation': 'truncation',
        'metric': 'metric',
        'failure': 'failure'
    }

    def __init__(self, median=None, diff=None, truncation=None, metric=None, failure=None, local_vars_configuration=None):  # noqa: E501
        """V1EarlyStopping - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._median = None
        self._diff = None
        self._truncation = None
        self._metric = None
        self._failure = None
        self.discriminator = None

        if median is not None:
            self.median = median
        if diff is not None:
            self.diff = diff
        if truncation is not None:
            self.truncation = truncation
        if metric is not None:
            self.metric = metric
        if failure is not None:
            self.failure = failure

    @property
    def median(self):
        """Gets the median of this V1EarlyStopping.  # noqa: E501


        :return: The median of this V1EarlyStopping.  # noqa: E501
        :rtype: V1MedianStoppingPolicy
        """
        return self._median

    @median.setter
    def median(self, median):
        """Sets the median of this V1EarlyStopping.


        :param median: The median of this V1EarlyStopping.  # noqa: E501
        :type: V1MedianStoppingPolicy
        """

        self._median = median

    @property
    def diff(self):
        """Gets the diff of this V1EarlyStopping.  # noqa: E501


        :return: The diff of this V1EarlyStopping.  # noqa: E501
        :rtype: V1DiffStoppingPolicy
        """
        return self._diff

    @diff.setter
    def diff(self, diff):
        """Sets the diff of this V1EarlyStopping.


        :param diff: The diff of this V1EarlyStopping.  # noqa: E501
        :type: V1DiffStoppingPolicy
        """

        self._diff = diff

    @property
    def truncation(self):
        """Gets the truncation of this V1EarlyStopping.  # noqa: E501


        :return: The truncation of this V1EarlyStopping.  # noqa: E501
        :rtype: V1TruncationStoppingPolicy
        """
        return self._truncation

    @truncation.setter
    def truncation(self, truncation):
        """Sets the truncation of this V1EarlyStopping.


        :param truncation: The truncation of this V1EarlyStopping.  # noqa: E501
        :type: V1TruncationStoppingPolicy
        """

        self._truncation = truncation

    @property
    def metric(self):
        """Gets the metric of this V1EarlyStopping.  # noqa: E501


        :return: The metric of this V1EarlyStopping.  # noqa: E501
        :rtype: V1MetricEarlyStopping
        """
        return self._metric

    @metric.setter
    def metric(self, metric):
        """Sets the metric of this V1EarlyStopping.


        :param metric: The metric of this V1EarlyStopping.  # noqa: E501
        :type: V1MetricEarlyStopping
        """

        self._metric = metric

    @property
    def failure(self):
        """Gets the failure of this V1EarlyStopping.  # noqa: E501


        :return: The failure of this V1EarlyStopping.  # noqa: E501
        :rtype: V1FailureEarlyStopping
        """
        return self._failure

    @failure.setter
    def failure(self, failure):
        """Sets the failure of this V1EarlyStopping.


        :param failure: The failure of this V1EarlyStopping.  # noqa: E501
        :type: V1FailureEarlyStopping
        """

        self._failure = failure

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
        if not isinstance(other, V1EarlyStopping):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1EarlyStopping):
            return True

        return self.to_dict() != other.to_dict()
