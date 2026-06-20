package main
import "fmt"
func main(){
	arr:=[5]int{10,20,30,40,50}
	slice:=arr[:2] //[10,20]
	fmt.Println(slice)
	arr_2:=[5]int{5,15,25,35,45}
	slice_2:=arr_2[:2]
	fmt.Println(slice_2)

	new_slice:=append(slice,slice_2...) //These 3 dots are used for variadic functions:functions that can take arbitrary number of arguments
	fmt.Println(new_slice)
}