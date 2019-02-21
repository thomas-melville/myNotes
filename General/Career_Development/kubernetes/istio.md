# istio

An open platform to connect, manage, and secure microservices.

* Intelligent Routing and Load Balancing
* Resilience Across Languages and Platforms
* Fleet-Wide Policy Enforcement
* In-Depth Telemetry

Istio provides an easy way to create a network of deployed services with load balancing, service-to-service authentication, monitoring, and more, without requiring any changes in service code.
You add Istio support to services by deploying a special sidecar proxy throughout your environment that intercepts all network communication between microservices, configured and managed using Istio’s control plane functionality.

Currently supports:

* service deployment on Kubernetes
* service registration with Consul or Eureka
* services running on individual VMs

It puts a control plane on top of the data plane to proxy all communication, monitor, ...