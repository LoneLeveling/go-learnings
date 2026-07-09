// Sorting in GO using that "sort" package.
package main

import (
	"fmt"
	"sort"
)

func main() {

	//Sorting an Integer slice
	vars := []int{5, 2, 0, 3, 6, 8, 6}
	sort.Ints(vars)
	fmt.Println(vars)

	//Sorting a String slice
	var1 := []string{"cat", "ball", "apple"}
	sort.Strings(var1)
	fmt.Println(var1)
}

//In termincal run: go doc sort , to check the docs for sorting package
