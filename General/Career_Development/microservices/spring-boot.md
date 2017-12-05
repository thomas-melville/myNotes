# Spring Boot

[Tutorial](https://dzone.com/storage/assets/4855140-microservices-for-java-developers.pdf)

An opinionated Java framework for building microservices, based on the Spring Dependency Injection framework.

Favours automatic, conventional configuration by default.
Curates sets of popular starter dependencies for easier consumption.
Simplifies application packaging.
Bakes in application insight.

## Starter Dependencies

Curated sets of libraries to enable some functionality:
	JPA persistence
	NoSQL DBs
	Redis caching
	Tomcat servlet engine
	JTA Transactions
	...

Convenient dependency descriptors you can include in your application to get some functionality.
[List of starters](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/#using-boot-starter)

Adding a submodule to your project brings in the required modules for that functionality, with the versions that are known to work together
Versions of these dependencies are managed by spring for you in a BOM, spring-boot-dependencies
	You can still override Springs provided version by specifying properties in your pom
		if you are using the spring-boot-starter-parent 
		if using spring-boot-dependencies you can't do this, you need to specify the dependency with the updated version
You can depend on sensible default by using spring-boot-starter-parent
	java 6 as default compiler
	utf-8 source encoding
	dependencyManagement section (no need for versions in your pom)
	sensible resource filtering
	sensible plugin configuration
	sensible resource filtering for props/yml files for spring-boot application


## Application Packaging

Self-contained JAR packaging of the application.
java -jar will bring it up
all dependencies and application code are in the jar with a flat class loader (no nesting of class loaders, just a single class loader)
This makes it easier to understand start-up, dependency ordering, and log stmts.
Helps reduce the number of moving parts required to take a service into production, safely.

## Bake in application insight

Spring boot ships with a module called actuator, enables metrics and stats about the application.
HTTP, JMX or SSH

## Auto-configuration

Let's developers codify things and have it stood up as part of application.
If a module is on the classpath
is a bean of a certain type defined
is an environment property specified.
is a class annotated with a certain annotation
...

## Creating a Spring Boot project

spring init --build maven --groupId com.redhat.examples --version 1.0 --java-version 1.8 --dependencies web --name hola-springboot  hola-springboot

--build
	can be maven or gradle
--groupId
	of the project to create
--version
	of the project to create
--java-version
	to use
--dependendices
	what starter dependencies to include
	web: Spring MVC and embed an internal servlet engine
--name

### what this does:

1. pom parent is spring-boot-starter-parent (you can also put spring-boot-dependencies as scope import in depMgmt instead)
2. java 8
3. utf-8 encoding
4. spring-boot maven plugin! This is important to get the application for run!!!
5. src/main folders and packages

## Creating a REST Endpoint

### In the example
@RestController
	Tell spring this a HTTP controller capable of exposing HTTP endpoints
@RequestMapping
	Map specific parts of the URI to classes, methods and parameters

### In allure-service
Class annotated with @Configuration which extends ResourceConfig & @Paths annotation in generated interface

Why? Is it because we are using raml? No, it's because we are using jax-rs

## Externalize Configuration

Spring makes it easy to use external property sources.
Entire classes of properties can be bound to objects.
This makes it easy to configure the service based on the environment it's running in, local, dev, test, stage, production, location, ...

### In the example
@ConfigurationProperties(prefex="myApp") on the class will bind all properties that start with the prefix to objects in the class.
	So, if there is a variable in the class named myVar the property myApp.myVar will be searched for.
	getters and setters are also required!
The default location for properties is src/main/resources/application.properties

### In allure-service
@Value on the field is used
and we use a yaml file instead of a properties file

## Application Metrics and info

Starter package, actuator, makes it a breeze
just add the dependency and the metrics end points are exposed
you also have to expose the management http port in the application configuration file.

## Calling another service

Use @ConfigurationProperties to externalize the host & port of the other service.

Use RestTemplate to do the invocation of the remote service.
	a convenient wrapper abstraction which handles all connection, (un)marshalling results 
	It's generic so you can give it the class of the object you want to return

[Complete Reference Guide](https://docs.spring.io/spring-boot/docs/current/reference/htmlsingle/)

Entry point for a spring-boot application is a class which you write which has a static main method
It calls SpringApplication.run(YourClass.class) which will bootstrap your application, start Spring and then anything else depending on what you have on the classpath

spring-boot compiles the application into an uber jar, by nesting jars!
To enable this you have to add the spring-boot-maven-plugin to your project

## Annotations

@RestController
	Tells Spring that this class is a web @Controller so Spring will consider it when handling incoming requests.
	This is a stereotype annotation,
		it tells Spring that this class of a certain thing and Spring will register as such
	Also, render the resulting string back to the caller.
	In allure-service this is done progammatically in **JerseyConfig**
@RequestMapping
	Provide routing information for Http Requests
@EnableAutoConfiguration
	Tells spring-boot to guess how you will want to configure spring based on the jars on the class path
@Profile
	takes a string which is the profile to have this configuration active or not active (use !) for
@AutoWired
	Marks a constructor, field, setter/config method to be autowired by Springs DI facilities
	Basically @Inject + @ImplementedBy
	But has the added advantage that it can inject an implementation based on the active profile
@Configuration
	This class contains variables which are mapped to properties.
@Value
	the value of the key in the properties which maps to this variable.
@ComponentScan
	Scan this class and recursively all it's fields to inject instance variables.
@SpringBootApplication
	convenience annotation which encompasses @Configuration, @EnableAutoConfiguration & @ComponentScan

## Structuring your code

No specific structure, but best practices

* don't use the default package
* Locate the main class in a root package above all other classes
* delegate execution of app in main class to SpringApplication#run
* use java based configuration, i.e. annotate with @Configuration. Usually the main class (but not in allure-service, because @SpringBootApplication contains @Configuration)
* No need to put all the configuration into one class
* Add @EnableAutoConfiguration/@SpringBootApplication to allow to spring to automatically configure everything based on whats on the classpath.
	* If you put in your own configuration the auto config will back off.
* Replace auto-config gradually
* You can disable specific auto-config with the exclude parameter in @EnableAutoConfiguration
* Use @ComponentScan and @Autowired for all dependency injection. (@SpringBootApplication is annotated with @ComponentScan)
* @SpringBootApplication == @Configuration, @EnableAutoConfiguration, @ComponentScan

## Dev tools

maven dependency to add developer tools to a running application.
Only active in dev mode by default
flag the dependency as optional to prevent transitive issues
disables caching by default
auto-restart app when it detects changes to files on the class path
	This works for in IDE and mvn spring-boot plugin
	resources are excluded as they don't require a restart
	lots of properties to configure when a restart happens
LiveReload for reloading browsers, plugins for the browsers available

You can enable dev tools remotely by removing the exclusion in the spring-boot plugin
	never do it in production, it is a security risk

## Profiles

Segregate parts of application configuration, and make them only available in certain environments
@Profile annotation is used with a string to specify what profile the configuration should be active for.
use spring.profiles.active to activate profiles, this overrides any other profiles which have already been set.
	For example a profile set in the application.propeties file will be overridden by a profile set on the command line
use spring.profiles.include to add more profile(s) on top of whats already there. 

profile specific properties/yml files can be created:
	application-{profile}.properties/yml

spring-boot automatically injects the properties from application.properties/yml and then application-{profile}.properties/yml based on the active profile(s)

## Spring Cloud Config

{application} - maps to spring.application.name on client side, this means that the config-server can be used for multiple services
{profile} - maps to spring.profiles.active on client side , this means the config-server can hold configs for multiple environments that a service is deployed in
{label} - this is on the server side, versions of configs

## packaging your code for production

jars are self contained so they are suitable for cloud based deployment
actuator dependency is recommended for health checks, ...

## Spring Boot Features

You can custumize the banner which is displayed on the console when you start the application
create a banner.txt file, can be a picture or a gif