Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-03-23T08:45:36+00:00

###### Patterns ######
Created Thursday 23 March 2017

https://dzone.com/refcardz/design-patterns

https://dzone.com/articles/design-patterns-overview

file:///home/ethomev/Downloads/w_jour01.pdf

##### Why #####
In programming all programs boil down to common problems. Creating things in a certain way, making things behave in a certain way, creating certain relationships between things.
Overtime, these common problems were identified and a set of generic solutions were created, these are design patterns.
These patterns give a consistency between software projects.
Makes software more flexible, reusable and extendable.

GoF book written in the 90's has stayed relevant right up until now.

3 categories
	1. Creational
	2. Structural
	3. Behavioural

Describing a pattern, according to GoF
	name
	problem
	solution
	consequence

purpose
	what category does it fit into
scope


#### Scope ####
Does the pattern deal with classes at compile time or objects at runtime

### Classes at compile time ###
relationships between classes and their sub classes
established through inheritance, so they are static/fixed at compile time
Creational class patterns defer some part of creation to subclasses

### Objects at runtime ###
relationships between objects and their delegates
established through composition at runtime
Creational object patterns defer some part to another object

Only one pattern can be implemented in both scopes, Adapter


##### My Goal #####
My goal is to understand all the patterns in the GoF book. Specifically I want to 
1. understand what each one does
2. how it does it
3. how to implement it
4. when to implement it.
5. The pros & cons of each one
6. know the indicators for using a pattern

#### Break it down ####

The objective of patterns
The different types of patterns
Each of the points above per pattern
