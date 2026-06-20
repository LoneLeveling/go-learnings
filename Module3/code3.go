package main
import "fmt"
func main(){
grades:=[...]int{88,69,78,70,90}
fmt.Println(grades)

//Iterating over the array 
for i:=0;i<len(grades);i++{
	fmt.Println(grades[i])
}
}