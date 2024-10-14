package main

import "fmt"

func main() {
	defer fmt.Println(4)
	defer fmt.Println(3)
	defer fmt.Println(2)

	fmt.Println(0)

	defer fmt.Println("A")
	defer fmt.Println("B")
	defer fmt.Println("C")

}
