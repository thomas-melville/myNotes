# pods

smallest unit of deployment on k8s.
pods let us augment containers in a lot of ways
  probes, affinities, restart policies, ...

## phases of lifecycle

* Pending
  * Accepted but no container created.
  * This includes waiting for system to pick it up and downloading images over the network
* Init: N/M
  * init Containers completion status
* Init: Error
* Init: CrashLoopBackOff
* Running
* Succeeded
* Failed
* Unknown

https://confluence.lmera.ericsson.se/display/AA/Approach+for+ADP+microservice+LCM?focusedCommentId=145454279#comment-145454279

## multi container pods

kind of goes against the idea of decoupling as much as possible, but there are certain needs in which a 2nd/3rd container in the pod makes sense
There are two types of multi container pods, single node and multi node.

each container in the pod should have it's own requests & limits

### init

guaranteed to start, run once and finish before main app container
can be as many as you want
all execute in sequence
but if one errors out the remainder don't execute and the pod is restarted
init containers must be idempotent
used to prepare environment

keep resource requests for init containers low, as it's wasted when init container is finished.
if there are multiple containers then it goes through the requests and takes the biggest value.
All of the requests are not summed

```yaml

#...
spec:
  initContainers:
    name: ...
    image: ...
    command: ...
  containers:
    ...

```

### sidecar

add some functionality not present in the main container. extend and work with primary container.
All remaining patterns are types of sidecar patterns.
Runs along side main container, with no scheduling so order starting up of containers is non-deterministic
work ongoing to tag side car containers so they start up before main container
Best used when there is a clear difference between a primary container and any secondary tasks that need to be done to it.
Rather than bloating the main container with functionality that may not be required in all deployments.
Examples of sidecar containers are:
* logging (fluentd)
* monitoring (prometheus)
* service mesh

### adapter

modify/transform data, either on ingress or egress, to match some other need.
transform output of primary container into output that fits the standards across your applications.
**Would this be a case for ETSI problemDetails**
normalize the monitoring interface is an example given.
Would adding headers be another example?
Examples: prometheus, adapter that transform metrics coming out of app into format prometheus likes. expose it on a different port

### ambassador/proxy

simplify the access of external services for the main app, acts as a service discovery layer
allow access to the outside world without having to implement a service or ingress
* proxy local connection
* reverse proxy
* limits http requests
* re-route from the main container to the outside world
The main app uses the external service on localhost. ambassador container forwards requests to external system
ambassador container takes care of connecting, re-connecting and updating configuration.

### multi node pods

#### leader election

replication is great to share load, but in some cases all those replicas need a leader / leaders

#### work queue pattern

#### scatter/gather pattern

split up request among a number of containers and gather the response into one.
