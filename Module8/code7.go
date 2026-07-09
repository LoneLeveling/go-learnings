// Writing to an already existing file and if file does not exit creating a new one and writing to it.
package main

import (
	"fmt"
	"os"
)

func main() {
	//Step1: Opening the file 'temp.txt' using OpenFile() func and we passed few flags using 'OR' operator because we want to append to the file, we provided the CREATE flag incase if the file is not already present then it will be created & then the 'WRONLY'; write only flag which specifies the right only mode.
	//And the permission is 600, Gave Read and Write permission only to the User.
	file, err := os.OpenFile("/Users/abhisheksharma/Desktop/Golang/Module8/temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600) //In terminal run: go docs os O_RDONLY

	//Step2: Checking if any error in opening the file or not.
	if err != nil {
		fmt.Println(err)
		return
	}
	//Step3: We delay the closing of the file using defer keyword until the end of the Program.
	defer file.Close() //this close func can be used on a file pointer, bcz if you run: go doc os OpenFile in terminal you can see the return type of OpenFile() its a pointer(*File) & an error.

	//Step4: Writing(appending) to the file using: WriteString() func to write to the file which is openend by the OpenFile() method.
	_, err = file.WriteString("Keep grinding mate that's all that matters")

	//Step5: Checking if we had any error in writing to the file.
	if err != nil {
		fmt.Println(err)
	}
}
