# volumes

Containers are considered transient, so data is lost when they are restarted.
A K8S volume shares the Pod lifetime, not the container lifetime.

A volume is a directory, which is mounted into the Pod.
As of v1.8 there are 25 volume types.
  anything from rbd to NFS & everything in between.

The CSI is standardizing how containers interface with storage systems.

Create a PersistentVolume, which is claimed by a Pod using a PersistentVolumeClaim.
A PersistentVolume is a storage abstraction used to retain data longer then the Pod using it.
Cluster-wide storage unit provisioned by an administrator with a lifetime independant of the pod
A Pod defines a volume of type PersistentVolumeClaim with various parameters.
The cluster then attaches the PVC to the PV.
The PV is backed by the Storage
Two Api Objects exist already to pass data to Pods.
* Secrets
* ConfigMaps

```yaml

apiVersion: v1
...
spec:
  volumes:
    - name: html
      emptyDir: {}
    - name: docker-socket
      hostPath:
        path: /path/on/host
        type: Socket
  containers:
    volumeMounts:
      - name: html
        mountPath: /mnt/www/html
        readOnly: true

```
A Pod spec can declare one or more volumes, each requires a name, type & mount point.
The same volume can be made available to multiple containers in a pod.
A volume can be made available to multiple pods, with each given an access mode.
  This can lead to data corruption as there is no concurrency checking done by K8S
  The application must handle this.

The cluster groups volumes with the same mode together, then sorts volumes by size, from smallest to largest.
The claim is checked against each in that access mode group, until a volume of sufficient size matches.
There are three access modes:
* RWO (ReadWriteOnce)   allows w by a single node
* ROX (ReadOnlyMany)    allows r by multiple nodes
* RWX (ReadWriteMany)   allows r-w by multiple nodes

RWO & ROX can be combined on the PV!

When a volume is requested, kubelet:
* map the raw devices
* determine and make the mount point for the container
* create symbolic link on host node FS to associate storage to the container
* If a particular StorageClass is requested: the API Server makes a request for the storage to the StorageClass plugin
    The specifics depend on the plugin used.

If no StorageClass is requested then access mode & size are the only consumed parameters.
The volume could come from any of the storage types available

There are several phases to persistent Storage
* Provisioning
  * of the persistent volume
  * can be in advance by the cluster administrator
  * or can be requested from a dynamic source, such as the cloud provider
* Binding
  * Occurs when a control loop (what controller is this?) on the master notices a PVC
* Use
  * begins when the bound volume is mounted into the Pod
  * continues as long as the Pod requires
* Releasing
  * when the Pod is done with the volume and an API request is sent, deleting the PVC.
* Reclaim
  * 3 options:
    * Retain
    * Delete
    * Recycle

## Volume types

emptyDir
* create a directory in the container
* but not mount any storage
* transient, it shares the pod lifetime

hostPath
* mounts a resource from the host node filesystem
* There is an option to create the directory if it doesn't exist!
* easy to setup
* but if worker node goes down data could be lost
* or if pod is re-scheduled on a different worker node the data would be inaccessible
  node affinity would need to be setup
* many types
  * DirectoryOrCreate
  * Directory
  * FileOrCreate
  * File
  * Socket
  * CharDevice
  * BlockDevice

NFS
* Good for multiple readers scenarios.

ConfigMap/Secret
* special types of volumes
* provide a pod with access to k8s resources

rbd, CephFS, GlusterFS
* good for multiple writers scenarios.

Cloud providers have their own specific volume types.

## StorageClass

Provides a way for administrators to describe the classes of storage they offer.
Acts as a storage template
Dynamically provision storage, so you don't have to manually create the PV
quality of service levels, backup policies, ...
PVC references StorageClass

```yaml

    provisioner
    volumeBindingMode: WaitForConsumer (Wait until first pod comes up, default is Immediate)

```
