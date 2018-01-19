# generics

Main benefit is compile time checks of types, so you don't end up with RuntimeExceptions which could be miles from the cause.

## Generic classes

MyObject<**type parameter**> myObject # new MyObject<>();

class MyObject<**formal type parameter**> {}

The invocation of a generic call is called a //parameterized type//

Inheritance doesn't work as expected with type parameters of a generic class.
If Foo is a subtype of Bar, and G is some generic class it is //not// the case that G<Foo> is a subtype of G<Bar>
The simple rule is that the type of the variable declaration must match the type you pass to the actual object

When a generic type is used without a type parameter it is called a raw type

## Wildcard
Wildcard Type
MyObject<?>

##### Bounded Wildcard #####
Bounded Wildcards
MyObject<? extends Shape>
This is an upper bound wildcard
This class can take any object which is a subtype of Shape

##### Generic methods #####
A method of a class can be generic
<T> void myMethod(T[]a, Collection<T> c) {}

When the type parameter is used only once, it is being used for polymorphism, in this case wildcards can be used. Supports flexibility
Generic methods allow type parameters to be used to express dependencies among types of one or more argument and/or it's return type.

Generic methods and wildcards can be mixed:
public static <T> void copy(List<T> dest, List<? extends T> src) 

? could be expressed as another formal type parameter but since it is used only once it is clearer and more concise.

##### Generics with legacy code #####

Raw types are very much like wildcard types but not type checked as stringently. Conscious design decision. Unchecked warning generated.
If you are using generics with legacy code then it is up to you to gaurantee safety.

erasure and translation?????

erasure: front end conversion. 
	Almost! like a source-to-source translation, generic version is converted to non-generic version.
	erases all generic type information.
	type replaced by Object, usually.
	and appropriate cast is included.
	
Legacy code with generic code generates unchecked warnings as the raw type is being used.

##### The Fine Print #####

**A Generic class is shared by all its invocations**

For example: 2 instantiations of list will have the same instance
All instances have the same runtime class, regardless of actual type parameter.

Static variables/methods are shared among all the instances.
It's illegal to refer to generic type in static context.

**Casts and InstanceOf**

Just don't do it.

**Arrays**
Component type of array may not be a parameterized type or type varialbe, unless unbounded wildcard.

##### Class Literals as Runtime-Type Tokens #####

Class is now generic, e.g. the type String.class is Class<String>
Can be used to improve type safety of reflection code.

##### More Fun with wildcards #####

#### Lower bound wildcards ####
<? super T>
An unknown type which is a super type of T (or T itself)

public static <T> T writeAll(Collection<? **extends** T>, Sink<T> sink) {...}
	legal but erroneous as inferred type could be higher than type
public static <T> T writeAll(Collection<T> coll, Sink<? **super** T> sink) {
	legal and correct as inferred type will only be lower than type.
	
This comes in handy in Comparators and Comparable

In general, if you have an API that only uses a type parameter T as an argument, its uses should take advantage of lower bounded wildcards (? super T). Conversely, if the API only returns T, you'll give your clients more flexibility by using upper bounded wildcards (? extends T).

**Whenever I have to cast should I think about using generics?**
