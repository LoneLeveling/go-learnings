//Go routine leak
package main
import(
	"fmt"
	"sync"
)
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go leak(&wg)
	wg.Wait()
}

func leak(s *sync.WaitGroup){
	ch:=make(chan int)

	//creating Anonymous go routine
	go func(){
		val:=<-ch
		fmt.Println("Received ", val)
		s.Done()
	}()

	fmt.Println("Exiting leak method")
	s.Done()
}