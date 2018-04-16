# ext4

## history

ext2 FS was linux's original native variety
    available on all
    but rarely used
ext3 FS was the first journaling extension of ext2
    same on disk layout
    except for presence of journal file
ext4 FS was officially included in version 2.6.28 of the kernel
    significant enhancements
        extents for large files, instead of lists of file blocks
            an extent is a group of contiguous blocks
                can improve large file performance
                and reduce defragmentation
    default FS on most enterprise Linux distributions

close correspondence of features between VFS and ext2/3/4 FSs, as each has been designed with the other in mind.

**Block size** is selected when the FS is built, 512, 1024, 2048, 4096 (by default)
    that's the best choice for an average file system
    minimizes disk access.
    unless there are many tiny files.

**Linux Kernel Memory Management System** requires only an integral number of blocks must fit into a page of memory, can't have 8kb blocks on x86 where pages are 4kb

**inodes** number can be tuned, to save disk space.

**inode reservation** consists of reserving several inodes when a directory is created, expecting them to be used in the future.
improves performance because inodes are already allocated for the new files.

**fast symbolic link** is created and stored in the inode if the link is less than 60 characters

## layout

all fields are written to disk in little-endian order, except the journal.
disk blocks are partitioned into block groups, each of which contains inode and data block stored adjacently. lowers access time.

layout of a standard block group is simple
group 0, first 1024 bytes are unused. to allow for boot sectors, ...
the super block will start at the first block, except for block 0
followed by the group descriptors and a number of GDT (Group Descriptor Table) blocks.
data block bitmap, inode bitmap, inode table and data blocks.

data blocks are pre-allocated to files before they are even created
so, when a file's size increases, adjacent space is available and fragmentation is reduced.

superblock contains bit-fields which are used to ascertain whether or not the FS requires checking when it is mounte:

* clean:    cleanly unmounted previously
* dirty:    usually means mounted
* unknown:  not cleanly unmounted

in the latter two cases fsck will be run to check the filesystem

first and second block are the same in every block group
under normal circumstances only those in the first block group are referenced by the kernel.
The rest are copies
which are only checked if there is a problem with the previous one

number of blocks is constrained by the fact that the block bitmap has to fit in a single block.

data block and inode bitmaps are blocks whose bits contain 0 for each free and 1 for each used

inode table contains as many consecutive blocks as are necessary to cover the number of inodes in the block group

## commands

### dumpe2fs

scan file system information

### tune2fs

change file system parameters
argument for each parameter which can be changed