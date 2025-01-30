# design patterns

https://www.geeksforgeeks.org/python-design-patterns/

https://python-patterns.guide/

https://github.com/onionmccabbage/furtherAdvancedPythonSept22

abc module - abstract base class
  adds a bit of self discipline to your code.
  Python figures it out, but by adding things from abc you will catch possible issues

@abstractmethod is a decorator in python

class Payment(metaclass = ABCMeta) - this will ensure that any implementing class have to implement any abstract methods


eval function. Give it a reference to a class and then you can instantiate it.

```python

from abc import ABCMeta, abstractmethod

class AbstractShape(metaclass=ABCMeta):
  # abstract class in Python. Declare properties and methods that we will need to implement in concrete classes.

  @property
  @abstractmethod
  def num_sides(self, num_sides):
    pass

```


Factory
  creates things
Facade
  groups related things together
  declare them as abstract non-inter-dependent classes
State(ful)
proxy
command
observable
  popular when there are large amounts of data in quick succession
  publish-subscribe model
  we create an observable stream of data then one or more subscribers can observe the data
  this way the observable is completely decoupled from all subscribers

## decorators

Interesting way to apply decorators in Python

Write a function which will be the decorating function
Add it as an "annotation" to the function to be decorated.

```python
@decoratingFunction
def myfunction(a, b, c):
  ...

def decoratingFunction(f):
  def new_func(*args, **kwargs):
    ....
    return f(*args, **kwargs)
  return new_func

```

### property

is a decorator in python. Gives special functionality to certain methods to make them act as getters, setters or deleters.

The validation in setters can also help protect against users corrupting variables after the fact.

```python

class House:

  def __init__(self, price):
    self.price = price # this will call the price property setter method!

  @property # getter
  def price(self):
    return self.__price

  @price.setter # setter
  def price(self, new_price):
    # This enables validation to be carried out on incoming value
    self.__price = new_price

  @price.deleter # deleter
  def price(self):
    del self.__price

```
