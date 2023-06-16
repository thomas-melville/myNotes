# types

### Numbers

Integers (int), Floats (float) and Longs (long).
floats are decimal.
floats have are imprecise, at about the 16th decimal point
if you required precision, use the decimal library

you can seamlessly add integers and floats and python won't complain.

if an int becomes too big it is dynamically changed to a float

### String

ascii text in python2
unicode text in python3

you can use ', " or """ to define strings
no difference
put r in front of the string to tell python to treat it as a raw string
i.e. don't try an interpolate any special characters

type(...)
String is a list of characters in Python and can be treated that way
can use array type indexing to get ranges of characters from strings: a[2:5]
	use negative values to index from end of string
There is no such thing as a character type in Python, only Strings of length 1

String interpolation introduced in python3.6
f"Hello there {name}"

there are a lot of useful utility methods you can call on a string

### printing

.format() is one of them.
think of logging using slf4j. "Hello there {0}".format("Dan")

Another way to print is to use printf syntax

print '%s is aged %d' % name, age
	%i,f,... they're all supported
	and alignment, size, decimal point in float

print is a "varargs" function, each parameter is printed with a space between it and the next one

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

.append() to add to the end of a list as a single element.
	This means that if you try an append a list of elements, it will be appended to the list as a single element of type list
.extend() to add to the end of a list as individual elements
del list[2] remove an entry from the list and reduce the size
len(list) will give the length

Not arrays, python lists
List element types can vary, not recommended though

you can "enumerate" over a list, this will give you a tuple of index & value

+ = concatentate two lists together

.remove() only removes the first occurrence

#### comparing lists

<,==,>

Compares each element with the corresponding element in the other list
numeric and alphabetic comparison.
This would be very useful for sorting in something like a phone book
each entry is a list where each index is the following:
[last name, first name, country, parish, village, street, house number]

#### list slicing

syntax: [x:y]

return a list which is a subset of the list based on x and y
x or y can be omitted and it will index from the start/end
negative indexes can be used also to index from the end

you can take a slice and assign a new list to it.
It can be <, ==, or > the existing slice

#### list comprehension

very powerful mechanism for generating lists
generate a new list by applying a function to elements in a source list, tuple or range

```python
def power(n, power):
    return n**power

squares = [power(n, 2) for n in range(20)]
```
For each element in the range 0 to 20 call the function and push the result into a new list

### dictionaries

similar to json
It is a hash table

```python
student = {
	"name": "Mark",
	"student_id": 432,
}

# another way to define a dictionary
student = dict(name:"Mark",student_id: 432)

name = student["name"]
name = student.get("name", "Unknown") # Unknown is the default value

student.keys()
student.values()
```

KeyError is thrown if the key doesn't exist
keys and values can be of any type, including dictionaries!
nested dictionaries

order of elements is sorted by python, not the way you enter them.
there is a sorted_dict

Advice from instructor:
"If you are writing code that does searching, ask the question would a dictionary be better?"
The keys are hashed which gives efficient lookup of items.

### tuples

similar to lists but are immutable
element types can be different

when creating a tuple with one element you must do the following:

```python

my_tuple = (4,)

```

there is a namedtuple object

```python

from collections import namedtuple

task = namedtuple('Task', ['summary', 'owner', 'done', 'id'])

```

### set

same as a set in java, can only have unique values
