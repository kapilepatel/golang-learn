package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4)
	fmt.Println("Sum is", x)

	fmt.Println()

	greeting("Hello", "John", "Mac", "Jenny")
	fmt.Println()
	greeting("no-one")

}

//three dots as prefix means it takes any number of arguments as input( same type)

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

func greeting(s string, name ...string) {
	fmt.Printf("Type of s is %T \n", s)
	fmt.Printf("Type of name is %T \n", name)

	fmt.Println(s)
	fmt.Println(name)

}
