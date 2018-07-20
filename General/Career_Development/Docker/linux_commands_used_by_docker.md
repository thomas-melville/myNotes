# linux commands used internally by docker

## File system

chroot

* change the root file system for a process
* this will create a separate, isolated file system for an app

## Namespace

unshare and nsenter

* create isolated share of resources for mount, pid, network, etc.

## Limit

cgroups

* limit resource usage
* uses systemd

## Linux Capabilities

app can do stuff that would have to be done as root on the machine