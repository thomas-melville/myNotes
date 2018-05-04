# containers

light weight method to isolate an application or set of applications.
unlike VMs multiple containers can run on an OS which itself can be virtualized

## basics

further integration between the hypervisor and the linux kernel allowed the creation of OS level VMS, or containers.
containers share many facilities in the kernel, such as namespace and cgroups
very light weight and reduce the overhead associated with whole VMs.

First flavour was the OS container
runs an image of an OS with the ability to run init processes and spawn multiple applications
LXC (Linux Containers) is one example

In a effort to further reduce the overhead associated with VMS, app virtualization is rising in popularity.
runs one application for each container
many containers initialized on a single machine
Docker is one example

## docker

app-level virtualization using many individual images to build up the necessary services to support the target app