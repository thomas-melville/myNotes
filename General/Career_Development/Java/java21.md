# java21

https://blogs.oracle.com/javamagazine/post/java-inside-21-features

## virtual threads

no longer preview, final release.
Now it's up to the java projects to start using virtual threads!
A replacement for CompletableFutures?
Don't use on CPU intensive code, great for blocking IO code.

## sequenced collections

a new set of interfaces for collections
SequencedCollection, which extends Collection.
Implemented by all lists and some sets and maps
new methods:
* add, get, remove First/Last
* reversed

## Generational low pause garbage collection.

improvements...

## pattern matching

finalized
using types in switch expression.
Type must be in a sealed class hierarchy

## key encapsulation mechanism api

Java now has an API for it.

## record patterns

add nested patterns
no longer preview, only minor improvements from now on

## minor api improvements

### http client

now implements auto closeable.

## preview JEPs

### Structured Concurrency

Once you get into Virtual Threads you can start to treat them as a set of tasks as if they're executing a single unit of work.
you can see them as children of the thread that created them.

Streamline error handling, cancellation, improve reliability and ehance observability.

It's a new API, which implements auto closeable.
fork tasks, new type SubTask
then join to wait
and throw exception if failed.

### Scoped Values

ThreadLocal API is for Platform Threads, it has some short comings.
...
Enter Scoped Values.
Solves thread local issues:
* within the scope, the bound value is immutable.
* no copies need to be created when inheriting
* visible only within the defined scope.

### Vector API

nothing new, waiting for project Valhalla

### Foreign function and memory api

better way of calling native code.

### unnamed classes and instance main methods

main method no longer needs to be public, static or have args.
Surrounding class is optional too.

### unnamed variables and patterns.

use _ to say I don't care about the name of this variable.
can be used multiple times in a single pattern.
type can be omitted too.

### String templates

embed expressions in Strings.

syntax: \{variableName}
