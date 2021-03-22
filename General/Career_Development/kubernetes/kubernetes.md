# kubernetes

http://www.baeldung.com/spring-boot-minikube
https://labs.play-with-k8s.com/#

1. open source system for the automated deployment, scaling and management of containerized apps.
2. system for orchestration of containerized apps across a cluster of nodes, including networking and storage infrastructure

Declarative way to define what you want and Kubernetes makes it so

Some important features are:

* Resource scheduling
  * ensures pods are distributed optimally across nodes
* Auto-scaling
  * with increasing load, the cluster can dynamically allocate additional nodes
  * and deploy new pods on them.
* Self-healing
  * cluster supervises containers and restarts them
  * if required, based on defined policies
* Service-discovery
  * Pods and Services are registered and published via DNS
* Rolling updates/rollbacks
  * support rolling updates based on sequential deployment of Pods and containers
* Secret/config mgmt
  * supports secure handling of sensitive data like passwords and API keys
* Storage Orchestration
  * several 3rd party solutions are supported
  * which can be used as external volumes to persist data

It is written in GoLang.
Configuration information is stored in JSON, and written in YAML normally

### why?

Accelerate Developer onboarding
elimiate app conflicts
Environment consistency
ship software faster

Orchestrate containers
zero-downtime deployments
self-healing
scale containers

emulate production locally
move from docker compoase to k8s
create an end to end testing environment
ensure app scales properly
ensure secrets/configs are working properly
performance testing
workload scenarios
learn how to leverage deplotment options
help devops create resources and solve problems

## understanding Kubernetes

The **master** maintains the desired state of the **cluster**
There can be more than one master for availability.
when we interact with our cluster, by using **kubectl** for example we are communicating with the clusters master
it's the brains, a collection of services:

* the API server, our point of entry for everything. It's over REST so we could write our own client to communicate with it.
* core services, scheduler / controller manager for example
  * scheduler decides on what node pods will be spun up.
    * it uses the PodSpec to decide this.
* * controller manager manages the threads which report the actual state of the cluster
    * also signals the controllers to update the actual state when it doesn't match the desired state
* etcd, a distributed highly available keystore. the db of kubernetes
  * b+tree, key/value store
  * no entry is edited, values are always appended to the end, previous copies are then marked for future removal by a compaction process
* cloud-controller-manager, a relatively new manager which interacts with cloud providers

**Nodes** in a cluster are the machines (VMs, physical servers, etc) that run our apps. The master controls each node.
A node needs a **container runtime**. Docker is the most popular with Kubernetes. It supports rkt also and will support any engines which conform to the Open Container Initiative reference spec.
has kubelet, the node-agent, talks with the container engine to ensure all containers that should be running actually are.
and kube-proxy, a necessary network component which allows user traffic into the node. It does this using iptables, but that may have been superceded by ipvs
It has a built in load balancer
Supervisord monitors the kublet and docker engine processes and restarts them if there are crashes.

**Minikube** is a Kubernetes distribution, which enables us to run a single node cluster inside a VM on a workstation for dev & test
**DockerDesktop** can install a kubernetes cluster
**kind** you can run kubernetes in docker
The **Kubernetes API** provides and abstraction of the Kubernetes concepts by wrapping them into objects / resources
**kubectl** is a command-line tool, we can use it to create, update, inspect and delete API objects
  interacts with api-server on master node

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
It encapsulates one or more closely related containers, storage resources, a unique network IP (clusterIP address) and configurations.
Typically, one container runs an application, while any other containers support the primary container.
containers in the pod share the same network namespace
  have the same loopback network interface, localhost
  this means containers need to bind to different ports in the pod
Pod never spans a worker node
The containers start in parallel so there's no way to determine which one will be ready first.
Represents a single instance of an application.

A **Service** is an abstraction which groups together logical collections of Pods, and defines how to access them.
You can't rely on a pods ip address as they are ephemeral
Services are an interface to a group of containers so that consumers do not have to worry about anything beyond a single access location
Labels are used to associate a pod with a service
kube-proxy create a virtual IP for a service
layer 4
Endpoints created behind services which sit between pod and service
load balances
when browsers connect they use the same connection to go to same pod everytime

Using **Volumes**, containers can access external storage resources for persistence. Also share files between containers.

With **Namespaces**, Kubernetes provides the possibility to run multiple virtual clusters on one physical cluster.
Provide scope for names of resources, which have to be unique within a namespace.

An **Ingress** is a set of rules that allow inbound connections to reach cluster services. It only supports http at the minute. It assigns a path to a service. An **IngressController** does the work of managing the rules (keeping the config in sync with the rules) and monitoring the rules and forwarding requests to the cluster service

A **NetworkPolicy** isolates groups of pods from one another. An example would be isolating test environments from production environments.

### Controllers

Higher level abstractions that build on basic objects and provide additional func.

A **Deployment** controller provides declarative updates for Pods and ReplicaSets.
wrap and simplify ReplicaSet
We describe the desired state in the Deployment object and the Deployment Controller changes the actual state to the desired.
This is the recommended controller to use over replica sets, as it encompasses replica sets and you can change the spec and Kubernetes can do a rolling update/grade of all the containers, history and rollback.
the labels and selectors between the deployment and the replica set must match.
zero downtime upgrading by creating new pods before destroying old
  Rolling updates
    create new pod with new version before deleting old one
    default
  blue-green deployments
    a/b
    multiple environments with different versions
  canary deployments
    a small amount of traffic goes to new versions
  rollbacks
    something went wrong so we roll back
rollback

minReadySeconds
  wait 10 seconds to make sure container doesn't crash

A **ReplicaSet** ensures that a specified number of Pods replicas are running at any one time
a declarative way to manage pods
self healing mechanism
ensure requested number of pods are available
provide fault tolerance
can be used to scale pods
no need to create Pods directly
used by deployments
unqiue name of replica set will follow through to pod name, with another identifier to make the pod unique

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

There are many other controllers for managing the resources in a cluster, and you can create your own.

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
* * can't be used in k8s commands
* * could be used by 3rd party agents or tools

#### labels

labels are very important and widely used in kubernetes.
They are used to group, identify and select any of the API objects.

### yaml

a text file composed of maps and lists

```yaml

apiVersion: <version>
kind: <the api object to create>
metadata: # metadata about the pod
    <...>
spec: # blueprint for the object
    <...>

```

## standardization

K8S is part of the CNCF (Cloud Native Computing Foundation)
K8S is standardizing on the CNI (Container Network Interface) spec.
K8S is standardizing on the CRI (Container Runtime Interface) spec.
  goal is to allow easy integration of container runtimes with kubelet
  docker-cri is done
  cri-o, rktlet & frakti are works in progress (Maybe they're done by now)
  rkt is part of CNCF, announced by CoreOS in 2014
  cri-o is part of K8S

## affinity

A loose way to assign pods to nodes.
can be based on nodes or pods already in the node.
can be required (hard) or preferred (soft)
and NoSchedule or NoExecute
NoExecute is stricter, it evicts any running pods in the node



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
