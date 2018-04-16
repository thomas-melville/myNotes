# mount/umount

All accessible files in Linux are under one directory tree.
They don't have to all be on the same partition.
A partition with a separate directory tree and fs type can be mounted at any point in the tree.

when mounting the directory to mount to must exist.
if it has files they will be hidden, but not deleted.

## mount

takes many options

-a          mount all filesystems mentioned in /etc/fstab
-remount    remount
,ro         read-only
...

mount -a is executed at boot time