//Iterating over an array using "range" keyword, 
//This is used to iterate over all the elements of an array, slice or map.

package main
import "fmt"
func main(){
	grades:=[...]int{98,39,56,78,86}
	for i:=0;i<len(grades);i++{
		fmt.Println(grades[i])
	}

	//Iterating Using "range"
for index,element:=range grades{
	fmt.Println(index,"==>",element)
}
}