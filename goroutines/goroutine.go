package main 
import (   
    "fmt"
    "time"
)

func main() {
   
   go task(1)
   go task(2)

   time.Sleep(1 * time.Second)

}

func task(msg int) {
    for i := 0; i < 10; i++ {
        fmt.Println("Task ", msg)
    }
}
