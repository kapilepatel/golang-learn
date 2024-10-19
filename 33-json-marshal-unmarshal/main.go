package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Friends []person
}

func main() {

	p1 := person{
		Name: "Kk",
		Age:  19,
		Friends: []person{
			{Name: "f1", Age: 20},
			{Name: "f2", Age: 21},
		},
	}

	fmt.Println(p1)

	j, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("Error while doing Marshal:", err)
	}
	fmt.Println(string(j))

}
