# DropWizard

around long before Spring

minimalist, production-ready, easy-to-use web framework

It's more prescriptive than Spring
Some components are part of the framework and can't be changed.

sweet-spot use case is writing REST-based web applications/microservices without too many fancy frills.

No dependency injection container like Spring.

Prefer to keep it simple and not hide all the magic like Spring does

Spring hides complexity in wiring things together but makes it more difficult when debugging issues

bundles entire project into one uber jar, like Spring Boot

flat class loader

Dropwizard uses the shade plugin to package everything up into an uber jar

Metrics are first class citizens, annotations are put on methods to make metrics available.

## Stack

Jetty
Jersey
Jackson
Hibernate
Guava
Metrics
Logback + SLF4j
JDBI

## Abstractions

Application
	contains the public void main method
Environment
	servlets, resources, filters, health checks, tasks, commands
Configuration
	how we inject env & app specific configuration
Commands
	Tells Dropwizard what action to take when starting a server
Resources
	REST/JAX-RS resources
Tasks
	admin tasks to be performend when managing the service


## Running

2 ports exposed

8080: application
8081: adminsitration

## Getting Started

### package structure

api
	pojos that define all objects used in your REST resources 
	dtos
cli
	dropwizrd commands go here
client
	client helpers classes
db
	any db related code
health
	microservice specific health checks go here
resources
	REST resource classes go here


Application.java is the entry point where we can put all out bootstrapping