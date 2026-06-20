//Array Initialization
package main
import "fmt"
func main(){
	var grades[3] int=[3] int {10,20,30}
	gradez:=[...]int{10,20,30}
	names:=[...]string{"roa","lok","susi"}
	fmt.Println(grades)
	fmt.Println(gradez)
	fmt.Println(len(names))
	fmt.Println(names[1])

	var fruits[5] string=[5]string{"apples","oranges","kiwi","banana"}
	fmt.Println(fruits)
	fmt.Println(fruits[2])
}
