package main

import (
	"fmt"
	"os"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	} else {
		fmt.Println(dir)
	}

	//Try to read file which does not exist
	content, err := os.ReadFile("file1.txt")
	if err != nil {
		fmt.Println("We caught error:", err)
		//return
	} else {
		fmt.Println(string(content))
	}

	//try to read another file
	content, err = os.ReadFile("file2.txt")
	if err != nil {
		fmt.Println("We caught error:", err)
		return
	}
	fmt.Println(string(content))
}
