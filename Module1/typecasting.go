package main 
import "fmt"
func main(){
	// Integer to float
	var i int=90
	var f float64=float64(i)
	fmt.Printf("%.2f\n",f)

	// float to Integer , but we loose precesion
var fl float64=45.89
var in int=int(fl)
fmt.Printf("%v\n",in)
} 