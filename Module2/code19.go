package main 
import "fmt"
 func main(){
	n,err:=fmt.Println("sp")
	fmt.Println("Bytes written:",n) //Output = 3 because "sp"=> 's','p','\n'
	fmt.Println("Error:",err)

	//In "fmt" package the return type of Println is 2 values "n":= number of bytes written & err = error object, this tells if printing was successful or not
 }