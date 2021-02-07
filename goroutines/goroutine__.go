package main

import (
   "fmt"
   "time"
)

func foo() {
  fmt.Println("foo", time.Now())
}


func bar() {
  fmt.Println("bar", time.Now())
}

func main() {

  for i := 0; i < 10; i++{
  	
  	go foo()
	go bar()
	time.Sleep(time.Millisecond)

  }

}