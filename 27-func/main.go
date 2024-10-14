package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a)

	b, c := bar()
	fmt.Println(b, c)

}

func foo() int {
	return -99
}

func bar() (int, string) {
	return 11, "Elevennnnn"
}
