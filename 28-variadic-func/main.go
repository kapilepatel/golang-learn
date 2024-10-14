package main

import "fmt"

func main() {
	si := []int{1, 3, 9, -3}
	a := foo(si...)
	fmt.Println(a)

	b := bar(si)
	fmt.Println(b)

}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
