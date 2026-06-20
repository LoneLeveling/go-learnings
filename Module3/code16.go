//Deleting from a slice
package main 
import "fmt"
func main(){
	//Objective: Delete element at index = 2
	arr:=[5]int{10,20,30,40,50}
i:=2 //Initializing i with 2
fmt.Println(arr)
slice_1:=arr[:i]//printing elements from o till i-1 
slice_2:=arr[i+1:]//printing elements from i+1 till last index
final_slice:=append(slice_1,slice_2...)
fmt.Println(final_slice)
}