package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//half-open interval [0,n), here 0 is inclusive but n is excluded
	x := rand.Intn(250)
	fmt.Println("X: ", x)

	//using if condition
	if x <= 100 {
		fmt.Println("Between 0-100")
	}
	if x >= 101 && x <= 200 {
		fmt.Println("Between 101-200")
	}
	if x >= 201 {
		fmt.Println("Between 201-250")
	}

	//using if and else
	if x <= 100 {
		fmt.Println("Between 0-100")
	} else if x <= 200 {
		fmt.Println("Between 101-200")
	} else {
		fmt.Println("Between 201-250")
	}

	//using switch

	switch {
	case x <= 100:
		fmt.Println("Between 0-100")
	case x <= 200:
		fmt.Println("Between 101-200")
	default:
		fmt.Println("Between 201-250")

	}

	y := 9
	switch y {
	case 8:
		fmt.Println("Case 8")
	case 10:
		fmt.Println("Case 10")
	case 9, 12:
		fmt.Println("Case 9 or 12, try by updating the variable")
	}

}
