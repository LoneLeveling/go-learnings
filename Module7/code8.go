//Closing a Channel: - when ok= true it means channel is open and when its false it means channel is closed and there are no more values to be received.
package main 
import("fmt")
func main(){
	ch:=make(chan int, 10)
	ch <-10
	ch <-11
	val,ok:=<-ch //ok = a boolean value to check if a channel is open or closed.
	fmt.Println(val, ok)
	close(ch)
	val,ok = <-ch //note how here we are not using the short assignment operator(:=) which create a variables and assigns a value, again as we already declared the val & ok variable above.
	fmt.Println(val, ok)
	val,ok=<-ch
	fmt.Println(val,ok)
	close(ch)
}