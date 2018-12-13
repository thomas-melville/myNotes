# file descriptors

abstract indicator used to access a file or other input/output resource
non negative integer

a process will generally have 3 file descriptors
    std_in, out and err

any time a process opens a "resource" a file descriptor is created!

File Descriptor points to the inode which points to the actual file on disk

lsof can be used to find groups of file descriptors in lots of ways

* lsof <path_to_file>   list all open file descriptors for a file
* lsof -u <username>    list all open file descriptors by a user
    it can take multiple users either in a comma separated list or multiple -u
* lsof -c <command>     list all open file descriptors by a command
* lsof -p <pid>         list all open file descriptors by a process id
* lsof -i               list all open network connections (sockets),
    add in the protocol to narrow it down
* lsof -i :<port>       find the process which is using the port
    add in the protocol before the : to narrow down

you can mix and match the above flags
-a AND any of ther above operations
negate with ^ before value

/proc/<pid>/fd shows the file descriptors the pid has opened/.