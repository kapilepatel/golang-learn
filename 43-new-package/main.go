package main

import (
	"fmt"
	"golang-learn/43-new-package/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	a := canine{
		name: "Fido",
		age:  dog.Years(10),
	}

	fmt.Println(a)
}
