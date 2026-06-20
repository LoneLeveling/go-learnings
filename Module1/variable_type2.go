//variable Type using reflect.Type0f()
package main
import(
	"fmt"
	"reflect"
)

func main(){
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Abhi"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(66.24))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}
