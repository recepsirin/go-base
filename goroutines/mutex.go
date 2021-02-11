package main

import (
    "fmt"
    "sync"
)

var (
    mutex   sync.Mutex // mutual exclusion 
    foobar int
    data int
)

func init() {
    foobar = 0
    data = 1
}

func foo(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    foobar += value
    mutex.Unlock()
    wg.Done()
}

func bar(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    foobar -= value
    mutex.Unlock()
    wg.Done()
}

func worker(value int, dest int,  wg *sync.WaitGroup) {
    mutex.Lock()
    for i := 0; i < dest; i++ {
        data *= value
    }
    mutex.Unlock()
    wg.Done()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

    go worker(100, 10, &wg)
    go foo(1000, &wg)
    go bar(500, &wg)

    wg.Wait()

    fmt.Println("Worker's Data", data)
    fmt.Println("Foobar's Data", foobar)
}