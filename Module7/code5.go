//	The output of this code is random and not in sequence because all gorotuines run concurrently.
package main 
import (
"fmt"
"sync"
"time"
)

func calculateSquare(i int, wg *sync.WaitGroup){
	fmt.Println(i*i)
	wg.Done()//this will be called by each go routine so counter will reduce by 1
}

func main(){
	var wg sync.WaitGroup
	start:= time.Now()
	wg.Add(10)
	for i:=0;i<10;i++{
		go calculateSquare(i, &wg)
	}
	wg.Wait() //Blocking execution of main go routine until all go routines finishes executing.
	elapsed:=time.Since(start)
	fmt.Println("Function took", elapsed)
}
