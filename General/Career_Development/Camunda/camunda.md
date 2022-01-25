# Camunda

## BPMN

A business process is a set of linked activities which accomplish a specific org goal.

Business Process Management is a systematic approach to making an orgs workflows more effective.
It has specific stages.

BPMN (Business Process Modeling Notation)

Map processes and procedures to a flow/map
so you can visualize it
Standard

We are shifting to Camunda as the orchestration workflow engine

ONAP use BPMN Workflows too

Apparently our customers want it too

https://camunda.com/

platform for workflow and decision automation

Execution engines for BPMN workflows and DMN decisions paired with essential applications for process automation projects.
java based

BPMN:   Business Process Model Notation
    structure processes
    potential high level of automation
    global standard
CMMN:   Case Mgmt Model Notation
DMN:    Decision Model and Notation
    decision automation based on decision rules.

Model -> Execute -> Improve

Camunda BPM is in the execute space
Camunda Modeler is in the Model & Improve space

Platform:

* BPMN Workflow Engine
  * for microservice orchestration and human task management
  * Remote REST service or embed in java app

* DMN Decision Engine
  * executes business driven decision tables
  * pre-integration with workflow engine

* Modeler
  * app for editing BPMN process diagrams and DMN decision tables

* Tasklist
  * web app which allows end users to work on tasks assigned to them

* Cockpit
  * monitor workflows and decisions in production to discover, analyse and solve technical problems.

* Admin
  * manage apps, grant permissions, ...

* Optimize
  * create beautiful reports and arrange them in a dashboard
  * configure alerts for performance goals missed
  * advanced analytics to identify process bottlenecks

## BPMN Workflow engine

automate BPMN 2.0 process diagrams for micro service orchestration, human task flows or both.
written in java
REST and Java APIs for interacting with the engine
the engine uses a Relational Database to store history of workflow execution, and I presume other stuff too
It's integrated with Spring boot, add the camunda dependency and away you go

* start process instances
* complete tasks
* ...

The engine is stateless which makes horizontal scaling easier.

External Task Pattern allows you to develop and operate your microservice completely decoupled from the engine
and let them pull the work via REST whenever it suits them
The BPMN contains a task of type External, which points to a topic on a message bus.
The External task service registers with the engine

Workflows are described in BPMN which is basically xml, you can visualize it using the Modeler
Create Classes which implement JavaDelegate and put their fqdn in the BPMN workflow to link them when the workflow is executed

### External Tasks Pattern

engine supports two ways of executing service tasks

* internal service task
* * sync invocation of code deployed along with process app
* external tasks
* * provide a unit of work in a list that can be polled by workers

Second way is the external task pattern

### using camunda in a spring boot app

1. add the camunda engine dependency:
    * camunda-engine-spring
2. add the camunda starter dependencies:
    * camunda-bpm-spring-boot-starter-[webapp/rest]
3. add an annotation to the application class
    * We have @EnableProcessApplication
3. create your bpmn diagrams using Camunda modeler
    * this generates xml
    * put this in the src/main/resources folder of your project
4. implement JavaDelegate interface to define the tasks in the diagrams
    * we have these separated out into modules specific to the type of task
