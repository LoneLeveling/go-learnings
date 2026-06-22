//Guess the output ?  O/P: Error
// Why err? Since the function has to return type it return nothing
// and we are trying to store that nothing inside a variable x hence we get error.
//So no result value, so compiler complains.
package main

import (
        "fmt"
        "strings"
)

func main() {
        x := func(s string) {
                fmt.Println(strings.ToLower(s))
        }("RacheL")
        fmt.Printf("%T \n", x)
}