# lvm

Logical Volume Management

(RAID: Redundany Array of Independent Disks)

lvm allows to takes bits of multiple partitions, even on separate hard disks, and put them together into one virtual partition.

+ives

* easily change size of logical partitions and FSs
* ,,     add more storage space
* ,,     rearrange things

one or more physical volumes are grouped into a volume group
volume group is subdivided into logical volumes
which mimic to the system nominal physical disk partitions
can be formatted to contain mountable FSs

lvm does impact performance
overhead of the lvm layer

lvm has better scalability than RAID
lvm can be built on top of RAID

space in the volume group is divided into extents

when resizing the number of extents is just in/decreased.
when expanding the volume is expanded first, then the FS
when contracting the FS is contracted first, then the volume
the -r option in lvextend and lvreduce will bundle both actions into one command

snapshots of a lvm can be taken
these are useful in backups, application testing and deploying VMs.
only use space for storing deltas

## commands

create and manipulate volume groups

vgcreate
vgextend
vgreduce
vg*

manipulate phyiscal partitions in volume groups

pvcreate
pvdisplay
pvmove
pvremove
pv*

create and manpilate logical volumes

lvcreate
lvextend
lvreduce
lvremove
lv*