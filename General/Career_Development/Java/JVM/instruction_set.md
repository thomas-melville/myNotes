Content-Type: text/x-zim-wiki
Wiki-Format: zim 0.4
Creation-Date: 2016-11-17T09:09:35+00:00

###### instruction set ######
Created Thursday 17 November 2016

##### Arithmetic #####
IADD	add integers
LADD	add longs
FADD	add floats
DADD	add doubles


JSR
RET
JSR_W

##### Loading & Storing #####
LLOAD x	take the long value at x in the local variable array and load it in th operand stack
LSTORE y	take the long value at y in the operand stask and store it in the local array, or is that store it at index y in the local array

##### Returning from methods #####
IRETURN		return value of type boolean byte char short int
LRETURN		long
FRETURN		float
DRETURN		double
ARETURN		return a reference to an object
RETURN		return void

##### Invoking #####
INVOKESPECIAL invoked the constructor of a class

##### Exceptions #####
ATHROW		throw an exception. The exception is popped off the operand stack.
