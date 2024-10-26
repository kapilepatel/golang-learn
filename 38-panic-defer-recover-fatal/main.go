package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")
	g()
	fmt.Println("Returned normally from g.")

}

func g() {
	fmt.Println("Panic in g.")

	panic("Dummy panic in g")

	defer fmt.Println("Defer in g.")

}
