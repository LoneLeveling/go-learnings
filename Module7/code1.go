//writing squares of numbers from 1 to 10,000

package main 
import (
	"fmt"
	"time"
)
func main(){
	start:=time.Now() // This line is to calculate the total time this code is going to take to complete.
	//The Now() method from the 'time' package is going to return us the current local time.
	for i :=1;i<=10000;i++{
		calculateSquare(i)
	}

	//Here at end we calculate the time elapsed while calculating all the squares so that gives us the total time taken by code.
	elapsed:=time.Since(start) //Since() method is from time package that returns the time that has elapsed since the time 't', where t is the argument we pass.
	fmt.Println("Function took  ", elapsed)
}

func calculateSquare(i int){
	time.Sleep(1*time.Second)	//Here we added a timeout which is going to sleep for 1 sec & will then print the result
	fmt.Println(i*i)
}

//NOTE: You will see that this code is going to calculate the squares 
//one by one, and since above line 21 have sleep time of 1 sec, we can assume that for the entire program to complete its going to take 10,000 seconds, 
//**Code2.go shows how we can implement the same using go routines.