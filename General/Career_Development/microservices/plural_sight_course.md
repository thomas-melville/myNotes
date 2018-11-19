# plural sight microservices course

https://app.pluralsight.com/library/courses/microservices-architectural-design-patterns-playbook/table-of-contents

## bounded contexts

Borrowed from Domain Driven Design
A strategic design pattern

### DDD

Design software which models real world domains
There are many tools & techniques in it
we'll just use bounded contexts in this course

microservice per bounded context

### BC Technique

"A specific responsibility enforced by an explicit boundary"
a microservice has a SRP which maps to a business function

* identify core domain concepts / functions
* each bounded context has an internal model
* each bc has an explicit interface
* communication between contexts is carried out using shared models for communication
* boundary becomes the contract

1. Round up the domain and software experts
2. identify core concepts of the domain
   * starting point for bounded contexts
3. within each core concept start identifying concepts which support it
   * these form a ubiquitous language
4. does each sub concept fit the language of the core concept
   * if not:
     * change the language
     * move it out
     * make it a core concept
5. if there are concepts in a context which don't relate to anything else in the context then it could be an integration point with another context.

### Ubiquitous Language

from DDD again
same language defined and used by everyone talking about a particular bounded context

## asynchronous communication

### why

fire and forget for long running job and you can't afford to wait
decouples client from service
better user experience because less time waiting, makes your system more responsive

### options

#### event based

don't call another api / service directly, send an event to a message broker
the other services are subscribed to the message broker
multiple services can take action on a single event

events are raised as messages, which are placed onto queues in message brokers
consumers of these messages are services which are going to carry out these tasks
  multiple consumers can pick up the same event

message broker can place the message back on the queue if it doesn't receive a positive response from the consumer

RabbitMQ is the most popular Message Broker.
I see the message has a correlation ID! when we have multiple services involved in a single user action we should make use of correlation IDs. They are headers.

##### Queueing patterns

###### Competing Workers Pattern

most microservices friendly pattern
allows scale out and load balance
only one of the consuming service will get a message
there is redundancy

###### Fanout Pattern

all consumers listening to queue/exchange get the message
and all consume it
useful when multiple things need to happen based on a single event

#### asynchronous API calls

request/acknowledge using callbacks
call made to service, but don't wait for the response
you get a default response
when making a request to give a callback address, which the service will use to notify you that your request is complete
The response the client wants is in the body of the request to the callback

## Architecting API based microservices

### Functional Requirements

* Autonomous microservices principle
* * loosely coupled
* * independently changeable
* * independently deployable
* * support and honour contracts
* * technology agnostic API
* * Stateless API

### Architectural Styles

#### Pragmatic REST

more refined and slightly extended version of REST, which is more suitable for microservices.
extended unofficial version of REST.

* not always about resources & CRUD
* action/task based endpoints
* * differentiate between resource & action/task by using verbs instead of nouns
* * they seem to flip the URI structure I've been using
* * * me:     /resource/id/generate
* * * course: /resource/generate/id
* parameters to action can be in query string or request body
* response body differs for resources & action/tasks
* * output V callback

#### HATEOAS REST

True REST, slightly idealistic
very hard to implement, and apparently hard to use also

#### RPC

Satisfy less requirements so not investigated further
Although, gRPC looks to satisfy more of the requirements so could be an option

#### SOAP

Satisfy less requirements so not investigated further

### Architectural patterns

#### Facade

A single interface to subsystems
The API we define in RAML and the Interface it generates is our Facade
each subsystem is a class in our microservice
the only way to access them is through the API

#### Proxy

Placeholder for another object to control access to it.
Other object could be an external API
Proxy object doesn't contain business logic
a wrapper which provides:

* simplification (encapsulates access to object)
* transformation (map retrieved object to local context)
* security
* validation

#### Stateless service

state is session id for example
no state is maintained in the microservice
clients maintain state
advantages:

* scalability
* performance
* availability

caveat

increased network traffic

## Composing Microservices together

how communication occurs within your micro service architecture

### Patterns

#### Broker

Introducing a Message Broker to your architecture
when you want to bring in asynchronous communication

#### Aggregate

Take data from multiple services and aggregate it into something useful to the client
For example a web app that pulls together info from multiple services and returns a page to the user
aggregator can also be a backend service
which could be an async call so that the user is not blocked while the data is aggregated

#### Chained

services calling downstream services synchronously
keep the chain small, and the responses small and quick

#### Proxy p

single gateway service which delegates all calls to backing services
This can be useful when you're exposing multiple services to external clients.
One entry point, instead of multiple
all security and authentication can be built into the proxy service

#### Branch

an aggregator which splits the request into multiple requests which are serviced by different backing services

## Data Consistency across Microservices

One interaction from the user may require interaction with multiple services which store data in multiple schemas
Traditionally this was done using transactions, ACID

Transactions can be complicated in a microservice architecture, because it is distributed, across databases.
and the network is unreliable
CAP Theorem, microservices can be AP or CP. Since it is distributed over the network it can never be CA
Availability at the cost of Data Consistency
Data Consistency at the cost of Availability
Different microservices in the product can be either or.

### Two Phase Commit

ACID is mandatory
CP, chooses consistency over availability

Tries to make distributed transactions as reliable as non-distributed
Uses a transaction manager

1. Prepare phase
* * asks services to prepare
2. Voting phase
* * all services vote to handle the transaction
* * if transaction mgr receives yes from all services it issues a commit
3. Rollback phase
* * if one service says no the mgr issues a rollback all other services undo the prepare phase

very complicated and error prone, stay clear of it if possible

caveats:

* reliance on a transaction mgr
* no voting response from service
* commit failure after a successful vote
* pending transactions lock resources
* avoid custom implementations, take an off the shelf product
* has scaling out issues
* overall reduced throughput

It's really an anti-pattern, so really think about it before using it.

### Saga Pattern

Trades Atomicity for availability and consistency

Splits a transaction up into many requests, all these requests make up the saga
track each request using the saga log, it drives the requests too
compromises atomicity
first described in 1987
also a failure mgmt pattern
what to do when one service fails
compensate requests, if one fails compensating requests are sent for all the requests which succeeded
SEC, Saga Execution Controller manages all this

#### Implementation

Saga log can be a service with a database to keep track of sagas
SEC listens and writes to the Saga log
SEC must be recoverable if it fails
compensation requests must be idempotent

### Eventual Consistency Pattern

Compromises ACID
AP, chooses availability over consistency

It's more an approach than a design pattern
what data updates don't have to happen immediately?

BASE model, in complete contrast to ACID

Basic Availability
Soft-State
Eventual Consistency

avoids resource locking
ideal for long running tasks
  improves user experience

eventual data consistency will normally happen within seconds

traditional approach was data replication

in microservices world it is event based
fire and forget
message brokers and queues

another approach is to write to a database and have other services read entries from it.
They then would be coupled together

## Centralizing access using an API gateway

central component to access all internal microservices
so don't expose all the micro services

### Why

a facade for all the services so a user makes one calls which triggers multiple calls in the background
pulling data together to return to the customer
consolidate security into one place
multiple network connections to multiple APIs

abstracting complexity from the client.

### API Gateway

component to access API
it's an API itself
RESTful API
security/authentication
pass on authorization tokens
traffic control
load balancing
logging & monitoring
transformations
caching
manage versions of internal APIs

API Gateway is placed in a DMZ zone
DMZ: an area in your network which is exposed to the internet

#### Kong

open source
lots of plugins
plugin for correlation ids!
retry policies built in
nice UI for configuring it

## Split monolithic databases

### why are they bad

coupling microservices so they can't be independently changed or deployed, ...
harder to scale out
performance bottleneck

### approach

database per service
micro database
only hold data owned by that service

avoid data first design, leads to tight coupling
traditional approach
it's a microservices anti-pattern
risk of creating anemic crud based services, lacks business logic
more likely to expose internal data structures

use function first design, hold off on database until context is more stable
for loose coupling
provide services that provide more than just CRUD, business logic
function defined by the scope of the service
Use **Bounded Contexts** design tool
separate internal models from external

Tools to support code first data design
Write the models, then use a tool to generate the database

### design patterns

#### sharing data using events

share data changes using events
it's a way to avoid shared databases
when a service stores data it sends an event to say this data has changed
interested parties subscribe to the event and either

* update local data to match it, for data joins
* trigger a task based

avoids anemic crud
push is better than pull for performance and user experience
asynchronous
well established message brokers

**calling another service for some data is another form of coupling.**

#### store data as events

traditional approach is to store the state of the data in the record in the database

event sourcing stores the state as a series of events which all sum up to the current state
replay the events to get the current state

why?
more efficient when broadcasting these events
records a history of the state to re-create the state at a certain point in time

challenges!!

replaying more events decreases performance
  consolidate to flat record older than a certain date
this also helps with data storage space

#### separate write model from query model

Command Query Response Segregation pattern
Separate command and query
two separate services

command service receives events from a queue, scalable

##### why

separate responsibilities
better performance
separate technologies
isolate read & write models

##### challenges

sync'ing databases
eventual consistency

### green field

* bounded context & code first tools
* * leave modeling the data until the last possible moment
* event notifications to share data
* event sourcing
* split to CQRS

### brown field

separate schemas??

* Pick a table which is heavily shared and develop a service for it
* refactor apps to go through service for data
* move the table out of the database to a micro database / separate schema
* repeat

very hard to retrofit greenfield patterns
event sourcing and CQRS

## Make microservices resilient

### Why

distributed architecture.
more moving parts.
any one of them could fail.
the network is unreliable!
fault tolerance is required.
avoid cascading failures.
avoid exhausting resource pools.

#### Types of failures

hardware.
infrastructure.
communication layer.
dependencies.
one of your own microservices.

### Patterns and approach

#### Timeouts

every call to an external resource should have a timeout set
timeout exception is thrown after timeout is exceeded instead of waiting forever for external resource to respond
reduces chances of cascading failures
do not rely on default value, the timeout value has to be just right!
can be used with retries

##### advantages

avoid waiting forever
threads aren't locked up
allows you to handle the failure
  degrade / default / bomb out back to user
  depends on the criticality of the external resources func

still a risk!
Multiple threads can be blocked on timeout, next pattern will help with that

#### Circuit Breakers

used with the timeout pattern.
when over usage is detected break the circuit.
an object which wraps calls.
monitors calls for failures.
trip after reaching threshold.
  returns an error without making the call.
  no longer waiting for timeouts.
thresholds can be different per exception from call.
reset after x number of tries.
test the water with a few calls before closing the circuit.
  the number of calls before retrying is the reset parameter.
you decide what to do when the circuit is open.
use for all synchronous calls
always report, monitor and log

hystrix is mentioned as a recommended library
thepollyproject.org is a .NET library

#### Retry

ideal for transient faults

* momentary loss of connection
* temporary unavailability of service
* timeouts when service is busy
* service throttles accepted requests

allows faults to self correct

Retry strategies

* retry
* retry after delay
* cancel

Log and monitor!!!!

they all can be used together!
3pps that handle all them

#### Bulkhead

High level design pattern when architecting your service.
how do you contain and handle a failure.
Separate failed components from the rest of the system.
ensure degraded functionality.

How?

* Separate by criticality
* isolated microservices
* isolated databases
* redundancy using load balancers
* circuit breakers
* Rejecting calls using load shedding

#### approaches

design for failures
embrace failures
fail fast
degrade or default functionality

## Questions

* When it's consumed by multiple consumers how does the message broker know to take it off the queue?