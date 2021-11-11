# what's new in java 15

## Deprecations & Removals

Removed JDK Ports
  Solaris
  SPARC

RMI Deprecations
  java.rmi.activation package is deprecated

Deprecated tool: rmid - listening for rmi activation requests

Removed tool: rmic

Nashorn Javascript engine removed
  hard to keep up with JS

removed tool: jjs

no drop in replacement, alternative is GraalVM

## Helpful NPEs

no need to add flag, it is enabled by default

## Text blocks

No long a preview feature! We can start using it in production code

All the details are in my notes on java 13 and 14

""" needs to be escaped if within a text block

## Preview features

### pattern matching

same notes as java 14

### records

same notes as java 14

a restricted form of a class to model data as data

record key word, like class key word

records can implement interfaces

you can define local records, within a method.
implicitly static
Can we useful to create types on the fly for use in streams

records are implicitly final

### sealed classes

Control the inheritance hierarchy of a class

sealed and permits keyword

```java

public sealed class MyClass permits Other, ThatLad{

}

```

Another option is to define the implementations inside the sealed class as static inner classes
Also, you can define sealed types in the same source file

```java

public sealed class MyClass {

  public final static class Other extends MyClass {}
}

```

you can mix sealed and non-sealed in a class hierarchy
non-sealed is another new keyword

## JVM improvements

### Garbage Collections

#### ZGC

low pause times, under 10ms. Goal is under 1ms
Can scale to multi terabyte heaps
should perform well without any tuning

now production ready!!!

Concurrent garbage collector, i.e. doesn't stop your app
Compaction and relocation, regions
  allows for efficient use of memory
encoding GC specific info into pointer. Coloured pointers. Only implemented on x64 platforms

-XX:+UseZGC

#### Shenandoah

developed by red hat
back ported to 8 & 11

goals, low pause times and large heaps

not included in all openjdk builds

### Hidden Classes

vm level feature
aids dynamic code generation, only can be used by the class that generated the code

### Edwards-Curve digital signature algorithm

cryptography feature in jvm
more secure connections

### reimplement datagram socket API

Allows to send UDP packects
