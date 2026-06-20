// Truncate a map	

	package main
	import "fmt"
	func main(){
			codes:=map[string]string{"en":"English",
	"fr":"French", "hi":"Hindi"}
	//truncate a map i.e. deleting a map's conten
	//so 2 ways to do that: 
	//1. iterate over map and delete each key one by one via delete() keyword

	for key,_:=range codes{
		delete(codes,key)
	}
fmt.Println(codes)
	

	// 2.way: Re-initialize the map using the make function
	codes1:=map[string]string{"en":"English",
	"fr":"French", "hi":"Hindi"}
	codes1=make(map[string]string)
	fmt.Println(codes1)
}