# fosdem 2019

It's been going 19 years
(First java room was in 2004)
This year it had 752 events, by 715 speakers
It's the biggest conference in the world

## links:

[fosdem](https://fosdem.org/2019/)
devtube(https://dev.tube/)

## talks

### The state of OpenJDK

Presenter: Mark Reinhold

Talked about changing the release cycle.
All the JEPs and how it'll end up being roughly the same amount just spread out every 6 months.
Take the red pill or the blue pill.
6 monthly / LTS every 3 years.
Even when on LTS test against latest build to spread out changes.
Talked about the core principles of the Java eco system, compatibility & readability.
Stewards & developers.

### Rumble in the java jungle

Presenter: Volker Simones

Talked a lot about JEPs, JSR, ...
Release cycle.
The different vm implementations.
Licensing change, complex.

### Weaving Computations

Presenter: Ron Pressler

Interesting new feature coming in java.
[Project Loom](https://openjdk.java.net/projects/loom/).
Still in prototype stage.
Fibers, lightweight threads in runtime.
Code like sync, run like async.
Continuation API.
  * remember where it was
  * when it calls yield it's stack is pulled out and stored.
  * when it is called again it's stacks is slotted back into the current threads stack
Integrates with Futures.
You can have nested Fibers.
They were given the freedom to rethink threading & concurrency
Structured concurrency, papers on it.
Should be done in a try-with-resources block (but it's not mandatory).

### Java Checkpoint & Restore

Presenter: Christine Flood

A jvm in comparison is slow to startup.
After startup it takes a while to warm up, GC, heap, ....
We want fast start up of a jvm.
Check point the jvm at this point in the time to the filesystem.
The jvm can be destroyed, but restored to it's previous state quickyl using the snapshot.
Applications:
  * backup & restore
  * replay crash
Uses the linux criu command, check restore in userspace?
This is barely a prototype, more a proposal.
Must be restored on the same machine with the same FD's otherwise it fails.
The files pointed to by the FDs must be at least the same size as the jvm was looking at a location in the file.
The talk turned into a design by committee meeting, people were throwing out ideas that she hadn't thought of.

### Java & Docker, making it work

Presenter: Adam Farley

Java was around before docker, the jvm is not container aware.
Jvm thinks it has resouces of host, when it only has resources of container.
Enter OOM Killer, bye bye container.
jlink to reduce java disk space usage.
He got it down to 30MB.
Ahead of time cache, persist this from the container so it can be stored and mounted again.
java11 resource constraint flags.
or openJ9 jvm.

### Securing the JVM, for fun or profit? Do we have a choice

Presenter: Nikolas Frankel

A private field can't be accessed, right?
Enter the Reflection API.
If you know the name of the field, setAccessible.
Java is Statically typed, right?
I forget the code but he changed the type & value of a variable.
Stenography, hide information in other data.
For example, hide code in images.
Image looks the same but the code can be executed, if the file is opend by a jvm.
Change bytecode on the fly, insert your own.
Restart the jvm & it's gone.
So, what can we do about it?
SecurityManager & Policy file.
jvm args to turn on SecurityManager.
start with most restrictive policy file,
only allow access that is required, everything else is blocked.
It's a lot of work!!!
It has to be noted as a Risk assessment.

### Performance tuning graal vm with ML & Twitter

Presenter: Chris ...

ML to tune vm parameters to increase Performance.
Bayesian Optimization as a service, Wheatlab (close source).
the graal vm is good with scala.

### Learning about Deep learning in Java

Presenter: Shelly Lambert

with a view to applying it to data from testing to improve the testing and predict failures.
functional coverage, instead of code coverage.
algorithm for what classes to increase testing on.
something like the number of changes for bug fixes.
services for pulling & analysing all data on failure.

### Misson Control & Flight Recorder

Presenter: Mario & Markus

Extend event & annotate with labels.
fire event to Flight Recorder.
Misson Control is used to analyse flight recordings.
The events can be exposed over JMX and Mission Control run remotely.
Opentracing is vendor neutral tracing API.
This can add contextual information to complement flight recording.
push tracing information in recorder.
Dapper, a google paper.

### Missing Benchmark

Presenter: Jens Wilks

JMH.
but no memory usage metrics.
object graph traversing.
jmap heap dump / histogram.
  invasive.
visualVM.
Microbenchmarks.
Linux OSS Metrics.

### Automate K8s workload with Ansible

Presenter: Michael Hrivnak

Ansible k8s module.
Ansible has templating engine too, jinja.
role, package related ansible code.
  * single app deployment
  * important is main.yaml and templates folder
Ansible playbook bundles (APB).
service Catalog.
Automation Broker.
  * uses APBs to serve up to Service Catalog.
Deploy, no day 2 mgmt.
svcat cli to sees available services.
some UI implementations, Openshift, can't remember the other one.
For day 2 mgmt you want an Operator.
An operator is a special purpose design controller.
  * with domain specific knowledge.
Use CustomResources to define your objects.
There is an Ansible operator.
The OperatorSDK, can write one in Go, using Ansilble or Helm.
Limited day 2 mgmt in Helm.

### Cloud Native Security 101

Presenter: Michael Hausenbian?

Know your attack vectors.
patch everything.
good base image that everyone uses.
secrets.

### OSS Security testing

Presenter: ...

Bug Bounty, funded by EU-FOSSA.
1 million euros.
Intigriti won the project to run it.
They picked a number of project, 15.
Examples: keepass, Kafka, 7zip, tomcat.

### TLS 1.3

Low latency, 1 RTT.
new features.
less algorithms.
OpenSSL.

### K8S new storage features coming

Presenter: John Griffith

The objects are PVC to PV to SC.
Persistent Volume Claim.
Persistent Volume.
Storage Class.
Container Storage Interface in 1.13.
Volume resize while in use.
Create snapshots from PVs.
transfer PVC across namespace.
  * while keeping the data
  * but will only transfer once original PVC is deleted.
Populators.
  * make PVCs declarative
  * say where to populate the data from, git repo, http?
  * this kicks off a job to do the work
  * requires taints on PV(C)s so it's not used until all the data is populated.
Custom Resource Definitions (External Controllers).
He used a tool to record and replace the terminal commands.
  https://github.com/asciinema/asciinema

### On observability

Presenter: Richard Hartmann

observable = inputs & outputs => deductions
1. Metrics & events
  metrics are numerical
  can create table, historgrams, ... from them
2. logs
3. traces
4. dumps
they are in order of when you should use them.
contain the complexity.
is it working -> monitoring.
why not -> observability.
SLI/O? Service Level Indicator/Object.
Current best practices:
  * blame free post mortem
  * (black box thinking - aviation industry)
He is working on open metrics.

### Grafana loki

Presenter: Tom Wilkie

Promethus, but for logs.
log aggregation, using alot of the existing prometheus functionality.
super early stages.
different direction to elastic.
It takes the information about the log and makes it the index, which doesn't scale.
Don't use high cardinality for labels.
log entries are lablled, with a few short labels, and timestamped.
include the namespace in the label.
They have promtail.
  * this is a daemonset which pulls the logs from each pod.
gRPC streaming.
Talked about reverser proxy for security?
Talked about cortex, https://github.com/cortexproject/cortex
  * "multitenant, horizontally scalable Prometheus as a service"
The demo was cool.
  * He gets an alert from Prometheus
  * looks at the data
  * pulls up another view, just like Prometheus
  * which has all the logs which are linked to the alert, same label
  * could quickly see what the cause of the alert was

### K8S security demystified

Presenter: Andrew Marlin

encrypt everything.
On the control plane.
  * PKI, PKC, X.509 & Diffie Hielman
  * self-signed certs in the cluster
  * beyondCorp
  * Zero Trust networks playbook  
On the data plane.
  * NetworkPolicy, L3&4
  * Istio is on L7
  * netassert tool
  * SPIFFE

### LXD Development

Presenter: ...

A run through of the new features in LXD in 2018.
LXD os one level up from LXC.
now on Chromebook.
There is a new version of cgroup coming out.

### Rootless K8s

Presenters: ... (from NTT in Japan)

"Usernetes"
running as an unpriviliged user.

User namespaces.
slirp.
vxlan.
fuse-overlayfs.
cgroup version 2.
docker 19.03 will support rootless.

### Nabla

Presenter: ...

container isolation.
horizontal attacks.
containers all share kernel.
one malicious container could access kernel syscalls and wreak havoc.
300 possible syscalls.
seccompo policy.
down to 7 syscalls.
lift up a lot of the calls into the userspace, libc.
strace, syscall trace.

### kubectl trace

Presenter: Alba ...

eBPF.
small program which can run in the kernel.
safe, not like modules, it has restrictions.
probes which can watch for calls.
k8s plugin.
runs on client, uses K8s objects to run.
it's a protype.
