# writing go

https://blog.stackademic.com/best-practices-in-go-golang-writing-clean-efficient-and-maintainable-code-dccf61542b57

all go files start with a package declaration

func main(){} in the package main is the entry point

functions can return multiple values

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

comments in go

one line //
multi line /* */

## go.mod file

defines module and go version

go get to install modules for your project. Adds them to go.mod
However, just import them to your files then call "go mod tidy" and it will add them to your go.mod file.

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

### Slices

Built on top of arrays, but slices are not fixed size!

```go

var arr []int // empty slice, size can change


```
Slice points to data that array is keeping so if you update a value in one it is reflected in the other.
A little bit like a pointer

Very similar to Python slicing syntax!
[begin:end]
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

### Maps

```go

m := map[string]int{"foo":42} // create a map and initialize it with one entry
m := make(map[string]int) // create an empty map. it is equivalent to map[string]int{}

m["foo"] // if the value doesn't exist then the default for that type is returned.

value,status = m["foo"] // status can be used to check whether the value is in the map or not

delete(m, "foo")

```

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

```go

type UserController struct {
  userIDPattern *regexp.Regexp
}

func (uc UserController) func_name (){

}
```

No constructor in GO, so use a special function

```go

func newUserController() *userController{
  return &userController{
    userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`)
  }
}

```

## goroutines

Lightweight threads in GO.
Use the keyword go and pass a function to it

```go

func gope(){
  ...
}

func main(){
  go gope()
}

```

Use channels to sync and share data across goroutines.

```go

ch := make(chan int)

go func(){
  ch <- fmt.Errorf("Something went wrong")
}

```

## context

Go has the concept of a context object which is passed around the place.
https://pkg.go.dev/context
It can be used for graceful shutdown to inform the application to finish requests before shutting down

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
type shape interface {
  area() float64
  perim() float64
}

// 2. define struct with properties
type rectangle struct {
  width, height float64
}

// 3. define methods which can be applied to the struct
func (r rectangle) area() float64 {
  return r.width * r.height
}

func main() {
  fmt.Println("Hello")

  r := rectangle{3, 4}
  fmt.Println("Area", r.area())
}


```

I'm sure there must be some value to interfaces but I don't see it.
They don't appear to give any type safety?!?
