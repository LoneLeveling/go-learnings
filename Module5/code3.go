package main
import "fmt"

func main(){
s:="Hey bud"

//Different ways to initialize a pointer
var b *string =&s
fmt.Println(b)

var a=&s
fmt.Println(a)

c:=&s
fmt.Println(c)

//NOTE HOW ALL 3 ABOVE RETURN THE SAME ADDRESS IN THE OUTPUT, 
//that's what pointers are meant for: to store the memory address of anothe variable.

fmt.Println("Value before derefrecing:", s)

//Now if we do defrencing --> "*" :
*c="hello"
fmt.Println("Value after derefrecing: ", s)
}