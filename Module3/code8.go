package main 
import "fmt"
func main(){
	arr:=[5]string{"a","b","c"}
	for index,element:=range arr{
		fmt.Println(index,"==>",element)
	}
}