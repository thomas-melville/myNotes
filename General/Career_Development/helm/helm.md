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

https://github.com/kubernetes/helm/blob/master/docs/charts.md

is the helm packaging format
A chart is a specific directory structure with files which is compressed for transfer over the network
contains yaml files
contains all the resource definitions necessary to run an application, tool / service inside a Kubernetes cluster: Pods, deployments, services, ingress, ...

directory name is the name of the chart

directory structure:

* Chart.yaml          # A YAML file containing information about the chart. **name in this chart must match directory name**
* LICENSE             # OPTIONAL: A plain text file containing the license for the chart
* README.md           # OPTIONAL: A human-readable README file
* requirements.yaml   # OPTIONAL: A YAML file listing dependencies for the chart
* values.yaml         # The default configuration values for this chart
* charts/             # A directory containing any charts upon which this chart depends.
* templates/          # A directory of templates that, when combined with values,
                      # will generate valid Kubernetes manifest files.
* templates/NOTES.txt # OPTIONAL: A plain text file containing short usage notes

#### versioning

helm charts follow the Semantic Versioning system

#### dependencies

two ways

1. in the charts/ sub directory
   * this is manual
2. in the requirements.yaml file
   * this is dynamic
   * you must do helm repo add, for each repo before trying to pull the dependencies
   * run **helm dependency update** to download all those charts as tgz files into the charts directory

#### templates

Kubernetes yaml files which contain markers for injecting values from the Helm values.yaml file.
The result, valid kubernetes files which can be used to create objects in kubernetes cluster

values for templates are supplied in three ways

* a values.yaml file in the chart which can contain default values
* a values.yaml file provided to the **helm install** command via the --values flag
  * these will override values in file inside chart
* the --set flag to the **helm install** command

Template files follow the standard conventions for Go templates

{{.Values.<specific attribute>}}

Templates can also pull from the Chart.yaml file at the root of the directory

{{.Chart.name}}

Templates can have required values:

{{ required "error message" .Values.<blah>}}

you can pipe the value returned into functions to format it

{{ .Values.<mine> | lower | quote}}

you can write your own functions and store them in the \_helpers.tpl file.
To use them you specify the template key word.

{{- template "my_function" }}

This places the output of the functions into the kubernetes yaml

If you want to do something with the output of the function use the include keyword

{{- include "my_function" $value | indent 2}}

This template passes $value to the function and then passes the output to indent

The tpl keyword allows you to process text and external files as go templates.

### Repository

The place where the charts can be collected and stored
any HTTP server which can serve yaml and tar files, and can answer GET requests can be used.
helm comes with a built in package server for testing, **helm serve**
the server requires an index.yaml file, that lists all the packages supplied by the repository

on the client side repos are managed with the **helm repo** command

### Release

instance of a chart running in a Kubernetes cluster
one chart can be installed many times in a cluster

## generating a chart

**helm create <my_chart>** to create a new chart folder with structure, basic information and simple deployment, service and ingress

once you have edited the chart directory use the **helm package** command to create the archive

use **helm lint <my_chart>** to check your chart for formating or missing info

**helm install** will then install the application in the cluster that kubectl is configured to talk to.

## hooks

Allows chart developers to intervene and perform operations at certain points in a releases lifecycle.
Declared as an annotation in the metadata section of the manifest.
The resource it's on contains what to do at that hook.
It could be a:
* Job
* ConfigMap
* Secret
Hooks can be weighted which determines execution order when there are multiple resources for a single hook

There are many hook points: https://github.com/helm/helm/blob/master/docs/charts_hooks.md#the-available-hooks

## plugins

helm supports plugins, on the client side.
Don't have to be written in Go.

### examples

* helm-diff
  * https://github.com/databus23/helm-diff
  * show the changes that helm will apply to a release
* helm-secrets
  * https://github.com/futuresimple/helm-secrets
  * integrates with helmfile
  * encrypted chart values files are decrypted locally
  * using a key from AWS,GCP KMS service!
