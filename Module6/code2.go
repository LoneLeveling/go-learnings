package main 
import "fmt"

type Student struct{
	name string
	rollNo int
}

func main(){
	st:=Student{
	name:"Joe",
	rollNo:12,
}
//Above st part can also be writeen like below:
sp:=Student{"Priya",22}
fmt.Println("%+v",st)
fmt.Println(sp)
}