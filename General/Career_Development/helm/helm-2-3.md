# helm 2 -> 3

Helm has moved to the CNCF

Provide an embedded scripting engine based on lua
Which means changing from YAML to code for charts
This enables the possibility of reusable libraries for defining charts
YAML and Lua will co-exist in Helm 3
It will be possible to dynamically update YAML charts using Lua

## migrating

https://github.com/helm/helm/blob/739afa7c97a4a9c08148ac3d0d9d00283ef5c6b2/docs/v2_v3_migration.md

it doesn't look like there'll be any blocking issues.
We start running helm 3, helm 2 charts can still be deployed using helm 3.
Slowly phase out helm 2 charts.
The easiest thing for us it to re-deploy on the cluster
If we have to do an upgrade and keep release history then it becomes complex.

## Getting rid of Tiller

With the introduction of RBAC, locking down Tiller in a production environment became more difficult.
Due to the vast number of security policies, the default was to be permissive.
This made it more difficult for operators in a multi tenant cluster.
With tiller gone the security model is simplified. It delgates it all to K8s.

It's goals can now be achieved in Kubernetes, see the release management section.

## chart repositories

issues which arose:

* hard time abstracting most of the security implementation details in a production environment.
* provenance tools used for signing and verifying charts are optional.
* In multi tenant scenarios the same chart can be uploaded by costing extra space.
* single index file for search, metadata info and fetch has made it difficult for security in multi-tenant scenarios.

Helm charts can be hosted on a Docker Registry V2 instance thanks to the efforts of the OCI.
The work is still experimental!

## release management

1. using Custom Resource Definitions to store information
   * This improves security as it is all delegated to Kubernetes
2. with two new records for storing information about releases
   * Release and ReleaseVersion

Release represents an instance of an application, and has a number of associated ReleaseVersion secrets.
ReleaseVersion secret represents an applications desired state at a particular instance of time.

The Release object points to one ReleaseVersion secret at a time.

`helm install` creates a release object and ReleaseVersion secret.
`helm upgrade` requires an existing release object, which it may modify, and creates a new ReleaseVersion secret which contains the new values and rendered manifest.
`helm rollback` requires an existing release object and updates it to point to the previous secret
  **Could this have an impact on our rollback setup?**

The release data is stored in the same namespace as the releases destination.
This means the release name only needs to be unique within a namespace.

## chart dependencies

charts packaged with helm 2 can be deployed using helm 3.
However, to continue developing the chart there are some updates required.

requirements.yaml is no more, all the requirements information is now in Chart.yaml

charts are still downloaded and placed in the charts directory for packaging

## library charts

a chart which is shared by other charts, but does not create any release artifacts of its own.
This allows users to re-use and share snippets of code across many charts, avoiding redundancy and keeping charts DRY.
declared in the dependencies section, installed and managed like any other chart.

and there's lots more: https://github.com/helm/community/blob/master/helm-v3/000-helm-v3.md

alpha release will be out soon
