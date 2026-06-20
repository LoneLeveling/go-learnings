package main 
import "fmt"
func main(){
	arr:=[5]int{10,20,90,70,60}
	slice:=arr[:3]
	slice[2]=900
	fmt.Println(arr)
	fmt.Println(slice)
}