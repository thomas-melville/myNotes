# troubleshooting

k8s is sensitive to network issues. The standard Linux tools and processes are a great help, I've lots of notes on these.
If your pod doesn't have a shell, consider deploying one that does (busybox)
dig, tcpdump, ...

Monitoring is not a core part of K8s so you'll have to include a 3pp
* Heapster, which is now retired
* Prometheus

Logging is another cross cluster concern
* Fluentd

## basic troubleshooting steps

if the pod is running execute 'kubectl logs pod-name' to view the std out of the container.
next is networking, including DNS, firewall and general connectivity.
Node logs for errors
RBAC & other security SELinux / Apparmor
API calls to kube-apiserver & controllers
inter-node network issues
master server controllers

get pods
  make sure all are running
  look for a lot of restarts
describe the troublesome pod

## monitoring

collect metrics from infra & apps.
Prometheus is part of the CNCF, as a K8s plugin it can collect metrics from all the resources in the cluster.

## Logging

Typically, logs are collected locally and aggregated before being ingested by a search engine and displayed via a dashboard.
ELK stack.
kubelet writes container logs to files.
Cluster wide you can use Fluentd, part of CNCF

OpenTracing
Jaeger

journalctl to view system logs
