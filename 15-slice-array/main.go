package main

import (
	"fmt"
)

func main() {
	//creating a data structure called Slice
	//this is zero-based indexing
	x := []int{42, 43, 44, 50, 61}

	for i, v := range x {
		fmt.Printf("Index: %v - Value: %v \n", i, v)

	}

	//array
	var a = [...]string{"Almond", "Banana", "Cookies"}

	fmt.Println("Array values", a)
	fmt.Println("Array length", len(a))
	fmt.Printf("Type %T \n", a)

}
