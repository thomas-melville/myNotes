# java 22

https://www.loicmathieu.fr/wordpress/en/informatique/java-22-quoi-de-neuf/

## JEP 461 Stream Gatherers

Custom intermediate stream operations.
Define your own using Stream::gather(Gatherer).

represents the transformation of an element in a Stream, one-to-one, one-to-many, many-to-one, or many-to-many.
Can also stop if necessary.
Gatherers can be combined: stream.gather(...).gather(...).collect(...)

Gatherer interface defines the followings methods:

* integrator - used to integrate a new element from the incoming stream, and if necessary emit an element in the downstream
* initializer - optional, can be used to maintain a state when processing elements.
* combiner - optional, can be used to evaluate the gatherer in parallel
* finisher - optional, called when the stream has no more elements

Stream API has been enhanced with the following gathers:

* fold - stateful many to one gatherer that builds an aggregate incrementally and emits it when there are no more elements.
* mapConcurrent - stateful one to one gatherer that concurrently invokes a provided function for each input element, up to a provided limit
* scan - stateful one to one gatherer that applies a provided function to the current state and the current element to produce an output
* windowFixed - stateful many-to-many gatherer that groups input items into lists of a supplied size, outputting windows when full.
* windowSliding - stateful many-to-many gatherer that groups input items into lists of a supplied size. After the first window, each subsequent window is created from a copy of its predecessor by deleting the first element and adding the next element from the input stream.

## JEP 458 Launch multi-file source code programs

Since Java11 it's been possible to execute a single java source file without compiling it first.
Now it's possible to execute a file which references another file by including the file in the command.

## JEP 447 Statement before super

Up to now the call to super had to be the fist command in the constructor of a sub class.
If it's omitted and the compiler finds a matching constructor it automatically includes it.

Now it's possible to include statements before the call to super so long as they do not access the instance being created.
Examples would be parameter validation, argument pre-calculation, etc.

## JEP 457 Class-file API

A standard API for parsing, generating and transforming Java class files.

## ListFormat

## Features coming out of preview

## Features that remain in preview
