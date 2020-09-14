# Java Module System

## why?

A class may make use of other classes at
  compile time (compiler checks dependency on compiling classpath)
  or at runtime (JRE tries to resolve dependency on the runtime class path)

Modularize the JDK itself
JRE delivers many classes, which may not all be used. Bloating final executable size
  rt.jar to many small jar files (more than 90)

For a class to be seen across packages it must be public, which means everyone can see it.
3PPs had no way of hiding internal classes

Java 9 tried to address this problem by providing another way to manage accessibility, modules.

increased security
  explicitly expose some packages
  and strongly encapsulate other packages
  reducing the surface area of attack

easy deprecation
  modules are isolated so can be removed over time
  corba has been marked deprecated in java9

easier to ship new features
  incubator modules can be added
  the new httpclient is an example of that

reduced footprint
  when you explicitly use the module system to bring in only what you want

module system is optional

## what is a module

A high level grouping construct which refers to other modules when necessary
has a name, it groups related code and is self-contained or explicitly states the modules it depends on for functionality it needs

A jar file is a little bit like a module:
  have (file)names
  group related code
but
  no explicit dependencies
  weak boundaries

## kinds of modules

1. Java and Jdk
    Provided by the JDK
      it has been completeley modularized
      unless you want modularization everything is brought in
    2 sets of sub packages:
      java, java.base for example.
        java.base is most important, every module has a dependency on it (transitively)
      jdk, httpserver, incubator & unsupported for example
        these are jdk specific
    no cyclic dependencies
2. Explicit application modules
    a jar file which has a module-info class at it's root
    Explicit, because of the metadata definition class
    application, because they don't come from the jdk
3. Automatic modules
    a non-module jar which is placed on the module path
    Add a Automatic-Module-Name entry to the jar's MANIFEST.MF file.
    this will be used for the module name
    by default it is derived from the jar's name
4. unnamed modules
    when using the classpath, every jar that is not modularized will be part of the unnamed module.
    this happens when you migrate to java9 and don't make use of the module system
    automatically gets dependencies to jdk modules

## module-info class

module descriptor

adds new keywords to the language

located at each JAR's root
a module definition:
  containing module name
  declaring required module dependencies
  exposed APIs
  provided services

at runtime loader reads the manifest file to load only modules that are necessary

```java

// module-info.java
module ch.frankel.jlink {
    requires org.slf4j;
    exports ch.frankel.blog.jlink;
}

```

```bash

java --list-modules # to list all modules
java --describe-module <module_name> # to describe module
java --illegal-access=deny # fail if a class is using an API encapsulated inside a module, IllegalAccessError
    # it's set to permit by default
    # logs warnings for each reflective illegal access!
    # permit mode is likely to be removed at some stage

java(c) --add-exports <module>/<package>=<module> <class>
   # if you are using an encapsulated API you can export it on the fly to your module

```

### keywords

module:
  defines a module
exports:
  APIs which are visible outside the module
requires:
  Modules which this module depends on
transitive:
  a module depending on this module will also get any dependencies marked as transitive
mandated:
uses:
  hooks the module into the Java ServiceLoader mechanism

## migrating to java9

module system is optional so compile and run as is
unless:
  you use jdk APIs which are now encapsulated inside a module
   once you compile with java9 you will get an error
   you can add a flag to allow it --illegal-acces=permit
   you can add a flag to export the module to your module
  you use types from non-default java modules
    java.xml.ws
    java.xml.bind
    java.activation
    java.corba
    java.transaction
    java.xml.ws.annotation
    --add-modules flag to add the missing module, or higher level module to add them all
      at compile and runtime
