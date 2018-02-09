# Processes

A process is an instance of a program in execution and associated resources.
Same program can be executed simultaneously by multiple processes.
tasks/threads within a process can share resources, the process is then multi-threaded.

not much difference between processes and threads in linux, they take about the same amount of time to start

first user process on the system is init, with the PId 1
second process on the system is kthreadd
* a daemon thread
* job is to create processes for maintenance work

every process has a:
pid     process id
ppid    parent process id
pgid    process group id

if the parent dies before the child then the childs ppid is changed to 1

init process allows zombied processes to die gracefully

max pid is 32768
can be changed by editing /proc/sys/kernel/pid_max

## process attributes

program being executed
context (state)
permissions
associated resources

kernel does the context switching of process states in and out.

## controlling processes

ulimit command

can change the max number of file descriptors, ...

## process permissions

process has permissions based on user who invoked program
It can also have permissions based on owner of program
s execute bit - effective user id

## process states

Running
    currently executing on the CPU or in the queue
Sleeping
    waiting on a request, usually i/o
Stopped
    process has been suspended (ctrl+z)
    the state of the program can then be examined.
Zombie
    the process has been terminated and it's parent hasn't enquired about it's state.

## execution modes

user / system (kernel)
this is down at the hardware  level

### user mode

lesser privileges than system mode
process is isolated in it's own user space
protects it from other processes

### kernel mode

cpu has full access to all hardware on the system
If an app requires access to these it must issue a system call which switches the context from user to kernel
file i/o, creating processes, ...
app never runs in kernel mode, it issues the system call

## daemon process

background process which provides some service to users
many are started at boot
most of the time end with d

## creating a process in the command shell

process creation in Linux is called "fork and exec"

* create a new child process with the program which the parent process is running
* make the exec system call to change the program the child is running from the parents one to the one wanted by the child.

1. a new process is forked from the command shell process (the command shell process becomes the parent) the command is the shell command
2. a wait system call puts the parent process to sleep
3. the command is loaded onto the child processes space via the exec system call, replacing the shell command
4. the command completes and the process dies via the exit system call
5. the parent shell process is re-awakened by the death of the child process

use & to put the new process into the background allowing the shell to continue to issue commands.
some commands are built into the shell so don't require loading program files, echo and kill

## kernel processes

not all processes are forked from user parents.
kernel creates 2 kinds of processes
* Internal Kernel processes
** maintenance work: flushing buffers, load on cpus is balanced.
* External user processes
** rare and short lived

## process priorities

can be controlled through the nice and renice commands

the higher the nice value the lower the priority as the process will yield to others
-20 to +19

decrease nice can only be done by a superuser
increase can be done by anyone

child process also gets updated nice value from parent