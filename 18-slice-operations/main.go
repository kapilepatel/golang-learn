package main

import "fmt"

func main() {

	a := []string{"kk", "rr", "ll"}
	//here same items is being referenced
	b := a

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	//looks like we are updating a but we can see in the printed logs a and b both reflect new change
	a[0] = "z"
	fmt.Println("Updated a")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	fmt.Println("\n")
	//in case we don't want the items to refer to same entity, we do as given below
	m := []string{"kk", "rr", "ll"}

	n := make([]string, 3)
	//func copy(dst, src []Type) int
	numElementsCopied := copy(n, m)
	fmt.Println("numElementsCopied: ", numElementsCopied)

	fmt.Println("m: ", m)
	fmt.Println("n: ", n)

	m[0] = "z"
	fmt.Println("Updated m")
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)

}
