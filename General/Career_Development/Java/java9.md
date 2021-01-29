# java9

September 2017

what's new?

Lot's!!!! http://openjdk.java.net/projects/jdk9/

* modularization of the jdk
    * This means that not all APIs available in java8 will be available in java9
    * javax.xml being a prime example
    * https://www.oracle.com/corporate/features/understanding-java-9-modules.html
    Separte folder!!
* new JsShell
    Separate file!!!
* Process API improvements
* lot's a small improvements
  Optional
  Streams
* new doclet API
* new versioning scheme for java
* javadoc search in html5
* new jvm diagnosis commands
* new keystore features
* multi release jar files
* G1 as default garbage collector
* Stack walking API
* Reactive Stream Flow API

## Collection Factory Methods

List.of(...)
  create a list containing the elements passed to the method
  returns *ImmutableCollections#ListN* type
    there are 2 more types: List1 & List2
      super optimized for tiny lists

List.of()
List.of(E e1)
List.of(E e1, E e2)
overloads all the way up to 10 elements!
List.of(E... elements)
  then the varargs method kicks in
  avoids intermediate array allocation

Set.of(...)
  if you pass in a null it will throw an NPE
  if you pass the same value twice it will throw IllegalArgumentException

all the same overloads as List

Iteration order is not guaranteed

Map.of(k, v, k, v, k, v, ...)
  won't compile if uneven number of parameters
Map.ofEntries(Map.entry(k, v), ...)

no nulls for keys or values
no duplicate keys

multiple overloads list List & Set

Iteration order is not guaranteed

## Stream API Improvements

### 4 methods added:

takeWhile(Predicate)
  take from the stream until the predicate is fales.
  then drop everything else
  works well on ordered streams,
  can lead to unusual results on unordered streams
dropWhile(Predicate)
  drop from the stream until the predicate is false.
  then take everything else
  works well on ordered streams,
  can lead to unusual results on unordered streams
ofNullable(T)
  create a stream of 0 or 1 elements
iterate(T seed, Predicate, UnaryOperator)
  start at seed, apply unaryoperator, stop at Predicate
  previously predicate was not there so iterate would continue
  removes the need for calling filter after creating stream

### new Collectors

very useful when used as downstream collectors in groupingBy calls

Collector filtering(Predicate, Collector)
  use when collecting to apply a filter to what gets collected.

Collector flatMapping(Function, Collector)
  map an element from a stream to another stream of elements
  then collect

## Optional Improvements

3 added methods

void ifPresentOrElse(Consumer, Runnable)
  if value is present call Consumer
  otherwise execute Runnable
Optional<T> or(Supplier)
  if the optional is empty execute the Supplier and return an Optional
  much better than using orElse and get!
Stream<T> stream()
  turn the optional into a stream
  if the optional is empty then it's an empty stream
  very useful when a stream contains optionals

## Small language changes

underscore as an identifer is now illegal
  deprecated in java8
  can still be used in variable names
  but a variable name _ is now illegal
improved try-with-resources
  can now accept effectively final variables
  if you already have a stream you can pass it to the try() without creating a new instance
better generic type inference for anonymous classes
  in java 8 overriding a method of a generic type in an anonymous class would fail
  now it's ok
private interface methods
  create private methods in an interface
  can be used by default methods
  but not by implementing classes
  reduces code duplication in interface default methods

## Other Improvements

Javadoc has been overhauled to be HTLM5 compliant
Search is in javadoc now
Modules in javadoc

Localization
  up to Unicode 8.0
  properties files to UTF-8
  Common Locale Data Repository used to get information
    instead of own storage

java.time.Duration
  long dividedBy(Duration)
    how many times the duration is divided by the given duration
  Duration truncatedTo(TemporalUnit)
    truncate the duration at the given unit
    if minutes is passed, the seconds will be dropped and rounded to the nearest minute

java.time.Clock
  Clock systemUTC()
    give back best and most percise system clock available

java.time.LocalDate
  datesUntil
    gives a stream of days from the date until the end date
    has overloaded method which takes a Period
      which results in dates which are that period apoart

## ProcessHandle

Represents any native process running on the OS.
Existing Process API Represents a native process created from Java

Process#toHandle() creates process handle for the given process

ProcessHandle exposes new info about the Process
  parent
  descendants
  destroy
    throws IllegalStateException if you try and kill current process, i.e. jvm
  pid
  onExit
    returns CompletableFuture
      which completes once process has quit
  ...
  info
    ProcessHandleInfo
      command
      arguments
      user
      startInstant
      totalCpuDuration


ProcessHandle.of(pid)
ProcessHandle.current()
ProcessHandle.allProcesses()
  stream of ProcessHandle

## HttpClient

Replacement for java.net.HTTPURLConnection
Supports Http/2, Websockets
in an incubator module
likely standard api in java10

### http/2

request/response based
methods are the same

binary protocol
mandatory tls
multiplex over a single tcp connection
server push capability

HttpClient
  send
    request, body handler (many available in HttpResponse.BodyHandler factory class)
  sendAsync
    CompletableFuture

  Builder for creating an instance

HttpRequest
  uri
  headers
  method
  Builder for creating instance

HttpResponse
  uri
  status
  body

## Reactive streams

Like normal streams + back pressure
back pressure: client of the stream can single to the producer that it can't keep up
  producer can back off until consumer says it can keep up again
RxJava
Flow API added to JDK
  only the interface though!
  You still need one of the implementations
    rxJava, Akka Streams, ...
Not meant as an end user API

Publisher
  onSubscribe
  onNext
  onComplete
  onError
Subscriber
  subscribe
  cancel
Subscription
  request

java.util.concurrent.Flow

HttpClient implements Pub/Sub interfaces
BodyHandler

projects announcing support
  rxjava 5
  spring 5
  akka streams

## stack walker

Before Java 9

```java

new Throwable().getStrackTrace();
Thread.getStrackTrace();

```
low perf
no guarantee all stack elements are returned
no partial handling possible

StackWalker
  walk
    pass in a lambda which is applied to the stream
    don't work on the stream directly to ensure a stable & valid stack trace
  forEach
  getInstance
StackFrame
  getMEthodName
  getDeclaringClass
  getLineNumber
  getByteCodeIndex

## desktop java enhancements

...

## performance and security enhancements
