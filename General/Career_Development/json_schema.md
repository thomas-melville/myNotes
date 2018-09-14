# json schema

http://json-schema.org/

A vocabulary that allows you to **annotate** and **validate** JSON documents

The document being validated it called the instance.
The document containing the description is called the schema.

The schema itself is a json object

```json

{}

```

You can apply constraints on an instance by adding validation keywords to the schema

```json

{ "type": "string"}

```

When defining your own schema you need to start with 5 properties

```json

{
    "$schema": "", // The schema to validate this json document
    "$id": "", // URI for this schema
    "title": "", // show the intent of this schema
    "description": "", // show the intent of this schema
    "type": "" // It has to be a json object!
}

```

The above schema introduces some terminology:

## Schema Keyword

* $schema
* $id
* * the unique identifier for this schema
* * it looks to always be a URI, but it's not pointing to the URI location
* $ref
* * reference another schema from this schema
* * the referenced schema can be embedded in this schema or external

## Schema Annotation

     title / description

## Validation keyword

* type
* * the type of the properties.
* * string, boolean, number, array, object, ...
* properties
* * the properties that this object contains
* required
* * the instance must have this property
* * there are two ways to use this, a required attributed on each property
* * or a required attribute on the object which is an array of required properties
* * when applied to an object the scope is limited to that object
* exclusiveMinimum
* * an exclusive lower limit for a numeric instance
* minimum
* * an inclusive lower limit for a numeric instance
* items
* * what is the valid type for this array
* * it is an object which has one property, type
* minItems
* * the minimum number of items that must be in the array
* uniqueItems
* * all items must be unique, think Set in Java
* additionalProperties
* * I think this allows properties not mentioned in the schema to be added to an instance and not fail the validation
* oneOf
* * allow a type of object to be one of many specific types.
* * value is an array which points to other schemas
* enum
* * the value is an enumerated type
* pattern
* * the value for a string must match the regular expression specified

```json

{
    // ...
    "properties": {
        "...": {
            "description": "...",
            "type": "..."
        }
    },
    "required": [ "..." ]
}

```