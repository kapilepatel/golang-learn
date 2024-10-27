package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	var p1 = person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"ABC", "DEF", "GHI"},
	}

	//using underscore means we are ignoring the error (if any)
	var bs, _ = json.Marshal(p1)
	fmt.Println(string(bs))

	//let's try to act in case of an error
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal", err)
	}
	fmt.Println(string(bs))

}
