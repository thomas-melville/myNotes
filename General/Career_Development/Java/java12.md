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

## micro benchmarking with JMH

## JVM changes
