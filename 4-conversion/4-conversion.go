package main

import "fmt"

func main() {
	num1 := 21
	num2 := 21.0

	fmt.Printf("num1 %v of type %T \n", num1, num1)
	fmt.Printf("num2 %v of type %T \n", num2, num2)

	var num3 float32 = 40.123
	fmt.Printf("num3 %v of type %T \n", num3, num3)

	//below will fail with error - cannot use num2 (variable of type float64) as float32 value in assignment
	//num3 = num2

	//conversion
	num3 = float32(num2)
	fmt.Printf("num3 %v of type %T \n", num3, num3)

}
