---
title: "Helm Reference"
sub_link: "agent/reference"
meta_title: "Polyaxon Agent Helm Reference - Polyaxon References"
meta_description: "Polyaxon Agent chart is a Helm chart for creating reproducible and maintainable deployments of Polyaxon with Kubernetes."
visibility: public
status: published
tags:
  - specification
  - polyaxon
  - yaml
  - helm
  - reference
  - kubernetes
sidebar: "setup"
---

Polyaxon Agent chart is a Helm chart for deploying Polyaxon Agent on Kubernetes.

## Introduction

This chart bootstraps a Polyaxon Agent deployment on
a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

This chart can be installed on a single node or multi-nodes cluster,
in which case you need to provide some volumes with `ReadWriteMany` or cloud buckets.

> **Tip**: The full list of the default [values.yaml](https://github.com/polyaxon/polyaxon-chart/blob/master/agent/values.yaml)

## Prerequisites

- Kubernetes
- helm

## deploymentChart

| Parameter                       | Description                  | Default
| --------------------------------| ---------------------------- | ----------------------------------------------------------
| `deploymentChart`               | The deployment chart to use  | `agent`


## deploymentVersion

| Parameter                       | Description                                                                                                                                             | Default
| --------------------------------| ------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------
| `deploymentVersion`             | The deployment version to use, this is important if you are using polyaxon-cli to avoid accidentally deploying/upgrading to a version without noticing  | `latest`


## Namespace

| Parameter                       | Description                                                                                          | Default
| --------------------------------| -----------------------------------------------------------------------------------------------------| ----------------------------------------------------------
| `namespace`                     | The namespace that will be used by Polyaxon to create operations and communicate with other services| `polyaxon`

## Gateway

The gateway is the service fronting the traffic to/from Polyaxon. By default it's deployed with `ClusterIP`.

You can use port forwarding to access the service on localhost:

```bash
kubectl port-forward -n polyaxon svc/agent-polyaxon-gateway 8000:80
```

Or using Polyaxon CLI

```bash
polyaxon port-forward -r agent
```

> **Note**: `[-r/--release-name]` could be different if you deployed with a different release name 

## Ingress and Gateway service

| Parameter                    | Description                                                                    | Default
| ---------------------------- | ------------------------------------------------------------------------------ | ----------------------------------------------------------
| `ingress.enabled`            | Use Kubernetes ingress                                                         | `true`
| `ingress.path`               | Kubernetes ingress path                                                        | `/`
| `ingress.pathType`           | Kubernetes ingress pathType (e.g. on GKE `ImplementationSpecific`)             | `Prefix`
| `ingress.hostName`           | Kubernetes ingress hostName                                                    | ``
| `ingress.tls`                | Use Ingress TLS (Secrets must be manually created in the namespace)            | `[]`
| `ingress.annotations`        | Ingress annotations                                                            | `{}`
| `gateway.service.type`       | Gateway Service type                                                           | `ClusterIP`
| `gateway.service.annotations`| Gateway Service annotations                                                    | `{}`


This chart provides support for an Ingress resource.

If you enable ingress, please set the gateway service type value to:
 * ClusterIP or NodePort
 * NodePort or LoadBalancer on GKE

Note: using TLS requires either:
 - a preconfigured secret with the TLS secrets in it
 - or the user of [cert-manager](https://github.com/helm/charts/tree/master/stable/cert-manager) to auto request certs from let's encrypt and store them in a secret.

It's also possible to use a service like [externalDNS](https://github.com/helm/charts/tree/master/stable/external-dns) to auto create the DNS entry for the polyaxon API service.

### Ingress Annotations

Example annotations:

```yaml
ingress.kubernetes.io/ssl-redirect: "false"
ingress.kubernetes.io/rewrite-target: /
ingress.kubernetes.io/add-base-url: "true"
ingress.kubernetes.io/proxy-connect-timeout: "600"
ingress.kubernetes.io/proxy-read-timeout: "600"
ingress.kubernetes.io/proxy-send-timeout: "600"
ingress.kubernetes.io/send-timeout: "600"
ingress.kubernetes.io/proxy-body-size: 4G
```

## Securing API server with TLS

If you have your own certificate you can make a new secret with the `tls.crt` and the `tls.key`,
then set the secret name in the values file.

### Automating TLS certificate creation and DNS setup

To automate the creation and registration of new domain name you can use the following services:

* [cert-manager](https://github.com/helm/charts/tree/master/stable/cert-manager)
* [externalDNS](https://github.com/helm/charts/tree/master/stable/external-dns) (Route53 / Google CloudDNS)

once installed, you can set the values for `ingress.tls`:

```yaml
ingress:
  enabled: true
  hostName: polyaxon.acme.com
  tls:
    - secretName: polyaxon.acme-tls
      hosts:
        - polyaxon.acme.com
```

TLS can have more than one host.

In order to get the domain registration to work you need to set the value of `api.service.annotations`
to the annotation needed for your domain:
i.e

```yaml
annotations:
  domainName: polyaxon.my.domain.com
```

### GKE ingress

 1. Create a static IP using `gcloud compute addresses create IP_NAME --global`
 3. Make sure to create a DNS record to point to the IP address (`xx.xx.xx.xx`) created in step 1 to a subdomain (`polyaxon`) of type record `A`.
 2. Deploy a managed certificate with `kubectl apply -f polyaxon-certificate.yaml`, the content should be:
    ```yaml
    apiVersion: networking.gke.io/v1
    kind: ManagedCertificate
    metadata:
      name: INGRESS_TLS_CERT_NAME
    spec:
      domains:
        - polyaxon.acme.com
    ``` 
 3. The gateway must be changed to use a node port service type in GKE for the ingress to function correctly:
    ```yaml
    gateway:
      scheme: https
      service:
        type: NodePort
    ```
 4. The configuration for the GKE ingress should look like:
    ```yaml
    ingress:
      enabled: true
      hostName: polyaxon.acme.com
      pathType: ImplementationSpecific
      path: /*
      annotations:
        kubernetes.io/ingress.class: "gce"
        kubernetes.io/ingress.global-static-ip-name: IP_NAME
        networking.gke.io/managed-certificates: INGRESS_TLS_CERT_NAME
        kubernetes.io/ingress.allow-http: "false"
        ingress.kubernetes.io/rewrite-target: "/"
    ```

### NGINX ingress

To use Https in Polyaxon on Kubernetes you can set an ingress-nginx for cluster running on Kubernetes.

Polyaxon's helm chart comes with an ingress resource that you can use with an ingress controller where you should use TLS so that all traffic will be served over HTTPS.

 1. Create a TLS secret that contains your TLS certificate and private key.

    ```bash
    kubectl create secret tls polyaxon-tls --key $PATH_TO_KEY --cert $PATH_TO_CERT
    ```


 2. Add the tls configuration to Polyaxon's Ingress values. (**Do not use CluserIP on GKE**)

    ```yaml
    gateway:
      service:
        type: ClusterIP
    ingress:
      enabled: true
      hostName: polyaxon.acme.com
      tls:
      - secretName: polyaxon.acme-tls
        hosts:
          - polyaxon.acme.com
    ```

    For more information visit the [Nginx Ingress Integration](/integrations/nginx/)

## dns

| Parameter | Description                                                             | Default
| ----------| ------------------------------------------------------------------------| ----------------------------------------------------------
| `dns`     | DNS configuration for cluster running with custom dns setup             | `{}`


Several users deploy Polyaxon on Kubernetes cluster created with [kubespray](https://github.com/kubernetes-sigs/kubespray),
and by default Kubespray creates Kubernetes with CoreDNS as a cluster DNS server.

### Update DNS backend

Although we could provide logic to detect the DNS used in the cluster, this would require cluster wide RBAC that we think it's unnecessary.
The default DNS backend used by Polyaxon is KubeDNS, to set it to a different DNS, you can provide this value in your Polyaxon's deployment config:

```yaml
dns:
  backend: "coredns"
```

### Update DNS prefix to different system

Since the DNS service is generally deployed on `kube-system` namespace, the default DNS prefix is `kube-dns.kube-system` or `coredns.kube-system` if you update the previous option.

You can also provide the complete DNS prefix, and not use the DNS backend options:

```yaml
dns:
  prefix: "kube-dns.other-kube-system"
```

### Update the DNS prefix for OpenShift

OpenShift has a different DNS configuration, the default prefix is:

```yaml
dns:
  prefix: "dns-default.openshift-dns"
```

### Update DNS cluster

The default dns cluster used in Polyaxon to resolve routes is `cluster.local`, you can provide a Custom Cluster DNS, by setting:

```yaml
dns:
  customCluster: "custom.cluster.name"
```


## Proxy

This section allows admins to set the configuration for a corporate proxy.


| Parameter          | Description                                                       | Default
| ------------------ | ----------------------------------------------------------------- | ----------------------------------------------------------
| `proxy.enabled`    | To use enable proxy section                                       | `false`
| `proxy.httpsProxy` | To set the environment variables `https_proxy` and `HTTPS_PROXY`  | ``
| `proxy.httpProxy`  | To set the environment variables `http_proxy` and `HTTP_PROXY`    | ``
| `proxy.noProxy`    | To set the environment variables `no_proxy` and `NO_PROXY`        | ``
| `proxy.host`       | Proxy host or ip value                                            | ``
| `proxy.port`       | Proxy port value                                                  | ``
| `proxy.kind`       | Proxy kind (transparent or connect)                               | `connect`

Example:

```yaml
proxy:
  enabled: false
  useInOps: true
  httpProxy: "12.12.12.12:8080"
  httpsProxy: "12.12.12.12:8080"
  noProxy: "12.12.12.12:8080"
  host: "12.12.12.12"
  port: 8080
  kind: transparent or connect
```

> **Note**: make sure that the in-cluster traffic, including access to the `streams` and `gateway` services is accessible by name `NAMESPACE-RELEASE-streams` and `NAMESPACE-RELEASE-gateway`, is specified in the `noProxy` field.

## Auth

When you deploy an agent all traffic will be checked by the control plane to ensure that:

 * Users are allowed to view the content.
 * Users are authenticated to access private projects and their content.
 * That any service: Notebooks, Tensorboards, custom dashboards and apps, can be accessed following the team roles and access control configured.


| Parameter         | Description                                                   | Default
| ----------------- | ------------------------------------------------------------- | ----------------------------------------------------------
| `auth.enabled`    | To use Polyaxon auth system                                   | `true`
| `auth.useResolver`| To resolve and authenticate all traffic for managed services  | `true`


## Time zone

To set a different time zone for the agent.
you can can provide a [valid time zone value](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)


| Parameter | Description         | Default
| ----------| --------------------| ----------------------------------------------------------
| `timezone`| The timezone to use | `UTC`


## Node manipulation

| Parameter                          | Description                                                  | Default
| -----------------------------------| -------------------------------------------------------------| ----------------------------------------------------------
| `nodeSelector`                     | Node selector for core pod assignment                        | `{}`
| `tolerations`                      | Tolerations for core pod assignment                          | `[]`
| `affinity`                         | Affinity for core pod assignment                             | Please check the values


### Node Selectors

Polyaxon comes with a couple of node selectors options to assign pods to nodes for polyaxon's core platform

By providing these values, or some of them,
you can constrain the pods belonging to that category to only run on
particular nodes or to prefer to run on particular nodes.

```yaml
nodeSelector:
  # Global nodeSelector
  ...

gateway:
  nodeSelector:
    # If null, the global nodeSelector will be used
    ...

streams:
  nodeSelector:
    # If null, the global nodeSelector will be used
    ...

operator:
  nodeSelector:
    # If null, the global nodeSelector will be used
    ...

agent:
  nodeSelector:
    # If null, the global nodeSelector will be used
    ...
cleaner:
  nodeSelector:
    # If null, the global nodeSelector will be used
    ...
```

### Tolerations

If one or more taints are applied to a node, and you want to make sure some pods should not deploy on it.

```yaml
tolerations:
  # Global tolerations
  ...

gateway:
  tolerations:
    # If null, the global tolerations will be used
    ...

streams:
  tolerations:
    # If null, the global tolerations will be used
    ...

operator:
  tolerations:
    # If null, the global tolerations will be used
    ...

agent:
  tolerations:
    # If null, the global tolerations will be used
    ...
cleaner:
  tolerations:
    # If null, the global tolerations will be used
    ...
```

### Affinity

It allows you to constrain which nodes your pod is eligible to schedule on, based on the node's labels.
Polyaxon has a default `Affinity` values for the agent's core components to ensure that they deploy on the same node.

Polyaxon's default affinity:

```yaml
affinity:
  # Global affinity
  podAffinity:
    preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 100
        podAffinityTerm:
          labelSelector:
            matchExpressions:
            - key: type
              operator: In
              values:
              - "polyaxon-core"
          topologyKey: "kubernetes.io/hostname"

gateway:
  affinity:
    # If null, the global affinity will be used
    ...

streams:
  affinity:
    # If null, the global affinity will be used
    ...

operator:
  affinity:
    # If null, the global affinity will be used
    ...

agent:
  affinity:
    # If null, the global affinity will be used
    ...
cleaner:
  affinity:
    # If null, the global affinity will be used
    ...

```

## Control Plane

This is the control plane that the agent will connect to. The control plane can be Polyaxon Cloud or a Polyaxon EE managed by the user.

```yaml
externalServices:
  api:
    host:
    port:
```

## Agent configuration

Set the agent token and instance value:

```yaml
agent:
  instance: acme.agents.UUID
  token: "TOKEN_VALUE"
```

## Security context

| Parameter                               | Description                                                                                         | Default
| --------------------------------------- | --------------------------------------------------------------------------------------------------- | ----------------------------------------------------------
| `enabled`                               | enable security context                                                                             | `false`
| `runAsUser`                             | security context UID                                                                                | `2222`
| `runAsGroup`                            | security context GID                                                                                | `2222`
| `fsGroup`                               | security context FS                                                                                 | `2222`
| `allowPrivilegeEscalation`              | Controls whether a process can gain more privileges than its parent process                         | `false`
| `runAsNonRoot`                          | Indicates that the container must run as a non-root user                                            | `true`
| `fsGroupChangePolicy`                   | defines behavior of changing ownership and permission of the volume before being exposed inside Pod |

Polyaxon runs all containers as root by default, this configuration is often fine for several deployment,
however, in some use cases it can expose a compliance issue for some teams.

Polyaxon provides a simple way to enable a security context for all core components.

Default configuration:

```yaml
securityContext:
  enabled: false
  runAsUser: 2222
  runAsGroup: 2222
  fsGroup: 2222
  allowPrivilegeEscalation: false
  runAsNonRoot: true
  fsGroupChangePolicy:
```

### Enable security context

```yaml
securityContext:
  enabled: true
```

Or enable with custom UID/GID other than 2222/2222:

```yaml
securityContext:
  enabled: true
  user: 1111
  group: 1111
```

Define behavior of changing ownership and permission of the volume before being exposed inside Pod:

```yaml
securityContext:
  enabled: true
  fsGroupChangePolicy: OnRootMismatch
```

This will enable a security context to run all containers using a UID/GID == 1111/1111.

> **N.B.** If you are using a host path or a volume for the artifacts store, make sure to allow the UID/GID to access it.

## Connections

```yaml
artifactsStore: {}
connections: []
```

You need to configure the connections to authorize for the agent. Please check [connections section](/docs/setup/connections/).

Each `Agent` is responsible for the artifacts store and connections it manages on the namespace where it's deployed.
