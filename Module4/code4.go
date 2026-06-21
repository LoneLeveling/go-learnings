package main 
import "fmt"

func operation(a int, b int)(sum int, diff int){
	sum=a+b//note how we do not have to write the short hand annotation(:=) here since Go allows us to write var name in return type itself.
	diff=a-b
	return 
}

func main(){
	sum,difference:=operation(20,10)
	fmt.Println(sum," ",difference)
}