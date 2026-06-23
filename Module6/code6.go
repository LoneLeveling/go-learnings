//Comparing structs
package main
import "fmt"

type s1 struct{
x int
}

type s2 struct{
	x int
}

func main(){
	c:=s1{x:5}
	c1:=s2{x:5}
	if c==c1{
		fmt.Println("yes same")
	}
}

// O/P : we get compile time error, because here the struct instances are not of the same type, i.e., one is of s1 and othe of s2 type and 
//we cannot compare structs of different types else we get CTE.

//See code7.go, where we have the same struct type.