# primtetime JMS

## complication before jpmn

### dependencies

unexpressed dependencies
transitive dependencies
cyclic dependencies

JPMS: exit with error when it encounters cyclic dependencies

### Shadowing

issues with different versions of transitive dependencies

JPMS: more than one version exits with an error

### version conflicts

JPMS: leave it to build tools

### java 8

compact profiles, 4 fixed levels

JPMS: cherry pick required jdk classes

### weak encapsulation

across packages there is only public
reflection!

JPMS: class can be private to a module (in the same way a private field is private to a class)

## jpms concepts

set of packages designed for reuse
module-info.java
  name
  exposed APIs, exports
  who it depends on, requires

exports only exports classes in that package, it is not recursive

module path, coexists with classpath

module graph created at runtime

explicit & autmoatic modules

jpms has the concept of transitive deps between modules, it's called implied readability
  keyword transitive in module which requires it

"export to" only export to specific modules

"opens" package to reflection
  hibernate & spring would require this a algorithms

jlink to build custom jdk distributions
  *this is where the real benefit for us is*
