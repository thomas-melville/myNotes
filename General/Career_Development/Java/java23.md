# java 23

## JEP 455 - primitive types in patterns, instanceof and switch

Still in preview

Does what it says on the tin, can use primitives in patterns, instanceof and switch.

## JEP 476 - Module import declarations

Instead of importing individual packages, can now import a whole module

## JEP 477 - Implicitly declared classes and instance main methods

Simple programs now have much less boiler plate code to get up and running

## JEP 482 - Flexible constructor bodies

Allow certain statements before super/this in constructors.
Enable the likes of validation of inputs before intiatializing the whole hierarchy tree.

## JEP 466 - Class-file API

An API for parsing, generating and transforming Java class files.
Will ease the tension between frameworks which use relfection to parse and transform class files.

## JEP 469 - Vector API

... 8th preview

## JEP 473 - Stream gatherers

Enhances the API to support custom intermediate operations. 
Allow stream pipelines to transform data in ways not easily achievalbe with the existing built in intermediate operations.
Provides as much flexibility to intermediate as the the collectros API does to terminal operations.

## JEP 480 - Structured Concurrency

Allows developers to treat a group of related tasks running in different threads as a single unit of work.
Thereby streamlining error handling and cancellation, improving reliability and enhancing observability.

## JEP 481 - Scoped Values

Share immutable data both with calles within thread and child threads.
Easier to reason about than thread local variables.
Lower space and time costs, especially when used together with virtual threads.

## JEP 474 - ZGC Generationl mode by default

Now generational mode by default. Feedback says it performs better than non-generational.
Non-generational is not marked as deprecated, for removal.

## JEP 467 - Markdown document comments

Now it's possible to use markdown in Javadoc comments

## JEP 471 - Deprecreate methods in Unsafe for removal

A few remaining internal APIs are still available because they were so heavily in use.
They are in heavy use because they are the only way to complete these tasks.
JEPs are in progress to provide official APIs for these tasks.
With the release of the Foreign function and Memory JEP, the corresponding methods in Unsafe are now deprecated and marked for removal.