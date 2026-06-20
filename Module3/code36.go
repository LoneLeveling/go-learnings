package main
import "fmt"
func main(){
	arr:=[5]int{}
	my_map:=make(map[string]int)
	my_map["A"]=65
	my_map["B"]=66
	fmt.Println(my_map)
	i:=0
	for _,value:=range my_map{
		arr[i]=value
		i+=1
	}
	fmt.Println(arr)
}