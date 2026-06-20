//Creating a MAP
package main 
import "fmt"
func main(){
	codes:=map[string]string{"en":"English","fr":"french"}
	fmt.Println(codes)
	fmt.Println(len(codes))
	fmt.Println(codes["fr"])
}