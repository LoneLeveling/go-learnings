package main
import(
	"fmt"
	"strconv"
)

func main(){
	// Integer to String conversion
	var i int=42
	var s string = strconv.Itoa(i) //convert int to string
	fmt.Printf("%q",s)

	// String to Integer conversion
	var str string="200"
	num,err:=strconv.Atoi(str) //storing return values return by Atoi function and below printing the same
fmt.Printf("%v, %T\n",num,num)
fmt.Printf("%v, %T",err,err)

//Non int String to Integer, seeing the error occuring
var str1 string="200abhi"
num1,err:=strconv.Atoi(str1)
fmt.Printf("%v,%T\n",num1,num1)
fmt.Printf("%v,%T",err,err)
}