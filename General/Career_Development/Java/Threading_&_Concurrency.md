Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2015-08-12T19:37:42+01:00

###### Threading & Concurrency ######
Created Wednesday 12 August 2015

A thread in java is a thread of execution in a program / light weight process
JVM allows an application to have multiple threads running concurrently
Threads have priority
Threads can also be marked as daemon
JVM starts with a single daemon thread - which typically calls main
extends Thread / implements Runnable

Making code thread safe.
	synchronized keyword on class, method or block of code
	blocks are guarded by a key (object you specify)
		String or object
		This is called a lock
		use the same lock for multiple blocks to make only one thread access any of them at any one time
			What does synchronized on a method use as a lock
				the instance of the object
				no 2 synchronized methods execution can interleave
	When a sync'd method exits it creates a "happens-before" relationship so that any waiting threads has all the latest changes to the object.
	Don't do any work can should be synchronized in the constructor
		the constructor can't be synchronized, it doesn't make sense

Java Memory Model
	Atomic operations
		single unit of work without possibility of interference from any other threads
		i++ is not atomic, it's a read followed by a write

Interrupts
	an indication that a thread that it should stop what it's doing and do something else.
	
