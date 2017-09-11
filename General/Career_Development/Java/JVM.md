Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2016-11-17T08:46:36+00:00

###### JVM ######
Created Thursday 17 November 2016

https://docs.oracle.com/javase/specs/jvms/se7/html/jvms-2.html

##### Class File #####
Compiled code to be executed by the jvm is hardware & OS independent

##### Data Types #####
JVM operates on 2 kinds of types: primitive & reference, just like the language
Correspondingly, there are 2 kinds of values that can be:
1. stored in variables
2. passed as arguments
3. returned by methods
4. operated upon

1. primitive values
2. reference values

jvm expects nearly all type checking is done prior to run time, by a compiler
jvm distinguishes operands using instructions intended to operate on specific types: iadd (int), ladd (long)...

jvm contains explicit support for objects, an object is either a dynamically allocatec class instance or an array.
A reference to a an object has type //reference//.
values of type reference can be thought of as pointers to objects. multiple references can point to the same object.
Objects are always operated on, passed and tested via values of type reference

##### Primitive Types and Values #####
primitive data types: numeric types, boolean type and returnAddress type.
numeric types consist of integral types and floating point types
integral types:
	byte	8-bit signed two's complement
	short	16-bit ...
	int		32-bit ...
	long	64-bit ...
	char	16-bit unsigned integers representing unicode characters
floating point types:
	float
	double
	
returnAddress:
	values are pointers to the opcodes of jvm instructions


#### Floating point types ####
conceptually assocaited with 32-bit single precision and 64-bit double precision format IEEE 754 values and operations
standard includes positive, negative magnitude numbers, but also possitive & negative zero, positive & negative inifinities and a special Not-a-number value (NaN) which represents the result of certain invalid operations (divide by zero)

//loads in this section//

#### returnAddress type ####
used by instructions: jsr, ret, jsr_w
values of type are pointers to opcodes of jvm instructions **(an opcode is a code which maps to a particular jvm instruction)**
does not correspond to an any java prog language type and cannot be modified by a running program.

#### boolean type ####
jvm defines a boolean type but only limited support for it.
no instructions solely dedicated to operations on boolean values
int instructions are used for boolean types

##### Reference Types and Values #####
3 kinds of reference types
1. class
2. array
3. interface

values are references to dynamically created class instances, arrays or class instances arrays that implement interfaces

array type consists of component type and eventually element type. (the component type of an array can itself be an array type, but you eventually get to element type)

default value of reference type is null, null initially has no run-time type and can be cast to any type

##### Run-Time Data Areas #####
various run-time data areas. Some are created on jvm start up and exist through out the life of the jvm, destroyed only when the jvm is destroyed. Others are per-thread, created and destroyed within the life time of a thread.

#### pc Register ####
program counter register
jvm can have many threads of execution at once.
each thread has its own pc register.
At any point in time each thread is executing the code of a single method
pc register will contain the address of the jvm instruction currently being executed, if it not native. If it's native then the value is undefined. **(native method - need to call a library written in another language, less common as java gets faster)**
This address points to a location in the method area, where the compiled code sits.

#### jvm stacks ####
each thread has it's own stack, created when the thread is created.
a stack stores frames.
it holds local variables and partial results and plays a part in method invocation and return
The stack is never manipulated directly except to push and pop frames, frames may be heap allocated.
Memory for a stack does not need to be contiguous.
stacks can be fixed size or can dynamically expand and contract as required by computation
jvm implementation can give user control over fixed size or max / min values for dynamic

Exceptions involving the stack:
1. StackOverflowError
	a. if the computation requires a larger stack than is permitted.
2. OutOfMemoryError
	a. if there isn't enough memory for the stack expansion.

#### heap ####
shared among all threads
run-time data area from which memory for **all class instances and arrays is allocated. **Different threads can have an instruction in the pc register which points to a single object in the heap.
created on jvm start up
garbage collection occurs on the heap for objects.
storage management system and garbage collection is left to the implementer of the jvm.
can be fixed or dynamicn
and memory doesn't need to be contiguous

Exceptions involving the heap:
1. OutOfMemoryError
	a. computation requires more heap space than can be made available by storage management system.

#### method area ####
shared among all threads.
analoguous to the storage area for compiled code of a conventional language
stores per-class structures such as the run-time constant pool, field and method data and the code for methods and constructors
created on jvm start -up
logically part of the heap, simple implementations may choose not to either garbage collect or compact it.
can be fixed / dynamic and doesn't need to be contiguous.

#### run-time constant pool ####
per-class/interface run-time representation of the constant_pool table in a class
contains several kinds of constants:
1. numeric literals known at compile-time to
2. class, method and field references resolved at runtime.
each pool is allocated from method area
constructed when the class/interface is created.

#### native method stacks ####
jvm impl may use conventional stacks, called C stacks, to support native mthods.

##### Frames #####
Is used for:
1. Store data and partial results
2. perform dynamic linking
3. return values for methods
4. dispatch exceptions

Created when a method is invoked and destroyed when method invocation completes either successfully or abruptly

allocated from the jvm stack of the thread creating the frame.

each frame has its own array of local variables
its own operand stack
and a reference to the run-time constant pool of the class of the current method

size of local variable array and operand stack are determined at compile time.

Only one frame is active at any point in a given thread of control, this frame is referred to as the current frame, the class is the current class.
A frame ceases to be come the current when it invokes another method or if it completes.
When a method is invoked, a new frame is created and becomes the current frame when control transfers to it.
On method completion it transfers control back to the previous frame which then becomes the current frame again. once control is transferred back the frame is discarded.
**(In the debugger in Intellij you'll see the "hierarchy" of frames at a particular point in the code)**

#### local variables ####
each frame contains an array of variables.
size determined at compile time
supplied in the binary representation of a class/interface along with the code for the method associated with the frame
single local variable can hold a value of type: boolean, byte, char, short, int, float, reference, returnAddress
pair of local variables can hold long or double.
addressed by indexing, first is zero.
long/double is addressed using the lesser index, since it takes up to spaces in the array

jvm uses local variables to pass parameters on method invocation.
Class method invocation: parameters are passed in consecutive local variables starting from local variable 0
Instance method invocation: local variable o is always used to pass reference to the object on which the instance methosd is being invoked, parameters are passed starting at 1

#### operand stacks ####
each frame contains an LIFO stack known as the operand stack (last in first out)
max depth is determined at compile time and supplied along with code to the frame
it's empty on frame creation.
jvm supplies instructions to load constants / values from local variables or fields onto the operand stack.
**(This is the variables pane on the debugger???)**
other instructions take operands from the stack, operate on them and push the result back.
also used to prepare parameters to be passed to methods and to receive method results.

values on the operand stack must be operated on in a way that is appropriate to their type.

#### dynamic linking ####
of method code.
class file code for a method refers to methods to be invoked and variables to be accessed via symbolic references
dynamic linking translates these symbolic method references into concreate method references.

##### Special Methods #####

#### <init> ####
every constructor appears as an **instance initialization method**
special name <init>

#### <clinit> ####
A class or interface has at most one **class or interface initialization method**
takes no arguments
returns void.

##### Exceptions #####
represetend by an instance of Throwable or one of it's subclasses
results in immediate nonlocal transfer of control from the point where the exception was thrown.
The jvm searches for an exception handler to handle the exception.
If none is found in the current frame then the frame is discarded and the previous frame is reinstated and the exception is thrown there, so on and so forth until it reaches the top of the stack.

##### Instruction Set Summary #####

jvm instruction consists of a one-byte opcdeo specifying the operation to be performend, followed by zero or more operands supplying arguments or data that are to be used in the operation

#### inner loop of jvm interpreter ####
do {
	atomically calculate pc and fetch opcode at pc
	if ( operands ) fetch operands;
	execute the action for the opcode;
} while ( there is more to do)

most instructions encode type information about the operations they perform
	iload loads the contents of a local variable, which must be an int
	
i - int
l - long
s - short
b - byte
c - char
f - float
d - double
a - reference

a lot of instructions don't have boolean, byte, char and short specific instructions: they use the int instruction:

Actual type 	Computational type 	Category
boolean		 	int 						1
byte 				int					 	1
char 				int					 	1
short 				int					 	1
int				 	int					 	1
float 				float 					1
reference 			reference 				1
returnAddress 	returnAddress 		1
long 				long 					2
double 			double 				2

#### Load & Store instructions ####
transfer values between the local variables and the operand stack of a jvm frame

load a local variable into the operand stack				<T>load, where T # type
store a value from the operand stack into a local variable <T>store, where T #type
load a constant on to the operand stack:					bipush, sipush, ldc, <T>const_<I> i is a number and T is a type

#### Arithmetic Instructions ####
compute a result which is typically a function of 2 values on the operand stack, pushing the result back onto the operand stack
2 types, integer and floating point

<T>add
<T>sub
<T>mul
<T>div
<T>rem
<T>neg
...

where T is a type

no care is taken for overflow results

only exceptions are divide by zero on div and rem

#### Type conversion instructions ####

### widening ###
int to long, float ot double
long to float or double
float to double

<T>2<U> T # type and U # type

loss of percision may occur

### narrowing ###
int to byte, short or char
long to int
float to int or long
double to int. long or float

same instruction syntax

#### Object creation and manipulation ####
Although both class instance and arrays are objects, the jvm creates and manipulates class instances and arrays with different instructions

new class instance
new com/../..///ClassName	

new array
newarray, anewarray, multianewarray 

access fields of class and instance
get/setfield, get/setstatic

load an array component onto the operand stack
<T>aload and the operand stack must hold a reference to the array and the index to retrieve

store a value from the operand stack as an array component:
<T>astore, and the operand stack must hold a reference to the array, the index to store the value at and the value

arraylength

instanceof & checkcast

#### Operand Stack Management Instructions ####

pop pop a value off the operand stack
pop2
dup duplicate the top erand stack value and push to the top of the operand stack
dup2
...

#### Control Transfer instructions ####

ifeq, ifne, iflt, ifle, ifgt, ifge, ....

#### Method Invocation & Return Instructions ####

invokevirtual invokes an instance method of an object

invokeinterface invokes an interface method, searching the methods implemented by the particular run-time object to find the approrpriate method

invokespecial invokes an instance method requiring special handling, instance initialization, private or superclass

invokestatic

invokedynamic ???

T<>return return a value of type T

return void

#### Throwing exceptions ####

athrow

#### Synchronization ####

method_info in run_time constants pool, flag ACC_SYNCHRONIZED

monitor construct is used

monitorenter and monitorexit
