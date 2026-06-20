package main

import "fmt"

func main() {
        my_map := make([int]int)
        my_map[2] = 4
        my_map[4] = 16
        my_map[8] = 64
        delete(my_map, 4)
        fmt.Print(my_map)
}