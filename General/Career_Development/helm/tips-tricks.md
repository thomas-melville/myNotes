# tips & tricks

## get a container to restart when a configuration/secret changes

When you do a helm upgrade, a configMap / secret may have changed.
This will be replaced but the running application will not be aware of it, if nothing in it's spec hasn;t changed.
i.e. it won't be restarted.

Using an annotation is a way to do this.

the key of the annotation can be anything, the value must be a Go template which process the specific configMap/secret and makes a checksum out of it.

```YAML

metadata:
  annotations:
    checksum/config: {{ include (print $.Template.BasePath "/configmap.yaml") . | sha256sum }}

```

Put this annotation on the Deployment

When an upgrade happens the template will be processed, if the yaml file pointed to has changed the checksum will be different which means the deployment will be upgraded.

## keep resources across helm deletes

When you delete a release, all resources are deleted.
If you'd like to keep a resource for the next installation you can annotate it.

helm.sh/resource-policy: keep/delete

this can cause issues if you use the --replace flag on the next helm install
