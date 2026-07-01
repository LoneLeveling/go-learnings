//Implementing code using goroutines to print squares from 1 to 10,000
//Note: For a method to be executed in a go-routine we just need to add the "go" keyword before it.
//so we took the code from code1.go below and just added "go" keyword at the begining of the method call.

package main 
import (
	"fmt"
	"time"
)
func main(){
	start:=time.Now() 
	for i :=1;i<=10000;i++{
	go calculateSquare(i) //So putting here go means all these 10,000 function calls are going to be executed in 10,000 different go routines.
	}

	elapsed:=time.Since(start)
	time.Sleep(2*time.Second) //adding a timer of 1 or 2 seconds in the main func itself so that it does not finish running before the 10,000 threads executes there result
	//Above line means the main() func is going to wait for 2 seconds before it exits.
	fmt.Println("Function took  ", elapsed)
}

func calculateSquare(i int){
	time.Sleep(1*time.Second)	
	fmt.Println(i*i)
}
//So using go routines all the  10,000 squares gets calculated and printed concurrently,
//Now, there's also a better and a more reliable way to do this, which is using "waitgroups".