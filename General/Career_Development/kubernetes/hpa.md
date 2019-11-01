# Horizontal Pod Autoscaler

https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
https://learnk8s.io/autoscaling-apps-kubernetes/

automatically scale the number of pods in a repl controller, deployment, repl set based on:
  CPU utilization
    v1
  custom metrics (https://git.k8s.io/community/contributors/design-proposals/instrumentation/custom-metrics-api.md)
    v2beta2!!!
Does not apply to objects that can't be scaled:
  DaemonSet

Implemented as a K8S API Resource and a controller.
The resource determines the behaviour of the controller.
The controller periodically adjusts the number of replicas to match the observed average usage
  default period is 15 seconds
  can be adjusted with --horizontal-pod-autoscaler-sync-period
  queries the resource utilization against the metrics specified in the HPA object.
  Gets the utilization of all pods, calculate the mean, and then decides how many to scale by.
    algorithm: desiredReplicas = ceil[currentReplicas * ( currentMetricValue / desiredMetricValue )]
  uses the scale sub-resource on the repl set / deployment / ... to scale the pods

To kubectl it is like any other object, it has all the CRUD operations
There is also a *autoscale* sub command.

```bash

kubectl autoscale deployment my-dep --cpu-percent=80 --min=1 --max=10

```

## hpa resource

contains following parameters:

* resource to scale
* min & max replicas
* scaling metrics
* target value for scaling the metric

if scaling a custom metric the hpa will talk to the metrics registry.
The metrics registry is a central place in the cluster where metrics of any kind are exposed to clients.
hpa is one of these clients
a standard interface for clients to query metrics

three apis:
* resource metrics
  * cpu & memory of pods and nodes
* custom metrics
  * associated with a k8s object
* external metrics
  * not associated with a k8s object

So, you need to expose your metric from the metrics registry.
install and configure additional components on the cluster, cadvisor and prometheus & prometheus adapter

Steps:
1. instrument your app to expose the metric
2. install & configure prometheus
3. install & configure prometheus adapter. Take the metric from prometheus and turn it into a rate (Prom-QL)
4. expose this metric through custom metrics api
5. create your hpa
