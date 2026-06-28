package main 
import "fmt"
func main(){
asii_codes:=make(map[string]int)
asii_codes["A"]=65
asii_codes["F"]=70
fmt.Println(asii_codes)
modify(asii_codes)
fmt.Println(asii_codes)
}

func modify(m map[string]int) {
	// fmt.Println(s)
m["k"]=90
}