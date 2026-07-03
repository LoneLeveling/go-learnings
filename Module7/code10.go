//Below is and unbuffered channel
package main 
import "fmt"
func main(){
	ch:=make(chan int)//Creates an unbuffered channel. capacity=0
	
	go func()/*this go routine(also called worker goroutine i.e., receiver) starts and it executes fmt.Println(<-ch), which then starts waiting to receive data and only 1 value*/{
	for val:=range ch{
		fmt.Println(val)
	}
	}()
	ch<- 100//Main goroutine executes, and there is a reciver waiting above.
	ch<- 200


}