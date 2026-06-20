package main
import "fmt"
func main(){
var a string
var b int
fmt.Print("Enter a string and a number:")
count,err:=fmt.Scanf("%s %d", &a, &b)
if(err!=nil){
	fmt.Println("ERROR OCCURED!!")
}
fmt.Println("count: ",count)
fmt.Println("error: ",err)
fmt.Println("a: ",a)
fmt.Println("b: ",b)
}