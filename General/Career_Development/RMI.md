Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-03-08T08:41:44+00:00

###### RMI ######
Created Wednesday 08 March 2017

Remote Method Invocation

Make calls from one JVM to another, client server architecture.

protocol underlying it is Java Remote Method Protocol (JRMP)

There is also a CORBA version of RMI (RMI over IIOP)

The basic ideas come from the "network objects" feature of Modula-3

#### Implementation ####

Client <-> Stub <-> Network <-> Skeleton <-> Server

### Server ###

Create some remote objects
Make references to these objects accessible
wait for clients to invoke methods on these objects.

### Client ###

Obtain a remote reference to one or more objects
invoke methods on them.

### Getting references ###

1. RMI Registry 
2. JNDI
3. jgroups????
4. jboss?

Idenitifiers/names are "bound" to remote objects.
