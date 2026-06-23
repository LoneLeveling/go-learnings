//Accessing fields: <variable_name>.<field_name>
package main 
import "fmt"

type Circle struct{
	x int
	y int
	radius int
}

func main(){
	var c Circle //creating a struct variale 
	c.x=5
	c.y=10
	c.radius=20
	fmt.Printf("%v \n",c)
}