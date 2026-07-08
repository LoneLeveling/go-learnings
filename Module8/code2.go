package main

//Important methods from String package
import (
	"fmt"
	"strings"
)

func main() {

	//Method 1.Contains()
	learning := "Learning standard library in Go"
	fun := "Library in Go"

	result := strings.Contains(learning, fun)
	fmt.Println(result)

	//Method 2. Count()
	str1 := "Hi learning this go is so cool so learning it"
	str2 := "learning"
	fmt.Println(strings.Count(str1, str2))

	//Method 3. ReplaceAll()
	str3 := "Hi there learning and learning is so cool"
	str4 := "learning"
	fmt.Println(strings.ReplaceAll(str3, str4, "grinding")) //Replace.All(original string, string containing word that you want to replace in original string, word you want to Replace with in original string)
}
