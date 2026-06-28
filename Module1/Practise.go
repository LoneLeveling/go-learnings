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

func printResult(radius float64, calcFunction func(r float64) float64){
	result:=calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you!")
}

func getFunction(query int) func (r float64) float64{ //"getFunction takes an integer and returns another function that takes a float64 and returns a float64."
	query_to_func:=map[int]func(r float64)float64{
		1:calcArea,
		2:calcPerimeter,
		3:calcDiameter,
	}
	return query_to_func[query]
}
func main(){
var query int
var radius float64
fmt.Println("Enter the radius of the circle: ")
fmt.Scanf("%f\n",&radius)
fmt.Println("Enter \n 1 - area \n 2-Perimeter \n 3-Diameter")
fmt.Scanf("%d\n",&query)

printResult(radius, getFunction(query))
}