# testing in python

## doctest

we write doctests within the documentation string

```python

import doctest

def method(a, b):
  '''

  >>> method(2,3)
  5
  '''
  return a + b

if __name__ == '__main__'
  doctest.testmod()

```

## unittest

```python

import unittest

class testMyClass(unittest.TestCase):
  def setUp():
    # override the default setup with your own

  # test methods


if __name__ == ' __main__'
  unittest.main() # run all methods which start with test.
```


## pytest

https://docs.pytest.org/en/latest/

tests are written separate to production code
convention of starting test function names with "test_"
