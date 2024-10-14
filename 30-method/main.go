package main

import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) speak() {
	fmt.Println(p.first, p.age)
}

func main() {
	p1 := Person{
		first: "john",
		age:   23,
	}
	p1.speak()

}
