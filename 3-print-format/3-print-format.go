package main

import "fmt"

func main() {

	///Printf - "Print Formatter" this function allows you to format numbers, variables and strings into the first string parameter you give it
	//Print - "Print" This cannot format anything, it simply takes a string and print it
	//Println - "Print Line" same thing as Print() however it will append a newline character \n at the end.

	num1 := 32
	fmt.Printf("Print 32 as default base 10 - %d \n", num1)
	fmt.Printf("Print 32 as binary - %b \n", num1)

	fmt.Println()

	fmt.Println("Print 0-9 in decimal and binary")
	fmt.Println("Decimal - Binary")
	fmt.Printf("%v       -  %b \n", 0, 0)
	fmt.Printf("%v       -  %b \n", 1, 1)
	fmt.Printf("%v       -  %b \n", 2, 2)
	fmt.Printf("%v       -  %b \n", 3, 3)
	fmt.Printf("%v       -  %b \n", 4, 4)
	fmt.Printf("%v       -  %b \n", 5, 5)
	fmt.Printf("%v       -  %b \n", 6, 6)
	fmt.Printf("%v       -  %b \n", 7, 7)
	fmt.Printf("%v       -  %b \n", 8, 8)
	fmt.Printf("%v       -  %b \n", 9, 9)

}
