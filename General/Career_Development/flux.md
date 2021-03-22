# Flux

CNCF Technology Radar, it's in the adopt space, beside Helm!

GitOps!
first to clearly define GitOps practices

1. describe your system declaratively
2. keep configuration under source control
3. approved changes can be automatically applied to the system
4. use software agents to reconcile and ensure correctness or alert for drift

benefits

all have adopted git
access control
aufitable history
dirft correction
clear boundaries between dev & k8s

Flux aims to provide a complete CD platform on top of k8s, supporting all the common practices and tooling in the field, e.g. kustomize, helm ,...


Flux v1, monolith
  sync single git repository to local cluster

Flux v2, k9s native microservices
  sync multiple git repositories to local or remote clusters ++++

  several controllers
    flux2
    source-controller
    kustomize
    helm

interesting flow

put your declarative specs in version control (helm chart)
point Flux at it (flux is running in the cluster)
Flux will then sync the state every time the code changes


## Who is Flux for?

Cluster Operators who automate provision and configuration of clusters
Platform Engineers who build continuous delivery for developer teams
App Developers who rely on continuous delivery to get their code live


## Source Controller

can pull information from a number of sources
  git
  s3 compliant buckets
  Helm repos

sync information between git repo and cluster.
It does not apply information to cluster

## Helm Controller

specialized for Helm Operations
  roll back if helm tests fail

Helm Controller has a depends on feature
  depend on another chart instead of a large integration chart

can install from any bucket, S3, ...

add helm repo without restarting

Can manage multiple clusters from a single management cluster
  using the ClusterAPI

## Kustomize Controller

Kustomization objects link back to folders in the sync'd git repo
applies manifests to cluster

## using flux

flux cli to bootstrap a cluster
  ends up in CRDs & Deployments installed on the cluster
  this also links the controllers to a git repository
once that's done you can drop manifests into the git repo and they will be applied to the cluster

They recommend to have manifest completely rendered before going into git repository
CRDs:
  GitRepositories
  Kustomization
  HelmRepository
  HelmRelease

HelmRepository points to a helm registry somewhere
HelmRelease depends on a HelmRepository and points to a chart in the helm registry
