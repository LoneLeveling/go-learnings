package main 
import "fmt"
func main(){
	for i:=1;i<=5;i++{
		if i==3{
			break //break immediately ends the loop when its encountered
		}
		fmt.Println(i)
	}

	for j:=1;j<=5;j++{
		if j==3{
		continue //continue skips the current iteration of loop and continues with the next iteration.
	}
	fmt.Println(j)
}
}