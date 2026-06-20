// Bitwise Operator
 package main
 import "fmt"
 func main(){
	var x,y int =12,25
	z:=x&y //"&:Bitwise AND operator takes 2 numbers as operands and does AND on every bit of two numbers"
fmt.Println(z) 

	var x1,y1 int =12,25
	z1:=x1|y1 //"|:Bitwise OR operator takes 2 numbers as operands and does OR on every bit of two numbers"
fmt.Println(z1) 


	var x2,y2 int =12,25
	z2:=x2^y2 
fmt.Println(z2)
//NOTE: The result of XOR is 1 if two bits are opposite

	var x3 int =212
	z3:=x3<<1 //"<<:Left shift operator: Shifts all bits towards left by a certain number of specifies bits.
fmt.Println(z3) 
//NOTE: The result of XOR is 1 if two bits are opposite


	var x4 int =212
	z4:=x4>>2 //"<<:Right shift operator: Shifts all bits towards right by a certain number of specifies bits.
fmt.Println(z4) 
//excess bits shifted to the right are discarded

 }