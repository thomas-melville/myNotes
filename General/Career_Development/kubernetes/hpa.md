# Horizontal Pod Autoscaler

https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/

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
