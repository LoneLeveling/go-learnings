//Arithmetic Operators
package main 
import "fmt"
func main(){
	var a,b string = "John","Doe"
	fmt.Println(a+b)

	//fmt.Println(a-b) //This results in error as substarction operation is invalid on string 

var a1, b1 float64=89.2,33.4
fmt.Printf("%.2f\n",a1-b1)

var n1,n2 int=2,4
fmt.Println("Product of",n1,"&",n2,"is:",n1*n2)

var num1,num2 int =12,4
fmt.Println(num1/num2)
// NOTE: Dividend operator("/" return the quotient)

var num3,num4 int=24,9
fmt.Println(num3%num4)
//NOTE: Modulus(%) returns remainder

//Unary Operator- acts on single variable to produce new value so called unary
var i int=1
i++
fmt.Println(i)

var j int=4
j--
fmt.Println(j)


	var f1,f2 float64=24.4,3.0
	fmt.Println(f1/f2)
	fmt.Println(int(f1)%int(f2))
}