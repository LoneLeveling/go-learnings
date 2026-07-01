package main 
import (
	"fmt"
	"time")
func main(){
go func(){
	fmt.Println("In anonymous method")
}()
time.Sleep(1*time.Second) //Adding 1 sec timer 
}