# working with file systems

## inode

is a data structure on disk that describes and stores file attributes, including location.
Every file is associated with it's own inode

information stored:

* permissions
* user & group ownership
* size
* timestamp
* * last accessed time
* * last modified time
* * Change time

filenames are not stored in inode, stored in a directory file.

## directory file

particular type of file used to associated file names with inodes.

2 ways to do this

* hard links point to an inode.
* soft (symbolic) links point to a file name which has an associated inode.

links can be created using the 'ln' command

## extended attributes

associate metadata not interpreted directly by the FS with files.
Four namespaces exist:

* user
* trusted
* security
* system

system namespace is used for ACL (Access Control Lists)
security namespace is used for SELinux (Security Enhanced Linux)

Flag values are stored in the inode file and can only be set/modified by root
viewed with 'lsattr' and 'chattr'

user name space flags:

* i: Immutable       can't be modified, even by root. No hard link can be created and no data written. Only super user can update this attribute
* a: Append-only     can only be opened in append mode for writing. only super user
* d: no-Dump         file is ignored when the dump program is run. Useful for swap and cache files that you don't want to backup
* A: no atime update file will not modify it's access time record when the file is accessed but not modified.

## making a filesystem

mkfs is the name of the utilities
this is just a front end for filesystem specific programs
mkfs figures out the program to use from the filesystem

mkfs -t <type> <options> <device_file>
each file system type has their own options

## checking for errors

fsck
again, it's just a front end for fs type specific programs

journaling systems are much quicker to checker than older systems
manually fix issues one by one by specifying -r
or have them fixed automatically as best as the os can using -a

## Network File System

NFS is the most commonly used method for mounting remote FSs. Created by Sun microsystems.
Built upon the Open Network Computing Remote Procedure Call system. RPC connections are managed in Linux using the portmap service.
NFS relies on daemon processes (portmap, nfsd, mountd) and depends on the NFS version used.

mounting a remote FS should be done with a timeout to handle occasions when the FS is not available, either down or no network.

timeo=14

### configuration

#### server

main config done in /etc/exports file.
Two elements in the file:
1. a directory to share
2. host with mount options
    common options include root squashing, read size, write size and read/write.

#### client

client mounts the remote filesystem on to the local file system.
showmount command queries the mount daemon on the remote server for info.
portmap daemon is a dynamic port mapping daemon designed to reduce usage of well known port numbers.
mount command has the filesystem type NFS, which links to the mount.ns command.

### security

default security is to use UID and GID. The challenge is that this must match across the two systems.
Single sign on system is anothe option, NIS, LDAP, Kerberos. Which will remove the UID match option.

### performance

Properly setting the values of rsize and wsize will allow for greater speed in a file transfer.
Ethernet bandwidth.
async mode trades speed for lake of robustness

### Server Message Block (SMB) protocol

SMB was originally designed at IBM and later incorporated as the de factor n/w file/printing sharing in Microsoft Windows.
Renamed Common Internet File System (CIFS) in 1996, as new features were added.
The Samba project started as a reverse-engineered impl of the SMB protocol for Solaris.
Samba is built to be an SMB/CIFS server which will run on any unix system.

## Other Network FileSystems

### AFS

Andrew File System.
distributed, network FS built for performance.
OpenAFS is an open source branch as AFS is owned by IBM.

### Gluster

distributed NAS FS now maintained by Redhat.
Scalable network FS which uses COTS HW

### GFS2

Global File System.
shared-disk or clustering file system.
multiple nodes can share the same disks, all nodes are peers.
Maintained and supported by Redhat.

### DRBD

Distributed Replicated Block Device.
distributed replicated FS built for HA clusters.
Maintained by LINBIT.

## Mounting Network Filesystems

Three methods:
1. immediate mount
2. always mounted
3. mount on demand

All use a common configuration file, /etc/fstab
