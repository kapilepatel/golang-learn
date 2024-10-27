package main

import "fmt"

func main() {

	var n int = 5
	for i := 0; i < n/2; i++ {
		for j := i; j > 0; j++ {
			fmt.Print(" ")
		}

		for k := i; k > 0; k++ {
			fmt.Print(" ")
		}

		fmt.Println()
	}

}

/*
   *
  ***
 *****


     *
    **
   ***
  ****
 *****
*/
