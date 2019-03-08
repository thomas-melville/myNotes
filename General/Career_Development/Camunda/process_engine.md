# Process Engine

https://docs.camunda.org/manual/7.5/user-guide/process-engine/

A java library responsible for executing BPMN 2.0 processes.
A lightweight POJO core with an RDB for persistence.
ORM mapping is provided by the MyBatis mapping framework

## bootstrapping

via Java API
xml config
via Spring

The engine uses the ProcessEngineConfiguration bean to configure and construct a standalone ProcessEngine.
Multiple subclasses available, represent different environments and set defaults accordingly.

* StandaloneInMemProcessEngineConfiguration
  * useful for unit testing!
  * H2 in memory database used by default
* SpringProcessEngineConfiguration
  * Config to be used in a Spring environment.

## Process Engine Concepts

### Process definitions

A process definition is the Java counterpart to the BPMN 2.0 process.
A representation of the structure and behaviour of each of the steps of a process.
Each process has a key, a logical identifier for the process, it is the id attribute in the process element in xml
You can deploy several process definitions with the same key, different versions.

### Process instances

An individual execution of a process definition. Same relation as Object and class in OO programming.
The process engine is responsible for creating process instances from process definitions and managing their state.
You can suspend a process instances

### Executions

An execution can be imagined as a token moving through the process instance.
The runtime state of a process instance is represented by a tree of executions.
Executions are heirarchial and span a tree from the root, the process instance itself.
Executions can execute in parallel.
Executions are variable scopes. Thery hold the variables which have been added, variables can shared across tasks using the execution

### Activity instances

similar to execution concept, but a different perspective.
an activity is an individual instance of a task, sub-proces, ...

The execution and activity concepts are perfectly aligned
n-1, executions to activity instance

## Process Engine API

### RepositoryService

Offers operations for managing and manipulating deployments & process definitions.
Suspend and activate process definitions.
A deployment is the unit of packaging in the engine, it is all the BPMN xml diagrams packaged up and deployed in the engine.

### RuntimeService

*This is the first service I see in our repo, we start instances by key*
Deals with starting new process instances of process definitions.
Retrieves and stores process variables, data specific to this process instance.
Also allows to query on process instances and executions.
An execution is a pointer to where the process instance currently is.
Lastly used for waiting for external triggers

### TaskService

Manages tasks which need to be performed by human users
A workflow can be made interactive and allow user intervention with user tasks by setting the isInteractive execution variable to true

### IdentityService

mgmt of groups and users
you can link into an LDAP

### HistoryService

exposes historical data gathered by the engine about process instances

## Process variables

Variables in process instances, consists of a name & a value.
All entities that can have variables are called variable scopes.
Child executions can access the parents variables, but not vice versa.
If there are two child executions, if one updates a parent variable the other child will see the update.
2 Java APIs to set & retrieve variables.
* Java Object API
  * directly store primitives and objects.
* Typed Value API
  * Store a "typed" value
  * Wraps a variable in a Typed Value
  * The object/value and some metadata
  * Accessed through the *Variables* class of static methods
  * Useful when it's important to access a variables serialized representation
    * or in which the engine has to be hinted to serialize a value in a certain format.
  * Also useful for storing files as BLOBs with metadata:
    * type, encoding, ...
  * Also useful when serializing a custom object which the calling application does not possess the actual classes.
There are limitiations on the values for primitives and strings that can be stored.
* It is a DB constraint, for example text or 4000 characters can't be stored in out Postgres configura
When an object value is passed to the engine, a serialization format can be specified.
There is a Camunda plugin, Spin, for serializing to XML and JSON

## Delegation

Allows you to execute code/scripts outside the workflow

3 types
* *Java Delegate* implementations can be attached to a BPMN Service Task
* *Execution Listeners* can be attached to any event within the normal token flow
  * start, enter, exit, stop
* *Task Listeners* can be attached to events within the user task lifecycle

### Java Delegate

implement JavaDelegate and provied the required logic in the execute method
Process instances information cam be accessed and manipulated through the *DelegateExecution* API
Each time a delegate is executed a separate instance of the class will be created, i.e. keep the stateless

There's another interface that can be implemented which will be called at runtime, *ActivityBehaviour*
Avoid it unless you know what you're doing!

You can inject values into fields of the delegated classes from the Workflow, BPMN xml. Fixed string values or expressions
If available the value is injected through a public setter method following the java bean convention.
  If the public setter method is not available the value of the private member will be set. (Not recommended)

### Execution Listener

Execute external java code, a script or evaluate an expression when a certain events occur.

* start/end process instance/activity/gateway/intermediate events
* taking a transition
* ending a start event or starting an end event

Implement the ExecutionListener interface.

### Task Listener

Not really interested in user tasks so I'll skip for now.

## Expression Language

Camunda BPM supports Unified Expression Language (EL), specified as part of JSP 2.1(JSR-245).
Uses the [JUEL](http://juel.sourceforge.net/) implementation.
Can be used in BPM to evaluate small script like expressions.
Methods can be called on beans in the Spring Container and variables passed into them from the workflow
Input/Output parameters can be mapped using this language

## Scripting

Camunda BPM supports scripting with JSR-223 compatible script engine implementations.

* Groovy
* JavaScript
* JRuby
* Jython

To use one you need to add the dependency to the classpath.
The first time the engine is needed it's brought up, then cached for subsequent requests.

## Templating

https://docs.camunda.org/manual/7.5/user-guide/process-engine/templating/
