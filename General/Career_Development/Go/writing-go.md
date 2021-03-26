# writing go

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

comments in go

one line //
multi line /* \*/

## go.mod file

defines module and go version

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
