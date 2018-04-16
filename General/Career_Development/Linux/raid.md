# raid

Redundant Array of Independent Devices
spreads i/o over multiple disks
increase performance
can be implemented in software or hardware
hardware raid is more efficient than software raid
os is not aware of hardware raid

one disadvantage in hardware raid
if the disk controller fails, it must be replaced with a compatible disk controller which may not be easy to obtain

3 essential features

* mirroring     writing the same data to more than one disks
* striping      splitting of data to more than one disk
* parity        extra data is stored to allow problem detection and repair, yielding fault tolerance

one of the main purposes is to create a FS that spans multiple disks

+ives
better performance, redundancy or both
    striping allows multiple parallel writes
    mirroring gives better redundancy

## RAID specifications

0       uses only striping
        data is spread across multiple disks
        no redundancy
        no stability / recovery capabilities
1       uses only mirroring
        each disk has a duplicate
        good for recovery
        at least 2 disks required
5       uses rotating parity stripe
        single disk failure will cause no data loss
        at least 3 disks required
6       striped disks with dual parity
        can handle loss on 2 disks
        at least 4 disks required
10      mirrored and striped data set
        at least 4 drives required

## commands

### mdadm

used to create and manage RAID devices