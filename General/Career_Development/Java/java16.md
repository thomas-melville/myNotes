# what's new in java 16

openjdk has switched from mercurial to git
  repo size is now smaller
openjdk has moved to host the repositories on github

## preview features & incubating apis

Sealed Classes
Vector API
  for working with high precision numbers
Foreign Linker and memory access APIs
  statically typed pure native access to native code

## api changes

### Stream API updates

collect(Collectors.toList())

can now be shortened to:

toList();

it is an unmodfiable list, collectors is modifiable

mapMulti method
  return 0 or more objects of type R from the input stream T
  similar to flatMap, but you need to output streams from it
  mapMulti is a good alternative to flatMap
  it's a bit more imperative, but can be more efficient

### Date formatting

DateTimeFormatterBuilder

parse/format the textual name of the day period

10:05  -> in the morning

.appendDayPeriodText(TextStyle.FULL)

## records

no longer preview!!!

same notes as previous versions

behave like tuples, but access by name and not index

record can have static fields

compact constructor form, leave off parenthesis only need curly braces

Records are compiled to normal java classes!

You can override methods

implicitly final
can't extend a class, because under the hood it is compiled to extend the Record class
Can't have native methods
compatible with serialization
Components can have annotations
  can be propagated to field, accessor method and constructor parameter

You can define a constructor with fields which aren't the components.
However, this constructor must delegate to the canonical constructor (i.e. the one generated by the jdk)

Records and reflection.
Class new methods isRecord(), getRecordComponents()

use cases for records

API layers, i.e. DTOs
Jackson already supports records

JPA records is not going to happen!

Compound map keys

multiple return values.


## pattern matching

no longer preview!

same notes as before

type must be the exact type, subtype will fail

Compiler is smart enough to know which branch of the if the type can be used in!
Flow analysis, specifically flow scoping

Compier will show red if check will always be true

type is in scope for && logical checks in the if statement

```java

if( string instanceOf String s && s.startsWith("R")){
  //...
}

```

## jpackage

no longer a preview!

jpackage bundles your application and a java runtime image + platform specific scripts to install the application
works with modular & non-modular apps

windows
  .exe or msi
mac
  .dmg or .pkg
linux
  .rpm & .deb

jpackage on each OS has some prerequisites

geared towards desktop applications

can't update the application or jvm, you would need to generate a new jpackage

2 step process
  1. creates a runtime image
      uses jlink underneath
  2. create native package

you can add cross-platform metadata
  version, copyright, ...

you can configure application start up
  including jvm options!!!
    --java-options

## jvm changes

### ZGC improvements

alot of the same notes

goals for java 16

sub-millisecond pause times!

Concurrent Thread Stack processing! JEP-376

### Strong encapsulation

Of JDK Internals

Encapsulate internal types
  in sun or oracle specific package
Block deep reflection on public types

It is blocked if the application is modular
However, up to java 15 it was still possible in classpath applications

It is now blocked for classpath apps in java 16
Private really is private now!

This will improve maintainability for JDK

--illegal-access is no longer permit by default, it is now deny
this flag is deprecated and will be removed so setting the flag to allow is only short term