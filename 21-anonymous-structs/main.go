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

	as := struct {
		first string
		last  string
		age   int
	}{
		first: "aa",
		last:  "bb",
		age:   19,
	}
	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n\n", as)

	fmt.Println(p1)
	fmt.Println(as)

}
