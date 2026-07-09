package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Println("Hi there!!")

	//So if you see the output you will see "INFO" in output which is one our 7 log levels that this logger framwork 'logrus' supports.

	//Log levels are as following(The severity increases in order as u go down):
	/*
	   Trace
	   Debug
	   Info
	   Warn
	   Error
	   Fatal
	   Panic
	*/

	// If you want to print a FATAL log you can write:
	log.Fatal("Hello there!!")

	// If you want to print a Panic log you can write:
	log.Panic("Nothing to worry about")
	//Panic means something went unexpectedly wrong and it is just and excpetion that occured at Runtime.

	log.Warn("Nothin chill")

	//To Debug:
	log.SetLevel(log.DebugLevel)
	log.Debug("Hi")

	//To print Trace logs:
	log.SetLevel(log.TraceLevel)
	log.Trace("HI")
}
