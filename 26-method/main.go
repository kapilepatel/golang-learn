package main

import "fmt"

type person struct {
	first string
}

// This is a method, not a function, here we define receiverType receiverName before method name
func (p person) speak() {
	fmt.Println("My name is", p.first)
}

func main() {
	p1 := person{
		first: "John",
	}
	p2 := person{
		first: "Kate",
	}

	p1.speak()
	p2.speak()

}
