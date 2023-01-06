# context manager

methods can be decorated with contextmanager to make them available to manage contexts.
The method can then be the argument of the "with" keyword

```python

@contextmanager
def my_method():



main:
  with my_method():
    //...

```
