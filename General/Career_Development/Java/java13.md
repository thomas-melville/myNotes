# java 13

## API Updates

### ByteBuffer

All about reading and manipulating blocks of bytes which live together
Not a byte array
efficient api for reading and manipulating bytes
native io operations if possible
used in a lot of the File NIO APIs

bulk put and get methods
  index, byte[]
  overloaded with offset and length so subset of array is used

these methods can be used in a multithreaded environment as the index and offset is part of the api call.
However, care should still be taken not to write to the same area at once,

### javax.security

javax.security.cert package removed
  Certificate
  X509Certificate

just remove the x

### Switch expression update

using break to return a value from a case has been changed.
A new key word "yield" is used

similar use in python

### Unicode support

updates support from uniode 11 to 12.1
south east asian languages mostly
61 new emoji characters

CLDR is updated to version to 35.1
Common Locale Data Repository

### ZGC Garbage Collector

Experimental, low pause times
now uncommits unsed memory back to the OS, by default 5 minutes

## Text Blocks

preview feature

String text = """gjreiog
nerkonero
jbvieroneo
nvkoernvejko""";

If you want a new line at the end then put the ending delimiter on the next line

Solve the problem of multi line strings, example json.
You had to be very mindful of escaping characters.

New lines are factored in
no need to escape special characters

Text blocks can distinguish between non essential and essential indentation.
non essential =  indenting for code formatting and white space between first """ and first character in the string
amount of non essential white space removed is decided by where the closing """ is placed on the next line.
The closer it is to the left hand side the less white space that is removed.

Compiler will
1. normalize line endings. Windows V Linux
2. Trim whitespace
3. Translate escape sequences

### String API Updates

3 new methods connected to text block feature

stripIndent, removes non essential / incidental whitespace from a multi line string
  already marked as deprecated
  because it's linked to a preview feature
  once the Text block featrure is finalized the deprecated will be removed.

translateEscapes, take a String which has been read in somehow and apply the special characters

formatted, achieves the same functionality as String.formt(string, string...)

## Platform Changes

### Socket API

Old API with lots of issues

Socket and ServerSocket

more maintainable and modern implementation

getting ready for loom

not using stack as i/o buffer

IT's not preview, it's just there.

NioSocketImpl is the new class.
Reuse native nio code
using modern locking constructs

compatible with old behaviour

old implementation can be used with a java switch

1-3% better performance

### Dynamic AppCDS Archive

Class Data Sharing

improve start up times across different runs of a jvm
reduce foot print when running the same application on one machine

start the jvm with a flag -Xshare:dump
generates a classes.jsa file

start the jvm with flag -Xshare:on
it makes use of the file generated in the previous run

IT started as only System class Sharing

then it was possible to do with Application classes

-XX:DumpLoadedClassList=myapp.lst

but you had to conver lst file to jsa file
and could not use custom class loaders

That was introduced in Java 10

In Java 13 new flag -XX:ArchiveClassesAtExit=myclasses.jsa now produces the file

To use it on a subsequent run use the flag -XX:SharedArchiveFile=myclasses.jsa
also picks up classes loaded from custom class loaders.

The class path for the two runs must be the same

## deprecated flags

-Xverify:none   
-noverify

use AppCDS as the files are pre verified
