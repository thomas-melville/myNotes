# BPMN

## best practices

* when naming tasks, adhere to the OO design principle of verb + object
* when naming events, use the object and make the verb passive
* always create your process with start and end events

## Events

Events refer to something that has happened:
* regardless of the process (catching events)
* as a result of the process (throwing events)

### start event

  * defines where a process / sub-process starts
  * The engine requires at least one start event to instantiate a process
  * a start event may be declared as asynchronous
    * the process instance is created.
    * A job is created and asynchronously processed by the job executor
    * **All our flows are async.**
    * It allows us to kick off the workflow and return a 202 accepted.
  * start events are always *catching events*
  * Different types of start events:
    1. blank/none
    2. timer
    3. message
    4. signal
  * There can be at most one blank/timer event per process.
  * There can be multiple message/signal events per process.

### intermediate event

  * represent a status that is reached in the process.
  * used infrequently
  * can be useful to call out certain milestones, and measure time
  * intermediate events are always *throwing events*

### end event

  * mark the end of the process path.
  * end events are always *throwing events*

### boundary events

Catching events that are attached to an activity.

### Specific event types

#### none event

Unspecified events.
Can be a start/intermediate/end event.
A none start event technically means that the trigger for starting the process is unspecified.
This means the engine can't anticipate when the process instance must be started.
It is used when the process instance is started through the runtime service API
**Which is exactly how we start our processes**.

A none end event means that the result thrown when the event is reached is unspecified.
The engine won't do anything besides ending the current path of execution.

A none intermediate event means the engine just passes through it, not doing anything.
Can be useful to indicate that some state has been achieved.
A good hook to monitor KPIs, you can add an execution listener with your own logic.

#### message event

Events which reference a named message.
A message has a name and a payload.
Always directed at a single recipient.
The event in the bpm will have a child element, messageEventDefinition, which will have an attribute messageRef, which will refer to the message which is defined at the root of the bpm, definitions element.
The Camunda process engine fires the event, the receiving part is environment dependent.
* JMS Topic/Queue
* Webservice
* REST
You have to implement this as part of the application/infrastructure.
After you receive a message you can choose to
* employ the engines builtin correlation
* explicitly deliver the message to start a new proces instance
* trigger a waiting execution.
These are all done using the RuntimeService API, which looks like needs to be done before the process is started. As it correlates the message to the action

A message start event can be used to start a process using a named message.
You can deploy a process definition with one or more message start events.
The message name must be unique in and across the deployed process definitions.
Message start events are only supported in top-level processes.

A message intermediate catching event can be used to pause a process until a specific message is received.
A receive task can be used for this purpose also and can be combined with boundary events.
Combine this with a event based gateway.

Message boundary event can be used to wait for a message from an activity.
Can be interrupting or non-interrupting.

A Message intermediate throwing event sends a message to an external service

#### timer event

Events which are triggered by a defined timer.
They can be used as
* start event (create a process instance at a given time)
* intermediate catching event (start a timer, after a specific interval the flow continues)
* boundary events. (starts a timer, which can interrupt this activity if it fires before the activity is finished)
They are only fired when the JobExecutor is enabled.
They are configured using the ISO 8601 time format.
They can be a:
* fixed time & date
* duration
* repeating interval
The values can be pulled into the event from process variables.

#### error events

events triggered by a defined business error.
A BPMN error is meant for business errors, java exceptions are technical exceptions.
You need to define the error first, it consists of an id, errorCode and name.
  (You can also set the errorCode to be the package and class name of an exception)
then an activity in the process can specify it as an end event.
The JavaDelegate can then throw the error.

An error start event can only be used to start a sub-process.

## Tasks

* Tasks are at the heart of the process
* part of the activities category.

https://docs.camunda.org/manual/7.5/user-guide/process-engine/transactions-in-processes/#transaction-boundaries

### service task

used to invoke services
* either by calling java code
  * 4 ways of declaring how to invoke java logic:
    1. Specify a class that implements JavaDelegate / ActivityBehaviour
      * The package and class of the implementation
    2. Evalute an expression that resolves to a delegation object
      * ${className}
      * This is how we do it.
    3. invoke a method expression
      * ${className.doWork()}
    4. evalute a value expression
* or providing a work item for an external worker
  * specify the task type and the topic to publish it on

### Script Task

is an automated activity, the script is executed.
Specify the script and script format.
Script format must be compatible with JSR-223.
you need to add the corresponding engine onto the classpath.
the DelegateExecution object is available to the script so it can pull in variables from the process

## gateways

Gateways work as routers based on input(s)
No decision is made in the router, that is made in the activity before hand and passed on to the gateway.
A gateway without an icon defaults to an exclusive gateway.
You can also leave out a gateway and connect an activity directly to two subsequent activities and both will execute in parallel.
You can put a condition on each condition which is evaluated independent of the other connections.
All gateways and tasks can have a default flow.
It's defined by the default attribute on the gateway / connector.
You can see this in the modeler with a / on the line to the default flow

1. parallel gateway
  * split the process into two parallel flows.
  * you can have a synchronizing parallel gateway after both flows to ensure they both complete before the process finishes.
  * a single parallel gateway can both join and fork flows.
  * It waits at the gateway until all joining flows get there, then it execute the forked flows in parallel.
  * Any conditions defined are ignored.
  * It doesn't need to be balanced, i.e. there can be different numbers of joining and forking flows.
2. exclusive gateway
  * only one of the flows will be executed.
  * It can be if / else if / else
  * when the execution arrives at the gateway all outgoing flows are **evaluated in the order they were defined**
    * The first one that evaluates to true is selected.
  * If no flow evaluates to true a run time exception is thrown unless you have a default.
3. inclusive gateway
  * fork: evaluate all conditions and follow all flows which evaluate to true
  * join: what at the gateway for all flows, which have been executed. Slightly different to parallel!
  * A single inclusive gateway can both fork and join
4. event based gateway
  * allows you to make a decision based on events.
  * each outgoing flow needs to be connected to an intermediate catching event
  * there must be at least 2 outgoing flows.
  * The intermediate catching event must have a single flow.
  * outgoing flows are different to ordinary flows.
  * The sequence flows are never actually executed.
  * They allow the proces engine to determine which events an execution arriving at the gateway needs to subscribe to.

## subprocesses

A Subprocess is an activity that contains other activities, gateways, events, ... which itself forms a process which is part of a bigger process.
Subprocesses have 2 major use cases:
* hierarchial modelling
  * collapse the subprocess and hide all the detail.
* new scope for events
  * events thrown during the subprocess can be caught by a boundary event on the subprocess.
  * thus limiting scope.

Some constraints:
* only one none start event
* at lease one none end event
* sequence flows can not cross subprocess boundaries

a subprocess is visualized as a typical activity, rounded bold rectangle with a + sign

### embedded subprocesses

in the case of an embedded subprocess the + can be clicked to expand the subprocess.

### call activity

conceptually the same as an embedded process, a subprocess is called at execution time.
In the case of a call activity, it references a process which is external to the process definition.
The main use case is to have reusable process definitions
Visualized the same way as a collapsed embedded process, but with a thick border
The sub process can be deployed independently.
The calledElement attribute in the call activity element contains the process definition key, this means the latest version is always called.
you can pass variables in and out of the subprocess, we pass all. A subset of variables can be passed into the subprocess also.
variables are also passed out when a boundary error event is thrown.

### event subprocess

A subprocess that is triggered by an event.
The event used to trigger an event subprocess is configured using a start event.
none start events are not supported.
supported events are:
* message
* error
* signal
* timer
* compensation
The subscription to the start event is created when the scope hosting the event subprocess is created.
The scope hosting the event subprocess can be a process instance or another subprocess.
subscription is removed when the scope ends.
the event subprocess may be interrupting or non interrupting.
interrupting cancels any executions in the current scope. and can only be triggered once
non interrupting spawns a concurrent execution. can be triggered multiple times
can't have any incoming or outgoing sequence flows
visualized as an embedded process with a dotted outline
interesting use for error events, the event subprocess is still in the scope of the triggering process so the variables are all available.

### transaction subprocess

group multiple activities to a transaction (a logical unit of work which allows grouping of individual activities so that they either succeed or fail collectively)
