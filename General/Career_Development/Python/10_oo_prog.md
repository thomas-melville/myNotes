# OO Programming

## 14. Object Oriented Design in Python

**Python supports multiple inheritance.**

```python
class Student: # this is the old way, don't do this anymore
	pass

class Student(object): # this is the new way to do it. You get some new features this way
	pass

#...
student = Student()
```

pass is a keyword in the interpreter which tells python to do nothing.

constructors are just special methods

```python
class Student:
	def __init__(self, name, age):
		self.__name = name
		self.__age = age

	# this is a method override
	def __str__(self):
		return "..."

```

double underscores signifies special methods in python

self refers to the instance of the class.
is this the same as this in java? **Yes**
In Java this is implicit whereas in Python you have to explicitly specify it
self is a parameter of the function but not passed to it. It is inferred by python at runtime

```python

# you type
my_class.my_method('hello')
# python interprets it as
MyClass.my_method(my_class, 'hello')

```

#### instance comparison

Python uses the same OO techniques as java,

```python

	def __hash__(self):
		# ...

	def __eq__(self, other):
		# ...
```

### dynamic typing

A variable of a type can be reassigned to a subtype.
The variable will then have the behaviour of the sub type.
This is much more dynamic than java

### dynamic objects

it is possible to add attributes and methods to an object at runtime
since attributes are stored in a dictionary it is easy to add a new entry to the dictionary

One caveat with adding a new method
it needs to be called on the object and the object passed as the first parameter

```python

class Account(object):
	pass

a1 = Account();
a1.acc_number = 1234
a1.show_number = lambda acc: acc.number

a1.show_number(a1)

```

### methods to override

#### object equality

a class must implement the __eq__ method to override the default object equality behaviour with value equality
synonymous with equals() in Java

#### string representation

override __str__ to give a meaningful string represntation of the class
synonymous with toString() in Java

### instance and class attributes/fields

instance fields are defined in methods using the self instance
class (static) fields are defined outside any method

```python
class Student:
	school = "Annaghdown" # class field

	def __init__(self, name, id):
		self.name = name # instance field
		self.school = Student.school
		#...
```

instance fields of a class become a dictionary which can be accessed and passed across the network

When accessing a class field python doesn't look in the class scope so you have to qualify a class/static field with the class name

**There are no access modifiers in python.**
Back to its philosophy, you should know what you are doing and not do stupid shit!
If you want to hide fields of a class prepend them with '\__'
Pycharm won't see it but if you know it you can use it

### inheritance

``` python
class HighSchoolStudent(Student):
# HighSchoolStudent now inherits from Student
```

super keyword is used in python also
python has no access modifiers, everything is public
the convention is to use __ to denote that a method is private / should not be overridden

to gain access to methods in super type you use the class name. It's the same for constructors and overridden methods

```python

#...
	def __init__(self, ...):
		MySuper.__init__(self, ...)

```

**When using multiple inheritance you must invoke all super constructors explicitly.**
