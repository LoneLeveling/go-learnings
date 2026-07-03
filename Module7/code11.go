package main 
import (
	"fmt"
	"sync")

func receiver(ch chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for val:=range ch{
		fmt.Println(val)
	}
}

func main(){
	ch:=make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go receiver(ch, &wg)//receiver starts here
	ch<-100
	ch<-200
	ch<-300

	close(ch)//RMBR: channel must always be closed in sender and never in receiver , because only the sender knows when it's done producing values.
	wg.Wait()
}