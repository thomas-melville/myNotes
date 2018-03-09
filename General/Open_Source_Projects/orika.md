# orika

Java bean mapping framework
recursively copy data from one object to another

uses byte code generation to create fast mappers with minimal overhead.

## Implementation

corner stone of the framework is the **MapperFactory**
Use this to configure any mappings and obtain the **MapperFacade** (which performs the actual mapping work)
Once you've created the MapperFactory you can register:

* field mappings for classes you'd like to map
* register converters
* custom mappers
* concrete types to be used for abs/if types
* obtain a MapperFacade instance, which will be used to perform the actual mapping

mapperFactory.classMap(class1, class2)...
    fluent api for creating mapping between fields
    .register()

mapperFacade.map(object, dest.class)

**BoundedMapperFacade** improved performance when you have a specific pair of types to want to map.
This is a generic class which takes to generic types, source and dest types.

If you have a mapping which is not supported out of the box then you can implement your own **CustomConverter<T, U>**
Extend CustomConverter and implement the convert method.
This is a single direction conversion, there is another class for implementing a bi-directional custom converter. **BidirectionalConver<T, U>**
These can be registered with the MapperFactory through, getConverters.register(...)