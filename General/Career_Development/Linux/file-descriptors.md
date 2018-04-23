# file descriptors

abstract indicator used to access a file or other input/output resource
non negative integer

a process will generally have 3 file descriptors
    std_in, out and err

any time a process opens a "resource" a file descriptor is created!
    **need to verify this**

File Descriptor points to the inode which points to the actual file on disk

lsof can be used to find groups of file descriptors in lots of ways

/proc/<pid>/fd shows the file descriptors the pid has opened/.