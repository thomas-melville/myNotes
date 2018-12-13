# Spring: The Big Picture

## what is spring

It's a lot of things:

* Spring framework
* Spring boot
* Spring data
* Spring cloud
* Spring batch
* ...

When people say "spring" they are often referring to the entire family of projects / eco-system.

It all started with "The Spring Framework"
which was built to combat the complexity of building apps with the J2EE framework.
aimed at reducing boiler plate code.

It's the foundation to all other projects in the eco-system.

Then spring-boot came along on top of them all which sped up app development even more.
It took an opinionated view of building an app and did a lot of the configuration for you.

spring cloud is built on top of spring boot.
great for microservices because of service discovery, ...

### why spring

* simplify creating java applications
* flexible and modular
* backwards compatible
* large and active community
* continually innovating and evolving
* open source but funded by Pivotal

## spring boot

it's a key project, I have a whole other file on this.

Makes learning spring both quick and easy, then once you know more you can circle back and find out how it works.
Which is exactly what I'm doing now, :-)

can be used to build web and non-web apps

### auto-configuration

best guess configuration for your app based on the dependencies you've added.
contextually aware and smart.
really easy to disable also, made as non-invasive as possible

### stand alone

because of the container that the java web app requires spring boot packages it as part of the app through a dependency, see auto-configuration

so then you can do java -jar on the command line to start your app

### opinionated

so many choices
we often just follow the standards we see online
spring boot does this automatically
"soft" opinion because you can override them

spring initializer is great for creating a spring boot project

## The Spring Framework

universal, reusable environment
provides particular functionality: web development, data access, ...
part of a larger software platform, the java platform
facilitate development of software apps, products and solutions, remove complexity of jee

2003:   enterprise java development complex
        Creator Rod Johnson
        simplify dev process
        grew and evolved

first it was apps using the spring framework
then it was apps using the spring projects and the spring framework
simplifying their development
then along came spring boot to simplify development even more

modular architecture

6 key areas:

* Core
* Web
* AOP
* Data access
* Integration
* Testing

### core

the most important piece of the spring framework

Internationalization support
data binding support
validation support
type conversion support
...

at the centre is Dependency Injection!
it's about the way objects get their dependencies

fulfilling dependencies

* object fulfills it's own dependencies
* * seems easy
* * the object is then responsible for the lifecycle of those objects
* * tight coupling
* * testing becomes difficult
* object declares what it depends on and something else fulfills the dependency
* * more flexible
* * loosely coupled

Spring core is considered a dependency injection container

* creates and maintains objects and their dependencies
* less for the developer to manage

often referred to as the glue of an app

### web

fwk for handling web requests

* Spring Web MVC
* Spring Web Webflux

#### Spring Web MVC

Based on the Java Servlet API, a low level API.

A request comes into the web server and is passed to the Java Servlet API, which then passes it to the Spring Web MVC which brings the request up to a higher level of abstraction and passes it onto the business logic of the application

#### Spring Webflux

reactive programming!

asynchronous execution, code asks to be notified when the action is complete.
Implementation of asynchronous API calls which Sam talked about in the microservices course

### aop

Aspect Oriented Programming: "A programming paradigm that aims to increase modularity by allowing the separation of cross cutting concerns"

examples:

* security
* logging
* monitoring
* resiliency

without it: scattered and duplicated code across the code

Spring AOP

an implementation of AOP:

* used to implement features of Spring: security, data
* valuable tool for developers to handle cross-cutting concerns

### data

make it easy to store and retrieve data in a relational database
remove a lot of boiler plate code
easy to use database transactions

Use a template and annotations

Exception Translation of database vendor specific codes in a standard set

makes testing easier, by allowing you to easily switch out the database for an embedded db

### integration

all about making different systems and apps work together
how does one app talk to another?

RMI
Messaging System
**Web Services**

Spring Framework makes it easy to expose and call web services

expose: annotations I'm used to
call: RestTemplate

### testing

Spring framework focuses on two areas of testing:

* unit
* integration

explicit support for unit testing is minimal
the support comes as a side affect of using dependency injection.
inject mocks, stubs, fakes, ...
pre written mocks can be used

in integration testing the side affect of dependency injection is again an advantage
There is more support built in:

* testing data access
* testing web services
* clean up after tests which isolate tests

## Other Spring projects

http://spring.io/projects

### spring data

provide a familiar & consistent, Spring-based programming model for data access.
extends the capabilities provided by the Spring Framework data.
adds support for NoSQL data bases.
It's a series of sub-projects, one per data store.
Combines with java JPA to translate objects to data.

### spring cloud

built on top of spring boot
good for distributed systems, microservices.
another sub project
There are a number of concerns which all microservices must solved for example:

* service discovery
* * with spring cloud add an annotation and it can discover other microservices and be discovered
* many more

### spring security

secure spring based apps
authentication and authorization
protects against attacks

add an annotation and implement a method to setup what happens on particular urls

## Is Spring a good fit

+ives

* rock solid and well engineered
* stood the test of time
* huge community
* well liked by many developers
* large talent pool
* wealth of knowledge: books, ...
* very actively developed
* a lot of built in IDE support
* Scalable

-ives

* too much magic
* steep learning curve
* increases the size of the final deliverable
* hard to debug
* adds memory overhead
* complexity has grown over time
* can be too configurable
* it's big
* some of the community projects can be hit or miss
