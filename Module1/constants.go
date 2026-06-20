package main
import "fmt"
func main(){
	const name="John Doe"
	const is_muggle=true
	const age=26
	fmt.Printf("%v: %T\n",name,name)
	fmt.Printf("%v: %T\n",is_muggle, is_muggle)
	fmt.Printf("%v: %T\n",age,age)
	
}