package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("Outer Loop %v - Inner Loop %v \n", i, j)
		}
	}

}
