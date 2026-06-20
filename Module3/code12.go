package main 
import "fmt"
func main(){
	arr:=[10]int{10,20,30,40,50,60,70,80,90,100}
	slice:=arr[1:8]
	fmt.Println(cap(arr))
	fmt.Println(cap(slice))
//capactiy of Array = length of the array
//capacity of slice=capacity of slice from start index till last index of array,
//so here as startIndex=1 and Array capacity=10 so slice capacity=9
}