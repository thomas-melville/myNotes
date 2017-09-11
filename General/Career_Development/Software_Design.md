Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2016-08-10T11:35:40+01:00

###### Software Design ######
Created Wednesday 10 August 2016

##### Thoughts on patterns #####

### From my usage of Decorator & Observer they could be used for the same problem? ###

* They seem to both add new functionality to an existing API.
* A difference is that the class being decorated knows nothing about it's decorator but the Subject must know about it's Observer.

### When to use Strategy or Template ###

Template seems to be able to handle commonality between algorithms whereas Strategy looks to be all or nothing.
