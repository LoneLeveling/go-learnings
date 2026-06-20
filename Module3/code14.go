package main
import "fmt"
func main(){
arr:=[4]int{10,20,30,40}
slice:=arr[1:3]
fmt.Println(slice)
fmt.Println(len(slice))
 fmt.Println(cap(slice))

 slice=append(slice,-90,899,900)
 fmt.Println(slice)
 fmt.Println(len(slice))
 fmt.Println(cap(slice))
//Imp Note: If element count exceeds the slice capacity then slice 
//capacity is doubled.
//Same as her in above example/; cap(slice) was =3 and but we 
//exceeded the append limit of slice capacity so the new slice capacity became 2*3=6
}

