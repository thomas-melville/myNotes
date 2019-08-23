# CompletableFuture

https://www.baeldung.com/java-completablefuture
https://mahmoudanouti.wordpress.com/2018/01/26/20-examples-of-using-javas-completablefuture/

CompletableFuture implements the Future and CompletionStage interfaces
A Future is a result of a callable which will be present in the future (the result of an async computation)
CompletionStage defines the contract for an async computation step that can be combined with other steps.
  It represents a stage of a certain computation which can be done either synchronously or asynchronously.

CompletableFuture is a building block and a framework, 50 different methods.

I'm working off Java8 but there have been lots of improvements in java9

## use cases

1. as a simple Future

Create an instance of it, hand it out to consumers and call complete on it at some other stage
The consumers can call the get method and block until the result is there.

2. execute some code asynchronously

*runAsync* and *supplyAsync* allow us to create a CompletableFuture instance out of a Runnable and Supplier correspondingly
The most generic way to process the result of a computation is to feed it to a function, for example *thenApply*.
If you don't need a return value use *thenAccept*, which takes a Consumer.
If you don't need to process the result or return a value then you can use *thenRun*

```java

CompletableFuture<String> cf = CompletableFuture.runAsync(() -> {
  //...
});

// Sync
cf.thenApply(s -> {
  return s.toUpperCase();
});

// Async
CompletableFuture<String> async = cf.thenApplyAsync(s -> {
  return s.toUpperCase();
});

async.join();
```

3. combine into a chain of steps

The API allows you to chain multiple CompletableFuture instances together, the result of which is a CompletableFuture.
This approach is ubiquitous in Functional Languages, and is often referred to as a monadic design pattern.
*thenCompose* method takes a function which returns a CompletableFuture instance. The argument is the result of the previous function.
*thenCombine* allows you to execute two independent CompletableFutures and do something with their results and pass it down the chain.
  This works for both synchronous and asynchronous.
*thenAcceptBoth* allows you to execute two independent CompletableFutures, but don't need to pass it down the chain

the difference between thenApply and thenCompose is analogous to the difference between map & flatMap
thenApply passes the result of the stage along.
thenCompose passes on the completionStage.

4. execute multiple futures in parallel.

Use the static method *CompletableFuture.allOf*
Does not return the combined result of all the futures.
You can put them into a stream and use *CompletableFuture.join*

5. create a completed Future.

```java
CompletableFuture<String> cf = CompletableFuture.completedFuture("vnrjek");
cf.isDone();
cf.getNow(null); // returns the result if completed or null

```

6. wait for multiple stages to complete

This works for both synchronous and asynchronous.
*whenComplete*

```java

CompletableFuture.allOf(futures.toArray(new CompletableFuture[futures.size()])).whenComplete((v, th) -> {
        futures.forEach(cf -> assertTrue(isUpperCase(cf.getNow(null))));
        result.append("done");
    });
```

## Exceptions

ExecutionException: encapsulates an exception that occured during a computation
The try/catch block does not fit so there is a separate method which you call *handle*
It's a BiFunction which takes the type of the Future and the Throwable
There's also the *completeExceptionally* method which returns an ExecutionException which has the cause you throw

## Async methods

Most methods in the API have two variants with Async postfix.
Used to run the step in another thread.
The methods without the postfix run in the calling thread.
The methods with the postfix run in one of two thread pools.
1. The common ForkJoinPool
2. The thread pool created by the executor if it is passed.

## Gotchas

https://4comprehension.com/completablefuture-the-difference-between-thenapply-thenapplyasync/

thenApply/thenApplyAsync
  both can be used to execute a callback after the source CompletableFuture completes.
  both return new CompletableFuture instances
  and seem to run asynchronously

  in thenApply the callback is executed on the callers thread.

  in thenApplyAsync the callback is executed on a thread in the ForkJoinPool.commonPool, the CompletableFuture#defaultExecutor
  which is not the pool given to the original CompletableFuture
  solution is to supply the ExexcutorService to each CompletableFuture call
