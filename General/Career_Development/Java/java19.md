# java19

https://blogs.oracle.com/javamagazine/post/java-19-gems-jeps

## Virtual threads

First preview.
Part of Project Loom aimed at lightweight concurrency.

VirtualThreads behave almost identically to the Threads we know today.
And barely any changes to the API.
One Thread == One OS Thread, which could create a bottleneck in applications.

Pool of carrier threads onto which a virtual thread is temporarily mapped.
Once the virtual thread goes into a blocked state, it is switched off the carrier thread.

A carrier thread is an OS thread.

## Structured concurrency

Help simplify error management and subtask cancellation.
Treats concurrent tasks operating in distinct threads as a single unit of work, improving observability and dependability.

```java

try(var = new StructuredTaskScope.ShutdownOnFailure()){

  // execute tasks
  scope.fork(() -> /* method which returns Future */);

  scope.join(); // wait for all to complete, or at least one to fail or cancel.
  scope.throwIfFailed();
}
```

As soon as one task throws an exception all tasks are cancelled

## Record patterns

Can be use in instanceof calls and switch expressions.
Also works for Records which contain nested record types.

```java

if(object instanceof Point(int x, int y)){
  System.out.println("object is a point, x = " + x + ", y = " + y);
}

```

Instead of passing an identifier for object and executing point.getX(), you can specify the fields you want and calls them directly.

## pattern matching for switch

We're used to switch checking on a literal value.
You can now use it to check on a type and specific parameters of the type.
this release introduces a new contextual keyword to improve readability, when, instead of &&

```java

switch (obj) {
  case String s when s.length()>8 -> // execute code
  case String s -> // execute code
  case Integer i -> // execute code
  //...

}

```

## foreign function and memory API

Replacement for JNI (Java Native Interface)
Promoted from incubator to preview.

Enables access to native memory and native code.

## Vector API

Make it easier for native code and JVM code to communicate with each other.
fourth incubation of the API.
Low level stuff :-D
