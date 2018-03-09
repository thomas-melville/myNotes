# i/o scheduling

jiffies:    acronym is HZ
            term used by the kernel as a rough measure of time
            5 * HZ is 5 seconds

## scheduler

provides interfaces between the generic block layer and the low-level physical device drivers.
Both the VM and VFS layers submit i/o requests to block devices.
it's the job of the i/o scheduling layer to prioritize and order these requests before they are given to block devices.

### algorithm

has to satisfy certain requirements:

* hw access times should be minimized
* requests merged if possible
* low latency as is possible
* write can usually wait to migrate from caches to disk without stalling processes. Reads must always require a process to wait.
* Share i/o bandwidth in a fair way

These demands can be conflicting so there are different i/o schedulers for different workloads

At least one i/o scheduling algorithm must be compiled into the kernel:

* Completely Fair Queue
* Deadline Scheduling
* noop

SSD's are affecting the i/o scheduling algorithms used
to find out if the device is conventional hard disk or SSD check: /sys/block/<device>/queue/rotational

algorithm used by a device can be seen in the pseudo filesystem
    /sys/block/<queue>/scheduler
    the one is brackets is the active one.

the algorithm can be changed by echoing the wanted available algorithm into the file

## tuning

each scheduler exposes parameters which can be tuned
accessed through the pseudo file system /sys/block/<device>/queue/iosched
parameters depend on active algorithm, if you change the algorithm the parameters change

## Completely Fair Queue

goal of equal spreading of i/o bandwidth among all processes submitting requests.
Theoretically each process has it's own i/o queue
dequeuing of requests is done round robin with a FIFO order.

### tunables

quantum     max queue length in one round service
queued      minimum request allocation per queue
...

## Deadline Scheduler

aggressively reorders requests to simultaneously:
    improve overall performance
    prevent large latencies (limiting starvation)

each request has a deadline
reads get higher priority than writes

### tunables

read_expire     how long a read request is guaranteed to occue in
write_expire    same thing for a write
...