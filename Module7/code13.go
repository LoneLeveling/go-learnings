//Select in goroutines, its like switch statement of goroutines.
package main 
import(
	"fmt"
	"time"
)
func main(){
	ch1:=make(chan string)
	ch2:=make(chan string)

	//Below we have 2 goroutines
	go goOne(ch1)
	go goTwo(ch2)

	select {
	case val1:= <-ch1:
		 fmt.Println(val1)
	case val2:= <-ch2:
		fmt.Println(val2)
	}

	time.Sleep(1*time.Second)
}

func goOne(ch1 chan string){
 ch1 <- "channel-1" //here in this line the goroutine is sending data to the channel
}
func goTwo(ch2 chan string){
 ch2 <- "channel-2"
}