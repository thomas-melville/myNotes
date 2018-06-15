# generator

like an iterator
it's a function, which when called iterates
very scalable
each value is only generated when required
good for huge datasets
once called it's exhausted

```python
roots = ( n**0.5 for n in range(6))
for root in roots:
    print root,

0.0 1.0 1.41421356237 1.73205080757 2.0 2.2360679775

```