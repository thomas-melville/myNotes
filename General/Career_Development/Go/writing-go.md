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

nice compile time error if you declare a variable and don't use it
  simplicity, compiler helps us

defer keyword, after the main function exits execute command passed to it!


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

  w.Write([]byte("Hello " + name)

  err := http.ListenAndServe(":3000", nil)

  if err != nil {
    log.Fatal(err)
  }
})

```

Write works with bytes, so the string is written as an array of bytes
