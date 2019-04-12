# CRDs

https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/

A resource is an endpoint in the K8S API that stores a collection of API objects of a certain kind.
  example: the pod resource contains a collection of Pod objects.
A custom resource is an extension of the K8S API that is not necessarily availble on all clusters.
  a cutomization of a particular cluster.
  For example Istio has it's own CRDs
They can appear and disappear in a running cluster through dynamic registration.
Once a custom resource is installed, users can create and access its objects with kubectl

On their own custom resources simply let you store and retrieve structured data.
It's only when they are combined with a controller do they become a true declarative API.
A declarative API alllows you to declare the desired state of your resource and tries to match the actual state to the desired state

A custom controller is a controller users can deploy and update on a running cluster, independently of the clusters own lifecycle.
*I think all those yaml files I applied during the PoC were for custom resources and controllers: istio, networking, load balancing*
  I think I'm actually wrong in some of this, there are operators and then there are plugins.
  and within plugins there are infrastructure plugins (storage, networking, device) and kubectl plugins (add a new command, just like git)
  Istio does define it's own CRDs
They can work with any kind of resource, but are especially effective when combined with custom resources.

The **Operator** pattern is one example of such a combination.
It allows developers to encode domain knowledge for specific applications into an extension of the K8s API.

## define a CRD

Please see the crd.yaml file

You can put validation on the properties that people will define in the spec of your CRD

## creating a Custom Object from a CRD

use the API version from the CRD
and the kind

## API Aggregation
