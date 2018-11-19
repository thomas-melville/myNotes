# Spring Fundamentals

## What is Spring

Started out as an inversion of control container, to replace the complex config of earlier JEE config.
Then was ported to JSE
POJO based
helping us write cleaner better code
unobtrusive
AOP/Proxies
Built around Best Practices
Spring uses the Template Method pattern a lot
as well as Singleton, Factory, Proxy, ...

### What is it trying to solve

increases testability, maintainability & scalability
reduce complexity
puts the focus on the business
remove configuration from your code!
Eases moving to a different environment and testing
focus on the business logic
annotation / xml based configuration

### How it works

Everything is a POJO
Could be thought of as a big old HashMap
Registry
Instantiates all required beans and wires them together
so you don't have to instantiate any bean yourself, never use "new"

Model
Repository
Service
Application

## XML Configuration

first method available.
sometimes still simpler than using annotations / java apis.

applicationContext.xml
root of the application, name doesn't matter but location does! src/main/resources

```java

// You then add this file name as the application context to your main method
ApplicationContext appContext = new ClassPathXmlApplicationContext("xml_filename.xml");

// To get a bean you then call methods on the appContext
Service service = appContext.getBean("myService", Service.class);

```

Spring Context is sort of a HashMap
namespaces aid in configuration and validation
add it to the top of your file
namespaces & schema locations are useful for hinting what can go into the xml

### XML declarations

#### bean

replaces keyword new
define class but use the interface, best practice to always program to the interface

##### attributes

* name
* class
* autowire

##### setter injection

create a setter method in the class for the field.
create a property element as a child of the bean element in the xml

property attributes:

* name
* ref

##### constructor injection

define a constructor for the class, **and a no-arg constructor???**
create a constructor-arg element as a child of the bean element in the xml

attributes:

* index
* ref

```xml

<bean name="customerRepository" class="com.ericsson.....CustomerRepository"/>

<!-- setter injection -->
<bean name="customerService" class="com.ericsson...CustomerService">
    <property name="customerRepository" ref="customerRepository">
</bean>

<!-- constructor injection -->
<bean name="accountService" class="com.ericsson...accountService">
    <property index="0" ref="customerRepository">
</bean>

```

Setter and Constructor injection can be used in the same bean

## Annotations

3 main stereotype annotations for finding beans in Spring:

* Component
* * Any POJO
* Service
* * Business Logic Layer
* Repository
* * Data Layer

Semantically they are all the same, it's how they are used that makes them unique

When using annotation configuration and programming to interfaces you can put a string in the annotation to give it the bean name, usually the interface name

## Annotation Configuration using XML

Second form of configuration where the annotation scanning is configured in xml.
And all subsequent configuration is done through Annotations on the classes.

```xml
<!-- add the context specific namespace also -->

<!-- Hello, I'm configured using annotations -->
<context:application-config/>

<!-- Here is where you should start looking for annotations -->
<context:component-scan base-package="com.ericsson"/>

```

## Annotation Configuration using Java

No XML!!! :-)

Create a class to configure the app on the same level as the Application class.
Then add this class to the Application class as the ApplicationContext.

```java

// AppConfig is the class we created and annotated with Configuration
ApplicationContext appContext = new AnnotationConfigApplicationContext(AppConfig.class);

```

A class annotated with @Configuration replaces the appContext.xml file.
It contains @Bean annotations which replace the bean elements from the xml
**By default beans are Singletons**

```java

@Configuration
public class AppConfig{

    @Bean(name="customerRepository")
    // method name doesn't matter
    public CustomerRepository getCustomerRepository(){
        return new ...Impl();
    }
    }
}

```

Setter and constructor injection happens inside the method annotated with Bean, the setter method / constructor is just called.

## Autowire

Spring got a bad name for lots of xml / java configuration to wire beans together.
To they implemented an attribute on bean elements and annotation which would trigger autowiring based on a policy

Four options:

* byType
* * if exactly one bean of this type exists it will be autowired to the property/field
* * an exception is thrown if there are two.
* byName
* * if a bean exists with the same name then autowire it to the field
* * this solves the issues of multiple beans of the same type.
* constructor
* * behaves like byType
* no
* * no autowiring

It doesn't look like there are any major advantages to either way to autowire

on the field    -> no need for setter code
on the setter   -> can add logging when the wiring happens

With autowiring the above xml doesn't need any of the child elements of the bean element:

```xml

<bean name="customerRepository" class="com.ericsson.....CustomerRepository"/>

<bean name="customerService" class="com.ericsson...CustomerService" autowire="constructor"/>

<bean name="accountService" class="com.ericsson...accountService" autowire="constructor"/>

```

Autowiring can also be done using the @Autowired annotation, on a constructor, setter or field.
To enable it add the @Component-Scan annotation to the class that is annotated with @Configuration with the base package to start with

## Bean Scopes

5 scopes:

* Singleton
* * default
* * only ever one instance per Spring Container
* Prototype
* * guarantees a unique instance per request
* * opposite of singleton
* Request
* * only available in web projects
* * Bean per http request
* Session
* * only available in web projects
* * Bean per http session
* Global
* * only available in web projects
* * Bean per application

```java

@Scope(ConfigurableBeanFactory.SCOPE_SINGLETON)
public class OrderServiceImpl implements OrderService {
    //...
}

```

```xml

<bean ... scope="singleton">

```

## Properties

12 factor app, externalize the configuration

### XML config

```xml

<!-- load the properties from the location-->
<context:property-placeholder location="..."/>

<!-- inject the properties into a bean-->
<bean name="..." class="...">
    <property name="..." value="${my.property}"/>
    </bean>
```

Setter injection can be replaced with Java annotation driven:

```java

@Value("${my.property})
private String ...;

```
### Java Config

```java

// Tell Spring where to find the file
// Add this to the config class
@PropertySource("location")

// Configure a bean to load in the properties
@Bean
public static PropertySourcesPlaceholderConfigurer get...(){
    return new PropertySourcesPlaceholderConfigurer();
}

@Value("${my.property}")
private String mine;