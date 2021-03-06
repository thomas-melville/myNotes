Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2017-02-24T13:21:56+00:00

###### Algorithms ######
Created Friday 24 February 2017

https://www.coursera.org/learn/algorithms-part1/lecture/buZPh/course-introduction

##### Introduction #####
Book site: http://algs4.cs.princeton.edu/home/

Algorithm: method for solving a problem
Data Structure: method to store information

### Scientific method: ###
Observer
Hypothesize
Predict
Verify
Validate

##### Union-Find #####

developing an algorithm
	model the problem
	find an algorithm to solve it
	fast enough?
	if not, why?
	find a way to address the problem
	iterate until satisfied
	
Dynamic connectivity problem: How to know what objects are connect or not

Union command: connect 2 objects
Find/connected query: is there a path connecting the two objects

Applications:
	computers in networks
	fiends in a social network
	....
	

#### Modeling connections ####

"is connected to" is an equivalence relation:
	Reflexive
	Symmetric
	Transitive
	
Connected components
	maximal set of objects that are mutually connected.
	arrays which contain all the objects which are said to be connected as per the equivalence relation above.
	
Implement operations

	Find query
	union command

Solution for the dynamic connectivity problem	

##### Quick Find #####
It's an eager approach

#### Data Structure ####
integer array of size N
Interpretation: p & q are connect if, and only if, their entry in the array have the same id.

#### Find ####
Is quick and easy, do the 2 ojects have the same id?

#### Union ####
Is a bit more difficult.
if the entry at p already has an id, it and all the other entries that have this id must be changed to the entry at q.
The array is initialized with each entry being the index of it's location so that no objects are connected.

#### Costs & Effectiveness ####
find is very cheap
union is expensive and grows exponentially as the array and connections get bigger

### Cost model ###

N # the size of the array
algorithm	initialize	union	find
quick-find	N			N		1

algorithm takes quadratic time (increase by a factor of 4) and squared time (increase by a factor of 2).

**Quadraric time is much too slow. We can't accept quadratic time algorithms for large problems, because they don't scale**

##### Quick Union #####
It's a lazy approach

#### Data Structure ####
integer array of size N
interpretation: p & q are connect if, and only if the they share the same root.
It's a tree approach. Each entry in the array contains the reference to it's parent

#### Find ####
Check if p & q have the same root.

#### Union ####
set the id of p's root to the id of q's root.
Only one value in the array needs to change

#### Cost & Effectiveness ####
trees can get tall. Can be a linear tree which has the same size as the array.
find too expensive, can be array accesses upto the size of the array

N # the size of the array
algorithm		initialize	union	find
quick-union	N			N		N

##### Quick Union Improvements #####

#### Weighting ####
Keep track of tree size.
when performing union link smaller tree directly to root of other tree.

Keeps average distance of object from root much smaller.

### Cost & Effectiveness ###

Tree never gets taller than lgN, where N is the size of the array.
100 #> 10, ...
algorithm	initialize	union 	find
weighted	N			lg N	lg N

#### Path Compression ####
When doing a union, point each examined node directly to the root. All nodes in the path.

lg* N (the number of times you have to call to get to 1)
algorithm is linear, in practice but not in theory.
(No algorithm is linear in theory for the union find problem)

##### Union Find Applications #####

Percolation, games, dynamic connectivity, image processing, least common ancestor

#### Percolation ####

N*N site, a square is either black or white.
It has the probability of p of being white, 1 - p of being black
The site is percolated if there is a path from top to bottom in white.
Model for many physical systems:
	electricity	social interaction	fluid flow
	material	population		material
	conductor	person			empty
	insulated	empty				blocked
	conducts	communicates	porous

How do we know whether it's going to percolate or not?
	threshold between yes and no is very sharp

Monte Carlo Simulation.
Generate site that is blocked
randomly open sites
until grid is open

object for each site.
connect the objects if sites are connected

going top to bottom is quadratic.
Create 2 virtual sites, from bottom and top

open a site, union it with all open sites around it.
