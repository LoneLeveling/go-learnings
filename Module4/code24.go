
//Defer statement: Delays the execution of a function until 
//the surrounding function returns.
package main 
import "fmt"
func printName(str string){
	fmt.Println(str)
}

func printRollNo(rno int){
	fmt.Println(rno)
}

func printAddress(adr string){
	fmt.Println(adr)
}

func main(){
	printName("joe")
	defer printRollNo(23)
	printAddress("LA WEST-22")
}