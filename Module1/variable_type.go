package main
import "fmt"
func main(){
	var grades int=22
	var message string="Hi there"
	var isCheck bool=true
	var amount float32=45.2624
	fmt.Printf("variable grades= %v is of type %T \n", grades, grades)
	fmt.Printf("variable message=%v is of type %T \n", message, message)
	fmt.Printf("variable isCheck=%v is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount= %v is of type %T \n", amount, amount)
}

// %v: to print the variable
// %T to print the data type of the variable