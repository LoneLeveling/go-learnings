//Assignment Operators
package main
import "fmt"
func main(){
var x int=10
var y int
y=x //Assigning left operand(y) with the value to the right
fmt.Println(y)


var x1,y1 int=10,20
x1+=y1 //meaning: x1=x1+y1
fmt.Println("x1:",x1)


var x2,y2 int=10,20
x2-=y2 //meaning: x2=x2+y2
fmt.Println("x2:",x2)

var x3,y3 int=10,20
x3*=y3 //meaning: x3=x3*y3
fmt.Println(x3)

var x4,y4 int =200,20
x4/=y4 //Meaning: Assigns left operand with the quotient of the divison
//Above means: x4=x4/y4
fmt.Println(x4)

var x5,y5 int =200,20
x5%=y5 //Meaning: Assigns left operand with the remainder of the divison
//Above means: x5=x5%y5
fmt.Println(x5)
}