# linux commands used internally by docker

## File system

chroot

* change the root file system for a process
* this will create a separate, isolated file system for an app

## Namespace

unshare and nsenter

* create isolated share of resources for mount, pid, network, etc.

pid namespace: process isolation
net namespace: network isolation
ipc namespace: inter process communication isolation
mnt namespace: mount isolation
uts namespace: isolating kernel & version identifiers (UTS: Unix Timesharing System)

## Limit

cgroups

* limit resource usage
* uses systemd

## Union File system

file systems that operate by creating layers, making them very light weight and fast.
docker uses this to provide building blocks for containers
lots of variants: AUFS, btrfs, vfs, DeviceMapper

## Linux Capabilities

app can do stuff that would have to be done as root on the machine
