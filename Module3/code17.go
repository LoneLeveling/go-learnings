package main 
import "fmt"
func main(){
	src_slice:=[]int{10,20,30,40,50}
	dest_slice:=make([]int,3)
	//Above destination slice capacity = 3 so only 3 number are copied from the source slice.
	num:=copy(dest_slice,src_slice)

fmt.Println(dest_slice)
fmt.Println("Number of elements copied: ",num)
}