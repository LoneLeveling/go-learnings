//Guess the output ? O/P:[]

package main 
import "fmt"

func calcSquare(numbers []int) []int{
	squares:=[]int{}
	for _,v:=range numbers{
		squares=append(squares, v*v)
	}

return squares
}


func main(){
	nums:=[3]int{10,20,15}
	fmt.Println(calcSquare(nums[:]))
}