# eBPF

(Berkley Packet Filter).
A new efficient in-kernel programming language.
Highly efficient sandboxed virtual machine in the Linux kernel, making the linux kernel programmable at native execution speed.
Extend existing kernel components or glue them together in a new form without requiring to change the kernel itself.
Can be used to turn linux into a modern microservice-aware operating system.
[Cilium](https://cilium.io/) is a tool built on it.
Jointly maintained by Cilium and Facebook, with collaborations from Google, Red Hat, Netflix, ...


The linux kernel is fundamentally event driven.
Processes make System calls, hardware makes interrupts.

Run a BPF program on an event
* kernel functions: kprobes
* user functions: uprobes
* system calls
* trace points
* network devices
* socket level interactions
* ...

BPF programs can communicate with each other and store state.
Stored in BPF maps.
Maps can be shared between programs, exported,
Types:
* hash tables
* arrays
* lru (least recently used)
* ring buffer
* stack trace

BPF helpers exist which allow BPF programs to interact with the kernel.
all safe operations, for retrieveing information.

BPF programs can be chained together.

BPF has a JIT compiler.

write in a high level language such as C

## who is using BPF for what?

Facebook
* L3-L4 load balancing
* network security
* traffic optimization
* profiling

Red Hat
* Replacing iptables with BPF
* NFV & load balancing (XDP)
* profiling & tracing
