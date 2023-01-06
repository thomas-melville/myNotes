# functional programming

Python supports higher order functions.
i.e. functions can be passed around as parameters and executed when needed.

```python

def squareIt(x):
  return x*x


list(map(squareIt, range(1,12)))

```

squareIt is passed as a parameter to map, where it is applied to each number in the range.
This is then collected to a list.

map returns a generator

python has a filter method too, which also returns a generator

```python

list(filter(isOdd, range(1,27)))

```

reduce - apply a function to end up with a single value

```python

from functools import reduce

r = reduce( addThem, range(1,10))

```

addThem is a function which should be defined.

## lambda functions

anonymous functions, same as I've seen in other languages

```python
power = lambda x: x * 2

sorted(names, lambda a,b: len(a) - len(b) or cmp(a,b))
```

x is the value to pass in
x * 2 is the operation

called just like a normal function

higher order functions
