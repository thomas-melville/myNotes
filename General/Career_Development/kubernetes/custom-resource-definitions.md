# CRDs

A resource is an endpoint in the K8S API that stores a collection of API objects off a certain kind.
  example: the pod resource contains a collection of Pod objects.
A custom resource is an extension of the K8S API that is not necessarily availble on all clusters.
  a cutomization of a particular cluster.
They can appear and disappear in a running cluster through dynamic registration.
Once a custom resource is installed, users can create and access its objects with kubectl

On their own custom resources simply let you store and retrieve structured data.
It's only when they are combined with a controller do they become a true declarative API.
A declarative API alllows you to declare the desired state of your resource and tries to match the actual state to the desired state

A custom controller is a controller users can deploy and update on a running cluster, independently of the clusters own lifecycle.
*I think all those yaml files I applied during the PoC were for custom resources and controllers: istio, networking, load balancing*
They can work with any kind of resource, but are especially effective when combined with custom resources.

The **Operator** pattern is one example of such a combination.
It allows developers to encode domain knowledge for specific applications into an extension of the K8s API.

## API Aggregation
