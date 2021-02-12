# deployments

Deployment -> ReplicaSet -> Pod(s)

rolling updates
  gradually replace pods
blue-green deployment
  run multiple versions in parallel
canary deployment
  multiple versions
  but only small amount of traffic going to new version
rollbacks

default strategy is rolling updates

## Rolling Update

incrementally update pods with new one
Replica sets increase new pods while decreasing old pods
service handles load balancing traffic to available pods
new pods only scheduled on available nodes

bring up new pod
remove clusterIp of old
then remove pod
repeat until all pods are replaced

may not always be possible
  there are corner cases where user experience would be effected
  based on what pod they get, old or new

```yaml


...
  minReadySeconds: 1          # how long to wait until pod is ready, default 0
  progressDeadlineSeconds: 60 # how long to wait until reporting pod is stalled
  revisionHistoryLimit: 5     # number of replica sets that can be rolled back, default 10

  strategy:
    type: RollingUpdate       # by default
    rollingUpdate:      
        maxSurge: 1           # max pods that can exceed the replicas count, defaults to 25%
        minAvailable: 1       # max pods that are not operational, defaults to 25%
                                # rounds up to 1 if 25% is less than 1

```

## Canary deployments

run two versions at the same time
have most hit the older version

Canary? Into the coal mines, keep an eye bird as they would be more susceptible to carbon monoxid poisoning.
Whistle to the bird to make sure everything was ok

rollout a deployment, but only a small amount of traffic goes to it using the service

3 main resources:

1. Service
2. Stable deployment
3. Canary deployment

Create a new deployment with a fraction of the number of replicas
with the same label so the service picks it up
add an extra label to distinguish stable from canary deployment
