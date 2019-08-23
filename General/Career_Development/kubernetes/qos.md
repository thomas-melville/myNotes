# QoS

K8S has Quality of Service classes for Pods.
This helps it decide who to evict when.

Three classes:

1. Guaranteed
2. Burstable
3. BestEffort

You don't set the class explicitly.
You give the pod certain requests and limits and it is then given a QoS class by K8S

## Guaranteed

For every container in the pod memory & cpu requests and limits must be set and the same

## Burstable

It does not meet the criteria for Guaranteed.
At least one container in the Pod has a memory or CPU request

## BestEffort

No container in the pod should have a cpu or memory request or limit

## Best Practices

always give DaemonSet pods Guaranteed class, as they'll always be redeployed onto the same node
