# Microservices

decompose business domain models into smaller, consistent, bounded-contexts implemented by services.
services are isolated yet autonoumously communicate to provide business functionality.
A team is responsible for all aspects of a service from design to deployment to support.

Scoping the boundary of a service allows us to quickly:
	understand what the service is doing
		without getting tangled up in the larger application
	quickly build services locally
	pick the right technology
	test the service
	build/deploy/release at a suitable cadence for the business
	horizontally scale
	improve resiliency

## 5 things to keep in mind when developing microservices

### Design for faults

They will happen! And they can also be unrelated to your code, underlying infrastructure & hardware
A single fault can propagate to other parts of the system and cause cascading failures that can take down an entire system.
Historically we've tried to predict faults, which is hard in a complex system (I presume they mean multiple services communicating)
*Handle failures, don't just prevent it*
Lot's of networking involved in a micro services system!
And what if a service you are using gets bogged down?

### Design with Depenencies in mind

loose coupling everywhere
depend on services, but if they're degraded or not available, we need to be able to handle this gracefully.
interacting with legacy systems.
team dependencies also, the team developing the service must be autonomous

### Design with the Domain in mind

Use models to simplify and understand a problem.
The real complexity in software is not the technology, but the ambiguous, circular, often conflicting models the business folk have in their heads
models and context must be baked into the software.
Identify and explicitly separate different models and ensure they're cohesive and unambiguous within their bounded contexts.
Bounded context: a set of domain objects that implement a model that tries to simplify and communicate a part of the business, code and org.

### Design with Promises in mind

The relationship between service provider and consumer is very important.
One team can't place obligations on another, all you can do is accept their promise of functionality.
As a service provider all you can do is promise a certain amount of functionality
Promise theory, first proposed by Mark Burgess in 2004 (In search of Certainty)
Promises help articulate what a service may provide, and make clear what assumptions can and can't be made.
	if a service I depend on is unavailable, don't just throw an exception back to my consumer, strive to give them something, a default list, ... but make this clear in the promise
Strive to keep a promise as it shows empathy for the rest of the system.
Consumer driven contracts is another form of Promises!
	These can be tests which make sure we are honouring our promise with every change

### Distributed Systems Management

...


## Promises

Teams who developer services communicate through promises.

A promise is a way a service can publish intentions to other components/systems that may wish to use the service.
They specify these promises with interfaces of their services and via wikis/sites that document their services.

## Deploying Microservices at scale with docker & kubernetes

### Immutable Delivery

Reduce the number of moving parts into prebaked images as part of the build process.
No reliance on an environment being in a certain state, you ensure it is in the state you want.
Once an environment is deployed it is never changed, the prebaked image is changed and the environment is re-rolled.
	No configuration drift or snowflake servers
Infrastructure as code.
VM's could work, but they are big and virtualize and entire OS.
Docker, docker, docker :-)

#### Docker

lightweight, layered image format.
Runs these images inside Linux containers with isolated cpu, ram
application/process virtualization
Allows multiple containers to be run on a single host, and they are all isolated.

cgroups, namespaces and chroot from the linux kernel are used by docker, so it's nothing revolutionary from a technology point of view
Linux containers have been around for 10 years.
cgroups was contributed by Google in 2008
Docker simplified the user experience
chroot was introduced to the linux kernel in 1979

layered
	images are layered one on top of another to provide additional functionality.
	If a lower level image needs to be changed, just redeploy the top level image to get the new lower level image

#### Kubernetes

Open-sourced from Google
run clusters of microservices inside Linux containers at scale.

##### Pods

a grouping of one or more docker containers.
Services that must be co-located are usually done this way, affinity containers
Kubernetes orchestrates, schedules and manages pods.
A pod has it's own IP address and all containers inside the pod share this IP address.
volumes mounted to the pod are also shared between containers
pods are fungible, they can disappear at any time.

##### Labels

k/v pairs we assign to pods
group and categorizes the pods in a loosely coupled fashion
Allows groups of pods to be selected

##### Replication Controllers

Multiple instances of a microservice need to be managed
A ReplicationController manages the number of replicas for a given microservice.
For example, create a ReplicaController for a certain label and control the number of pods with that label.

##### Services

Groups pods with a label and abstract them with a single virtual IP.
That way we never have to manage the IP's of the individual pods.