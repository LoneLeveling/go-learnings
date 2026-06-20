//Iterating over a map
package main 
import "fmt"
func main(){
		codes:=map[string]string{"en":"English",
	"fr":"French", "hi":"Hindi"}
	
	for key,value:=range codes{
		fmt.Println(key, "=>", value)
	}

}