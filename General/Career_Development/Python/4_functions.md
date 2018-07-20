# functions

syntax:

```python
def func_name(arguments):
	# indent 4 spaces
	return # this is optional, functions can be void just like java/javascript
```

### optional

assign the argument a value in the method definition
the optional arguments must be last

### named arguments

when calling a method you assign the value to the parameter name

```python
def func(name, id):

func(name="Tom", id=42)
```

you can change the order if you want

### variable number of arguments

python supports it too

```python
def func(*args):
	# args is a single list of all the parameters to the method
```

if you pass a list it will become the first element in the args list in the function.
To expand the list into it's individual elements and into the args list precede it with a *

```python
numbers = [4,3,5,2,2]
print average(*numbers)
```

### keyword arguments

access all the parameters as a dictionary not a list

```python
def func(**kwargs):
	kwargs["desc"]

func(desc="Loves Python", age=43)
```

** tells python to use a dict

### nested functions

you can nest functions in other functions
delegate a task within a function without polluting the namespace of the functions enclosing file
nested function has access to variables that are defined in the outer function
this way you could create a truly private method in Python, will only work if the private method is only used by one method!

## Variable scope

Python supports local & global scope