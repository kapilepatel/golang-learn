package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	//composit literal
	p1 := person{
		first: "kk",
		last:  "pp",
		age:   18,
	}

	p2 := person{
		first: "aa",
		last:  "bb",
		age:   19,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p2.first)

}
