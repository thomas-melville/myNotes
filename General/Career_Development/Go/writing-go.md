# writing go

https://blog.stackademic.com/best-practices-in-go-golang-writing-clean-efficient-and-maintainable-code-dccf61542b57

all go files start with a package declaration

func main(){} in the package main is the entry point

functions can return multiple values

the case of the first letter of a function decides whether it is exported outside the package or not.
Capital means exported, lower case means hidden.

functions can also return pointers to values, this is identified by a *
when referencing the variable you also need to use the *

```go

f, err := os.Open("myapp.log")
```

import keyword for packages

```go

import (
  "fmt"
  "os"
)
```

nice compile time error if you declare a variable and don't use it
  simplicity, compiler helps us
  also if you define a module but don't use it

opening brace for function must be on same line.
  Go enforces automatic ; insertion

defer keyword, after the main function exits execute command passed to it!
  can be used anywhere in the function, but the method will only be executed at the end.
  If there are multiple returns it will save putting the function call before each return
  it's evaulated where it's coded, but only executed at the end
  you can stack defer calls, last one in first one out.

comments in go

one line //
multi line /* */

## go.mod file

defines module and go version
(kind of like pom.xml im maven)

go get to install modules for your project. Adds them to go.mod
However, just import them to your files then call "go mod tidy" and it will add them to your go.mod file.

A module can have multiple packages

when importing it's <module-name>/<package-name>

keyword internal to create packages which should not be exported

## go cli apps

flag package for command line flags
since go is strongly typed methods to get flags are typed
command line flags require default values

```go

path := flag.String("path", "myapp.log", "Help text")

// ...

flag.Parse()

```

## go web services

http package

```go

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  names := r.URL.Query()["name"] // qeury parameters
  var name string
  if len(names) == 1 {
    name = names[0]
  }
  m := map[string]string{"name": name} // map declaration in Go
  enc := json.NewEncoder(w) // working with json
  enc.Encode(m)

  w.Write([]byte("Hello " + name))

  err := http.ListenAndServe(":3000", nil)

  if err != nil {
    log.Fatal(err)
  }
})

```

Write works with bytes, so the string is written as an array of bytes

## Primitive data types

```go

var i int
i = 42

var f float32 = 3.14

firstName := "Arthur" // implicit initialization syntax
```

All the usual data types, they also include a complex number type (real and imaginary parts for mathematics)

## Pointers

Go is pass by value, so if you want to update a variable passed into a method you need to pass a pointer to it.

Instead of holding the value, hold a pointer to the address in memory where the value is
Add an asterix before the type when declaring, and the name when setting/getting
need to be initialized to a new(<type>)

```go

var firstName *string = new(string)

firstName = "Arthur" // won't work, as you'll replace the pointer address

*firstName = "Arthur" // "dereference" the value pointed to by the pointer

```

<nil> = Go's terminology for a empty pointer

No pointer arithmetic in Go

& "addressof" operator

```go

firstName := "Me"
ptr := &firstName

```

ptr now holds address of firstName value

## Constants

const keyword

have to be initialized when declared
value has to be known at compile time

if you don't declare the type, Go will implicity define it when it's used. int can be used with float arithmetic for example

```go
  const pi = 3.14


```

### IOTA

definition has to be outside function

can create const block at start

```go

const (
  a = 3
  b = 5
)

```

you can use iota as a value, everytime it is used it is incremented
if used in a const block values don't have to be assigned, previous value will be used and incremented iota

```go

const (
  a = iota + 6 // 6
  b // 7
  )
```

iota resets between constant blocks

## Collections

### Arrays

fixed sized collection of similar data types.
Once size is set, it can't be changed without creating a new array

```go

var arr [3]int
arr[0] = 1
//...

arr := [3]int{1,2,3}

```

when passing an array to a function we need to specify the size of the array in the function definition

### Slices

Built on top of arrays, but slices are not fixed size!

```go

var arr []int // empty slice, size can change


```
Slice points to data that array is keeping so if you update a value in one it is reflected in the other.
A little bit like a pointer

Very similar to Python slicing syntax!
[begin:end]
begin is inclusive, end is exclusive
indexes are optional

```go

var arr [3]int
arr[0] = 1
arr[1] = 2
arr[2] = 3

slice := arr[:]

```

You can use the dynamic array initialization to reduce the boiler plate code

Note that the size is not defined. This tells the compiler that we want it to manage the size

```go

slice := []int{1,2,3}

slice = append(slice, 4, ...)

```

Underlying Go will handle array size. Once max size of the array is reached Go will create a new bigger array and copy all elements to it

length of a slice: len()
capacity of a slice: cap() - this counts from the start of the slice to the end of the array!

When passing a slice to a function you pass a pointer to the original array.

You can create a slice with make:

make([]int, 8, 5) - 8 is length, 5 is capacity


### Maps

```go

m := map[string]int{"foo":42} // create a map and initialize it with one entry
m := make(map[string]int) // create an empty map. it is equivalent to map[string]int{}

m["foo"] // if the value doesn't exist then the default for that type is returned or if it's assignment then it will be added

value,status = m["foo"] // status can be used to check whether the value is in the map or not

delete(m, "foo")

```

### range

easily iterate over an array, slice or map

for i, v := range s {
  i - index
  v - range
  are available in the the scope of each iteration
}

### Structs

Associate disparate data types together. Like a data object
fields are fixed at compile time
2 steps
  1 define struct
  2 use it

When initialized to a variable default (zero) values are used, boolean false, int 0 ...

use . operator to interact with fields of struct

```go
 type User struct {
   ID int
   FirstName string
   LastName string
 }

var u User
u.ID = 1
u.FirstName = arthur

u2 := User{
  ID: 1,
  FirstName: "Arthur",
  LastName: "Dent",
}
```
End struct definition with a , to allow you to put the } on the next line

GO has scope so if struct is defined within a function it is only available within that function

To define functions which act on the a struct they are defined outside a struct.

```go
func (user User) getFullName() string{
  return user.FirstName + " " + user.LastName
}

```

To update the struct in the function you need to pass a pointer to the function

```go

func (user *User) updateLastName( newLastName string){
  user.LastName = newLastName
}

```

Structs can embed other structs to make complext objects.

```go

type Name struct {
  firstName string
  lastName string
}

type User struct {
  name Name
}

user.name.firstName = "Tony"

```

You can also make the embedded struct anonymous

```go


type User struct {
  Name
}

user.firstName = "Tony"

```

When a struct is passed to a function it is passed by value, so any changes made are not reflected in the original struct
To update original a pointer to it must be passed around

pointer dereference can be explicit or implicit!

Another feature in structs is tagging each type in the struct.
Tags are annotations that appear after the type in a struct.

These tags are then examined by other Go code which can then react to them.
In the example below the tag is for json, so the json module will react to this.
In this case when marshalling the struct to json it will give the value for ItemID the key item_id in the json body

```go

type LineItem struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity uint `json:"quantity"`
	Price    uint `json:"price"`
}

```

## functions

func keyword, name of function, () {}

GO is like javascript in how it treats functions
a function name without the () is the function itself and it can be passed around the place

```go

func main(){
  myvar := ""
  mymethod(myvar)
}
func mymethod(inputvar string) bool { // bool is return type
  //...
  return true
}

```

If you have multiple function parameters of the same type you can omit the type from all but the last one and GO will know all have the same type

GO has an error type which can be returned from functions
GO does not have the concept of exceptions you can catch, instead error values are returned
Check for the presence of an error using the nil check

In Go you can return multiple values

```go

port, err := mymethod()

func mymethod() (int, error){
  // ...
  return 1, nil
}

```

To ignore one of multiple return values from a method use _ in the variables definition

```go

_, err := mymethod()

```

## methods

Object Oriented programming in GO is a bit different
1. Define a Struct
2. create methods and tie them to the Struct
    By specifying the struct as a receiver of the function
    directly after the keyword func

```go

type UserController struct {
  userIDPattern *regexp.Regexp
}

func (uc UserController) func_name (){

}
```


Initializing an object

me := BacnkAccount{}

or 

var me = new(BacnkAccount) - fields are initialized to zero


No constructor in GO, so use a special function
convention is New<Struct-name>

```go

func newUserController() *userController{
  return &userController{
    userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`)
  }
}

```

returns a pointer to the created object

function or method?
  if the logic only depends on the input parameters use a function
  if the logic depends on values that are changed while program is running

## goroutines

Lightweight threads in GO.
Use the keyword go and pass a function to it
goroutines are multipled onto multiple O/S threads
GO hides complexitie of thread creation and management

```go

func gope(){
  ...
}

func main(){
  go gope()
}

```

Can also pass a function literal (closure) to goroutine

```go

func main(){
  go func(n int){
    ...
  }(5)
}

```

We use wait groups to wait for all goroutines to finish

Create a wait group
specify the number of goroutines to wait for
add wg.Done as defer to each func
call wg.Wait()

```go

wg := new(sync.WaitGroup)
wg.Add(2)

```

Need to be careful with closures and variable states! Especially in loops which create goroutines.

Use channels to sync and share data across goroutines.
A channel can share one type of data.
Only one piece of data on the channel at a time
Nothing on it, read waits.
Something on it, write waits

To write/read we use the <- operator


```go

ch := make(chan int)

go func(){
  ch <- fmt.Errorf("Something went wrong") // write
  mine := <- ch // read
}

```

Only one go routine has access to a variable in the channel at a time.
No worrying about synchronzing!

If one is reading, any other has to wait until its completes before it gets its go

Multiple goroutines can read and write to a channel

There is a select keyword which allows a go routine to read from multiple channels.
  like switch, whichever channel has a value first is used
  If you don't want it to block specify a default case
  if you want it to timeout use a case with a time

### buffered channels

When creating the channel specify a size to it.
This allows multiple values to be placed in it at a time.
Buffered channel is same effect as semaphores (protect access to resources)

## context

Go has the concept of a context object which is passed around the place.
https://pkg.go.dev/context
It can be used for graceful shutdown to inform the application to finish requests before shutting down

## Struct embedding

This enables polymorphism in GO!!!

```go

type BankAccount struct{
  ...
}

type SavingsAccount struct {
  BankAccount
  ...
}

sa := SavingsAccount{ BankAccount{...}, ...}


```

sa now has access to all methods of BankAccount

## interfaces

Define your own interfaces and pass them around

```go

type Logger interface{
  Log(message string)
}

func DoSomething(log Logger){
  log.Log("...")
}

```
but where's the implementation?
It is defined like a method on a struct

1. define the interface methods
2. define the struct with the properties
3. define the methods which can be applied to the struct

```go

package main

import "fmt"

// 1. define interface methods
type Shape interface {
  area() float64
  perim() float64
}

// 2. define struct with properties
type Rectangle struct {
  width, height float64
}

// 3. define methods which can be applied to the struct. Once the methods match the interface it will work
func (r *Rectangle) area() float64 {
  return r.width * r.height
}

func (r *Rectangle) Perim() float64 {
  return 2 * ( r.width + r.height)
}

func main() {
  fmt.Println("Hello")

  r := rectangle{3, 4}
  fmt.Println("Area", r.area())
  processShape(r)
}

func processShape(shape Shape){
  shape.area()
}


```

interface types can be used as cases in switch!

Need to be careful with interface methods that have the same signature.
There is no type safety like in Java which fails if there are two interfaces in the hierarchy with the same signature

## loops

only one, for
same syntax as Java, only addition is you can leave out the  ( )
Like other languages you can omit parts of thhe for loop to simulate while loop
leave out all 3 parts to have an infinite loop.

break and continue keywords are present with same function as other languages.

You can give the break/continue a label to go to a certain point, like a go to in Java

## conditionals

if, same as Java, again ( ) are optional

You can declare a variable inside the condition.
These variables are visible only within the if, else

## switch

switch, case and default keywords
Same as java
Can put multiple values comma separated in one case

switch can also have no expression and each case has a condition

expression can be integer, float, string.

Go is the opposite of Java, there is no need for break as it does not fall through.
If you want fall through you use the keyword fallthrough