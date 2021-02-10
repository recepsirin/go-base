package main

import (
	"fmt"
	"time"
)


func worker(done chan bool) {

	fmt.Println("working...")
	time.Sleep(time.Second * 10)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 3)

	go worker(done)

	messages := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		messages <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		messages <- 2
	}() 
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 3
	}()
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 4
	}()
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 5
	}()

	go func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second * 5)

	<-done

}
