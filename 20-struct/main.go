package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Embedded struct
type secretAgent struct {
	person
	licencetokill bool
}

type dualSecretAgent struct {
	p1            person
	p2            person
	licencetokill bool
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
	fmt.Println("---------")

	sa1 := secretAgent{
		person: person{
			first: "GI",
			last:  "Joe",
			age:   25,
		},
		licencetokill: true,
	}
	fmt.Println(sa1)
	//we can access the property directky or by providing full path
	fmt.Println(sa1.first, sa1.person.first)
	fmt.Println(sa1.last, sa1.person.last)
	fmt.Println(sa1.licencetokill)

	fmt.Println("---------")

	da1 := dualSecretAgent{
		p1: person{
			first: "DA1F",
			last:  "DA1L",
			age:   25,
		},
		p2: person{
			first: "DA2F",
			last:  "DA2L",
			age:   26,
		},
		licencetokill: true,
	}
	fmt.Println(da1)
	fmt.Println(da1.p1.first)
	//will not work as we have more than 1 "first"
	//fmt.Println(da1.first)

}
