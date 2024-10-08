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

}
