// - name of the function: `returnCube`
// - input parameters: `int`
// - output parameters: `int`
// the function calculates the cube of the number and returns it.
// (cube of n = n*n*n)

package main
import "fmt"
func returnCube(n int) int{
	return n*n*n
}

func main(){
fmt.Println(returnCube(2))
}