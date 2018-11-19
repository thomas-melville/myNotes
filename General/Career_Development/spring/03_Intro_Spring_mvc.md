# Introduction to Spring MVC

## Framework Overview

History of Web Frameworks:

* Struts
* tapestry
* Apache Wicket
* GWT
* JSF from Oracle
* Something from Jboss, couldn't hear the name
* Stripes
* Spring MVC

## What is Spring MVC

* A web framework built around the principles of Spring
* Isn't built for one particular type of web app:
* * REST
* * JSP
* * Freemarker
* * Velocity
* * ..
* POJO & Interface Driven
* Based on a Dispatcher Servlet / Front Controller Pattern
* Light weight and unobtrusive
* Built from the Shortcomings of Struts 1
* Support for:
* * Themes
* * Localization
* * Restful Services
* * Annotation Based Configuration
* * Integration with other Spring services

## Architecture

Spring MVC is built on top of the Java Servlet API.
We then build our apps on top of Spring MVC.

Request Response Lifecycle:

* Request comes into the Front Controller, the DispatcherServlet in Spring.
* Delegates the request to one of our Controllers
* Our Controller handles the request using the other classes in the backend
* Backend returns a model to the Controller
* Our Controller delegates the rendering to the Front Controller
* Front Controller passes the model to a view template
* which renders the response and returns it to the front controller
* which returns the response to the client

### MVC design Pattern

Traditional (In rich client apps)

* A request comes into the **View** from some user event
* This triggers an event in the **Controller**
* Which will trigger a change in the **Model**
* Which the **Model** will report to the **View**
* the **View** will query the **Model**

MVC Web

View and Model go through the Controller to update each other.

### N Tier Apps

Enterprise apps tend to be built in a tiered architecture.

* Presentation Layer
* Business Logic
* Data Layer

* Separation of Concerns
* Reusable Layers
* Maintenance / Refactoring

#### Tiers in a spring mvc application

* View
* * represented by **JSP pages**
* Controller
* * represented by classes annotated with **Controller**
* Service
* * represented by classes annotated with **Service**
* Data Model / Database (Model Object)
* * represented by classes annotated with **Repository**

## Using Spring MVC

### Maven dependencies

* spring-webmvc
* servlet-api
* * provided
* jstl
* * provided

### Configuration

5 parts:

* Configure web.xml in src/main/webapp/WEB-INF
* * create a servlet element which has a servlet-name and servlet-class child elements
* * * servlet-class is the Spring MVC DispatcherServlet class
* * best practice to specify context config location
* * * init-param element with a param-name and param-value children
* * * contextConfigLocation & /WEB-INF/config/servlet-config.xml
* * setup the servlet-mapping
* * * child element servlet-name must match servlet-name from above
* * * url-pattern child element also
* Configure servlet-config.xml in src/main/webapp/WEB-INF/config
* * add extra namespaces for ease of configuration
* * add element mvc:annotation-driven
* * add element context:component-scan with attribute base-package="..."
* Add Controller
* * POJO with some annotations
* * * Controller & RequestMapping
* * Add some attributes to the model, this will be used by the view
* Add View in src/main/webapp/WEB-INF/jsp
* * This is html with some attribute injection from the model
* * using the ${} format
* Configure servlet-config to find jsp pages
* * bean element with class attribute whose value is InternalResourceViewResolver
* * The configuration of the location can be done 2 ways
* * * child properties of bean element
* * * * prefix : /WEB-INF/jsp
* * * * suffix : .jsp
* * * attribute of the bean element which is in the p (property) name space
* * * * prefix & suffix

### Standards

* views go in the webapp/WEB-INF folder
* * security
* * control user experience, they can't go directly to pages
* * breadcrumbing

InternalResourceViewResolver class finds pages in this folder

## View

Many different templating tools can be used in Spring MVC

* JSP
* Velocity
* Freemarker
* JSF
* ...

### Resolving a View

* The Controller method returns a string
* which is the name of the view file without the file extension
* We could also return a view object which wraps the string
* The controller builds the model which is used by the view

There are many types of ViewResolvers and you can create your own

* Templating tool specific resolvers
* Content negotiation resolvers
* ...

### Chaining

* in the return string from the controller
* prepend with "forward:" and append with ".html"
* there are other key words:
* * redirect, starts a new request
* * * whereas forward, forwards on the request

## Displaying Static Data

* configuration in servlet-config
* * <mvc:resources location=".." mapping="...">
* configuration in web.xml
* * add a new <servlet-mapping> with the url-pattern element to match the mapping above

Sets up a FileResolver and Controller in the background

This forwards the request to another request mapping.
It goes outside of the framework and back into the new request

## Controller

* handle incoming requests and build responses
* interpret user input and transform into a model
* provide access to the business logic
* handles exceptions
* routes views accordingly
* business logic should not be handled here
* Request and response objects should not go into business logic
* get the information required and hand it off to the business logic

@Controller
@RequestMapping
@<HTTP_METHOD>Mapping
@PathParam
@RequestParam
@RequestBody
@ModelAttribute
    Used to link data in the model and view
    can be used for all HTTP Methods
    the view has extra html tags added which are used to map elements to models
@SessionAttributes
    save attributes from the model for the length of the session
@ResponseBody
    This is what I want to return to the client
    instead of returning a string which is the name of the view to render

## Service

* describes the (verbs) actions of our system
* where the business logic resides
* should not bleed into the controller or repository layer
* ensure the business object is in a valid state
* transactions should begin and end here
* often has same methods as repository but different focus

## Repository

* describes the (nouns) data of our system
* focused on persisting and interacting with databases
* one-to-one mapping to an Object

## Tags

Used to make interacting with data in views easier

2 tag libraries:

* spring
* * setting themes
* * internationalization
* * outputting errors
* * some examples:
* * * hasBindErrors
* * * * is there an error in our binding?
* * * url
* * * * invoke a url
* * * * escapes special characters in the url
* * * message
* * * * externalize strings out of pages
* * * * useful for internationalization
* * * * and repeating text
* spring-form
* * based around processing form data
* * most are same as the html form tags but have bindings for working with data
* * and can carry error messages with distinguishing css
* * you can use events also, onfocus, ...
* * some examples:
* * * checkbox
* * * form
* * * option

to add a tag library to a page:

```html

<%@ taglib prefix="spring" uri="http://www.springframework.org/tags" %>


<spring:message code="unique_idenitifer_in_properties_file"/>

```

It's not actually pulling from the uri, it is just the unique identifier for it.
It's packaged into the jar

We need a bean to look up the properties used by spring:message
in servlet-config.xml

```xml

<bean id="messageSource" class="org.springframework.context.support.resource.ResourceBundleMessageSource" p:basename="messages">

```

basename must match name of file

### Interceptors

not part of the tag library, but used with it.
registered as part of request lifecycle
intercepts data which is coming from a our page and going to our controller
**The ability to pre & post handle web-requests**
The have callback methods
more of an edge case for creating your own
Commonly used for Locale Changing

* Register a bean SessionLocaleResolver in the servlet-config
* Register a bean LocalChangeInterceptor inside an mvc:interceptor element the servlet-config
* * give it a property paramName with the value language
* append underscore and the language code to the name of the properties file
* same properties in each file, the values are in each language

## Validation

2 types of validation:

* Constraint validation
* * ensure inputs meet requirements
* Business logic validation
* * null values
* * the rest handled in service / controller classes

All form tags have an error class for showing an element in a particular font / class
There is also a specific error tag for displaying validation errors

Two ways to handle validation in Spring:

* Validator Interface & Validation Utils class
* * implement the interface and use the methods from the class for all kinds of field validation
* * Uses the BindingResult class
* * Uses the SimpleFormController
* JSR 303
* * Java standard for annotation based validation
* * can be extended to use our own validation rules
* * Hibernate Validator is a implementation of the spec
* * * but has nothing to do with data Hibernate, it started there but has become more generic
* * POJO Based

* add dependency, org.hibernate:hibernate-validator
* Use annotations on fields in POJO to validate
* * @Range(min = 1, max = 120)
* Put @Valid on parameter to Controller method validate
* Add a BindingResult parameter to the controller method
* Add logic in controller method to route to correct view based on errors
* * bindingResult.hasErrors()
* Update the view to show the error message
* * add css for error messages
* * add the error tag to the form, path="*", cssClass="...", element="div"
* * add the error tag to the fields in the form, name attributes as the input element
* * both display the same error in a different place
* * you can also add a cssErrorClass to the form input element

## Ajax

Spring MVC allows us to return json or html quite easily.
Not as feature rich as jersey but we only have to learn one framework

ContentNegotiatingViewResolver

* Uses the Accepts header: XML,JSON, HTML
* Can combine with other ViewResolvers
* Requires additional jars for marshalling
* * jackson for json

* add dependencies
* * org.codehaus.jackson:jackson-mapper-asl
* * com.thoughtworks.xstream:xstream
* * org.springframework:spring-oxm

* update servlet-config
* * make sure InternalResourceViewResolver is the last resolver by order
* * add a new bean, ContentNegotiatingViewResolver
* * * set order = 1
* * * set contentNegotiatingManager
* * * * with a bean ContentNegotiationManager
* * * * * which has constructor injection of PathExtensionsContentNegotiatingStrategy
* * * * * * which has constructor injection of a map
* * * * * * * json = application/json
* * * * * * * xml = application/xml
* * * set defaultViews
* * * * list of beans: MappingJacksonJsonView, MarshallingView (constructor-arg: XStreamMarshaller ( property: autodetectAnnotations = true))

* Add a new model and request methods

* configure web.xml
* * add a new servlet-mapping for json