# monitoring

linux distros come with a number of monitoring tools

most of the them make use of the psuedo file system /proc and /sys

## process and load monitoring tools

* top
* uptime
* ps
* pstree
* mpstat
* iostat
* sar
* numastat
* strace

## memory monitoring tools

* free
* vmstat
* pmap

## i/o monitoring tools

* iostat
* sar
* vmstat

## network monitoring tools

* netstat
* iptraf
* tcpdump
* wireshark

## tuning the system

/proc/sys contains information which can be changed to tune the system.
        vm is one of these folders which contains files that can be changed to tweak the system
                flush, swap, ...
        these values can be changed by directly writing to the files as root, or the sysctl utility.
        /etc/sysctl.conf you can set them on boot time.
editing these files is the same as executing sysctl commands to change configuration options.

## OOM Killer

Out of memory killer, used by linux to kill processes when the system runs out of memory.

each process has a badness score, in /proc/[pid]/oom_score.

the order of killing is decided by this score.

a user with root privileges can change this score using /proc/[pid]/oom_adj_score

## Disk bottlenecks

disk performance can be strongly linked to other factors.
insufficient memory or inadequate network hardware.
Admins can be misled, too slow i/o can be the cause too.
buffers filling up waiting for i/o to complete.

## tools

### sar

System Activity Report

backend is sadc ( System Activity Data Collector) which accumulates stats and stores them in /var/log/sa

sar [options] [interval] [count]

lots of options for what information to display

ksar, is a java based program for generating graphs from sar data

### ps

workhorse for process information retrieval
ps information is populated from /proc directory

there are 3 different option types:
UNIX    preceded by - and can be grouped
BSD     not preceded by - and can be grouped
GNU     long options, --

show all processes
BSD: aux
UNIX: -elf

processes surrounded by [] live entirely in the kernel
an integer after the name specifies the cpu they're on

#### output fields

most are self explanatory

VSZ:    process virtual memory size in KB
RSS:    resident set size, non-swapped phyiscal memory used in KB
STAT:   describes the state of the process
        S Sleeping
        R Running
        < high priority, not nice
        N low priority, nice
        L has pages locked in memory
        s session leader
        l multi threaded
        + being in the foreground process group

f option add ancestry to output of aux

-o lets you customize the output fields

### pstree

Gives a visual description of the process ancestry and multi-threaded applications

Note: child processes can be seen in the proc/<pid>/task folder

#### arguments

a       show command line arguments
A       Use ASCII characters to draw tree
p       show pids

### top

Interactive tool which shows you whats using up system resources currently.

top is quite old so there are newer more graphical tools built on top of top.
htop, ntop, atop

#### arguments

1       shows cpu's separately.
i       only show active processes.
q       quit
h/?     help
k       kill
r       renice
p       track the process id specified

handy trick:

* execute your process as a back ground process
* pipe the output to dev/null
* then echo the pid
* pass this into a call to top

```bash

top -p `{ <my_command> /dev/null & } && echo $!`

```

### vmstat

displays info about memory, paging, i/o, processor activity and processes

vmstat [options] [delay] [count]
repeat the report [count] times with [delay] seconds between each report.

#### output

r       number of processes waiting to be scheduled in
b       number of processes in uninterruptible sleep
swpd    virtual memory used
free    free memory
cache   cached memory
si      memory swapped in
so      memory swapped out
bi      blocks written to devices
bo      blocks read from deviuces
in      interrupts/second
cs      context switches/second
us      cpu time running user code
...

#### arguments

-S      memory metric, KB, MB, GB, ...
-a      show active and inactive memory
-s      table of memory stats (subset of memory stats which reside in /proc/meminfo)
-d      table of disk stats
-p      specify partition to get stats on

### dmesg

print kernel ring buffer
real time message buffer which eventually gets "persisted" to /var/log/messages

-H      human readable format
-L      colour the output
-w      continue printing messages as they come in

### iostat

basic workhorse utility for monitoring i/o device activity on a system.
can generate a report

#### output

brief CPU utilization summary
then table per device

tps     transactions per second
dm      device mapper

#### arguments

-k      output in KB
-m      output in MB
-N      show by device name
-d      no brief summary
-x      more detailed report

### iotop

table of current io usage and updates periodically, like top

prio column:
be      best effort
rt      real time

when you open it the > character shows what column is sorting on.
use the left and right arrows to change this

### arguments

-o      avoid clutter, only show active i/o

### ionice

set the i/o scheduling class and priority for a process

class refers back to the prio column in iotop

None/Unknown    0
Real time       1
Best Effort     2
Idle            3

#### arguments

-c      class
-n      priority, priority can range from 0 - 7. 0 is the highest
-p      pid
