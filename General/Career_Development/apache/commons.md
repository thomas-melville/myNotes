# apache commons

http://commons.apache.org/

if you have to write code which you think is generic and some might have wrote before go to apache commons!

Composed of 3 parts:

1. Proper   reusable components
2. Sandbox  development
3. Dormant  currently inactive

## Active libraries

### Lang

and lang3, same APIs with a different package structure

#### lang3

Static utility methods

<type>Utils, there's a class for every type.
  the majority of operations are null safe

Functions
  attempts to address the fact that lambdas are not meant to throw exceptions
  It has Failable<Functional_Interface> class for all functional interfaces from the java.util.function package

RegExUtils
  Helpers to process Strings using regular expressions
  <remove/replace><All/First/Pattern>()

StringUtils!!!!
  Remove the chance of NPE!!!
  It's thread safe

Validate
  validate arguments provided to a method

#### lang3.builder

Assists in creating consistent equals, hashCode, toString and compareTo methods

CompareToBuilder
  useful methods to use in the compareTo method
  can compare object on multiple fields consequtively

EqualsBuilder
HashCodeBuilder
ToStringBuilder
ToStringStyle

#### lang3.time

Stopwatch
  very useful and precise for timing events in your code.
  if you write System.currentTimeMillis() consider moving to Stopwatch

A number of classes for working with Calendar, Date and SimpleDateFormat objects.
  all these classes are immutable

#### lang3.tuple

Mutable/ Immutable Pairs and Triples

### Exec

External process execution and env mgmt in java
last release was in 2014!!!

#### commons.exec

DefaultExecutor
  The default class to start a subprocess
  you can prepare a lot in the environment before executing the command
ExecuteStreamHandler
  capture the subprocess stdout and stderr
ExecuteWatchDog
  kill long running processes
ProcessDestroyer
  terminate any started processes when the main process is terminated.
CommandLine
  The command to execute
LogOutputStream
  connect a logging system to the out/err of the external process
OS
  Conditions that test the OS Type.
PumpStreamHandler
  copies standard out/err of sub-process to the out/err of the parent process
DaemonExecutor
  runs daemon processes async
DefaultExecuteResultHandler
  default impl of ExecuteResultHandler used for async process handling

### IO

utilities to assist with developing IO func.
Latest version requires java7+
Last release date was 2 years ago!
6 main areas.

#### Utility classes

Charsets
  charsets required of every impl of java platform
DirectoryWalker
  walk a directory structure
  subclasses with convenient hooks to add specific behaviour
FileCleaningTracker
  keeps track of files waiting for deletion
  deletes them when the associated marker object is garbage collected.
FileDeleteStrategy
  strategy for deleting files
FilenameUtils
  General filename and path manipulation utilities
FileUtils
  General file manipulation utilities
IOUtils
  General IO stream manipulation utilities
LineIterator
  An iterator over the lines in a Reader

#### Input

Loads of different input streams

AutoCloseInputStream
  close and discard underlying stream when end of input is reached.
CharSequenceInputStream
  read from String, StringBuffer, StringBuilder or CharBuffer
CloseShieldInputStream
  prevent underlying stream from being closed.
MessageDigestCalculatingInputStream
  Calculates a checksum using a MessageDigest
Tailer
  simple implementation of the "tail -f" functionality

#### Output

Loads of output stream and write implementations

ChunkedOutputStream
  breaks larger output blocks into chunks
ChunkedWriter
  breaks larger output blocks into chunks
CloseShieldOutputStream
  prevents underlying output stream from being closed
LockableFileWriter
  create and honor lock files to allow simple cross thread file lock handling
StringBuilderWriter
  Writer impl which outputs to a StringBuilder
TeeOutputStream
  Classic Splitter of OutputStream

#### Filters

FileFilter and FilenameFilter implementations for filtering a list of files

AgeFileFilter
  filter files based on cutoff time, newer/older
CanRead/WriteFileFilter
  accept files that can be read/write
Directory/Empty/File/HiddenFileFilter
  accept files that are directories/empty/files/hidden
NameFileFilter
  accept files with a certain name
RegexFileFilter
  accept files based on a regex
Size/SuffixFileFilter

#### Comparators

Various implementations of java.util.Comparator for files

Size/Path/Name/LastModified/Extension/DirectoryFileComparator

And they all have reverse versions

#### File Monitors

Monitor file system changes

FileAlterationMonitor
  A runnable that spawns a monitoring thread triggering any registered
FileAlterationObserver
  represents the state of files before a root directory, checking the file system and notifying listeners of create, change or delete events.
FileEntry
  state of a file at a point in time

### Collections

A recognized standard for collection handling in Java.
Builds on the Collections API in the JDK.
New interfaces, implementations and utilities

Majority of the implementations are not thread safe

#### Interfaces

Bag
  collections that have a number of copies of each object.
  Could be useful to know how many times an object has been found?
BidiMap
  maps that can be looked up from value to key and vice versa
MapIterator
  simple and quick iteration over maps
Transforming
  alter each object as it is added to the collection
Composite
  make multiple collections look like one
Ordered
  maps and sets that retain the order elements are added in
Reference
  allows key and/or values to be garbage collected under close control
Many comparator implementations
Many iterator implementations
Adapter classes from array and enums to collections
Utilities to test or create union, intersection, closure, ...
  CollectionUtils
  MapUtils
  <CollectionInterface>Utils

### Compress

defines an API for working with all manner of compressed files
Split into two components compressors and archivers.
Compressors work with streams that store a single entry.
Archivers deal with archives that contain structured content.
Some limitations when it comes to specific compression algorithm implementations.

It is recommended to wrap the stream in a buffered stream
You can limit the amount of memory that it can use during compress/archive
Wrap streams in streams when archive and compress are combined, for example tar.gz

Compressor/ArchiveStreamFactory
  Factory class for creating in/output streams with/without the algoritm.
  It will try and guess the format if not provided.
InputStreamStatistics
  Can be used to track progress while extracting a stream
ArchiveEntry
  Meta data about the individual archive entries
ArchiveInputStream
  Abstract class which all implementations extend
  call getNextEntry until it is null
ArchiveOutputStream
  Abstract class which all implementations extend
  iterate a list of files
    createArchiveEntry
    putArchiveEntry
    closeArchiveEntry
  finish


http://commons.apache.org/

Math

Compress
Configuration
CSV
Numbers
RNG (Random Number Generators)
Statistics
Text
VFS (Virtual File System)
