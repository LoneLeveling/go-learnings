package main 
import "fmt"
func main(){
	i:=10
	fmt.Printf("%T %v \n", &i, &i)//Note the output of "%T", i.e. data type of &i = *int, which is basically how the pointers are declared in golang And rmbr this * is different from the dereference operator
	fmt.Printf("%T %v \n", *(&i), *(&i)) //de-  refencing the address using the dere ference operator.
  //Here above "%T" for *(&i) gives int because the dereference operator gives us the value at an address and the value here is an int type so we get "int" in the output.
}