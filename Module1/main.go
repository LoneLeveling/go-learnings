package main
import "fmt"
var dest string="online"
const PI float64=3.14
func main(){
var radius float64=5.154
var area float64
area=PI*radius*radius
fmt.Printf("Radius: %.2f \n PI:%.2f \n",radius,PI)
fmt.Printf("Area of circle: %f",area)
fmt.Println("Thank you")
// var s,t string,int="Priya",90
// fmt.Println(s," ,",t)
var name string
fmt.Print("Enter your name:")
fmt.Scanf("%s", &name)
fmt.Println("Hey, ",name)
// name:="Joe"
// i:=89
// fmt.Printf("Hey, %v! %T You have scored %d/100 in Maths \n",name,name,i)
// var s,t string="Joe","Bites"
// fmt.Println(s,t)
// city:="SF"
// {
// 	country:="US"
// 	fmt.Println(city," ",country)
// }
// fmt.Println(city)
// name,city:="Miky","SF"
// fmt.Println(name,city)
// fmt.Println(dest)

// var str string
// fmt.Printf("%s",str)
// fmt.Println("------------")
}