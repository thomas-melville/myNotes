# java18

https://blogs.oracle.com/search.html?q=java18

## UTF-8 by default

variable width character encoding and considered the web standard.
pre-18 could vary based on runtime env language settings and OS.

Will help with Input/OutputStream and Reader/Writer code.

## Simple Web Server

jwebserver command, serves the local folder on port 8000
Only GET & HEAD methods supported
HTTP/1.1
No HTTPS

You can also run it from the Java API.

```java

SimpleFileServer.createFileServer(..)
server.start();
```
## Code Snippets in Javadoc

@snippet annotation for inside javadoc comment.

```java

/**
  *
  * {@snippet :
      try(var reader. . .)
  * }
  *
**/

```

## Reimplement core reflection

Two ways to print the private field of a class using reflection.
  core reflection
  method handles.
A third hidden option, leveraging native methods after the JVM has been running for a while.

Maintenance would be a lot of effort.
JDK team have decided on method handles for it all.

## Vector API

## Internet Address resolution SPI

Developed SPI for hostname and address resolution.
Allows InetAddress to use resolvers other than the operating systems native one.
Benefits:
1. Enables seamles integration of emerging network protocols
2. more control to frameworks and easier to retrofit
3. allow Project Loom to work more efficiently, current API resolution operation blocks an OS call

## Foreign function and memory api

## pattern matching for switch

improvements on the previous preview releases.
Solving edge cases of multiple conditions in a case, which are also encompassed in a more general case.

## deprecate finalization for removal

An object can implement the finalize() method, which will be called by the GC before freeing the memory.
It has had it's issue so it's marked deprecated for removal.
