package main

import "fmt"

var x = 40

const y = 5

func main() {
	fmt.Printf("%v %T \n", x, x)
	fmt.Printf("%v %T \n", y, y)

	x = 41
	fmt.Printf("%v %T \n", x, x)

	//below fails- cannot assign to y (neither addressable nor a map index expression)
	//y = 6
	//fmt.Printf("%v %T \n", y, y)

}
