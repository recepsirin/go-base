package main

import (
    "fmt"
)

func main() {

    messages := make(chan string)

    go worker(messages, "dsadasd")
    go worker(messages, "asddsa")
    go worker(messages, "zxczx")
    go worker(messages, "wqfqw")

    for msg := range messages {
    	fmt.Println(msg)
     	// error: all goroutines are asleep - deadlock
    }

}

func worker(messages chan<- string, msg string){
    messages <- msg
}