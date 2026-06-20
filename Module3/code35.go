package main
import "fmt"
func main(){
	arr:=[5]string{"one","two","three"}
	slice:=arr[:3]
	my_map:=make(map[int]string)
	for i,el:=range slice{
		my_map[i+1]=el
	}
	fmt.Println(my_map)
}