package main

import(
  "fmt"
  "time"
)

func sayHello(){
  fmt.Println("Hello")
}
func main(){
  chan := make(chan int)
  go sayHello()
  go sayHello()

  time.Sleep(time.Second)
}
