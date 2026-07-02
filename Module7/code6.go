package main 

import(
	"fmt"
	"time"
)
func main(){
	ch:=make(chan string)

	go sell(ch)
	go buy(ch)
	time.Sleep(2*time.Second) //adding timer of 2 sec so that main() does not exit
}

//Below we create 2 methods sell and buy, which we call above as go routines.

//sell() sends data to the channel
func sell(ch chan string){
ch<-"Furniture"
fmt.Println("sending data to the channel...")
}

//receives data from the channel
func buy(ch chan string){
fmt.Println("waiting for data...")
val:=<-ch
fmt.Println("Received data =",val)
}

// NOTE: A send channel on an unbuffered channel is blocked until a receive happens on that channel is some other goroutine
//and vice versa i.e., the receive on another channel is also blocked unless there another goroutine to send to that channel.