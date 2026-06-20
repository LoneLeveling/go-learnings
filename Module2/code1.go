// Comparison Operators
package main
import "fmt"
func main(){
	var a int=10
	var b int=20
	fmt.Println(a!=b)

	var str1 string="John"
	var str2 string="John"
	fmt.Println(str1==str2)

	var a1,b1 int =5,10
	fmt.Println(a1<b1)

	var a2,b2 int =10,10
	fmt.Println(a2<=b2)

	var a3,b3 int=10,20
	fmt.Println(a3>b3)

	var a4,b4=40,60
	fmt.Println(a4>=b4)
}