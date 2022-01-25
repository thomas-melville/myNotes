# jenkinsfiles

https://www.jenkins.io/doc/book/pipeline/

definition of CI/D/D in Jenkins
can be modelled as code, so it can be version controlled

## Declarative or scripted

Declarative is newer with richer syntactical featuers

declarative - pipeline block
scripted - node block

## Pipeline shared libraries

It it possible to write jenkinsfile code in a git repo and make it available to jenkinsfiles in a Jenkins instance.

1. add the repo as a shard library in jenkins
2. "import" the library into the jenkins file

## Concepts

### Pipeline

user-defined model of a CD pipeline

pipeline keyword and code block are a key part of the declarative pipeline syntax.
defines all the work done through out the entire pipeline

written in a limited form of groovy syntax

### Node

A machine which is part of the Jenkins environment and is capable of executing a pipeline

a key part of the scripted syntax.

### Stage

A block which defines a conceptually distinct subset of tasks performend through the entire pipeline.
Common to declarative and scripted

### Step

A single task within a stage which executes a command
Common to declarative and scripted

## Syntax

agent, the jenkins machine to execute the pipeline on
  used in declarative pipelines
  It can take a number of parameters
    any
      execute on any agent
    none
      when applied at the top level, each stage will need to acquire it's own agent
    label
      use an agent associated with the label
    node
      similar to label, but allows for additional options
    docker
      execute within a given docker environment
  node can be used within an agent
  there is a slight difference when using a top level agent and a stage level agent
    when the options are executed!
    top level, after
    within a stage before


node, the jenkins machine to execute the pipeline on
  used in scripted pipelines
  can also take a closure with code to be executed on that node.

https://www.jenkins.io/doc/pipeline/steps/workflow-scm-step/#checkout-check-out-from-version-control
checkout(
  class - the name of the plugin to use to check out the code
  branches - the branch(es) to checkout, best to checkout one
  doGenerateSubmoduleConfigurations -
  extensions
  )

dir

if

load


lots of plugins provide steps for a jenkinsfile:
  sh, junit, ...

https://www.jenkins.io/doc/book/pipeline/syntax/

### post

what to do after all the stages of the pipeline are finished

always
changed - if the result is different from the previous run
fixed - if the pipeline is a success and the previous was a failure
regression - if the pipeline is failure, unstable or aborted and the previous run was a success
aborted
failure - pipeline is red in the ui
success
unstable - pipeline is yellow in the ui, caused by test failures, static check violations, ...
unsucessful - pipeline result is not a success
cleanup - always run after all other post steps

### stages

contains a sequence of one or more stage directives
no parameters, just a list of steps

### steps

one or more steps to be executed
no parameters

### environment

sequence of key value pairs
available to all steps/stages

credentials function is available, which when given the string looks for the credentials in the jenkins environment

### options

configure pipeline specific options for the job in the pipeline itself.
examples:
timestamp in console
when to discard builds
...

This can also be achieved in the jobs-dsl code.
