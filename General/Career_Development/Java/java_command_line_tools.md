# java command line tools

https://docs.oracle.com/en/java/javase/11/tools/jdeps.html#GUID-A543FEBE-908A-49BF-996C-39499367ADB4

## jconsole

monitor a running jvm process

threads, memory, classes, mbeans

### mbeans

you can do a lot with mbeans
for example, take a dump of the running threads and heap

## jhat

analyse the heap of a dump

parses the dump file and starts a web server.

### OQL

Object Query Language, similar to SQL but on objects in the heap

examples: find all the class of a certain type.

I'm sure there's more complex stuff that could be done.
Like returning all the maps which contain a certain value???

## jps

print running jvms on the system
experimental and unsupported

-l display FDN of main class
-d display arguments to jvm
-m display parameters passed to app

## jinfo

generate java config information about a specified java process

system properties, vm flags and vm arguments

## jmap

print class related data about a jvm
requires certain arguments to be provided to the jvm on startup

## jstack

print threads of a jvm

## jdeps

Show package/class level dependencies of a class/jar
