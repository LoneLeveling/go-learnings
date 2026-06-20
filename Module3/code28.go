//Delete a key-value pair from map
package main
import "fmt"
func main(){
	codes:=map[string]string{"en":"English",
	"fr":"French", "hi":"Hindi"}

	fmt.Println(codes)
	delete(codes,"en")
	fmt.Println(codes)
}