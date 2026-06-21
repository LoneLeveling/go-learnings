package main
import "fmt"

//variadic function
func sumNumbers(number ...int) int{
	sum:=0
	for _,value:=range number{
		sum+=value
	}
	return sum
}

func main(){
	fmt.Println(sumNumbers())//no argument so we will get a empty slice.
	fmt.Println(sumNumbers(10))
	fmt.Println(sumNumbers(10,20))
	fmt.Println(sumNumbers(10,20,30,40,50,60))
}