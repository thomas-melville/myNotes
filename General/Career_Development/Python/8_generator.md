# generator

like an iterator
it's a function, which when called iterates
very scalable, lazily loaded. Only loads one value at a time into memory!
each value is only generated/yielded when required
good for huge datasets
once called it's exhausted

```python
roots = ( n**0.5 for n in range(6))
for root in roots:
    print root,

0.0 1.0 1.41421356237 1.73205080757 2.0 2.2360679775

```
roots is a generator in the snippet above

Writing your own generator:
1. It's a function
2. use yield instead of return
3. create a variable which is the instance of the function.
4. call __next__() on the instance
