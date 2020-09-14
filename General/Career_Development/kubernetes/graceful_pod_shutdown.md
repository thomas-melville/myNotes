# gracefully shutdown pods in a k8s cluster

https://twitter.com/danielepolencic/status/1280087657968627713?s=09

What happens when a pod is removed from a cluster?

1. kubelet starts the shutdown sequence
2. kube-proxy daemon on all worker nodes will remove ip address from iptables
3. endpoints controller will remove the pod from the list of valid endpoints

kubelet on node invokes preStop hook in the pod.
once the hook completes issue SIGTERM signal
wait for grace period, default 30 seconds
if the container hasn't shut down issue a SIGKILL

The container spec has a lifecycle section

```yaml

  lifecycle:
    preStop:
      exec:
        command: [ "/usr/sbin/nginx/", "-s", "quit"]

```

The above hook can result in the service sending traffic to the pod after the shutdown sequence has been initiated.
Due to the distributed nature of k8s.
The solution is to add a sleep for a few seconds at the start of the command in the hook

## PodDisruptionBudget

The number of disruptions that can be tolerated at a given time for a class of pods (can use labels)
We can ensure X number of pods are available
