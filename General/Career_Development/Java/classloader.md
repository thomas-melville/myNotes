# classloader

An object which is responsible for loading classes.
A typical strategy is to read a class file from the filesystem.

Every class object contains a reference to the ClassLoader that defined it.

```java

public class MyClass {

  public static void main(String[] args){
    ClassLoader loader = MyClass.class.getClassLoader();
  }
}
```

The ClassLoader class uses a delegation model to search for classes and resources.
Each instance of class loader has an associated parent class loader.
A class loader instance will delegate the search to its parent class loader before attempting to find it itself.

The JVM has a built in class loader, boot strap class loader.
This class loader has no parent.

Normally, the JVM loads classes from the local file system in a platform dependent manner.
On UNIX, the JVM loads classes from the directory defined by the CLASSPATH environment variable.

There are 3 class loaders:

1. bootstrap/primordial class loader
  * loads core java libraries
  * JAVA_HOME/jre/lib
  * written in native code
2. extensions class loader
  * loads the java extensions
  * JAVA_HOME/jre/lib/ext
  * sun.misc.Launcher$ExtClassLoader
3. System class loader
  * loads code found on the CLASSPATH
  * sun.misc.Launcher$AppClassLoader

Class loader implemenations are useful for:

* modifying byte code, for example load-time weaving of aspects when using AOP

## api

getResource - returns a url
getResourceAsStream - returns an InputStream
