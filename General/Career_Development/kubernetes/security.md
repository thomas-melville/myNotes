# security

All communication with the API server is over HTTPS (TLS & SSL).
This is automatically configured when using kubeadm.
To perform any action in a K8S cluster, you need to access the API and go through 3 main steps:

* Authentication
* Authorization (ABAC or RBAC)
* Admission Control

The action can come from a human user or a pod (service account)

## Transport Security

Communication with the cluster is on port 443.
The API server presents a certificate. It is often self-signed so the user config typically contains the root certificate for the APIs server cert.

## Authentication

*Authentication* is handled by any Authentication modules that have been configured.

* is done with certificates, tokens or basic Authentication (user & pass) or JWT tokens ( used for service accounts)
* Users are not created by the API, managed by an external system.
* Service accounts are used by processes (inside Pods) to access the API (https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/)

There are two more advanced Authentication mechanisms:

* Webhooks, which can be used to verify bearer tokens
* OpenID provider, connection with an external provider

The type of Authentication used is defined in the kube-apiserver startup options.
One or more Authenticator modules are used: x509 client certs, static tokens, bearer / bootstrap tokens, static password file, service account and OpenID connect tokens.
Each is tried until successful, the order is not guaranteed.

https://kubernetes.io/docs/reference/access-authn-authz/authentication/

## Authorization

*Authorization* is checked against existing policies to see if the user has the permissions to perform the requested actions.

There are 3 main Authorization modes and two global deny/allow settings:

* ABAC (Attribute Based Access Control)
  * first Authorization model in K8S
* RBAC (Role Based Access Control)
  * All resources are modelled API objects in K8S
  * They also belong to API groups
  * These resources allow CRUD operations
  * Operations are called verbs inside YAML files
  * Rules are operations which can act upon an API group
  * Roles are a group of rules which affect a single namespace (unless it's a ClusterRole)
  * Each operation can act upon one of 3 subjects
    * User Accounts
    * Service accounts
    * Groups
  * RBAC is then writing rules to allow/deny operations by users, roles or groups on resources
* Webhook

Each mode implements policies to allow requests.

https://kubernetes.io/docs/reference/access-authn-authz/authorization/

## Admission Control

*Admission Control* checks the actual content of the objects being created and validate them

* pieces of software which can access the content of the objects being created by the requests.
* They can validate, modify and deny requests.
* They are required for certain features to work properly.
* Here are some example Admission Controllers:
  * Initializers
  * NamespaceLifecycle
  * ServiceAccount
  * DefaultStorageClass

https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/

## Security Context

Pods and containers can be given specific security constraints to limit what processes running in containers can do.
This limitation is called a security context.
Can be for the entire pod or per container
Defined in the resource manifest, i.e. the yaml.

```YAML

...
spec:
  securityContext:
    runAsNonRoot: true
```

To automate the enforcement of security contexts you can define PodSecurityPolicies
Cluster level rules which define what a pod can do.

## Network Security Policies

By default, all pods can reach each other, all ingress & egress traffic is allowed.
You can isolate portions of the network in a cluster by defining a NetworkPolicy.
You can narrow down traffic to a namespace, you can add a label seletore to isolate pods in a namespace.
This is achieved using a plugin, some providers are:
* Calico
* Romana
* Cilium
* Kube-router
* WeaveNet
