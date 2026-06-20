//Multidimensional arrays
package main
import "fmt"
func main(){
	nums:=[3][2]int{{1,2},{4,5},{7,8}}
	// for index,element:=range nums{
	// 	fmt.Println(index,"==>",element)
	// }
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums[i]);j++{
			fmt.Println(nums[i][j])
		}
	}

	//Printing using range:
	for a,row:=range nums{
		for b,value:=range row{
			fmt.Println("(",a,",",b,")","==>",value)
		}
	}
}