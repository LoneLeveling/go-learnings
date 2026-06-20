// Logical Operator
package main
import "fmt"
func main(){
	var x int=10
	fmt.Println((x<100)&&(x<200))
	fmt.Println((x<300)&&(x<0))

	var y int=10
	fmt.Println((y<0)||(y<200))
	fmt.Println((y<0)||(y>200))

	var a,b int =10,20
	fmt.Println(!(a>b))
	fmt.Println(!(true))
	fmt.Println(!(false))

	var num bool=false
	result:=10>50 //false
	fmt.Println(!(num&&result))
}