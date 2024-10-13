package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4}
	//in case we already have a slice, we can use ... to unpack/unfurl the slice
	x := sum(slc...)
	fmt.Println("Sum is", x)

}

func sum(items ...int) int {
	//when T type of arguments are passed it becomes slice of T
	fmt.Printf("Type is %T \n", items)
	fmt.Println(items)

	s := 0
	for _, v := range items {
		s += v
	}
	return s
}
