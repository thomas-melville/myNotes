# good api design

https://medium.com/97-things/lets-make-a-contract-the-art-of-designing-a-java-api-854441cd42f5

consistency
minimal, for ease of use

"principle of least astonishment"

break apart large interfaces into smaller pieces
  consider implementing a fluent API

never return null
  use empty collections and Optional instead

limit usage of exceptions and possibly avoid checked ones.

avoid long method argument lists, especially of the same type
use the weakest possible type
keep them in consistent order in overloads
