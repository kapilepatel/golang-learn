package main

import (
	"fmt"
)

func main() {
	//for loop with initialization, condition and increment
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	// for loop with just the condition
	j := 10

	for j < 20 {
		fmt.Println("j: ", j)
		j++
	}

	k := 0
	for {
		if k > 5 {
			break
		}
		fmt.Println("k: ", k)
		k++
	}
}
