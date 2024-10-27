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

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal", err)
	}
	fmt.Println(string(bs))

}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in toJSON: %v", err)
	}
	return bs, nil
}
