# Introduction to Spring MVC 4

## What's new in version 4

* Java Configuration is on feature parity with XML configuration
* support for java8
* JEE 6 and Servlet 3+
* Better support for  REST and WebSockets
* Improvements to allow you to go container less
* * don't create a war which is deployed in a web container
* * package the web container into the jar
* * this is spring boot

## Configuration

no need for any xml configuration files now! as long as you are using Servlet 3+
Every element from the xml files now has a corresponding Java Annotation

* component-scan
* interceptors
* handlers
* formatters
* converters

Also some convenience annotations similar to namespaces

Servlet 3+ negates the need for a web.xml file
Exposes some interfaces for initializing which Spring implements and makes callable through annotations

You can still have a web.xml if you want, just point it to the config class instead of the config location

Without web.xml:

* Implement WebApplicationInitializer, which builds your application context ( the same thing the web.xml did)
* * normally put it in a config package
* * has one method
* * onStartup, which hooks into the servlet engine
* Create a WebApplicationContext, which is an instance of AnnotationConfigWebApplicationContent
* * set the config location using setConfigLocation(<class_name>) or register(<class_name>) (if the spring mvc version is newer)
* Create an instance of a ContextLoaderListener with the context above
* Add that to the servletContext as a listener
* Create a DispatcherServlet with the context above
* Add the DispatcherServlet to the servletContext as a dispatcherServlet
* Return that servlet as a ServletRegistration.Dynamic
* set the load on startup = 1 in the dynamic dispatcherServlet
* add a mapping to the dynamic dispatcherServlet

```java

public class WebAppInitializer implements WebApplicationInitializer {

    @Override
    public void onStartup(ServletContext servletContext) throws ServletException{
        WebApplicationContext context = getContext();
        servletContext.addListener(new ContextLoaderListener(context));
        ServletRegistration.Dynamic dispatcher = servletContext.addServlet("DispatcherServlet", new DispatcherServlet(context));
        dispatcher.setLoadOnStartup(1);
        dispatcher.addMapping(".html);
    }

    private AnnotationConfigWebApplicationContext getContext(){
        AnnotationConfigWebApplicationContext context = new AnnotationConfigWebApplicationContext();
        context.setConfigLocation("com.ericsson.app.config.WebConfig");
        return context;
    }
}

```

* Still the same JSP pages with spring tags
* First class to create is the config class and annotate it with:
* * @Configuration
* * @EnableWebMvc - servlet-config.xml
* * These two annotations create the same thing that the appContext.xml and servlet-config.xml created

* The namespaces in the xml have an equivalent java class which can be extended: **WebMvcConfigurerAdapter**
* Then override methods from that class with your own implementation which is the equivalent of the xml configuration
* For example: serving static content
* * mvc:resources element -> override addResourceHandlers(ResourceHandlerRegistry registry)
* Extend from the WebConfig class

```java

@Override
public void addResourceHandlers(ResourceHandlerRegistry registry){
    registry.addResourceHandler("/pdfs/**").addResourceLocation("/WEB-INF/pdf/");
}

```

## Architecture

Architecture is pretty much the same as 3

With the move to javascript front ends updates have been made to support MV* patterns
This also helps with mobile apps and restful services
Controllers are more lightweight and added features for more REST based Controllers

### MVVM

* originated from .NET
* intended for XAML
* principles adopted to other languages
* usually tied to a RESTful backend and a rich javascript frontend
* * which means layers are cleanly separated
* * Not always, data can leak between the two ;-)

* View and Controller (View Model) are on the frontend
* Model is on the backend
* Controller communicates with the Model

## Controllers

@Controller

* Signifies a class is a WebController

@RequestMapping
@RestController
@Configuration
@EnableWebMvc

* Convenience annotation for WebMvcConfigurationSupport
* Only used for the Java Config of Spring MVC apps
* Customizable by extending the WebMvcConfigurerAdapter

@ComponentScan

* Exact same purpose as component-scan in xml
* has a field basePackages

## Service

## Repository

Get's a little complicated with the java config
There are different types of repositories if you are using Spring Data JPA
Convention over Configuration will save you a lot of time

## Views

returning the full filename will find the jsp file directly under WEB-INF
But we want to control user experience so it should be put under WEB-INF/jsp
which means we need to configure the app

* Add a method to our @WebConfig annotated class which is annotated with bean
* create a new instance of InternalResourceViewResolver
* set the suffix
* set the prefix
* return it

```java

@Bean
public InternalResourceViewResolver getInternalResourceViewResolver(){
    InternalResourceViewResolver resolver = new InternalResourceViewResolver();
    resolver.setPrefix("/WEB-INF/jsp/");
    resolver.setSuffix(".jsp"):
    return resolver;
}

```

The java and jsp files are the same as in the first spring mvc course

## Internationalization

* getResourceBundle Bean
* * Autowired by name so the method name must match, over else it won't work.
* * This bean loads in the properties file which starts with the basename messages
* getSessionLocaleResolver Bean
* * also autowired by name
* * stores the language as a session variable instead of reading from the header each time
* addInterceptor override
* * the method from the base class
* * detects if we change the language
* Create the resource bundle
* * messages_<language>.properties in src/main/resources
* Spring messages tag in jsp page
* * add spring taglib
* * <spring:message code="attendee.name"/> code is property name in file
* add href for changing language on the page
* * <a href="?language=en">

```java

public class WebConfig extends WebMvcConfigurerAdapter {

    @Bean
    public MessageSource messageSource(){
        ResourceBundleMessageSource messageSource = new ResourceBundleMessageSource();
        messageSource.setBasename("messages");
        return messageSource;
    }

    @Bean
    public LocaleResolver localeResolver(){
        SessionLocaleResolver resolver = new SessionLocaleResolver();
        resolver.setDefaultLocal(locale.ENGLISH);
        return resolver;
    }

    @Override
    public void AddInterceptors(InterceptorRegistry registry){
        LocaleChangeInterceptor changeInterceptor = new LocaleChangeInterceptor();
        changeInterceptor.setParamName("language");
        registry.addInterceptor(changeInterceptor);
    }
}

```

## Validation

* Tags in JSP page
* maven dependencies
* BindingResult in Controller method signature

* add dependencies
* * javax.validation:validation-api
* * org.hibernate:hibernate-validator
* Add annotations to fields of model
* * @Size, min & max
* * @NotEmpty
* * @Email
* Add @Valid to parameter of controller method
* Add BindResult parameter to controller method
* * We can check this in the method for errors.
* Add Model parameter to controller method
* * We can use this to pass any information back to the view
* Parameter order in method signature is important!!!

Error messages in Validation are already internationalized

### Customizing validation error messages

Add entries in the message.properties file.

format in file is as follows:

```properties

<annotation>.<class>.<field>

```

Only available with certain annotations, not all.
This works with custom validation annotations you create.

This is then used to internationalize the cutom messages by placing corresponding entries in other language files

### Creating custom validations

Why?

* Cleaner implementation
* More descriptive

* Create your annotation
* * @Documented
* * @Constraint
* * * set field validatedBy to your class below
* * @Target
* * * Method & Field
* * @Retention
* * * Runtime
* * String message() default "{Phone}";
* * * **This is how the annotation is tied to the property in the message.properties file**
* * Class<?>[] groups() default {};
* * Class<? extends Payload>[] payload() default {};
* Implement the ConstraintValidator interface
* * It's generic with two types
* * * the annotation above
* * * The object to return
* * 2 methods to override
* * * initialize
* * * isValid
* * * This is where you do the validation

## RESTful Services

ContentNegotiatingViewResolver -> @RestController

@RestController

* A convenience annotation for annotating a Controller and the ResponseBody at the same time
* Assumes each RequestMapping is going to return the body of it's request

**Add .json mapping to the DispatcherServlet configuration.**