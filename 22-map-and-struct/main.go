package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
	fav   []string
}

func main() {

	p1 := person{
		first: "kk",
		last:  "pp",
		age:   18,
		fav:   []string{"Vanilla", "Hazelnut"},
	}
	p2 := person{
		first: "aa",
		last:  "bb",
		age:   20,
		fav:   []string{"Mix", "milk based"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("---")
	//Here our key is last name and the value is person object
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Printf("%v %v likes ", v.first, v.last)
		for _, item := range v.fav {
			fmt.Printf("%v ,", item)
		}
		fmt.Printf("\n")
	}

}
