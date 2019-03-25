# configmaps

Similar API resource to Secrets, except the data is not encoded.
Using a ConfigMap decouples the container image from the configuration artifacts

They store data as:

* sets of K/V pairs
* plain configuration files in any format

The data can come from:

* a collection of files
* all files in a directory
* populated from a literal value

A ConfigMap can be used in several different ways

* A Pod can use the data as env vars from one or more config maps
* Values in the map can be passed to commands in the pod
* Populate Volume from ConfigMap
* Add ConfigMap data to specific path in volume
* set file names and access mode in Volume from ConfigMap data
* can be used by system components and controllers

```yaml

kind: ConfigMap
apiVersion: v1
metadata:
  name: ...
data:
  ...: |
  {
  ...
  }
```

They must exist prior to being used.
They reside in a specific namespace.

```yaml

env:
  - name:
    valueFrom:
      configMapKeyRef:
        name: ...
        key: ...
```
