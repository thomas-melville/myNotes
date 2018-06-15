# Overview

* **Object-oriented** and **Functional** programming language
* All the normal features
* Dictionaries ~> json
* Supports exceptions
* pass by reference
* Dynamic language, in two ways. types, and objects
* Strong and dynamically typed. Variable creation doesn't require a type, but once assigned if you try and do something specific to another type you'll get an exception
* * You can reassign the variable to a new type without any issues. Python only complains when you try to do an operation which is not supported by the current type.
* * Because it's dynamically typed it's not clear at compile time if there will be an issue with the operations on a variable.
* * * Java being completely strongly typed will fail compilation if an illegal operation is attempted.
* An object can be updated at runtime to include new attributes and

**Multithreading in Python is still weak compared to Java.**

Python cites it's development speed as one of it's main advantages
It's processing speed is slower than other languages like C
A lot of the libraries Python calls are written in C
So if processing time is critical you can write your critical code in C and using some thing, I forget the name, you can call it from your python module

## Fundamentals of the language

### indentation

indentation is very important, 4 spaces! every time
this is used for defining blocks of logic, instead of {}
indentation does not have to be consistent across a file, only within a block
but it is recommended to be consistent
It can be any number but the recommendation is 4 spaces

### Getting help

dir(<object>) will print the available functions and objects of an object
dir() on it's own prints default available stuff
__builtins__ are the built in functions

help(<object>.<func>) will print the documentation about a function

without the brackets python will print what an object / function is

### Conventions

__xxx__ implies the function is special

Classes and Exceptions are PascallCase
	MyClass
Modules, functions and variable are snake case
	my_class

### Gotchas

if you give a variable the same name as a built in function, python won't complain.
It will silently reassign the builtin function identifier to the variable
and you'll get an error when you try to call methods on it.
Moral of the story, don't give anything a name which is a builtin function / type

**Everything in python is an object!!!**

x + 5 translates to x.__add__(5)

There is no such thing as a void return method in python
Something is always returned, if you don't return something None is returned

When accessing static/class fields from within a method you have to qualify it with the class name as Python skips the class when searching for the variable.
It looks in the method and then in the global space, not like Java which looks in the method, then class then global (I don't think there is a global in java)

No concept of method overloading in python
The closest you can get is by using named arguments and default values

### Commenting

/ for single line
''' for multi line (3 single quotes)

' or " can be used, there is no difference.

## Documenting

First unassigned string in a function / module becomes the __doc__ property
when the help function is called the __doc__ property is returned.
Make it a multi line string for readability

## Variables

Type of variable figured out when it is assigned a value
no type is given

```python
name = "Tom"
age = 33
```

casting exists, similar to java

```python
int(pi) = 3
```

Everything is an object in Python, even ints!
Python stores a cache of integer objects on the heap, from -5 to 257

**All objects are immutable in Python.**
a = 3
a = 5
This creates a new object, 5, and points a's reference to it

## Type hinting

```python
def add(a: int, b: int) -> int:
    return a + b
```

helps reduce runtime errors with wrong types
it's not a compile time check though