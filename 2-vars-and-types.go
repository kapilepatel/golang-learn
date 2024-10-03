package main

import "fmt"

func main() {

	// ":=" this is called Short Declaration Operator
	num1 := 99
	fmt.Println(num1)

	fmt.Println()

	name := "Sachin"
	fmt.Println(name)

	fmt.Println()

	num2, name2, name3 := 11, "Rahul", "Krishna"
	fmt.Println(num2)
	fmt.Println(name2)
	fmt.Println(name3)

	fmt.Println()

	//Below will fail if it is declared but not used anywhere,
	//Golang enforces to use a VARIABLE if it is DECLARED [declared and not used: a]
	//a := 1

	//this underscore is called Blank Identifer, it has multiple uses,
	//it use only "=" operator for assignmnet
	//it does not throw "declared and not used" error like other variables
	_ = 1

	//var is used to DECLARE a VARIABLE along with it's TYPE, in other statement we INITIALIZE the variable
	var x int
	fmt.Println(x)
	x = 10
	fmt.Println(x)

	fmt.Println()

	var y int = 22
	fmt.Println(y)

}
