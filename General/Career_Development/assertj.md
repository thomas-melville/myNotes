# assertj

http://joel-costigliola.github.io/assertj/

Thoughtworks recommended to adopt it in their May 2018 tech radar

fluent assertion library with powerful assertion methods, which also allows chaining of assertions
easily include descriptive error messages
easily assert on exceptions thrown

## extracting

Extract one or more fields from a list of objects.
When extracting more than one field use the tuple method to group expected values

## flatExtracting

just like flatMap.
Take a list of objects that have a property which is a collection.
Each collection gets merged into one and you can assert on the contents

## assertion generator

you can create your own assertion classes, or have assertj create them for you