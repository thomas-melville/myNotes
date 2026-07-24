# Go

https://golang.org/doc/

A tour of go!

Created as an attempt to provide a single language that combined:

* Efficient compilation
* Efficient execution
* ease of programming

## Philosophy & Values

* Simplicity
	Go ecosystem should be easy to approach
	adding new features tends to drown this out
	example:
		pre/post increment expressions
		easily missed in code
		expression is a component of a statement
		only post increment is a statement in GO.
			no need for pre as it's its owbn statement
		a statement is evaluated all on it's own
	it means extra lines of code, but intent is clearer
	another example
		all loops in go are for loops
* Network aware and concurrent apps
	a must in todays world
	net & net/http packages in core GO
	goroutines, lightweight threads which allow for massive concurrency
		1000s of concurrent tasks with minimal resources
		abstraction over a processor thread
		scheduler to keep track of goroutines
	channels, for communication between concurrent goroutines
* out-of-the-box experience
	related to simplicity
	all tools should be there
	many languages require a lot of setup
	go want a single place to start developing
	standard library
		string manip, compression, files, net, testing apis, ...
	Go CLI
		project init, build, code generation, retrieve deps, test, profiling, docs, report language bugs
* cross-platform
	are the norm
	built into the language
	GOOS = android
	GOARCH = arm
	just set these to compile for an OS
	program with a singel style and compiler will change to OS
		file separators for example
* backward compatibility
	very focused on it
	honour it, within limit!
	any app written in v1 will work for the life time of the v1
	some exceptions!
		security
		unspecified behaviour
		spec errors
		bugs

## What is Go?

### Who Developed Go?

3 engineers, Ken Thompson, Rob Pike, Robert Griesemer
worked for Google at the time

They broke down what languages were good at:

Efficient compilation: Java, python, js
efficient execution: C++, Java
ease of programming: js, python

no language solved all 3

### Go's history & future

2007, designed started by the 3 above

2009, publically announced

2012, 1.0 release

6 month release cycle

discussions about version 2.0

### Language characteristics

* Strong, static type system
	focused on having well known types that are available at compile time
	allows for higher levels of tooling support
		IDEs to help you
		compiler to generate more efficient code
	typically associated with verbose syntax
		decision to put as much emphasis on the compiler helping you figure out the types
		rather than you as the programmer telling it all the types
* c-inspired syntax
	not just blindly adopted
	re-examined how we as engineers to write code
		to help with how we view the code
		and make it easier for the compiler
* multi-paradigm
	* procedural
	* object oriented
	* functional?
	choice is up to the programmer what style to use
* garbage collected
	no need to manage memory on your own
	makes it more robust
	makes application performance a challenge when GC is there
	They have worked hard to eliminate this performance hit
* fully compiled
	down to an executable binary
	different to java which is compiled to byte code which is ran by a jvm
* rapid compilation
	when compilation takes time its hard for engineers to stay in a productive flow
	even more important when you use TDD
* single binary output
	allow for simple deployment
	this is the default option
	you can write libraries and share it with others you import it

### who is using Go?

IBM
Amazon
DigitalOcean
Netflix
Google
docker
facebook
Kubernetes
Helm
Helmfile

### what types of apps

1. web services
	REST
2. web applications
	HTML
3. DevOps
	docker
	k8s
GUI / Thick client apps
Machine Learning
... & still growing

## go tools

### go

command line tool to allow you to fetch, build and install Go packages and commands
it requires you to organize your workspace in a particular way

### gofmt

format the code adhereing to strict rules from go

### go tool

see the list of and run a specific go tool

Example: pprof for profiling

##### workspace #####
all code is kept in a single workspace.
	the workspace can have many version controlled repos
	each repo can contain one or more packages
	each package consists of one or more go source files in a single directory
	**the path to a package's directory determines it's import path**
"go" folder in home directory
	under this is bin, pkg and src
		the convention in go is to place all programs in separate folders under <scm>/<userid> under source, for example github.com/ethomev
		**However, all folders under ethomev are separately versioned controlled repos**
	go build, builds executable beside src
	go install, install executable to bin folder


##### environment variables #####

#### GOPATH ####
	path to workspace on your machine
	add GOPATH/bin to the PATH for convenience


##### go code #####

functions in go can be higher order, i.e. take functions as arguments

it uses slicing of arrays like in python, e.g. [1:] give me everything after the first element to the end

importing other packages is done inside a bracket if there are multiple
e.g.
import (
	"..."
	"..."
)

when importing you can give a package an alias, the alias is placed between import and the package name

it's called a factored import statement
you can also write each import individually

for a go file to be executable from the command line it looks like it needs to have:
1. package main
2. func main

install a file with a main and the executable goes to bin
install a file without a main and the executable goes to pkg

first statement in a go source file must be package //name//
convention is that the package name should be the current folder name

convention is that packages are lower case single word names

case is important for fields and functions of a source file.
starting with lowercase -> private / unexported to use go terminology
starting with uppercase -> public / exported to use go terminology

the var keyword can be used to declare a list of variables, type is last
these variables can be initialized, type can be left out as it will be inferred
when not initiialized vars have default values, 0, false, ""

you can also define multiple vars inside () outside the function

when creating multiple vars in one line you can initialize them as well.
But the list of initial values comes after the type. 

:= "short assignment statement"
	can be used in a function in place of var with implicit type

no auto conversion between compatible types in GO, must explicitly cast.

constants are declared using the "const" key word
	can't use := with const
	should do same convention of constatns as upper case

Not many keywords in GO, 25

#### function signatures ####

func <pointer> <name>(<parameters>) <returns> {

#### function parameters ####
x int, y int
type comes after the variable name
if mulitple parameters in a row have the same type only the last type specificuation is required.
	the earlier ones will infer the type from the last one

#### return values ####

functions can return multiple values
variables that returned values are assigned to are specified in a command separated list
specify _ as the identifier for a particular returned value if you want to ignore it

to return mulitple values from a function specify them as a command separated list

you can also name the values to be returned
specify their name along with the type in the function signature
Then set the names in the function and use the return keyword on it's own at the end.

no default values for function parameters
no named parameters
no function overloading

##### go compiled code #####

the executables look to be quite large in size.
that is because all required dependencies are compiled along side your code

##### testing in go

append test file with _test.go
In the same package as the code, so it can access unexported functions.
But, if tests in folder then can't run go run *.go

import testing package
it should contain functions with signatures: func TestXXX(t *testing.T)

Arrange, Act & Assert in each test

no formal assert function.
verify outcome, if it's incorrect call t.Error() / t.Errorf() - function will keep running.
Call Fatal/Fatalf() if you want function to end

Parameterized tests supported
* define a structure containing inputs and expected outputs.
* create a slice of this data
* iterate through the slice and execute the test code

###### Setup & Teardown

TestMain function to manage state and run tests

```go

func TestMain(m *testing.M){

	...

	exitVal := m.Run()
}
```

###### Coverage & Benchmarking

integrated in Go testing framework
run the tests with go test -cover

benchmarking integrated into GO testing fwk also.

write BenchmarkXxx (b *testing.B){
	for loop with b.N for number of iterations
}

run go test -bench=.

##### go get #####
for downloading remote packages
go get <scm>/<import path>

##### pointers

Go is pass by value, so if you want to update a function parameter in the function and have it visible after you need to use pointers.

declare a pointer

var myPointer *int
you need to specify the type it will point to.
initialized to nil by default

To get the address of a variable use the & prefix.

var myPointer *int = &myVariable

To dereference a pointer and update the value use *

*myPointer = 34556

Need to use * in functions to update the value of the poibnted to address

No pointer arithmetic in GO!

## garbage collector

automatic memory management process
	identfies and reclaims memory no longer in use
Simplicity
Safety
Performance
mark and sweep algorithm
tri-color system
	white - unreachable objects collected during sweep phase
	gray - objects with references that have not been fully processed
	black - fully processed objects confirmed as reachable

each function gets a stack for variables which is discarded once the function returns.
Go has a heap for objects which live outside the scope of a single function.

reduce heap allocations
	return value types, which means a copy of the value is returned.
	inline functions, so that stack is discarded after
	object polling
		use sync.Pool for reusable objects
		instad of constructing and destroying objects many times use a pool of said objects