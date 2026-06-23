package main 
import "fmt"

type Student struct{
	name string
	rollNo int
	marks []int
	grades map[string]int
}

func main(){
	var s Student
	fmt.Printf("%+v \n",s)
	st:=new (Student)//we initialize a struct via new keyword but most commonly we initialize via short hand notation, see :code2.go
		fmt.Printf("%+v",st)

}