//This program shows how all go routines exist as an independent execution, 
//and that there is no parent-child relationship between them.
//We create 2 methods: start() & process(), and we call the start method in a goroutine from the main function.
//Note that we added a timer of 1 sec in main() to ensure that our main func does not exit before all our goroutines complete.
package main 
import
( "fmt"
"time")
func main(){
go start()
time.Sleep(1*time.Second)
}

func start(){
go process()
	fmt.Println("In start")
}

func process(){
	fmt.Println("In process..")
}

//Final conclusion: If you run the code different times, 
// you will get different output, basically the output is non-deterministic, 
//because sometimes "In process" is printed 1st and sometimes its printed after "In start", which shows that 
//goroutines do not have parents or children, and they exist as independent executions.