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

type alien struct {
	First   string
	Last    string
	Age     int
	Sayings []string
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

	fmt.Println("----------")

	aj := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken", "not stirred"]},{"First":"Jenny","Penny":"Bond","Age":22,"Sayings":["s1", "s2"]}]`
	var as []alien
	err = json.Unmarshal([]byte(aj), &as)
	if err != nil {
		fmt.Println("Error while doing Unmarshal:", err)
	}

	fmt.Println(as)
	fmt.Println("-----++++++++-----")
	for _, v := range as {
		fmt.Println("Alien First", v.First)
		fmt.Println("Age ", v.Age)
	}
}
