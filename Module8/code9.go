//Logging in golang
//In go lang we use inbuilt log package and a third party package (logging framework), 2 most famous ones: glog & logrus, for logging.
//We will be using logrus as it is better maintained and it is used in popular projects like Docker.
//What to log?
// - Timestamp for when an event occured, or a log was generated.
// - Log levels such as devug, error or info.
//And some contextual data that can later on help us recreate the same issue.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//We can also log to a file, we know we can use the "os" package to write to a file.
	file, err := os.OpenFile("/Users/abhisheksharma/Desktop/Golang/Module8/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	//this "file" var above is a file pointer (*File), you can check that running: go doc os OpenFile in terminal.
	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(file)
	log.Println("Hi there!!") //so this gives us : date time output

}

//see code10.log where we installed 3rd party logging framework(logrus) for logging.
