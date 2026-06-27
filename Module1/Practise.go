package main

import "fmt"

func main() {
    arr := [4]int{10, 20, 30, 40}
    i:=2//element that you want to delete
		fmt.Println("Original array:",arr)
		slice_1:=arr[:i]
		slice_2:=arr[i+1:]
		finalSlice:=append(slice_1,slice_2...)
		fmt.Println(finalSlice)
}