package main

import "fmt"

func main() {
        ascii_codes := map[string]int{}
        ascii_codes["A"] = 65
        _, found := ascii_codes["B"]
        if found {
                fmt.Println("key B was not found")
        }
}