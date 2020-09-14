# java 10

--release flag to compile with older version that currently on machine!

## what's in java 10

local variable type inference
  reduce boiler plate
  use keyword var and let the compiler determine type
Improvements to G1 garbage collector
Application class data sharing
improved docker container awareness

## Root Certificates

Oracle's Root CA Program to be ported to OpenJDK
jep319!

## deprecations and removals

Removed command line tools and options
  javah
    javac -h now
  policytool
    add policy files
  -X:prof
    print a warning intead of profiling
    jmap
    3pp profilers

Deprecated APIs
  java.security.acl
    java.security is the replacement
    Certificate,Identity,IDentityScope, Signer
  javax.security.auth.Policy
    java.security.Policy

## local variable type inference

all about variables
inside methods
replace type of variable with the new keyword "var"
compiler will infer type from value
  it's called type inference

why?
  type less code
  vertical alignment of variable names
  focus on variable names
  great for construction using new keyword
    type is on the RHS already!
  or constant

why not?
  keep the code readable
  assign return from a method as var, not clear what the type is.

*var is not a new keyword!*
  reserved type name!
  won't break backward compatibility

*must be initialized when declared*
  no initialization results in compiler error
  initializing an array!
    must use new int[], can't use {}

type inference
  this does not make java a dynamic type language
  type is still there, it's now infered by the compiler
  using var or explicit type results in the same bytecode!
  no runtime overhead
  diamond operator was type inference!
    but the direction was opposite
  lambdas have type inference

### limitations

can't have type inference on both sides!

```java

var p = s -> s.length() >3

```  
just not enough information for compiler to figure out types

lambdas must always have an explicit target type, no local varialbe type inference!

var + diamond operator

```java

var list = new ArrayList<>();

```

What will happen? *type will be inferred as ArrayList<Object>()*
  Not raw type! Parameterized type is Object!
  And you end up with concrete class of ArrayList, not List interface

var is only possible in local variables
  also as type in for each constructs

not on
  method parameters
  fields
  return types
  catch blocks

compiler would have to look at all invocations and figure out the best type to infer!!
  adding a new invocation of a method could change the supertype!

### Non-denotable types

1. Null
  compilation will fail

2. Anonymous inner class instances
  results in an inner class
  ```java

  var obj = new Object(){};

  ```
  results in Object$1 extends Object
  becomes an issue if someone does obj = new Object()
    Object does not match Object$1 extends Object


3. Intersection types
  combine several class types into a single type
  var list = List.of(1, 2.0, "3")
    results in List<? extends Serializable & Comparable>

## Performance Improvements

### G1GC Parallel Full Garbage Collector

default in java 9
heap is divied into cells
  eden
  survivor
  old
incremental gc
  can select sub set of cells
parallel marking
  scan the heap in parallel to your code
designed for large heaps
tunable pause fgoal
slightly more CPU intensive

stop-the-world garbage collection
  serial full GC in java 9
  parallal full GC in java 10

## Application Class Data Sharing

improve vn startup time
reudces memory foot print of jvm
shared lib of class metadata
persisted into a shared lib
mem mapped by a different jvm
original system classes
it;s now all classes
useful for multiple vms running the same code

run with -Xshare:dump flag to export system classes
  classes.jsa

run with -Xshare:on
  to get exported .jsa file

10-30% improvement in start up times

a lot more arguments for sharing application class data files
  -XX...

### Improved Container Awareness

Java now detects container
query container instead of Host OS
  available CPU & Memory

some flags
-XX:ActivePRocessorCount
-XX:InitialRAMPercentage
-XX:MaxRAMPercentage
-XX:MinRAMPercentage

can now easily attach to the java process running inside the docker container from the host os
Linux & Docker only
