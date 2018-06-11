# kubernetes

http://www.baeldung.com/spring-boot-minikube
https://labs.play-with-k8s.com/#

1. open source system for the automated deployment, scaling and management of containerized apps.
2. system for orchestration of containerized apps across a cluster of nodes, including networking and storage infrastructure

Some important features are:

* Resource scheduling
* * ensures pods are distributed optimally across nodes
* Auto-scaling
* * with increasing load, the cluster can dynamically allocate additional nodes
* * and deploy new pods on them.
* Self-healing
* * cluster supervises containers and restarts them
* * if required, based on defined policies
* Service-discovery
* * Pods and Services are registered and published via DNS
* Rolling updates/rollbacks
* * support rolling updates based on sequential deployment of Pods and containers
* Secret/config mgmt
* * supports secure handling of sensitive data like passwords and API keys
* Storage Orchestration
* * several 3rd party solutions are supported
* * which can be used as external volumes to persist data

## understanding Kubernetes

The **master** maintains the desired state of the **cluster**
There can be more than one master for availability.
when we interact with our cluster, by using **kubectl** for example we are communicating with the clusters master
it's the brains, a collection of services:

* the API server, our point of entry for everything
* core services, scheduler / controller manager for example
* etcd, a highly available keystore. the db of kubernetes

**Nodes** in a cluster are the machines (VMs, physical servers, etc) that run our apps. The master controls each node.
A node needs a **container runtime**. Docker is the most popular with Kubernetes
has kubelet, the node-agent
and kube-proxy, a necessary network component
**Minikube** is a Kubernetes distribution, which enables us to run a single node cluster inside a VM on a workstation for dev & test
The **Kubernetes API** provides and abstraction of the Kubernetes concepts by wrapping them into objects / resources
**kubectl** is a command-line tool, we can use it to create, update, inspect and delete API objects

### API Objects

An API object is a **record of intent** which is specified in json or yaml. Recommendation is yaml.
Once we create the object, the cluster system will continuously work to ensure that object exists.
Every object consists of two parts, the object spec and the object status

* The spec describes the desired state of the object. We provide this in json/yaml format
* The status describes the actual state of the object, and is supplied and updated by the cluster.

#### Basic Objects

A **Pod** is a basic unit that Kubernetes deals with.
It encapsulates one or more closely related containers, storage resources, a unique network IP and configurations
Represents a single instance of an application

A **Service** is an abstraction which groups together logical collections of Pods, and defines how to access them.
Services are an interface to a group of containers so that consumers do not have to worry about anything beyond a single access location

Using **Volumes**, containers can access external storage resources for persistence. Also share files between containers.

With **Namespaces**, Kubernetes provides the possibility to run multiple virtual clusters on one physical cluster.
Provide scope for names of resources, which have to be unique within a namespace.

### Controllers

Higher level abstractions that build on basic objects and provide additional func.

A **Deployment** controller provides declarative updates for Pods and ReplicaSets.
We describe the desired state in the Deployment object and the Deployment Controller changes the actual state to the desired.

A **ReplicaSet** ensures that a specified number of Pods replicas are running at any one time

With **StatefulSet**, we can run stateful apps.
Impl apps with unique nw id or persistent storage and guarantee, ordered graceful deployment, scaling, deletion, termination as well as orderd and automated rolling updates.

With **DaemonSet**, we can ensure that all or a specific set of nodes in our cluster run one copy of a specific Pod
app monitoring, log collection, etc.

**GarbageCollection** makes sure that certain objects are deleted which have no owners

A **Job** creates one or more Pods, makes sure that a specific number of them terminate successfully and tracks job completion.

### MetaData

All objects have metadata

Mandatory metadata is:

* Namespace
* Name, unique within the namespace
* UID, unique in space and time

Some optional ones are:

* Labels
* * K/V pairs which can be attached to objects to categorize them
* * Helps to identify a collection of objects which satisfy a condition
* Labels Selectors
* * Identify a set of objects by their labels
* Annotations
* * K/V pairs too
* * not used to identify objects
* * hold information about the respective object

## kubectl

Command line tool for interacting with Kubernetes cluster

### commands

#### run

create a deployment with the specified configuration

#### get

get information about particular types of API objects in the cluster
deployments, pods, services, ...

#### logs

get the logs for a pod, particular container within a pod or a group of containers grouped by a label.

#### expose

this creates a service which exposes a container outside the cluster
--type=NodePort is important here. It tells Kubernetes that this service should be exposed externally.
The port which is exposed will be in the range 30000 to 32767
Another value for type is ClusterIp, which means the service is internal

#### delete

delete any type of API object

#### create

create the API object specified by the file
can only take one file at a time