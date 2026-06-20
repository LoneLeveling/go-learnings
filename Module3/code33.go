// - create a map using make() function with key data type as string, and value data type as int.
// - add the following key_value pairs to it
// ("A", 65)
// ("F", 70)
// ("K", 75)
// - delete the key "F"
// - print the map

package main
import "fmt"
func main(){
	ascii_codes:=make(map[string]int)
	ascii_codes["A"]=65
	ascii_codes["B"]=70
	ascii_codes["F"]=75
delete(ascii_codes,"F")
fmt.Println(ascii_codes)
}