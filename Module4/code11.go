//Guess the output?? 
// O/p:
// Hey there
// Joe, Monica, Gunther, 

package main 
import "fmt"

func printStrings(s string, names ...string) {
	fmt.Println(s)
	for _, value := range names {
			fmt.Printf("%s, ", value)
	}

}

func main() {
	printStrings("Hey there", "Joe", "Monica", "Gunther")
}