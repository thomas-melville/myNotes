# configmaps

store configuration information and provide it to containers
  share across pods

Similar API resource to Secrets, except the data is not encoded.
Using a ConfigMap decouples the container image from the configuration artifacts

They store data as:

* sets of K/V pairs
* plain configuration files in any format
  key is the filename
  value is the file contents

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
  my-key: my value:
  or:
  mt-configfile: |-
    {
    ...
    }
```

They must exist prior to being used.
They reside in a specific namespace.

```yaml

# load individual
env:
  - name: my-env-var
    valueFrom:
      configMapKeyRef:
        name: config-map-ref
        key: key-in-cm

# load whole configmap
envFrom:
  - configMapRef:
    name: config-map-ref

# load into volume
volumes:
  - name: cconfig-map-volume
    configMap:
      name: config-map-ref
containers:
  volumeMounts:
    - name: config-map-volume
      mountPath: /etc/config
```

When using volume
  each key/value pair becomes a file in the mounted folder
  it allows any updated config data to get into the file automatically
  then your code must read it in though

env, fromEnv & volume can be mixed and matched

```bash

kubectl create configmap <name> --from-file=<path_to_file> / --from-env-file=<path> / --from-literal=key=value

```

from-file results in ConfigMap looking like one key value pair. File name = key, file contents = value
from-env-file results in multiple key value pairs which match key value pairs in file
from-literal you can use multiple times to enter k.v pairs individually
