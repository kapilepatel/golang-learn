package main

import (
	"fmt"
	"math"
)

func main() {
	//below will success and print 9223372036854775807
	var num1 int = math.MaxInt
	fmt.Println(num1)

	//below will fail with error: cannot use math.MaxUint (untyped int constant 18446744073709551615) as int value in variable declaration (overflows)
	//var num2 int = math.MaxUint
	//fmt.Println(num2)

	//below will success and print 18446744073709551615
	//we can see MaxUint is double in value compared to MaxInt
	var num3 uint = math.MaxUint
	fmt.Println(num3)

	//this is also valid
	var num4 uint = math.MaxInt
	fmt.Println(num4)

}
