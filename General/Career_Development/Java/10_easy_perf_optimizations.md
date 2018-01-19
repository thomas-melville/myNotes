# 10 easy optimization improvements

## Use StringBuilder instead of String concatenation

    String concatenation is compiled into a StringBuilder. If you just use it once it won't have much advantage.
    But, if you are concatenating a single string in multiple places uses a StringBuilder to reduce the overhead of creating an implicit StringBuilder every time.

## Avoid regular expressions

    Cache the pattern instead of creating it multiple times, store it as static
    Beware of JDK String methods which use regular expressions! split, replaceAll
    Use Apache Commons Lang instead for String Manipulation

## Don't use iterator deep in an algorithm

    Iterators and the foreach loop are useful for readability but they create a small new instance on the heap for each iteration
    use old school index based for loops if it's going to be called a lot

## Double check to make sure you have to call expensive methods

    There may be a default result for another method call which means you don't need to call the expensive one

## prefer primitives to wrappers deep in the algorithm

    Lots of Wrapper objects will be created in the heap space when primitives live in the stack
    Boolean, Byte are exceptions

## Avoid Recursion

    Always prefer iteration over recursion when you're deep in the algorithm

## use entrySet()

    when you want the keys and values out of a map prefer entrySet to keySet & get(key)

## Use EnumSet or EnumMap

	When the keys for a map are known beforehand then consider using an Enum

## Optimize your equals and hashCode methods

	A good hashCode implementations prevents calls to the more expensive equals method as it will produce more distinct has buckets per set of instances