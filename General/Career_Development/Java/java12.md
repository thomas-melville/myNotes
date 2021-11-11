# java 12

## api updates

### String class

indent
  add spaces to left of string
  adds new line character, as it's for multi line string
  for use with multiline
    indents every line
  can be called with a negative value
    unindent an indented string

transform
  take a string and transform it
  method takes a Function and returns an object
  can chain string functions together
  a bit like streams

### CompactNumberFormat

turn numbers into short text representations
1000 -> 1K
the short representations are locale specific

``` java

NumberFormat snf = NumberFormat.getCompactNumberInstance();

String short = snf.format(1000);

```
By default it rounds to integers! You can set the number of fraction digits

### Stream

New collector in Streams API, teeing collector

Allows you to do two things at once in the collection phase.
apply two collectors at the same time.
then combine them to one at the end

### mismatch on the Files

do two files have the same content
checks bytes of files, if they don't match it returns the first mismatching byte otherwise -1

## switch expressions

preview feature
  --enable-preview on compiler and runtime

improvements to Switch statements

Since it's an expression it can be on the right hand side of an assignment

: can be changed to ->, which means the break statements can be dropped
  no fall through!!!

any value on the right hand side of the switch case can be returned

can combine cases

each case is it's own scope so variable names won't conflict
  use break as the return keyword

```java

String month = switch (monthNum) {
  case 1 -> "January";
  case 2 -> "February";
  case 3, 4, 5 -> "W";
  //...
  default -> "unknown";
};

```

## micro benchmarking with JMH

Micro benching is about analysing performance of small pieces of code, execution time
down to a single method
performance regression testing

performance critical part
code which is called a lot

reuse objects?
catch exception or check condition first?
which library call is faster?

Java Micro-benchmarking harness

reproducibility
handles jvm warm up
consistent reports
multi threading support

in the jdk, JMH is used to test:
  BigDecimal
  forName method
  Threading
  crypto algorithms
  ...

Annotation based

@Benchmark, on a method
reporting is throughput by default

@BenchmarkMode to specify what measurement

@OutputTimeUnit, what time unit you want

@State, use when you input values to the benchmark

@Parse, parse the values into JMH to be used in each run

@Fork how many jvms to fork

@Setup & @TearDown

### Pitfalls

Dead Code elimination
  return a value
other compiler optimizations
assumptions

### execution

separate benchmark.jar is generated which you can execute

## JVM changes

### G1GC improvements

divides the heap into smaller areas
wants to minimize time spent on garbage collection

abortable mixed collections for G1
promptly return unused memory from G1

### Shenandoah

new GC implementation

support very large heaps
low-pause times
Brookes pointers
  if object moves
  pointer to new location is left in old location

Experimental

### JVM Constants API

how class files are represented

constant pool
  numeric literal
  string values
  method references
  class references
  no distinction between description & actual value
  entries growing
  java.lang.constant package
metadata
fields
methods
