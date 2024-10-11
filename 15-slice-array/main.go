package main

import (
	"fmt"
)

func main() {
	//Slice
	//creating a data structure called Slice
	//this is zero-based indexing
	x := []int{42, 43, 44, 50, 61}

	for i, v := range x {
		fmt.Printf("Index: %v - Value: %v \n", i, v)

	}
	/////////////////////////
	//array
	var a = [...]string{"Almond", "Banana", "Cookies"}

	fmt.Println("Array values", a)
	fmt.Println("Array length", len(a))
	fmt.Printf("Type %T \n\n", a)

	/////////////////
	//Slice
	s := []string{"Almond", "Banana", "Cookies"}
	fmt.Println("Slice values", s)
	fmt.Println("Slice length", len(s))
	fmt.Printf("Type %T \n", s)

	//for range loop over slice
	for i, v := range s {
		fmt.Printf("i %v v %v \n", i, v)
	}

	for _, v := range s {
		fmt.Printf(" %v \n", v)
	}

	//Access by index
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	//fmt.Println(s[3]) //Panic: runtime error: index out of range [3] with length 3

	//Slice of int

	si := []int{1, 2, 5, 10}
	fmt.Println(si)

	//variadic paramater
	si = append(si, 11, 13, 33)
	fmt.Println("After appemd, ", si)

	//delete from a slice
	//si 1 2 5 10 11 13 33
	//let's skip item 5 which is at index 2 (zero based indexing)
	//[includes:excludes]
	xi := append(si[:2], si[3:]...)
	fmt.Println("xi ", xi)

	//slice using make
	yi := make([]int, 0, 10)
	fmt.Println("yi ", yi)
	fmt.Println(len(yi))
	fmt.Println(cap(yi))
	yi = append(yi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("yi ", yi)
	fmt.Println(len(yi))
	fmt.Println(cap(yi))

}
