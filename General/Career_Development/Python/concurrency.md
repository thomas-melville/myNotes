# concurrency

Threads use the same memory heap so they can write to the same location in memory.
That is why the GIL - Global Interpreter Lock in CPython was created.
Some Python implementations don't have a GIL: IronPython, JPython
I'm nearly certain I use CPython, will this hinder our concurrency efforts?

https://www.tutorialspoint.com/python/python_multithreading.htm

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

only one thread will be executing at a time due to GIL, the code is concurrent but not parallel.
It can still be faster than sequential if it is an IO task, the majority of time is spent waiting for the network.

## map & pool from multiprocessing

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