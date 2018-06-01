# Python

## Overview

Object-oriented programming language
all the normal features
dictionaries ~> json
supports exceptions
strong and dynamically typed. Variable creation doesn't require a type, but once assigned if you try and do something specific to another type you'll get an exception

Because it's dynamically typed it's not clear at compile time if there will be an issue with the operations on a variable, Java being completely strongly typed will fail compilation if an illegal operation is attempted.

### Fundamentals of language

indentation is very important, 4 spaces! every time
this is used for defining blocks of logic, instead of {}

### Commenting

/# for single line
''' for multi line (3 single quotes)

no concept of "pythondoc" so convention is that method documentation is a multi line comment inside the method body

### Variables

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

### Type hinting

```python
def add(a: int, b: int) -> int:
    return a + b
```

helps reduce runtime errors with wrong types
it's not a compile time check though

## Operators

### Arithmetic

 * multiplication
 ** to the power of
 / division, include remainder
 // division, floor (exclude remainder / rounds down)

## Conditionals

if, elif, else

syntax: if condition:
		first line of block indented 4 spaces, this is the convention, you can use 2/3/5 but the standard is 4
		else:

use == for equality check
use in to check for existence of a value in a list

you can use not in place of !
you can use the words 'and' and 'or' for boolean algebra

python supports ternary statements
syntax is different though!

```python
"bigger" if a > b else "smaller"
```

python has the truthy/falsey concept, just like javascript

number other than zero is true
non empty string is true
non empty list is true

## Functions

syntax:

```python
def func_name(arguments):
	# indent 4 spaces
	return # this is optional, functions can be void just like java/javascript
```

### optional

assign the argument a value in the method definition

### named arguments

when calling a method you assign the value to the parameter name

```python
def func(name, id):

func(name="Tom", id=42)
```

### variable number of arguments

python supports it too

```python
def func(*args):
	# args is a single list of all the parameters to the method
```

### keyword arguments

access all the parameters as a dictionary not a list

```python
def func(**kwargs):
	kwargs["desc"]

func(desc="Loves Python", age=43)
```

### nested functions

you can nest functions in other functions
delegate a task within a function without polluting the namespace of the functions enclosing file
nested function has access to variables that are defined in the outer function
this way you could create a truly private method in Python, will only work if the private method is only used by one method!

## Variable scope

Python supports local & global scope

## Loops

range method, upper limit is less than

```python
range(10)
```

give me a list of integers starting at 0 and stopping before 10
return a list of 10 ints 0 - 9

```python
range(5, 10)
```

give me a list of integers starting at 5 and stopping before 10
return a list of 5 ints 5 - 9

```python
range(5, 10, 2)
```

give me a list of integers starting at 5, stopping before 10 and incrementing by 2
return a list of 5, 7, 9

syntax:

```python
for a in <list>:
	...(a)

while a < condition:
	....(a)
```

### break & continue

the same as other languages

## Types

### Numbers

Integers (int) and Floats (float).
floats are decimal.

you can seamlessly add integers and floats and python won't complain.

### String

ascii text in python2
unicode text in python3

you can use ', " or """ to define strings
no difference

type(...)
String is a list of characters in Python and can be treated that way
can use array type indexing to get ranges of characters from strings: a[2:5]
	use negative values to index from end of string

there are a lot of useful utility methods you can call on a string

.format() is one of them.
think of logging using slf4j. "Hello there {0}".format("Dan")

String interpolation introduced in python3.6
f"Hello there {name}"

### Boolean

True/False
first letter is always capital
conversion to an int will yield 0/1
a string converted to a boolean will evaluate to textual representation

### None

similar to null
variable has been defined, but not assigned yet.

### Lists

```python
names = ["Tom", "Dick", "Harry"]
```

indexing is similar to lists in other languages
With an addition
by using a negative number you can index from the end of the list
last element starts at -1

.append() to add to a list
del list[2] remove an entry from the list and reduce the size
len(list) will give the length

Not arrays, python lists
List element types can vary, not recommended though

#### list slicing

syntax: [x:y]

return a list which is a subset of the list based on x and y
x or y can be omitted and it will index from the start/end
negative indexes can be used also to index from the end

### dictionaries

similar to json

```python
student = {
	"name": "Mark",
	"student_id": 432,
}

name = student["name"]
name = student.get("name", "Unknown") # Unknown is the default value

student.keys()
student.values()
```

KeyError is thrown if the key doesn't exist
keys and values can be of any type, including dictionaries!
nested dictionaries

### tuples

similar to lists but are immutable
element types can be different

### set

same as a set in java, can only have unique values

## Exceptions

Same concept as java exceptions.
with a similar mechanism to handle them.

```python
try:
    # code that may throw an exception
except KeyError as error:
    # when an exception of type KeyError is encountered handle it in this block.
finally:
	# just like java
```

Exceptions can be chained by adding more except blocks

to catch all use "Exception"

you can get more information by putting the exception into a variable and printing it

**raise** is Pythons word for throw from Java

## working with files

built in method open("filename", "access modes")
returns a file object

access modes

* a	append
* r	read
* w	write, will overwrite existing data
* b	open the file in binary mode, Windows only! but doesn't hurt when your developing on linux, makes the module platform independent

.write(...)
.close()

always remember to close files.
files have locks and if it's already open it won't be possible to open it again

## yield

key word
creates a generator function, which can be used as an iterator.
yield is used in place of return

## lambda functions

anonymous functions, same as I've seen in other languages

```python
result = lambda x: x * 2
```

x is the value to pass in
x * 2 is the operation

called just like a normal function

higher order functions

## pipenv

dependency manager for Python projects
similar to npm

manages dependencies on a per project basis

```bash
cd project
pipenv install <package>
```

this installs the package and creates/updates a pipfile
others can use this pipfile to install all the required dependencies

run the module with pipenv too!

pipenv creates a virtualenv for the module

to use the virtualenv execute "pipenv shell"

Looks to be more for creating the environment locally from a setup.py file

## virtualenv

create independent working environments for each application with specific versions of required modules.
create a virtual environment with the required modules installed to run your python script
lower level than pipenv

install with pip

create a virtual environment:

```bash
virtualenv <env_name>
```

you can specify python version using --python=

a virtual environment needs to be activated

source <path_to_created_venv>/bin/activate

at this point installed packages will only be installed to this virtualenv

deactivate to leave the virtualenv

```bash
pip freeze > requirements.txt
```
this generates your requirements.txt file with all the packages installed in the virtual env

then you can do

```bash
pip install -r requirements.txt on a new environment
```

## Object Oriented Design in Python

```python
class Student:
	pass

#...
student = Student()
```

pass is a keyword in the interpreter which tells python to do nothing.

constructors are just special methods

```python
class Student:
	def __init__(self, name, age):

	# this is a method override
	def __str__(self):
		return "..."

```

double underscores signifies special methods in python

self is another keyword in python, it refers to the instance of the class.
is this the same as this in java?
the keyword is a parameter of the function but not passed to it. It is inferred by python at runtime

### instance and class attributes/fields

instance fields are defined in methods using the self instance
class fields are defined outside any method

```python
class Student:
	school = "Annaghdown" # class field

	def __init__(self, name, id):
		self.name = name # instance field
		#...
```

instance fields of a class become a dictionary which can be accessed and passed across the network

### inheritance

``` python
class HighSchoolStudent(Student):
# HighSchoolStudent now inherits from Student
```

super keyword is used in python also
python has no access modifiers, everything is public
the convention is to use __ to denote that a method is private / should not be overridden

## modules

```python
import <filename_minus_extension>
from <filename_minus_extension> import <class_name>
```

/* is used as a wildcard for imports in python too

when defining a module the folder must contain an __init__.py file

the root of the module should define a requirements.txt file to specify what versions of other modules it requires.
This can lead to issues if you have multiple dependencies who depend on different versions of a dependency that don't overlap

## developing web apps in python

django is one library
flask is another
sanic is yet another, python 3.5+ only!

django tries to solve everything for you
flask is a minimalist library
	seven lines of code to serve a page!

## installing python packages

pip, package manager for python
package index at pypi.python.org

## creating an executable file

you can package the interpreter with the code

pyinstaller is the tool, install it using pip

pyinstaller main.py
and the executable file is created
--onefile flag packages everything required into one file

## creating a setup file

the tool he uses is only on windows