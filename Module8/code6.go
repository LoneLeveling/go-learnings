package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 1.Join()
	path := filepath.Join("dir1", "dir2//", "dir3")
	fmt.Println(path)

	//Join also normalizes the path say if we had:
	path1 := filepath.Join("dir1", "dir2/../dir3", "file.txt")
	fmt.Println(path1)

	// 2.Dir(): Return the directory path
	fmt.Println(filepath.Dir(path1))

	// 3.Base(): Returns the last element of the path
	fmt.Println(filepath.Base(path1))

	//4. To check if a path is absolute: isAbs()
	fmt.Println(filepath.IsAbs(path1)) //Retuns false since path is not absolute => Path:dir1/dir3/file.txttrue
	// fmt.Print(path1)
	fmt.Println(filepath.IsAbs("/dir/file"))

	//5.Extension function-> Ext(), this func return the file name extension used by the path
	fmt.Println(filepath.Ext(path1))

	//6. Stat() -> we use this when we want to get metadata from a file
	fileInfo, err := os.Stat("/Users/abhisheksharma/Desktop/Golang/Module8/temp.txt")
	if err != nil /*meaning error occured*/ {
		fmt.Println(err)
	}
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.IsDir())

	//7. ReadFile() -> ReadFile reads the named file and returns the contents. We pass the file path as the string param to the ReadFile() func
	path3 := "/Users/abhisheksharma/Desktop/Golang/Module8/temp.txt"
	data, err := os.ReadFile(path3)
	if err != nil {
		fmt.Println(err)
	}
	//if everything goes well we just print the contents of the file:
	fmt.Println(data) //This prints all the ASCII characters of the file content, as we know "byte is nothing but an array of unsigned integers hence we get the ASCII codes printed in the output",
	//To get the string value out of the ASCII codes we use the string func()
	fmt.Println(string(data))

	//NOTE: Sometimes when the file size if high we need to read the file in chunks so for that we use, os package's Open() method
	//Run: go doc os Open, to see the function definiton , so it returns us the file kind of Pointer(*File) & error.

	path4 := "/Users/abhisheksharma/Desktop/Golang/Module8/temp.txt"
	file, err := os.Open(path4)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	b := make([]byte, 4) //creating a buffer
	for {
		n, err := file.Read(b)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println(b[0:n])
	}

}
