# job dsl plugin

https://jenkinsci.github.io/job-dsl-plugin/
https://www.jvt.me/posts/2021/02/23/getting-started-jobdsl-standardised/
https://github.com/jenkinsci/job-dsl-plugin

A plugin that allows you to manage all your jenkins job configuration as code.
No manual creation and maintenance of jobs.
Changes to jobs go through code review.
Can standardize jobs/pipelines across teams

Install the plugin!
allow script execution, Jenkins has security on it: https://github.com/jenkinsci/job-dsl-plugin/wiki/Script-Security

## API

Complete API available at: https://your.jenkins.installation/plugin/job-dsl/api-viewer/index.html

DslFactory
  for creating jobs of different types
  #job create a freestyle job
    it has loads of method calls to create a complete job.

Write Builder classes to create the configuration of the jenkins job.
  Inheritance can be used to reuse configuration
  The build method in each class calls super then adds it's own configuration
  The root calls the abstract method create to create jobs of different types, although we only create pipeline jobs

The build method is called on all instances then to create the actual jobs

## Groovy

with method
This is a way of assigning the contents of a closure into an object.
