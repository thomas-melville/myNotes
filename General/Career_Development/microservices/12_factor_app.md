# 12 factor app

a methodology for building software-as-a-service apps:
	use declarative formats for setup automation
		to minimize time and cost for new developers
	have a clean contract with the underlying OS
		offering maximum portability between execution environments
	suitable for deployment on modern cloud platforms
		obviating the need for servers and sys admins
	Minimize divergence between dev and prod envs
		enabling continuous deployment for maximum agility
	can scale up without significant changes to tooling, arch or dev practices

## Codebase

**One codebase tracked in version control, multiple deploys**

Always use version control
one-to-one correlation between the codebase & app
if an "app" has multiple codebases then it's a distributed system and each component within the system is an app which complys with 12 factor app
many deploys of the app, dev, stage, test, prod
one-track

## Dependencies

**Explicitly declare and isolate dependencies.**

Never rely on the implicit existence of system wide packages.
declare all dependencies completely and exactly, via a dependency declaration manifest.
Use a dependency isolation tool during execution to ensure no implicit dependencies leak in
That's why microservices are bundled into uber jars, all runtime dependencies are in the jar.
Example: Python
	Declaration: Pip
	Isolation: VirtualEnv
		 Node
	Declaration: npm & package.json
	Isolation: ????
		 Java
	Declaration: some build tool?
	Isolation: use shade plugin to compile into uber jar / OSGi

_(would a container / vm also be regarded as a dependency isolation tool?)_

This also simplifies setup of the app for new developers
pull down the codebase and build/install using the build tool of the project

## Config

**Store config in the environment**

everything that is likely to vary between deploys
Separate config from code, no config as constants in the code!
There is config that won't vary between deploys so this is only concerned that config which will.
	Credentials, File System paths, database info, host information

Store config in environment variables.
	easy to change between deploys,
	won't be checked into version control

env vars are never grouped by environment, this can cause an explosion of configs.
env vars are granual controls. orthogonal to other env vars

This is all about reading the config from the environment, how you populate it is up to you
	The config has to be somewhere, just not with the app

## Backing Services

**Treat backing services as attached resources**

A backing service is any service the app consumes over the network as part of it's normal operation.
Examples: databases, message buses, smtp, caching.

_another service?????_

Treat them as attached resources, accessed via url or other locator/credentials in the config.
Should be able to swap out any of them without a change to the apps code.

## Build, release, run

**Strictly separate build and run stages**

3 stages:
	build
		transform the repo into an executable
	release
		combines the executable with the config and deploys it on the environment
		it is ready for immediate execution
	run
		run the app in the environment

Code can't be changed at runtime and propagated back to the code base.
Releases are append only and immutable.
	If there's an issue a new release

Builds are initiated by developers when there are code changes
Runs can happen automatically for example when there's a crash so should be kept simple

## Processes

**Execute the app as one or more stateless processes**

processes are stateless and share nothing
Any data that needs to be persisted is stored in a stateful backing service
memory/fs of process may be used as a brief single transaction cache
never assume anything cached will be available for a future job
With many processes of each type running, chances are the future job will be served by a different service
Sticky-sessions are a violation of the 12-factor
Session state data is a good candidate for a datastore that offers time expiration, memcached or redis. Store the data outside of the running code.
	Then the database becomes a single point of failure, unless you have some HA on that!

This seems to be more about state than processes, yes it is!

The file system is not a backing service!

## Port Binding

**Export services via port binding**

12 factor app is completely self contained, it doesn't run inside a web container provided by someone else.
Export http as a service by binding to a port, and listen to requests coming in on that port.
In dev that would be localhost:<port>, in production a routing layer would handle routing requests from the public hostname to the port-bound web processes
This is typically down by declaring a web server library as a dependency of the app (spring-boot webserver dependency for example which uses tomcat by default, but can be changed to use jetty or undertow)
	the app itself defines and includes the web server it will use

Apps aren't just web services, a service can listen on any port. Http is just the most common protocol.

This enables apps to be backing services for each other.

## Concurrency

**Scale out via the process model**

processes are a first class citizen
take strong cues from the unix process model for running service daemons
	processes should be managed, not just started up. service / upstart / ...
Each type of work can be assigned to a process type. web / worker / ...
A process is a service.
	This makes it easy to scale out the number of processes using a platform tool, kuberentes, ...
	since the process is stateless

I think this is stating that since processes will be stateless use an underlying manager to scale them out and in as needs be.

## Disposability

**Maximise robustness with fast startup and graceful shutdown**

processes are disposable, started and stopped at a moments notice. this facilitates fast scaling, rapid deployment of code/config changes, and robustness of production deploys

processes should strive to minimize startup time, a few seconds max!

processes should shutdown gracefully when they receive a SIGTERM signal from the process manager, (process manager is defined in Concurrency!)
	Example: Http
		1. stop listening for requests
		2. finish any current requests
	Example: Worker
		1. return job to queue
		2. release any locks

processes should also be robust against sudden death

## Dev/prod parity

**Keep dev, stage and prod as similar as possible**

Historically there have been gaps between dev & prod
3 areas:
											trad 	12 factor
	time:		between dev and prod 		weeks	hours
	personnel: 	dev devs and ops prod 		diff 	same
	tools:		tech stack is different		diverge similar

Backing services is where parity is important. There could be issues (func & perf) if a different backing service is used between envs
A light weight backing service can simplify development but inconsistencies and failures can crop up which disincentives CD

Lightweight local services are less appealing as the real things are easier to install now. (Homebrew, apt-get)
Alternative provisioning tools + virtual envs makes it easy to replicate the env on a laptop

## Logs

**Treat logs as event streams**

Logs are normally written to a file, but that is just an output format.

Logs are the stream of aggregated, time orderer events collected from the output streams of all running processes and backing services
Logs flow as long as the app is running

A 12 factor app should never concern itself with routing/storage of its output stream.
Write it unbuffered to stdout.
During dev the developer can view the logs in their terminal
In stage/prod, each processes stream will be captured by the execution environment, collated together with all other streams from the app 
and routed to one or more final destinations for viewing and archiving
The app has no idea where the stream goes
Log indexing, Splunk
Data warehousing, Hadoop

## Admin processes

**Run admin/mgmt tasks as one-off processes**

An app has an array of processes which run the apps regular business.
Separately, developers will wish to do one-off admin/maintenance tasks for the app:
	database migration
	run a console to execute commands/scripts. a REPL (Read, Evaluate, Print, Loop) shell

One off admin processes should be run in an identical environment as the regular running processes.
run against a release, using the same code base and config as any process run against the release.
admin code must ship with application code.
same dependency isolation techniques should be used.

12 factor apps strongly favour languages which provide a REPL shell out of the box