package main 
import "fmt"

//Below is a variadic function that takes in 2 kinds of input 
//1. string & 2. variadic parameter of string type, so 
//this means one string is mandatory i.e. one for student and for the other one 
//which is our variadic input i.e., subjects we can have 0 or more than 1 strings for our input.
func printDetails(student string, subjects ...string){
fmt.Println("hey", student,"here are your subjects-")
for _,sub:=range subjects{
fmt.Println("%s, ", sub)
}
}


func main(){
printDetails("Joe", "Physics", "CS", "Maths")
}