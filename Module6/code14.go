//Using method with structs
package main
import "fmt"

//Step1:Below we create a custom Struct of type Circle, every Circle object will have these 2 things: radius & area
type Circle struct{
	radius float64 //field
	area float64 //field
	perimeter float32//note this is not being used in our code it is taken just to understand that "%v" prints the fields of the struct, and if they are not assigned any value then they are assigned default value by go.
}
//Here above the struct Circle has 3 fields.


//Step2:This is a method, here the part right after "func" is called <receiver>
//so c *Circle means, method ko Circle ka address milega and inside the 
//function, c.area will modify the actual object.

// Pointer receiver:
// The method receives a pointer to a Circle.
// Any changes made through the receiver affect the original object.
func (c *Circle) calcArea(){
	c.area=3.14*c.radius*c.radius
}

// VV IMP NOTE: If above we had: func (c Circle) calcArea(){} i.e., 
//If the receiver for a method was an instance of circle struct i.e,. func (c Circle).
// rather than pointer then the value of area will be = 0 , i.e., value of original argument will not get modified.

//Step3: c:=Circle{radius:5}<- this means a Circle object is created, with 
// Below A Circle object is created.
// radius is explicitly initialized to 5.
// area and perimeter are not specified, so Go assigns their zero values.
func main(){
	c:=Circle{radius: 5}
	fmt.Println(c)//this line shows you the object of struct type i.e., c having values as 5,0,0 for radius, area & perimeter respectively.
//Step 4: method call to calcArea() & in step 6 the above method calcAreas current object is with radius=5 & area=0
	c.calcArea() // this line is same as writing:	(&c).calcArea(), because go inmternally makes c --> &c i.e., address
	//step 6: Print
	fmt.Printf("%+v\n", c)

}