package main
import "fmt"
func main(){
	arr:=[]int{10,20,30,40,50}
	for index,value:=range arr{
		fmt.Println(index,"=>",value)
	}

	//looping without the need of index, relacing index with underscore
	for _,value:=range arr{
		fmt.Println(value)
	}
}