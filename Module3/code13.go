package main 
import "fmt"
func main(){
	arr:=[5]int{10,20,30,40,50}
	slice:=arr[:3]
	fmt.Println(arr)
	fmt.Println(slice)
	slice[1]=9000
	fmt.Println("after modification")
	fmt.Println(arr)
	fmt.Println(slice)
	slice=append(slice,10,20,30)
	fmt.Println(slice)
}