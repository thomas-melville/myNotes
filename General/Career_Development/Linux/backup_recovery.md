# backup and recovery

backup is important as eventually your data will get corrupted / lost.
Due to hw failure, human error or malicious actors.

## what to backup

1. Definitely
* * Business related data
* * System Config files, /etc
* * User files, /home/<user>
2. Maybe:
* * Spooling directories, printing, mail, etc.
* * logging files, /var/log
3. Probably not:
* * sw that can be easily reinstalled
* * /tmp
4. Definitely not:
* * pseudo file systems, /proc, /dev & /sys
* * swap partitions

backup media have finite life times:

* magnetic tapes:   10-30 years
* CDs & DVDs:       3-10 years
* Hard Disks:       2-5 years

tapes are mostly used for long term archival now, as they are too slow for regular backing up
day to day backups are done with some form of NAS or cloud solution

## methods

always store backups separate from system being backed up.

* Full
* * backup all on files for the system
* Incremental
* * backup all files which have changed since the last incremental or full backup
* Differential
* * backup all files which have changes since the last full backup
* Multiple level incremental
* User
* * backup files only in a users home directory.

## strategies

* full backup once, and then incremental backups
* full backups take time
* restoring from incremental can be difficult and time consuming
* use a mix

## utilities

### archive

#### cpio

* Copy In and Out
* create/extract archive files
* around since the early days of unix
* designed for tape backups
* even though there are newer utilities it is still used today
* * rpm2cpio
* * kernel uses it internally
* lighter than successor
* not as robust

-o / --create   create an archive
-i / --extract  extract an archive
-t / --list     list the contents of an archive
takes a list of files, and can be patterns

#### tar

* create/extract archive files
* for each directory given as an argument it creates an archive with all directories and sub directories
* -f space separated list of directories
* when restoring it reconstitutes directories as necessary
* -N / --newer option which allows incremental backups
* --after-date is another possibility for creating incremental backups
* each one requires a date
* only considers modifications to the contents of the file
* not changes to the file name or permissions
* -M create the backup on multiple volumes
* -x for extract
* you can specify what files / directories you want to extract
* -p ensure files are restored with the same permissions
* -t list the contents of an archive

### compression

in order of increasing compression efficiency at the cost of longer compression times

* gzip
* * Uses Lempel-Ziv Coding
* * .gz file extension
* * add the z argument to tar to compress using gzip
* bzip2
* * Burrows-Wheeler block sorting text compression algorithm and Huffman coding
* * .bz2 file extension
* * add the j argument to tar to compress using bzip2
* xz
* * .xz file extension
* * add the J argument to tar to compress using xz

You can add the option when extracting to specify the compression type but tar is smart enough to figure it out

It's more efficient to do the compression as part of the archiving as

* there is no intermediate file storage
* archiving and compression happen simultaneously in the pipeline

### copying

#### dd

transfer raw data between media
one of the original UNIX utilities, very versatile.
very low-level raw copying of files or even whole disks
can do many kinds of data conversion during copying

### synchronization

#### rsync

remote synchronization
sync directory subtrees or entire filesystems across a network, or between different file locations on the same machine
rsync is clever, checks local files against remote files in small chunks
very efficient when copying across, only copies differences
use the -r option to make it recursive through a directory tree
rsync has embedded compression features like tar

this command is a simple and effective backup strategy.

### dump & restore

around since the early days
directly read the filesystem which makes them more efficient

+ives:

* can perform full/inc backups
* understand specific filesystem format and how to read & write it
* head movement reduced in full backups
* can specify output tape size and density, block size and count or both
* cam dump to any valid device
* parameters in /etc/fstab decide what gets dumped and when

-ives:

* multiple file system passes are required for backups
* only works with ext2,3,4
* can't be run safely on mounted filesystems
* must be restored only onto the same file system

dump level is set by integer passed to command, 0-9
some other command arguments

restore is used to read archives created by dump
-r restore everything
-t list contents
-x extract specific files and directories
-i interactive mode

### manage tape drives

#### mt

control magnetic tape devices
querying and positioning tapes before backup and restore

## Backup Programs

* Amanda
  * Advanced Maryland Automatic Network Device Archiver
  * uses native utilities, cpio & tar
* Bacula
  * automatic backup of heterogenous networks
* Clonezilla
  * robust disk cloning program