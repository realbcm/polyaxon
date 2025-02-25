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

from marshmallow import EXCLUDE, fields, validate

from polyaxon.api import STATIC_V1
from polyaxon.contexts import paths as ctx_paths
from polyaxon.env_vars.keys import (
    EV_KEYS_ARCHIVE_ROOT,
    EV_KEYS_DNS_BACKEND,
    EV_KEYS_DNS_CUSTOM_CLUSTER,
    EV_KEYS_DNS_PREFIX,
    EV_KEYS_DNS_USE_RESOLVER,
    EV_KEYS_K8S_NAMESPACE,
    EV_KEYS_LOG_LEVEL,
    EV_KEYS_NGINX_INDENT_CHAR,
    EV_KEYS_NGINX_INDENT_WIDTH,
    EV_KEYS_NGINX_TIMEOUT,
    EV_KEYS_PROXY_API_HOST,
    EV_KEYS_PROXY_API_PORT,
    EV_KEYS_PROXY_API_TARGET_PORT,
    EV_KEYS_PROXY_API_USE_RESOLVER,
    EV_KEYS_PROXY_AUTH_ENABLED,
    EV_KEYS_PROXY_AUTH_EXTERNAL,
    EV_KEYS_PROXY_AUTH_USE_RESOLVER,
    EV_KEYS_PROXY_FORWARD_PROXY_HOST,
    EV_KEYS_PROXY_FORWARD_PROXY_KIND,
    EV_KEYS_PROXY_FORWARD_PROXY_PORT,
    EV_KEYS_PROXY_GATEWAY_HOST,
    EV_KEYS_PROXY_GATEWAY_PORT,
    EV_KEYS_PROXY_GATEWAY_TARGET_PORT,
    EV_KEYS_PROXY_HAS_FORWARD_PROXY,
    EV_KEYS_PROXY_NAMESPACES,
    EV_KEYS_PROXY_SERVICES_PORT,
    EV_KEYS_PROXY_SSL_ENABLED,
    EV_KEYS_PROXY_SSL_PATH,
    EV_KEYS_PROXY_STREAMS_HOST,
    EV_KEYS_PROXY_STREAMS_PORT,
    EV_KEYS_PROXY_STREAMS_TARGET_PORT,
    EV_KEYS_STATIC_ROOT,
    EV_KEYS_STATIC_URL,
    EV_KEYS_UI_ADMIN_ENABLED,
)
from polyaxon.schemas.base import BaseConfig, BaseSchema


class ProxiesSchema(BaseSchema):
    namespace = fields.Str(allow_none=True, data_key=EV_KEYS_K8S_NAMESPACE)
    namespaces = fields.List(
        fields.Int(), allow_none=True, data_key=EV_KEYS_PROXY_NAMESPACES
    )
    gateway_port = fields.Int(allow_none=True, data_key=EV_KEYS_PROXY_GATEWAY_PORT)
    gateway_target_port = fields.Int(
        allow_none=True, data_key=EV_KEYS_PROXY_GATEWAY_TARGET_PORT
    )
    gateway_host = fields.Str(allow_none=True, data_key=EV_KEYS_PROXY_GATEWAY_HOST)
    streams_port = fields.Int(allow_none=True, data_key=EV_KEYS_PROXY_STREAMS_PORT)
    streams_target_port = fields.Int(
        allow_none=True, data_key=EV_KEYS_PROXY_STREAMS_TARGET_PORT
    )
    streams_host = fields.Str(allow_none=True, data_key=EV_KEYS_PROXY_STREAMS_HOST)
    api_port = fields.Int(allow_none=True, data_key=EV_KEYS_PROXY_API_PORT)
    api_target_port = fields.Int(
        allow_none=True, data_key=EV_KEYS_PROXY_API_TARGET_PORT
    )
    api_host = fields.Str(allow_none=True, data_key=EV_KEYS_PROXY_API_HOST)
    api_use_resolver = fields.Bool(
        allow_none=True, data_key=EV_KEYS_PROXY_API_USE_RESOLVER
    )
    services_port = fields.Str(allow_none=True, data_key=EV_KEYS_PROXY_SERVICES_PORT)
    auth_enabled = fields.Bool(allow_none=True, data_key=EV_KEYS_PROXY_AUTH_ENABLED)
    auth_external = fields.Str(allow_none=True, data_key=EV_KEYS_PROXY_AUTH_EXTERNAL)
    auth_use_resolver = fields.Bool(
        allow_none=True, data_key=EV_KEYS_PROXY_AUTH_USE_RESOLVER
    )
    ssl_enabled = fields.Bool(allow_none=True, data_key=EV_KEYS_PROXY_SSL_ENABLED)
    ssl_path = fields.Str(allow_none=True, data_key=EV_KEYS_PROXY_SSL_PATH)
    dns_use_resolver = fields.Bool(allow_none=True, data_key=EV_KEYS_DNS_USE_RESOLVER)
    dns_custom_cluster = fields.Str(
        allow_none=True, data_key=EV_KEYS_DNS_CUSTOM_CLUSTER
    )
    dns_backend = fields.Str(allow_none=True, data_key=EV_KEYS_DNS_BACKEND)
    dns_prefix = fields.Str(allow_none=True, data_key=EV_KEYS_DNS_PREFIX)
    log_level = fields.Str(allow_none=True, data_key=EV_KEYS_LOG_LEVEL)
    nginx_timeout = fields.Int(allow_none=True, data_key=EV_KEYS_NGINX_TIMEOUT)
    nginx_indent_char = fields.Str(allow_none=True, data_key=EV_KEYS_NGINX_INDENT_CHAR)
    nginx_indent_width = fields.Int(
        allow_none=True, data_key=EV_KEYS_NGINX_INDENT_WIDTH
    )
    archive_root = fields.Str(allow_none=True, data_key=EV_KEYS_ARCHIVE_ROOT)
    static_root = fields.Str(allow_none=True, data_key=EV_KEYS_STATIC_ROOT)
    static_url = fields.Str(allow_none=True, data_key=EV_KEYS_STATIC_URL)
    ui_admin_enabled = fields.Bool(allow_none=True, data_key=EV_KEYS_UI_ADMIN_ENABLED)
    has_forward_proxy = fields.Bool(
        allow_none=True, data_key=EV_KEYS_PROXY_HAS_FORWARD_PROXY
    )
    forward_proxy_port = fields.Int(
        allow_none=True, data_key=EV_KEYS_PROXY_FORWARD_PROXY_PORT
    )
    forward_proxy_host = fields.Str(
        allow_none=True, data_key=EV_KEYS_PROXY_FORWARD_PROXY_HOST
    )
    forward_proxy_kind = fields.Str(
        allow_none=True,
        data_key=EV_KEYS_PROXY_FORWARD_PROXY_KIND,
        validate=validate.OneOf(["transparent", "connect"]),
    )

    @staticmethod
    def schema_config():
        return ProxiesConfig


class ProxiesConfig(BaseConfig):
    SCHEMA = ProxiesSchema
    IDENTIFIER = "proxies"
    UNKNOWN_BEHAVIOUR = EXCLUDE
    REDUCED_ATTRIBUTES = [
        EV_KEYS_PROXY_GATEWAY_PORT,
        EV_KEYS_PROXY_GATEWAY_TARGET_PORT,
        EV_KEYS_PROXY_GATEWAY_HOST,
        EV_KEYS_PROXY_NAMESPACES,
        EV_KEYS_PROXY_STREAMS_PORT,
        EV_KEYS_PROXY_STREAMS_TARGET_PORT,
        EV_KEYS_PROXY_STREAMS_HOST,
        EV_KEYS_PROXY_API_PORT,
        EV_KEYS_PROXY_API_TARGET_PORT,
        EV_KEYS_PROXY_API_HOST,
        EV_KEYS_PROXY_API_USE_RESOLVER,
        EV_KEYS_PROXY_SERVICES_PORT,
        EV_KEYS_PROXY_SSL_ENABLED,
        EV_KEYS_PROXY_SSL_PATH,
        EV_KEYS_PROXY_AUTH_ENABLED,
        EV_KEYS_PROXY_AUTH_EXTERNAL,
        EV_KEYS_PROXY_AUTH_USE_RESOLVER,
        EV_KEYS_DNS_USE_RESOLVER,
        EV_KEYS_DNS_CUSTOM_CLUSTER,
        EV_KEYS_DNS_BACKEND,
        EV_KEYS_DNS_PREFIX,
        EV_KEYS_NGINX_TIMEOUT,
        EV_KEYS_NGINX_INDENT_CHAR,
        EV_KEYS_NGINX_INDENT_WIDTH,
        EV_KEYS_K8S_NAMESPACE,
        EV_KEYS_LOG_LEVEL,
        EV_KEYS_ARCHIVE_ROOT,
        EV_KEYS_STATIC_ROOT,
        EV_KEYS_STATIC_URL,
        EV_KEYS_UI_ADMIN_ENABLED,
        EV_KEYS_PROXY_HAS_FORWARD_PROXY,
        EV_KEYS_PROXY_FORWARD_PROXY_PORT,
        EV_KEYS_PROXY_FORWARD_PROXY_HOST,
        EV_KEYS_PROXY_FORWARD_PROXY_KIND,
    ]

    def __init__(
        self,
        namespace=None,
        namespaces=None,
        auth_enabled=None,
        auth_external=None,
        auth_use_resolver=None,
        gateway_port=None,
        gateway_target_port=None,
        gateway_host=None,
        streams_port=None,
        streams_target_port=None,
        streams_host=None,
        api_port=None,
        api_target_port=None,
        api_host=None,
        api_use_resolver=None,
        services_port=None,
        dns_use_resolver=None,
        dns_custom_cluster=None,
        dns_backend=None,
        dns_prefix=None,
        nginx_timeout=None,
        nginx_indent_char=None,
        nginx_indent_width=None,
        log_level=None,
        ssl_enabled=None,
        ssl_path=None,
        archive_root=None,
        static_root=None,
        static_url=None,
        ui_admin_enabled=None,
        has_forward_proxy=None,
        forward_proxy_port=None,
        forward_proxy_host=None,
        forward_proxy_kind=None,
        **kwargs
    ):
        self.namespace = namespace
        self.namespaces = namespaces
        self.auth_enabled = auth_enabled or False
        self.auth_external = auth_external
        self.auth_use_resolver = auth_use_resolver or False
        self.gateway_port = gateway_port or self.default_port
        self.gateway_target_port = gateway_target_port or self.default_target_port
        self.gateway_host = gateway_host or "polyaxon-polyaxon-gateway"
        self.streams_port = streams_port or self.default_port
        self.streams_target_port = streams_target_port or self.default_target_port
        self.streams_host = streams_host or "polyaxon-polyaxon-streams"
        self.api_port = api_port or self.default_port
        self.api_target_port = api_target_port or self.default_target_port
        self.api_host = api_host or "polyaxon-polyaxon-api"
        self.api_use_resolver = api_use_resolver or False
        self.services_port = services_port or self.default_port
        self.dns_use_resolver = dns_use_resolver or False
        self.dns_custom_cluster = dns_custom_cluster or "cluster.local"
        self.dns_backend = dns_backend or "kube-dns"
        self.dns_prefix = dns_prefix
        self.nginx_timeout = nginx_timeout or 650
        self.nginx_indent_char = nginx_indent_char or " "
        self.nginx_indent_width = nginx_indent_width or 4
        self.ssl_enabled = ssl_enabled or False
        self.log_level = log_level or "warn"
        self.log_level = self.log_level.lower()
        self.ssl_path = ssl_path or "/etc/ssl/polyaxon"
        self.archive_root = archive_root or ctx_paths.CONTEXT_ARCHIVE_ROOT
        self.static_root = static_root or "/{}".format(STATIC_V1)
        self.static_url = static_url
        self.ui_admin_enabled = ui_admin_enabled
        self.has_forward_proxy = has_forward_proxy
        self.forward_proxy_port = forward_proxy_port
        self.forward_proxy_host = forward_proxy_host
        self.forward_proxy_kind = forward_proxy_kind

    @property
    def default_target_port(self):
        return 8000

    @property
    def default_port(self):
        return 80
