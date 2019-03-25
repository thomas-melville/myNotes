# exposing your application

## services

Types of services:

* ClusterIp
  * default type
  * internal to the cluster
  * IP range is defined via an API server startup option
* NodePort
  * expose a static IP address & port
  * great for debugging
  * Port range is defined in cluster config
* LoadBalancer
  * created to pass requests to a cloud provider
  * private cloud solutions may also implement this service type, if there is a plugin
  * Even without the plugin the address is made public and packets are spread among the Pods in the deployment automatically.
* ExternalName
  * newer service, and a bit different
  * no selectors, ports or endpoints
  * It allows the return of an alias to an external service
  * Redirection happens at DNS level
  * Redirect traffic to outside the cluster
  * but then can easily be changed to inside the cluster when the time comes.

### services diagram

https://kubernetes.io/docs/concepts/services-networking/service/#proxy-mode-userspace

kube-proxy running on each worker node watches the Service and Endpoints resources on the API server.
Creates / updates rules in the iptables when resources are added / changed.
Runs a readiness probe on all pods before connection.

another mode ipvs is available instead of iptables.
works in the kernel space for greater speed.
Also allows for a configurable load balancing algorithm.
Can be useful when a cluster is larger than 5000 nodes.

You can give a service multiple ports, an obvious example is http & https

### services to pods

When a service is created it creates an Endpoint which points to the Pods and is updated as pods die or are created.
labels are used to determine which pods should received traffic from which service.
labels can be dynamically updated which may affect what pods service what service.
The default update pattern is a rolling update, which means new and old versions of pods will service the same service.
If the versions are incompatible consider making the label version specific.
If you create a service without a selector the endpoint is not created.
  You can add the endpoint later (list of ip address & port combinations)
    to connect to a remote db.
    or to a service in another namespace / cluster

### discovering Services

2 ways

#### environment variables

When a pod is running on a node kubelet adds a set of env vars for each active Service.
{SVC_NAME}_SERVICE_HOST & {SVC_NAME}_SERVICE_PORT

https://kubernetes.io/docs/concepts/services-networking/service/#discovering-services

#### dns



## ingress

is an object containing a list of rules against all incoming requests.
Only http rules are supported.
The request must match the host and path declared in the ingress
Routes traffic to existing services.
Easier to manage when there are lots of services
