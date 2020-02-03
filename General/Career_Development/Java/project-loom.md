# project loom

https://openjdk.java.net/projects/loom/
https://www.javaadvent.com/2019/12/project-loom.html

For asynchronous non-blocking code.
Run tasks as fibers (lightweight/virtual threads)

A thread as we know it can have multiple fibers.
Just like a thread is blocked waiting on IO so another thread can take the OS process.
When a fiber is blocked waiting for IO another fiber can take the thread.

This is beneficial to tasks which are IO intensive, computational intensive is no advantage.

This very much exploratory.

Currently represented as an object of the Thread class, created with a Thread.VIRTUAL argument.
You can use a ThreadFactory either with the virtual builder method, and pass this to the ExecutorService factory method.

From an API point of view it's the same as before. Except now virtual threads, or fibers are being scheduled.
Submit Callables and Runnables just like before.
