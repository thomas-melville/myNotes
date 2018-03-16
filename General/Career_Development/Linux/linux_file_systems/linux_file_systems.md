# Linux File Systems

programs write/read files, rather than dealing with physical locations.
files are an abstraction camouflaging the physical layer. Files could actually be stored in a fragmented way on the storage device
dangerous to directly write to physical layer.
File systems can be in many places.
    local disk partition
    logical partition
    networked
    ...

All linux systems use an inverted tree hierarchy branching off from root. /
While the whole tree can be located in one partition there is usually mounted FSs and removable media
VFS is mounted within the tree also.

The fact that Linux can work with so many file systems has been a factor in it's success.

## VFS

Virtual File System
implemented by all modern OSs.
programs interact with this file system which translates the calls to the physical storage
    The physical file system implementation can be one of many implementations

## File Systems

ext4/3/2    Native Linux Filesystem
proc        for /proc
vfat        Windows. FAT, FAT32, etc
ntfs        Windows NT NFS, read only
udf         CD R/W, DVD
hfs+        Mac
jffs/2      Journaling flash filesystem
iso9660     cdrom
tmpfs       RAM disk that is swappable
gfs2        Clustering file system on red hat
nfs         Network Filesystem
smb         Samba networking
ncp         Novell Netware FS using NCP (Network Control Protocol) Protocol
coda        experimental distributed FS
afs         Andrew distributed FS from Carnegie Melon
ocfs2       disk cluster file system from oracle

### What FS's your OS currently uses

cat /proc/filesystems

when you mount a new filesystem the required module to understand it is loaded.

### Journaling File Systems

recover from system crashes or ungraceful shutdowns with little or no corruption.
and do so rapidly
operations are grouped into transactions
these transactions are atomic
log file of transactions is maintained

ext3 & 4, reiserfs, JFS, XFS, btrfs

### Special File systems

Linux widely employs the use of special file systems for certain tasks.

rootfs      no mount point      during kernel load
hugetlbfs   anywhere            provides extended page access
bdev        none                block devices
proc        /proc               Pseudo FS access to many kernel structures and subsystems
sockfs      none                BSD sockets
tmpfs       anywhere            RAM disk with swapping, re-sizing
pipefs      none                used for pipes
usbfs       /proc/bus/usb       used by usb sub-systems for dynamical devices
sysfs       /sys                used as a device tree
