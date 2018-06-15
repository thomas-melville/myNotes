# helm

Package manager for Kubernetes
helps you manage Kubernetes applications
define, install, upgrade applications

information is in a chart

download a chart and deploy it, this results in a container in the cluster

it has two parts: client (helm) and a server (tiller)

charts can be downloaded from the central repo.

## concepts

Helm installs **charts** into Kubernetes, creating a new **release** for each installation. And to find new charts, you can search Helm chart **repositories**.

### Chart

is a helm package
archive file which contains yaml files
contains all the resource definitions necessary to run an application, tool / service inside a Kubernetes cluster

### Repository

The place where the charts can be collected and stored

### Release

instance of a chart running in a Kubernetes cluster
one chart can be installed many times in a cluster