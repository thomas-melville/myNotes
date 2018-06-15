# exceptions

Same concept as java exceptions.
with a similar mechanism to handle them.

Exception handling is optimistic approach to exceptions
low-level code / libraries can detect exceptions but shouldn't handle them as the context would be different for each user of the low-level code.
Some may want to propagate it, just log it, raise their own exception

```python
try:
    # code that may throw an exception
except KeyError as error:
    # when an exception of type KeyError is encountered handle it in this block.
finally:
	# just like java
else:
	# This block is only executed if no exception is raised
```

else is new for me! Could be useful for logging

Exceptions can be chained by adding more except blocks

to catch all use "Exception"

you can get more information by putting the exception into a variable and printing it

**raise** is Pythons word for throw from Java

It is possible to wrap exceptions in Python also, the technique is different though

### auto closing resources

pythons "with" keyword is synonymous with javas try-with-resources

```python
with open('file','r') as f:
	#....
```

this code snippet automatically closes the file regardless of an exception occurring or not.

f is a ContextManager, which has a hook which is called when the variable goes out of scope.