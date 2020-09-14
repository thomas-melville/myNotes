# jlink

https://blogs.oracle.com/javamagazine/containerizing-apps-with-jlink

available since java 9

assemble and optimize a set of modules and their dependencies into a custom runtime image
can also make this image executable so there is no reliance on target system having compatible JRE

```bash

jlink --add-modules ch.frankel.blog.jlink \
    --module-path ${JAVA_HOME}/jmods:target/jlink-ground-up-1.1.jar \
    --output target/jlink-image \
    --launcher hello=ch.frankel.jlink/ch.frankel.blog.jlink.Main

```

--add-modules
  define the modules required by their names
--module-path
  specify module path
  this is used instead of the class path
--output
  specify output folder of resulting jar
--launcher
  specifies entry point of the custom distribution
  consists of several parts:
    final executable name
    =
    module name
    /
    fully qualified class name of the Main class

Managing the dependencies to the jlink command would become very cumbersome
Enter maven dependency! Moditect plugin: https://github.com/moditect/moditect

When adding a dependency in the pom it must also be added to module-path in the maven plugin
this also becomes cumbersome
advice is to copy all dependencies into directory and use that directory as the module-path in the maven plugin
