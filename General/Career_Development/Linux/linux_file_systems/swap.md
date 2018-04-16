# swap

Linux employs a virtual memory system.
OS can function as if it has more memory than it actually does
2 ways

* many programs do not actually use all the memory they are given permission to use
* when memory becomes important, less active memory is swapped out to disk.

in most cases recommended swap size is RAM size on the system.

## commands

mkswap  format a swap partition or file
swapon  activate swap on a partition or file
swapoff deactivate swap on a partition or file

## quotas

control max space users or groups are allowed
can be assigned per FS

### commands

quotacheck  generates and updates quota accounting files
quotaon     enables quota accounting
quotaoff    disables quota accountinhg
edquota     used for editing user/group quotas
quota       reports on usage and limits

will only work on FS's mounted with user and group options