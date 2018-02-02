# RxJava

https://realm.io/news/intro-to-rxjava/

[Observable]	emits items
	|
[Operator]	transform items
	|		   operators are chainable
[Subscriber] consume those items, can be empty

## Functional Programming

### Pure Functions

A pure function is a function which has *no side affects*.
An example side affect would be changing the state of the object you're operating on.
Which is in contradiction to Object oriented programming where objects maintain state.

Impure functions make the outcome non-deterministic if there are multiple threads.

### Functions as first class citizens

java8 Lambdas are used as first class citizens that can be passed to other functions
java8 method references help with passing functions to functions as parameters
and there are a bunch of interfaces in the function package which help.
They are all FunctionalInterface's, an interface with one methods

### Higher order functions

Functions that take functions as arguments and can return functions
the returned functions can be executed when required, Supplier which is used in withParameter in Scenarios
enables lazy execution

### Recursion

In functional programming all looping is done through recursion.
It avoids variable assignment, which is a side affect!
And facilitates lazy evaulation
As long as everytime the function is called the set of inputs is one less than the previous time.
Otherwise it could end up in an infinite loop.

## Reactive Programming

### Reactive Manifesto

#### Event Driven

Observer Pattern: Observable & Observer

RxJava: Observable<T> & Observer<T>

Threading is handled internally
Observers are passive when not notified

RxJava has a Scheduler which defines the concurrency of event processing

#### Scalable

Seamless concurrency
The Scheduler API handles this
*Is it Futures and Javas ExecutorService internally?*

A stream of events can be passed to an observer concurrently and an individual event can be processed by multiple observers concurrently

#### Resilient

Graceful Error Handling
instead of checked exceptions and error messages try to handle the error and give the user some value
maybe delayed processing of the event

Managing Failure
If one service in a cluster goes down it shouldn't bring down the whole app
Detect issue and move away from service to reduce/remove chance of knock on affect
Bulkheads
**This isn't reactive specific, it's more cloud native**

#### Responsive

Slow response is considered a failure.

Observable Data Models
MVVM: Model View View Model
When Model is updated (event) View is notified and updated

User interacts with UI (View) triggering an update to the model (event), which triggers an event which is processed in the background and when done notifies the view and it updates.

## https://gist.github.com/staltz/868e7e9bc2a7b8c1f754

Reactive programming is programming with asynchronous data streams.
Nothing completely new, think EventBus or MVC
Reactive are these on steroids
You can create a stream from anything

Functional paradigm,
	a stream outputs a new observable so there are no side affects to the input observable

Observer Design Pattern
	stream is the subject
	operator is the observer

## Asynchronous Ideals

* Explicit execution
	* Create async tasks without needing to execute them immediately.
	* This allows you to build up a picture of what you want to execute.
	* Kick it off when it's just right
* Easy thread management
* easily transformable
	* things are interdependent
	* chaining tasks together
* as few side affects as possible

In traditional programming the fundamental unit is the object.
In Reactive Programming it is the stream of events.

## RX

* explicitly started, you have to subscribe for anything to happen
* subscribeOn, observeOn for thread management
* easy to transform
* good at minimizing side affects
* single threaded by default

### Fast facts

Reactive extensions
composing async events
javascript, swift, java, ..

### Key Points

Observables

* represent async data streams

LINQ queries

* compose streams with operators

Schedulers

* manage concurrency within streams

## Observables

streams of data
pull based
create, store and pass around
hide away all the complexity with threading sync & concurrency
Immutable, operators always produce a new Observable.
*If you are creating your own parameterized type for the observable make it immutable, No setters*

### Types

#### Non-blocking Observable

Observable<T>

Asynchronous execution supported
	any thread can handle the event
can unsubscribe at any time

#### Blocking Observable

BlockingObservable<T> extends Observable<T>
 
all onNext events are synchronous

no method takes a Schedule in this type

#### Connectable Observable

Allow an observable to have multiple event subscribers and control when they start retrieving events

ConnectableObservable<T> extends Observable<T>

.publish() on an observable returns a ConnectableObservable

.subscribe(...) does not start the listener receiving events.
can be called multiple times to register multiple subscribers to the observable
.connect() needs to be called on the observable to start event receiving/emitting

If no scheduling is defined then the subscribers are executed in the order they are added as subscribers
by specifying a schedule using observeOn for each listener (connectableObservable) multiple subscribers can be called in parallel

All subscribers will receive all events even if they are running slowly

### Creating Observables

Easiest way is static factory method

Observable.from(...)
overloaded:
* Object
* Iterable
* Array
* Future

With a future we need to specify a Scheduler and schedule the futures action. Then subscribe on the Observable to see the result.

BlockingObservable.from(Observable)

Observable.concat(...)
	Create one long observable from n observables.
	All the events from the first followed by the second followed by the third, etc.

### Factory analogy

start with raw material
conveyor belts
input is shaped into a product
using multiple conveyor belts

put in raw material, shape it in many ways, and we get our output
Except in a factory the same material is passed all the way along, in RX each operator creates a new observable

Put data in and get data out.
You create an observable from some data, a single element, a list, .....
It's like streams in java8, **I wonder is it built on streams**

### Observable Lifecycle Creation

When we want to get information out of an Observable we implement an Observer interface and call subscribe on the Observable

3 methods in Observer interface

onNext
onCompleted
onError

```java
    Observable.create<String> { s ->
        s.OnNext("vnwfjk")
        s.onCompleted() }
```

Subscription interface

return value for the Observable subscribe method

one method, unsubscribe

## Subjects

It is both a Subscriber and an Observable
It can subscribe to multiple observables and reemit the events to drive new subscribers

### Publish Subjects

emits events to subscribers from the point of subscribing.
any earlier events are not emitted to the new registered subscriber
good for multicast

### Behaviour Subjects

Provides a start state and last published state which will always be the first event the subscriber gets
Allows you emit events into it and any subscribers received the onNext, onError, and onCompleted messages
It is an observable itself

### Async subject

Will only emit the last event seen before the onCompleted call.

## Schedulers

This API allows you to specify what thread your code will run on

Schedulers class provides a number of predefined Schedulers

* computation() - matches number of cores
* currentThread() - next in the current thread
* immedaite() - now, in this thread
* io() - for io. backed by a thread pool
* newThread() - new thread for each unit of work
* executor() - wrap a Java ExecutorService interface in an RxJava Scheduler Interface *This is what's used in RxScenarios* 


what do we need from a stream:

* what's the next bit of data (onNext)
* is there anything left to do (onComplete)
* did any errors happen (onError)

onError and onComplete are terminal operations in a stream
only difference is why the stream is ending

## Operators

So many
combining, filtering / reducing, ...
Stream operations! just like the java8 tutorials
RxMarbles website

### parallel

execute the function against each event in parallel.
if no scheduler is defined, the computational() one is used

### filtering operators

#### Predicates

.filter(i -> {})

#### Positional

.first() (overloaded with a predicate)
.take(<first X>)
.last() (overloaded with a predicate)
.takeLast(<last X>)

.first/lastOrDefault(...)

.elementAt(offset)
.elementAtOrDefault(offset, default)

.distinct() don't emit duplicates, only emit event one time

#### Time based

.sample(interval, time unit) take an event from the observable once every time interval
.timeout(interval, time unit) if there are no events within the interval then throw an error, call onError

## Transformations

marbles diagrams

### Mapping

#### one to one

map(...)

map the input events to a different output event

#### one to many

flatMap(...)

map the input event to multiple different output events on the same stream.
**OR**
takes events from an Observable and maps each event to a separate observable, then flattens those observables into a single observable.

An example would be recursively listing all files in a directory tree.
each event can be a file, or a directory which will be expanded to a stream of files and directories.

### Scanning

carry forward state from event to event

scan(initialStateForAccumulatedState, lambda)
the lambda has 2 arguments, the accumulated value and the next value

### GroupBy

Events can be outputted to different streams

groupBy(lambda)
the lambda

### Buffer

buffer(timespan)	group incoming events in the time span into a list and emit the list at the end of the time span

### Conditional

defaultIfEmpty("default") if the observable is empty emit the default event
skipWhile(predicate) skip all the events while the predicate is true
takeWhile(predicate) take all the events while the predicate is true
takeWhileWithIndex(predicate) take all the events while the predicate for the index is true

#### ambiguous operator

Is a race conditional operator.

amb(obs1, obs2) emits all the events from the observable which is the first to emit an event

skipUntil(obs) skips all events from the source observable until obs passed to the method emits an event.
takeUntil(obs) takes all events from the source observable until obs passed to the method emits an event.

## Schedulers

**without subscribe nothing happens**
can take 0..3 functions (onNext, onError, onComplete)
always pass onError
subscribeOn & observeOn: manage where code is run

.subscribe returns a subscription
useful for clean up

subscribeOn
should only be used once, any downstream ones will be ignored
defaults to run on thread observable was created on
you can pass a Scheduler to this method to control threads

observeOn
use as many times as needed
affects everything down stream
switch thread for the subsequent stream operations
you can pass a Scheduler to this method to control threads

**Unit is the null type**
**String? #> lik Optional return the value or null**

Dan Lew series of blog posts

## Resource Management

using() operation
allows the resources to be cleaned up when the subscriber is done

In the subscription object every resource must be registered for closing using, registerResourceForClose(...)
For example, in jdbc all Statements, ResultSets and the connection need to be registered for closing.
We then need to implement what to do on unsubscribe on our subscription implementation.
registerResourceForClose is not part of the subscription interface

using takes a Func0 and Func1 generic objects.
Func0 parameterized type is the subscription
Func1 parameterized types are the subscription and the Observable

## Exceptions

Are treated as first class citizens in reactive programming.
No point to throw an exception so it is passed as an event in the stream.

## Implementation Patterns

### Event Driven Programming

In Imperative programming services must know about the details of the services they depend on
These services are then coupled.

In Event Driven Programming we get the service to emit an event, and create a new service which subscribes to the events.
This new service then drives the service that was previously depended on,
thus decoupling the services.
The dependent service is now pulling in the events instead of the other service pushing in the calls.

### Data Store Interaction

using
combining streams

## Back pressure

It's a protocol in the reactive framework to stop the consumers getting over run with events

### Timing streams of events
You can create an Observable that emits events at a regular pace.

Observable.interval(amount, unit)

## Example Operators
Apparently a great site to see how an operator will work

### range(i, n)
emit a stream of numbers beginning at i until n

### zip(Observable, Function)
combine the elements of the source stream with the supplied stream using the function.

### flatMap
takes events from an Observale and maps each event to a single obeservable, then flattens those observables into a single observable.

### distinct()
only pass on distinct events. (think sql distinct)

### sorted()
sort the events alphabetically / numerically

### merge()
merge upto 9 streams of events together into a single thread

### debounce()
Remember 2nd electronics, make sure you only get one click of a button
in rxJava it treats all events within a specified time delay as a single event, only returning the last one.

### amb()
ambiguous
conditional operator
given x streams it selects the first stream that emits as it's input and ignore all others.


https://www.infoq.com/articles/rxjava-by-example
