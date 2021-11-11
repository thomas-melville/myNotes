# what's new in java 14

## API Updates

### Removal of deprecated types

java.security.acl removed
java.util.jar.Pack200 removed

### deprecated for removal

Thread#suspend/remove

### additions

PrintStream#write(byte[] buf)
  write the whole byte array to the stream
  previously you had to specify the offset and limit
  but a parent class has the same method which throws an IO exception
  if you don't want an IO exception use:
PrintStream#writeBytes(byte[] buf)

new annotation @Serial
  apply to serialization related fields
  class implements Serializable
  and the serialVerionUID
  compile time check

### Helpful NPEs

not very informative, only gave line numbers. But what if there was chaining.

-XX:+ShowCodeDetailsInExceptionMessages

results in a much more informative NPE message which describes at what point there was a Null pointer
specific for calling methods on null objects and setting fields on null objects

## Switch expressions

No longer preview feature! Officially added to the java language

old switch statement
  fall through
  no return value
  no new scope

not a replacement for switch statements
they are a variant on switch!

1. a switch expression can be a variable and so assigned to something
    and a ; needs to be added at the end
2. change colon to arrow syntax
    ->
    no fall through
    return value
3. default case required if cases are not exhaustive
    compilation will fail otherwise
4. easily combine cases using ,
5. Block scope is possible by using {}
6. use the keyword yield to return the value

you can mix old and new switch variants
  expre can use old : but then yield has to be explicit
  statemwnr can use -> so there is no fall through and not exhaustive


## preview features

### pattern matching

instanceOf check followed by a cast

A pattern is a combination of a predicate which can be applied to an object and one or more binding variables

instanceOf now accepts a binding variable

```java

if( object instanceOf BigDecimal b){
  return b.precision();

}

```

can't be used in negation check

### text blocks

I've seen a lot of this in "what's new in java 13"

what's new in java 14?

add \ in text block -> tell compiler that this line ending should not be preserved

more subtle changes, can be seen in the jep

### jpackage

incubation phase
*platform specific* installer/packages for java applications
  no need to have java installed

supports windows, mac & linux (Deb, rpm)

### records

Behaviour V data
Java is everything about Behaviour
if you just want a data class you still end up with a lot of code
lombok minimized a lot of this, but all the code was still generated

```java

public record Person(String name, int age, String hobby){}

````

Components are the variables name, age and hobby
A Record is like an ENUM, an immutable class
Constructor is generated
equals//hashcode/string are available also

each component gets an accessor method, by the component name
So, no get/set

You can add a Constructor to the record which validates the values

```java

public record Person(String name, int age, String hobby){
  public Person {
    if(age < 0 ){
      throw new Exception();
    }
  }
}


```

## JVM improvements

### Garbage collection

#### G1

NUMA-aware memory allocation
Non Uniform Memory Access

goal is to improve performance on large machines
large heap on multiple processor machine

#### ZGC

Goals:
Pause times of under 10ms
scale to terabyte heap sizes

Coloured Pointers
  bits in pointer reserved to highlight GC status of memory

#### CMS

removed!
you will get G1 by default

### low-level improvements

Non-volatile mapped byte-buffers
  writing memory which is preserved across app restarts

foreign access access API
  incubator
  efficient access to off heap Memory
  project Panama

JFR event streaming
  real time insight
  instead of writing to file and using after the fact
  could integrate with your monitoring
  
