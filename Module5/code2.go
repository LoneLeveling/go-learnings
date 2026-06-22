package main 
import "fmt"
func main(){
	var i *int //declaring a pointer "i" of type int, NOTE: Do not confuse "*" with derefrence operator
	var s *string//declaring a pointer "s" of type string.
	fmt.Println(i)
	fmt.Println(s)

	//RMBR: Output of this code: 
	///nil since "0 value of a pointer is nil, rmrbr that."
}