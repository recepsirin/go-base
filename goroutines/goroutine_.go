package main
 
import(
    "fmt"
    "time"
    )
 
func main(){
    msg:=make(chan string)
     
    go foo(msg, time.Now().String())
    go bar(msg)
     
    var userInput string
    fmt.Scanln(&userInput)
    fmt.Println("All is well")
}
 
func foo(channel chan string,content string){
    time.Sleep(time.Second*3)
    fmt.Println("Foo...")
    channel<-content
}
 
func bar(channel chan string){
    fmt.Println("Bar...")
    ctx:=<-channel
    fmt.Println(ctx)
}