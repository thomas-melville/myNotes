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

It is written in GoLang

## understanding Kubernetes

The **master** maintains the desired state of the **cluster**
There can be more than one master for availability.
when we interact with our cluster, by using **kubectl** for example we are communicating with the clusters master
it's the brains, a collection of services:

* the API server, our point of entry for everything. It's over REST so we could write our own client to communicate with it.
* core services, scheduler / controller manager for example
* * scheduler decides on what node pods will be spun up.
* * controller manager manages the threads which report the actual state of the cluster
* etcd, a distributed highly available keystore. the db of kubernetes

**Nodes** in a cluster are the machines (VMs, physical servers, etc) that run our apps. The master controls each node.
A node needs a **container runtime**. Docker is the most popular with Kubernetes. It supports rkt also and will support any engines which conform to the Open Container Initiative reference spec.
has kubelet, the node-agent
and kube-proxy, a necessary network component which allows user traffic into the node.
It has a built in load balancer

**Minikube** is a Kubernetes distribution, which enables us to run a single node cluster inside a VM on a workstation for dev & test
The **Kubernetes API** provides and abstraction of the Kubernetes concepts by wrapping them into objects / resources
**kubectl** is a command-line tool, we can use it to create, update, inspect and delete API objects

### architecture

Everything on the master and nodes are containers themselves.
Every pod has at least one container, pause, which reserves the namespace for networking and IPC before your containers are started.
And if all your containers die, the pod will still be there running the pause container (even though it does nothing) so the namespace doesn't disappear too.

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

An **Ingress** is a set of rules that allow inbound connections to reach cluster services. It only supports http at the minute. It assigns a path to a service. An **IngressController** does the work of managing the rules (keeping the config in sync with the rules) and monitoring the rules and forwarding requests to the cluster service

A **NetworkPolicy** isolates groups of pods from one another. An example would be isolating test environments from production environments.

### Controllers

Higher level abstractions that build on basic objects and provide additional func.

A **Deployment** controller provides declarative updates for Pods and ReplicaSets.
We describe the desired state in the Deployment object and the Deployment Controller changes the actual state to the desired.
This is the recommended controller to use over replica sets, as it encompasses replica sets and you can change the spec and Kubernetes can do a rolling update/grade of all the containers, history and rollback.
the labels and selectors between the deployment and the replica set must match.

A **ReplicaSet** ensures that a specified number of Pods replicas are running at any one time

With **StatefulSet**, we can run stateful apps.
Impl apps with unique nw id or persistent storage and guarantee, ordered graceful deployment, scaling, deletion, termination as well as orderd and automated rolling updates.

With **DaemonSet**, we can ensure that all or a specific set of nodes in our cluster run one copy of a specific Pod.
This is used for infrastructure related work, app monitoring, log collection, etc.

**GarbageCollection** makes sure that certain objects are deleted which have no owners

A **Job** creates one or more Pods, makes sure that a specific number of them terminate successfully and tracks job completion.

An **IngressController** monitors an **Ingress** object and forwards any incoming requests to a **Service**
It is a daemon, deployed as a Kubernetes POD, that watches the apiservers /ingresses endpoint for updates to the ingress resource
When there is any updates it reloads the nginx configuration file, which holds the forwarding information to services.
The nginx configuration is stored on the IngressController POD

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

#### labels

labels are very important and widely used in kubernetes.
They are used to group, identify and select any of the API objects.

### yaml

```yaml

apiVersion: <version>
kind: <the api object to create>
metadata:
    <...>
spec:
    <...>

```

## affinity

A loose way to assign pods to nodes.
can be based on nodes or pods already in the node.
can be required (hard) or preferred (soft)
and NoSchedule or NoExecute
NoExecute is stricter, it evicts any running pods in the node

## kubectl

Command line tool for interacting with Kubernetes cluster

```bash

kubectl completion bash

```

Source this in your bashrc file to get auto completion of kubectl commands

### commands

#### create

create the API object, pod, replica set, controller, ... specified by the file
can only take one file at a time

#### run

create a deployment with the specified configuration
each pod will only have one instance of a single image.
If your app has multiple images then you'll need to define it in a yaml file and use the create command

#### get

get information about particular types of API objects in the cluster
deployments, pods, services, ...
there is also the option **all** to get information about everything in the cluster

#### exec

execute a command on a container / pod, ...
like docker exec you can start an interactive shell to the container

#### logs

get the logs for a pod, particular container within a pod or a group of containers grouped by a label.

```bash
kubectl logs <pod>  # will return the logs of the first container in the pod
kubectl logs <pod> -c <container> # will return the logs of the specified container
```

#### expose

This creates a service which exposes pod(s) outside the cluster
best to expose a replication controller or deployment
--type=NodePort is important here. It tells Kubernetes that this service should be exposed externally.
The port which is exposed will be in the range 30000 to 32767. Or you can override this an specify a port
Another value for type is ClusterIp, which means the service is internal

#### delete

delete any type of API object

--cascade=false, don't delete the underlying objects created by this object
it defaults to true. So if you delete a deployment you'll delete the replica sets and pods underneath that deployment by default

#### rollout

history

show the history of changes to a deployment, daemonset or statefulset
To see more information create the API object with the --record flag

undo

undo the last change in the history

#### label

add a label to an API object

#### (un)cordon

cordon, stop pods being put onto node
uncordon, allow pods to be put onto the node
useful for node downtime

#### drain

cordon node and remove all pods from it

## typical workflow

1. create a PersistentVolume and Claim
2. create a deployment
* * pod
* * * add claim to volume
* * replicas
3. expose it as a service
4. create an ingress to the service
5. create an ingress controller to forward requests from a rule to a service
6. create any daemon sets for monitoring / logging