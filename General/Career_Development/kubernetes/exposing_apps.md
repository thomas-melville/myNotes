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

This implies an ordering requirement, if the service isn't up when the Pod is started then it won't get the env vars.

#### dns

An optional, but strongly recommended, cluster add-on is a DNS server, CoreDNS
The DNS server watches the Service resource and creates new DNS records for each.
If DNS has been enabled throughout the cluster then a Pod should be able to do name resolution of Services.
kubelet configures each Pods /etc/resolv.conf to use the coredns pod as the nameserver

```bash

bash-4.4# cat /etc/resolv.conf
nameserver 10.96.0.10
search tom.svc.cluster.local svc.cluster.local cluster.local rnd.gic.ericsson.se
options ndots:5


```

nameserver: where queries are forwarded to
search: search pattern for a particular domain
ndots: threshold value of number of dots in a query name to consider it as a fully qualified domain name
        capped to 15, and 5 is the kubernetes default
From within the name space use the service name.
From another namespace qualify it with the namespace
  my-service.other-ns

K8s also supports DNS SRV (service) records for named ports.
If a service has a named port then it can be looked up
  \_http_.\_tcp.my-service.other-ns

##### A Records

Normal Services are assigned a DNS A record for a name of the form: my-svc.my-namespace.svc.cluster.local.
This resolves to the Cluster IP address of the Service

Headless Services are assigned a DNS A record of the same form as normal.
This resolves to the set of IPs of the pods selected by the service.
Clients handle the selection and load balancing algorithm.

##### SRV records

Create for named ports which are part of Servces of the form: \_my-port-name.\_my-port-protocol.my-svc.my-namespace.svc.cluster.local

When enabled Pods can be given records in the DNS also: pod-ip-address.my-namespace.pod.cluster.local
A pods hostname is taken from it's metadata.name value. The Pod spec has an optional hostname.

##### Pod DNS Policy

It is possible to define one of many dns policies per pod

* Default
  * inherits from node
* ClusterFirst
  * any dns query which does not match the configured cluster domain suffix is forwarded to the upstream nameserver inherited from the node
  * **This is the default policy if none is provided**
* ClusterFirstWithHostNet
  * for pods running with host network
* None
  * ignore all settings from k8s
  * need to provide settings in dnsConfig of Pod spec



## ingress

is an object containing a list of rules against all incoming requests.
Only http rules are supported.
The request must match the host and path declared in the ingress
Routes traffic to existing services.
Easier to manage when there are lots of services
