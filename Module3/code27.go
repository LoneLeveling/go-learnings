//Adding a new value to the map
package main
import "fmt"
func main(){
	codes:=map[string]int{"en":1,"fr":2,"hi":3}
	codes["mj"]=42
	fmt.Println(codes)

	//updating the value, inserting on existing key overrides the value
	codes["mj"]=88
	fmt.Println(codes)
}