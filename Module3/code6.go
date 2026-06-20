package main
import "fmt"
func main(){
	arr:=[5]bool{true,true,true,true}
	for i:=0;i<len(arr);i++{
		if arr[i]{
			fmt.Println(i)
	}
}
}