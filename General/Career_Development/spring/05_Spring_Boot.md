# Creating your first spring boot app

Spring Tool Suite IDE, based on Eclipse and geared towards Spring apps

* create basic maven project
* add spring-boot as parent
* * spring-boot-starter-parent
* * alternatively you can add spring boot to the dependencyManagement section
* * What are the differences again???
* add specific spring-boot-starter
* * spring-boot-starter-web
* * * add Spring Web Mvc and auto configure it for us
* Annotate the class with the main method with @SpringBootApplication
* * Scans the classes for spring components
* * auto wires everything based on dependencies in the pom
* Change the main method body to 'SpringApplication.run(Classname.class, args);'
* * fires up spring boot when the main method is called
* recommended package structure:
* * controller
* * model
* * repository
* * service
* Create first Controller
* * class in controller package
* * annotate it with @RestController
* * create a method and annotate it with RequestMapping, or the more specific annotations which are now available

And now you have a basic spring boot application up and running.

## Spring Boot Dependency Management

Provides a bill of materials
The spring developers have curated list of dependencies and versions for common functionality which work well together.
This gets brought in when you set the spring-boot-starter-parent as the parent of your project.
This parent also configures some generic stuff

* resource filtering
* java min version
* plugins
* ...

Spring Initializers

* start.spring.io
* spring boot cli
* * the website uses the cli
* copy one of the spring boot examples from gitub
* jhipster?
* yeoman?

## how does Spring Boot work

* Main method entry point
* * public static void main
* Spring boot initializes the spring context and honours any auto config initializers
* * @SpringBootApplication
* * * Under the hood this specifies the following annotations:
* * * * @Configuration
* * * * * Configure the Spring Context for this app
* * * * @EnableAutoConfiguration
* * * * * Configure any other frameworks found on the class path
* * * * @ComponentScan
* * * * * scan this and all child packages for instances which are annotated with Spring annotations to inject
* * SpringApplication.run();
* * * Starts Spring
* Embedded server container is started and auto configured

## Why move to containerless deployment

The question is more why move to an embedded container deployment

* less provisioning per environment
* complete isolation

## Creating Web Apps

spring boot scans 3 folders for static resources

* /static
* /src/main/resources/public
* /src/main/resources

when developing the running app doesn't have to be restarted to serve updated static content

Spring boot handles the integration of SpringMvc and sets up the jackson json library so that the data is marshalled into json

* Sets up ViewResolvers, ContentNegotiating!
* sets up static resource serving
* Sets up HttpMessageConverter
* Sets up customizable hooks, which allow you to change the default behavior

use application.properties file on classpath root to change properties and configuration.
can be YAML either, you need the YAML snake dependency on your classpath.

Spring places all file and folders in src/main/resources at the classpath root at runtime

You can make the properties file environment specific by adhering to a certain filename format:

application-{profile}.properties

Spring loads application.properties first and then overrides with any specified profiles.

To activate a profile:

-Dspring.profiles.active={profile}

## logging

add a line in the properties file:

logging.level.{package}={level}

## Databases

Will use Spring Data JPA with an in memory H2 database

* Add the h2 dependency to integrate the database, we can then configure it using application.properties
* Add the spring-boot-starter-data-jpa dependency to include and auto configure data access.
* * When it sees the h2 dependency on the classpath it will do some more configuration based on the database type.
* Add configuration to the properties file to allow spring access to the database.
* * url, username, password, driver
* Annotate the DTO's accordingly
* * @Entity
* * @Id
* * @GeneratedValue(strategy = GenerationType.AUTO)
* Create a repository
* * Create an interface which extends JpaRepository<{DTO_Type}, Long>
* Inject repository to Controller and call methods on the interface

### datasource pooling

if spring detects a pooling dependency on the classpath it will integrate it and pooling will be provided.
default pooling library is tomcat-jdbc, which is brought in by spring-boot-starter-data-jpa.
it can then be configured by specifying properties in application.properties: max-active, max-idle, max-wait, min-idle, ...

### h2 database

* in memory database useful for testing.
* Can also save the data to a file to persist it across restarts.
* you can configure it so that you have a console into the database from the browser:
* * spring.h2.console.enabled=true
* * spring.h2.console.path=/h2

### Version controlling your database

FlywayDB
Scripts are database dependent. This means different databases require their own scripts if the syntax is different.

* add the flyway dependency
* spring uses the primary datasource you have setup by default
* create migration scripts
* * they must be in src/main/resources
* * first script must be v2, as flyway baselines at version 1
* * v2__{outcome_of_script}.sql
* when the app starts the migration scripts are automatically run
* add a property to application.properties:
* * flyway.baseline-on-migrate=true
* turn off a hibernate feature which would auto pick up the flyway ddl files and try and create the tables
* * spring.jpa.hibernate.ddl-auto=false

## configuration

You can define Spring beans for configuration instead of the corresponding properties in application.properties

```java

@Configuration
class

@Bean
@ConfigurationProperties(prefix="spring.datasource") // picks up properties with this prefix from application.properties
public DataSource datasource(){
    return DataSourceBuilder.create.build();
}

```

@Primary annotation, if there is ever ambiguity use this bean by default

@FlywayDataSource, specify a particular datasource as the flyway datasource, separate from the primary datasource

## Testing

spring-boot-starter-test, scope test

Brings in:

* junit
* hamcrest
* mockito
* spring test tools

### unit testing

create instances of beans and assert on calls to methods

mocking & mockito:

* @InjectMocks, inject mock objects for all fields of class.
* @Mock, instantiate a specific mock which will be injected into the class.
* MockitoAnnotations.injectMocks(this), does the setup of the mocking based on the annotations in the test class

### integration testing

test all classes working together

* app / test startup time can be slow
* * each test boots up the app so it can get slow, quickly
* database consistency has to be considered

Change unit test to integration test:

* @RunWith(SpringJUnit4ClassRunner.class)
* @SpringApplicationConfiguration, and provide main class to annotation
* autowire class under test instead of instantiating it
* Use @MockBean to inject the mock instance into the application context

### web integration testing

call REST API

* @WebIntegrationTest
* & the annotations from the integration test above
* Use TestRestTemplate to make the REST call
