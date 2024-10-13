package main

import "fmt"

func main() {
	foo()

	bar(11)

	wish := "Happy new year"
	l := aloha(wish)
	fmt.Printf("Length of %v is %v \n\n", wish, l)

	ss, ii := dogYears("Duggu", 2)
	fmt.Printf("%v and returned int value is %v \n\n", ss, ii)

}

func foo() {
	fmt.Println("no param no return \n ")
}

func bar(n int) {
	fmt.Printf("1 param no return %v \n\n", n)
}

func aloha(s string) int {
	i := len(s)
	return i
}

func dogYears(name string, age int) (string, int) {
	dogAge := 7 * age
	return fmt.Sprintf("%v is %v old", name, dogAge), dogAge
}

/*
no param no return
1 param no return
1 param 1 return
2 params 2 returns
*/
