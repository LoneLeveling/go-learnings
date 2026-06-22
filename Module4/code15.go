//Anonymous Functions
package main
import "fmt"

func main(){

// x:=func(a int, b int) int{
// 	return a*b
// }

//Above function we can re-write like below without the need to store function in another var
// This is an Immediately Invoked Function Expression (IIFE).
x:=func(a int, b int)int{
	return a*b
}(20,30)

fmt.Printf("%T \n",x) //Printing the data type of x variable
fmt.Println(x) //passing args into anonymous functions
}