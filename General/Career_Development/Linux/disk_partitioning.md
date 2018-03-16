# disk partitioning

## hard disk types

characterized by the data bus they are connected to, as well as speed, capacity, how well multiple drives work simultaneously.

### SATA

Serial Advanced Technology Attachment

seen to the OS as SCSI devices
makes device drivers easier to implement
smaller cable size
hot swapping
faster and more efficient data transfer

### SCSI

Small Computer Systems Interface

main stay of enterprise servers for decades
lower capacity than SATA
but much faster and work better in parallel

### SAS

Serial Attached SCSI

newer point to point serial protocol

### USB

Universal Serial Bus

### SSD

Solid State Drives

have come down in price
no moving parts
use less power as no rotational media
faster transfer speeds

### IDE

Integrated Device Electronics

were the standard in consumer laptops and desktops for years
small and slow

## disk geometry

concept with a long history for rotational media
heads, cylinders, tracks and sectors.

## partitioning

disks are divided into partitions.
contiguous groups of sectors or cylinders
4 primary partitions
one of these can be extended into 4 logical partitions
SCSI supports up to 15 partitions
1-4 primary
5-15 logical
first SCSI/SATA drive is called sda, second sdb
the first partition on sda is called sda1, second sda2

### why partition

5 S's and a P

Separation
    separate user and app data from OS data
Sharing
    across partitions and machines
Security
Size
    some is consistent other is volatile
    separate them whic will improve performance also
Swap
Performance

### partition table

table is contained within the disks MBR (Master Boot Record)
the first 446 bytes are reserved for program code
    hold boot loader program such as GRUB

linux kernel interacts with disks through devices nodes, usually in /dev
normally only accessed through VFS

linux systems should use a minimum of 2 partitions
    /(root) for entire logical FS
    swap    extension of physical memory

## backup

dd command

## table editors

fdisk       menu driven editor
sfdisk      non-interactive, useful for scripting
parted      GNU program
gparted     graphical parted
