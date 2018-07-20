# tosca

https://searchcloudcomputing.techtarget.com/definition/TOSCA-Topology-and-Orchestration-Specification-for-Cloud-Applications

Topology and Orchestration Specification for Cloud Applications

An open source language used to describe the relationships and dependencies between services and applications that reside on a cloud computing platform.

Highly extensible language which allows us to add domain/vendor specific mechanisms

Supported by the Organization of Advancement of Structured Information Standards (OASIS)

## metamodel

It describes cloud services using templates and plans

A template defines the structure of a cloud service. It is a topology template.
A graph of node templates modelling the components a workload is made up of and as relationship templates modelling the relations between those components.
lifecycle operations can be defined in these nodes templates

Plans define the processes that start, stop & manage that cloud service over its lifetime

## tosca simple profile in yaml

http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.1/csprd01/TOSCA-Simple-Profile-YAML-v1.1-csprd01.pdf

A rendering of TOSCA which aims to provide a more accessible syntax as well as a more concise and incremental expressiveness of the TOSCA DSL
It is a subset of the TOSCA v1.0 XML spec