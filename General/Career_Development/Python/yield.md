# yield

key word
creates a generator function, which can be used as an iterator.
yield is used in place of return


in threads yield means the thread will remain runnable.
If return is used it means the thread is done and destroyed
if there is no return, one is implied at the end of the function

```python

def byteGen(start, stop_before):
  s = start
  while s < stop_before:
    yield bytes(start)
    s+=1

```
