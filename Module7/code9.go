// In this code we have 2 channels: buy & sell
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) //creating a weight group for our gorotuine
	ch := make(chan int)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, wg *sync.WaitGroup) {
	//Below we sending 3 lines as buffer is 3 above
	ch <- 1
	ch <- 2
	fmt.Println("Sent all data")
	//close(ch) /*commenting this line i.e., if we do not close the channel we end up creating a deadlock situation, this happens bcz for range is never going to finish until the channel is closed and hence it creates a deadlock situation for us*/<--
	// IMP NOTE: Hence when you are using for range while iterating over the values received from a channel, make sure to close the channel.
	wg.Done()
}

// buy() method we receive values from the channel sell
func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")
	//Using for range below to receive the values
	for val := range ch /*channel name*/ {
		fmt.Println("Received data:", val)
	}
	wg.Done()
}
