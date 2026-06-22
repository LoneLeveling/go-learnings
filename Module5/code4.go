package main 
import "fmt"

func main(){
	a:="hello"
	fmt.Printf("%T %v \n", a,a) //This line to print the data type and the value of string s.
	ps:=&a //Creating a pointer 'ps' via short hand notation, which will store the address of variable s
	*ps="Priya" //using derefrencing operator to change the value at address: &a
	fmt.Printf("%T %v \n",a,a)//Printing the data type and value of a
}