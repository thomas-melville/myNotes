package main

import(
  "fmt"
  "github.com/pkg/errors"
)

func throwError(inp int){
  if int < 5 {
    return errors.Wrap("Give me more")
  }
}

func main(){
  throwError(6)
  throwError(4)
}
