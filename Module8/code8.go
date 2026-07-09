// Error Handling in go
// NOTE: When an error occurs the execution of the program stops completely with a built-in-error message, hence its important to handle those errors in our program.
// Unlike in other languages we do not use try-catch blocks in go.
// In Go we use New() & Errorf() function to create our own and handle errors.
// In "Errors" package of go we have functions to manipulate errors. Run: go doc errors New , in terminal to see method details of errors package in GO.
package main

import (
	"fmt"
)

func main() {
	// err := errors.New("Custome error occured")
	// fmt.Println(err)
	err := checkNum(3) //passing odd number to checkNum
	checkError(err)    //passing the error received from checkNum

	err = checkNum(4) //passing even number to checkNum
	checkError(err)   //passing the error received from checkNum
}

// func checkNum(i int) error /*checks if num is odd or Even*/ {
// 	if i%2 == 0 {
// 		return errors.New("Only odd Numbers are allowed In here!!")
// 	}
// 	return nil
// }

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return //If we find any error we print the error and we return from the function.
	} else {
		fmt.Println("No error all good!!")
	}
}

// No.2:go doc fmt Errorf
//This is error formating function available in the fmt package.

func checkNum(i int) error /*checks if num is odd or Even*/ {
	if i%2 == 0 {
		return fmt.Errorf("Only odd Numbers are allowed In here, got:%d", i)
	}
	return nil
}
