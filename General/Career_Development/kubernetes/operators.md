# Operator Framework

There are two parts to the K8S operator framework:

* OperatorSDK
* Lifecycle Manager

An Operator is an application-specific controller that extends the K8S API to create, configure and manage instances of complex stateful applications on behalf of a K8S user.
Builds on the basic K8S resource and controller concepts but includes domain or app specific knowledge to automate common tasks.

The OperatorSDK makes it easier to build K8s native applications.

The Lifecycle Manager manages the lifecycle of operators running in your cluster.
**It must be installed on the cluster!**

*The stateful above is important.*
It is easy to manage stateless deployments (replica sets to scale out and in) as all you need to do is spin up more instances.
When state is involved it becomes more complicated.
This is about embedding knowledge about how to correctly scale, upgrade and reconfigure apps while protecting against data loss or unavailability.

Two example operators: Etcd operator & the prometheus operator.

Etcd operator

when scaling an etcd cluster the following has to happen:

* create a DNS name for new etcd member
* launch the new instance
* execute an admin task to tell the cluster about the new instance

With the etcd operator, increase the cluster size and all this happens automatically

They build on two central concepts: resources & controllers.
resource states the state we want.
controller tries to make that desired state match what is in the cluster.
  The controller watches a specific type of deployment, for example the etcd deployment
  The controller watchs a primary object, the CR, and a secondary object, the deployment (associated to the CR)

## common patterns

should be a single deployment.
create a new 3rd party type when installed.
leverage built in primitives when possible.

## OperatorSDK

A component of the Operator Framework.
Simplifies building operators
also for testing & packaging

https://github.com/operator-framework/operator-sdk/blob/master/doc/helm/user-guide.md
Steps to install the operator SDK

It is a command line tool to create operators crds and controllers, probably more too.

1. Create the operator
2. Create the crd
3. Create the controller, put th logic in it to manage the crds deployments
4. build the operator image
5. Set up the RBAC permissions for the operator
6. deploy the operator
  * Either manually or through the Lifecycle Manager
7. Apply the CR to kick the operator into action

*Implementing the domain specific knowledged in the Controller requires knowledge of the GO programming language*

There are two parts to the Controller:

1. Watch the CR
  * It watches two things:
    1. the CR for add/update/delete, the primary object
    2. Deployments, which are mapped to this CR
  * Reconcile Requests are sent
    * a namespace/name key for that object.
2. Reconcile it.
  * Each controller has a reconcile loop
  * The loop is passed the reconcile argument, which is used to look up the primary object.

## Lifecycle Manager

It runs as a K8S extension and you can use kubectl for all of the operations.
Create the ClusterServiceVersion manifest for the:
* operator container * permissions
* CRDs
* ... there's a lot in the example file.

https://github.com/operator-framework/getting-started
