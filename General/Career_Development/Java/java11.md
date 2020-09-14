# java 11

Http Client API
ability to run source code directly
two new garbage collectors
other improvements

## licensing

can't use oracle jdk in production without a license
must use openjdk

LTS support is only for oracle jdk, which means you need to buy a license
open jdk will only have 6 monthly support

AdoptOpenJdk is another option, that is built on top of open jdk

## launching single-file source code

pre-java 11

javac .java
java .class

java 11

java .java
compiler sees the file is .java so compiles and executes it
limited to a single source file

--source <version> version of java to compile to, assumed to be current
  this flag allows us to pass in multiple java files

This allows us to write java scripts, like bash scripts, not javascript!

```java

#!/user/bin/java --soource 11

import hello; //...

public class Test {
  public static void main(String[] args){
    //..

  }

}

```

## deprecations & removals

### removed enterprise APIs

JaxB, java.xml.bind
  replace with maven dependency, javax.xml.bind:jaxb-api & jaxb-impl
JAX-WS
JAX-annotations
CORBA
JTA
Javabeans Activation

jep 320 describes all the replacements

Enterprise applications servers already have these APIs so you shouldn't be affected.

If you're an SE application and you use these you'll have to make changes.
maven dependencies

### removed methods

Thread
  destroy
  stop(throwable)

runFinalizersOnExit has been removed from System & Runtime
  finalizers have been discouraged for a while!

SecurityManager
  checkAwtEventQueueAccess
  checkSystemClipboardAccess
  checkTopLevelWindow
  checkMemberAccess

this has a dependency from a core API, Security Manager, to a high level API, AWT
  proved problematic in modularization
  replace with checkPermission(java.security.Permission)

### JavaFX

GUI Technology in Java
moved to openJFX
  openjdk subproject
maven central dependency

java packager removed, no replacement but a proposal, jpackager

### removed & deprecated technologies

applets removed
  buy commerical support for java8
java web start removed
  jlink
  jpackager
nashorn deprecated
  consider graal.js

## Language & Library improvements

### Http Client API

incubator has become GA

HttpClient
HttpRequest
HttpResponse
BodyHandlers

### String

repeat(int number_of_times)
  repeat the string instance number of times
isBlank()
  equivalent to "".equals(string.trim())
strip()
  strip all white space from a String
  better than trim because it can deal with unicode
stripTrailing
stripLeading
lines
  for a multi line string returns a stream of strings

### Files

readString(Path p)
  pass in a path and get back the string contents of the file
  don't have to worry about opening/closing a file or converting bytes
  assumes file is encoded with utf-8
  overloaded method where you can pass in character set
writeString(Path p, CharSequence cs, OpenOptions oo)
  write a String to a file
  OpenOptions specifies how to open the file
    append/overwrite/...

### Optional

isEmpty

### Predicate

not
  useful for negating an existing predicate
  Predicate.not(String::isBlank)

### Unicode 10

aligned to unicode 10
16000+ new characters
localized apps for new languages

### Local-variable syntax for Lambdas

now you can use var as the type of parameters to a lambda

```java

(a, b) -> a.concat(b)
// is equivalent to
(var a, var b) -> a.concat(b)

```

why would you want this?
you can then annotate the variables when a type is provided

limitations
  must be on all parameters
  can't mix var and explicit types
  must use brackets with var

### Nest-based access control

in the jvm
compiled inner class has access to a bridge method in the outer class to access private methods/variables

In java 11 inner & outer class will be inside a nest
and called nest mates
outer class will be nest host
inner class will be nest member
runtime knows that inner is allowed access to outer
  so no need for compiler to generate bridge method

## Performance & Security improvements

### Garbage collection improvements

incremental improvements to default G1GC

new Epsilon GC
  no actual garbage collection
  just keeps on giving memory until system runs out and program crashes
  why?
    apps with predictable, bounded memroy usage
    set the heap size and wlka away
    no garbage collection to better Performance
    for short lived apps
    during performance testing, to see
  experimental

Z Garbage Collector
  Pause times under 10ms
  keep pause times constant
  multi terabyte heaps!
  coloured pointers
    reserves 4 bits on pointer tyo give infor about object reference
  experimental

### TLS 1.3

implemented in jdk
  legacy algorithms pruned
  all handshake messages are encrypted except first
  elliptic curve algorithms in base spec
not all features of TLS 1.3 implemented!
