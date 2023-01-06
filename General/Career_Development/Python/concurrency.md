# concurrency

Threads use the same memory heap so they can write to the same location in memory.
That is why the GIL - Global Interpreter Lock in CPython was created.
Some Python implementations don't have a GIL: IronPython, JPython
I'm nearly certain I use CPython, will this hinder our concurrency efforts?

https://www.tutorialspoint.com/python/python_multithreading.htm

multi threading runs one instance of python on one process, in one core.
multi processing runs an instance of python for each process on a separate core.

## thread lifecycle
                          running
                            yield
new Thread --- start ---> Runnable  <---> Not Runnable
                              |
                              | the run method terminates
                              |
                            Dead

Not Runnable
  The OS needs resources so asks Python to free up resources
  Python puts the thread into the Not Runnable state
  our code can't


## thread module

    very low level

## threading module

    at a slightly higher level of abstraction but no where near Java.
    builds on thread module
    has the Thread class that implements threading

writing your own Thread class

1. subclass Thread
2. override __init__ to add additional arguments
3. override run to implement what the thread should do

Create an instance and call the start method

To block and wait for all threads to finish iterate through them and call the join method

http://chriskiehl.com/article/parallelism-in-one-line/

### GIL affect

Only from accessing CPython code. For example import python code.
only one thread will be executing at a time due to GIL, the code is concurrent but not parallel.
It can still be faster than sequential if it is an IO task, the majority of time is spent waiting for the network.

## multiprocessing

unlike multi threading, each process has its OWN copy of Python!
Therefore there is no GIL degradation

Hard/Impossible to have one process interact with another process.
They are completely separate instances of Python.

### map & pool from multiprocessing

The multiprocessing module builds on top of the threading module.

map, taken from functional. Map is a function which maps a function on to a list of objects and returns each output in a list
versions of this function which can be used with a thread pool are available in two modules:

* multiprocessing
* * works on processes
* * better for cpu intensive tasks due to GIL
* * Overhead of the entire program being copied into the memory heap for each cpu
* multiprocessing.dummy
* * works on threads
* * better for io tasks due to GIL

1. import modules
2. create thread pool (int parameter sets number of workers in pool, for io tasks experiment)
3. call map function
4. close the pool
5. call join to wait for all threads to finish

**You can call the pool as a Context Mgr so there's no need to call join**

```python

with Pool() as p:
    p.map(func, array)

# done
```

each thread will run in the same python instance
each process can run on separate python instances

python -o argument when running a module.
  optimizes execution for threads

## locks and rlocks

Python has locks and rentrant locks

locks can be used in "with" statements so that the lock is automatically released

## semaphores

When more than one thread can access to the resource at one time, however there is a limit to the number of threads that can access it concurrently.
When creating a semaphore you specify the number of concurrent threads allowed

## async, await and coroutines

the word coroutine has come to be overloaded in python.
conventional coroutines
coroutine is a decorator
in python generators are coroutines
they can yield more than one value

modern python lets us manage asynchronous code using 'async' and 'await'
if we async we must use await or something similar
