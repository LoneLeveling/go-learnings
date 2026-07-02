package main 
import(
	"fmt"
	"sync"
)
func main(){
	var wg sync.WaitGroup
	wg.Add(2)//creating a weight group for our gorotuine
	ch:=make(chan int, 3) //buffered channel 
	go sell(ch, &wg)
	wg.Wait()
}

func sell(ch chan int, s *sync.WaitGroup){
	//Below we sending 3 lines as buffer is 3 above
	ch<-10
	ch<-11
	ch<-12
	go buy(ch, s)
	fmt.Println("sent all the data to the channel") //and since all the processing will be done we will be calling the Done() method on the weight group
	s.Done()
}

//buy() method we receive values from the channel sell
func buy(ch chan int, s *sync.WaitGroup){
	fmt.Println("Waiting for data")
	fmt.Println("Received data:", <-ch)
	s.Done()
}