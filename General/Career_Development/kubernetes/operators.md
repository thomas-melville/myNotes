# Operators

An Operator is an application-specific controller that extends the K8S API to create, configure and manage instances of complex stateful applications on behalf of a K8S user.
Builds on the basic K8S resource and controller concepts but includes domain or app specific knowledge to automate common tasks.

The stateful above is important.
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

## common patterns

should be a single deployment.
create a new 3rd party type when installed.
leverage built in primitives when possible.

## OperatorSDK

A component of the Operator Framework.
Simplifies building operators
also for testing & packaging
