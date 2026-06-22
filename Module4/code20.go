//High order Functions in GO: function that receives a function as an argument
//or return a function as output.
package main
import "fmt"

func calcArea(r float64) float64{
	return 3.14*r*r
}

func calcPerimeter(r float64) float64{
	return 2*3.14*r
}

func calcDiameter(r float64) float64{
	return 2*r
}

// Higher Order Function #1
// Accepts a function as an argument
func printResults(radius float64, calcFunction func(r float64) float64){
	result:=calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thanks you!")
}


// Higher Order Function #2
// Returns a function
func getFunction(query int) func(r float64) float64{
	query_to_func:=map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}

f,ok:=query_to_func[query]
if !ok{
	return nil
}
return f
}





//Taking inputs from the user
func main(){
var query int
var radius float64

fmt.Print("Enter the radius of the circle: ")
fmt.Scanf("%f\n", &radius)
fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: ")
fmt.Scanf("%d\n", &query)

// if query==1{
// 	fmt.Println("Result: ", calcArea(radius))
// }else if query==2{
// 	fmt.Println("Result: ",calcPerimeter(radius))
// }else if query==3{
// 	fmt.Println("Result:", calcDiameter(radius))
// }else{
// 	fmt.Println("Invalid query Input!!")
// }

//All the above commented "if's" part is replaced by below code:

f:=getFunction(query)
if f==nil{
	fmt.Println("Invalid query Input!!")
	return
}
printResults(radius, getFunction(query))
}