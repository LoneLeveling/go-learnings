package main
import "fmt"
func f()(int, int){
return 42,54
}

func main(){
// a,b:=f()
v,_:=f()
//Now if you just want the 1st variable output then we
//use blank identifier '_', which acts like anonymous place holders instead
//of declaring another variable which we do not actually need.
fmt.Println(v)
}
