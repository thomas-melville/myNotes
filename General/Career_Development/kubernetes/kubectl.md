# kubectl

version
cluster-info
get
run - simple way to create a deployment for an image
port-forward
expose - port of a Deployment/pod
create - create a resource - fail if it already exists
apply - create/modify a resource

Command line tool for interacting with Kubernetes cluster

```bash

kubectl completion bash

```

Source this in your bashrc file to get auto completion of kubectl commands

## commands

### create

create the API object, pod, replica set, controller, ... specified by the file
can only take one file at a time
Will fail if any resources already exist

--save-config
  store current properites in resources annotations
  when it comes to apply later it will use this to compare with whats already there.

can specify a file or cat and pipe into it ( add - flag)

-R flag can be used to recursively go through all files in sub folders

### run

kubectl run [podname] --image=...

create a deployment with the specified configuration
each pod will only have one instance of a single image.
If your app has multiple images then you'll need to define it in a yaml file and use the create command

#### arguments

--record
  records a history of the changes to the deployment

### get

get information about particular types of API objects in the cluster
deployments, pods, services, ...
there is also the option **all** to get information about everything in the cluster

you can change the format of the output using ther -o argument

wide            print more info
json            print all the information in JSON
custom-columns  print only selected fields under headers you define
                  =header:jsonPath,...
jsonPath        print a specific field from the output
                  you can do a lot with this!

--show-labels     show all labels associated with a resource

-l              show resources matching the label

### describe

information about the pod
including events, useful for troubleshooting

### exec

execute a command on a container / pod, ...
like docker exec you can start an interactive shell to the container

```bash
kubectl exec -it <pod_name> -- /bin/bash
```

### logs

get the logs for a pod, particular container within a pod or a group of containers grouped by a label.

```bash
kubectl logs <pod>  # will return the logs of the first container in the pod
kubectl logs <pod> -c <container> # will return the logs of the specified container
kubectl logs -l <label> # will return the logs for all pods which have the label
kubectl logs -p <pod> # will return logs for a previously running pod
kubectl logs -f <pod> # tail the logs
```

add the -f argument to follow the logs, like tail
add --max-log-requests X if you get an error following multiple pods

### expose

This creates a service which exposes pod(s) outside the cluster
best to expose a replication controller or deployment
--type=NodePort is important here. It tells Kubernetes that this service should be exposed externally.
The port which is exposed will be in the range 30000 to 32767. Or you can override this an specify a port
Another value for type is ClusterIp, which means the service is internal

### delete

delete any type of API object

--cascade=false, don't delete the underlying objects created by this object
it defaults to true. So if you delete a deployment you'll delete the replica sets and pods underneath that deployment by default

### rollout

rollout status deployment my.deployment

show the status of a rollout

history

show the history of changes to a deployment, daemonset or statefulset
To see more information create the API object with the --record flag

undo

rollback to a previous revision
by default rolls back to previous
--to-revision to specify revision to go to

#### arguments

pause
resume
undo
  undo the last change in the history

--to-revision=x
  roll back to a previous revision, need to run the deployment with the --record flag

### scale

scale the number of replicas in a deployment or replica set

--replicas=5

### edit

edit a particular object in a text editor.
The changes are applied to the cluster when the file is saved.

### set

edit certain resources by setting fields on them:

Fields are:
* env, of a pod
* image, of a pod
* resources, of a pod
* selector, of a resource
* serviceaccount, of a resource
* subject, in a binding

### patch

update a single property on a resource

### label

add a label to an API object

### (un)cordon

cordon, stop pods being put onto node
uncordon, allow pods to be put onto the node
useful for node downtime

### drain

cordon node and remove all pods from it

### proxy

create a local service to access a Cluster.
Useful for troubleshooting or development work.

### explain

print of the specification of any resource
you can get sub parts of the resource by chaining

for example: kubectl explain pod.spec

it also has a --recusive flag which goes down through everything!

### port-forward

kubectl port-forward [podname] <external_port>:<internal_port>
