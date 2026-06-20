package main
import "fmt"
func main(){
	var acii_codes map[string]int
	acii_codes=make(map[string]int)
	acii_codes["A"]=65
	fmt.Println(acii_codes)

	marks:=map[string]int{
		"A":69,
}
			fmt.Println(marks)
}