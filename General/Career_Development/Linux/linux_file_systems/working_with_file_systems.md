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

NFS is the most commonly used method for mounting remote FSs

mounting a remote FS should be done with a timeout to handle occasions when the FS is not available, either down or no network.

timeo=14