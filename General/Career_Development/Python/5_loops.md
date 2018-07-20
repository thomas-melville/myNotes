# loops

range method, upper limit is less than

```python
range(10)
```

give me a list of integers starting at 0 and stopping before 10
return a list of 10 ints 0 - 9

```python
range(5, 10)
```

give me a list of integers starting at 5 and stopping before 10
return a list of 5 ints 5 - 9

```python
range(5, 10, 2)
```

give me a list of integers starting at 5, stopping before 10 and incrementing by 2
return a list of 5, 7, 9

syntax:

```python
for a in <list>:
	...(a)

while a < condition:
	....(a)
```

for and while both support and else condition.
This is executed when the loop completes successfully, i.e. no break

### break & continue

the same as other languages