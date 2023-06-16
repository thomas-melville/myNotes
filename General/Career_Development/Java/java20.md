# java20

https://blogs.oracle.com/javamagazine/post/java-20-gems-jdks

## scoped values

Modern alternative to thread local variables.
maintain immutable and inheritable per-thread data.
reduce lifetime of data also. Once method which creates it has finished the variable should also.

## virtual threads

Fundamentally redefines the interaction between the Java runtime and the underlying OS.
However, virtual threads behave almost identically to familiar threads.
barely any additional API.
This Jep allows for more feedback on the API while in incubator.
Virtual threads will likely be production ready in the next release.

## structured concurrency

Aims to make multithreaded programming easier to help simplify error mgmt and cancellation.
Treats concurrent tasks operating in distinct threads as a single unit of work., improving observability and dependability.
API has not changed, allowing time for feedback

## record patterns

Provide an elegant way to access a records elements after a type check.
second preview with further refinements.
* add support for an inference of type arguments for generic record patterns.
* add support for record patterns to appear in the header of an enhanced statement
* Remove support for named record patterns

## pattern matching for switch

preview feature since 17, another preview this time.
* exhaustive switch, change exception if no matches.
* switch label's grammar has been simplified
* inference of type arguments for generic record patterns

## foreign function and memory api

Improve operability between JVM and well-defined foreign/non-java APIs.
Defines API through which Java programs can interoperate with code and data outside the runtime.

Removes the need for the brittleness and danger of Java Native Interface (JNI)

## Vector api

computations.
this version has a small set of bug fixes and perf enhancements.

## non-jep changes.

### lossy-conversions lint option

new javac lint option which generates warnings for implicit casts in compound statements

```java

a += b
```
If a is a different type to b that requires a cast then the compiler puts one in.
This lint option will add a warning.

### New APIs for TLS and DTLS key exchange named groups

javax.net.ssl.SSLParameters.getNamedGroups()
javax.net.ssl.SSLParameters.setNamedGroups()

### JMod command line option

--compress

### Garbage collection

G1 improvements.

### unicode 15.0 support

new characters

### Timz zone data updates

match IANA

### 
