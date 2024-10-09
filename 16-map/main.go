package main

import (
	"fmt"
)

func main() {
	//creating a data structure called Map
	//Funny here missing the last comma creates compile error
	//so we need to put comma or closing curly braces in same line
	//key has to be unique
	m := map[string]int{
		"item1":  10,
		"itemx":  12,
		"itemzz": -65,
		"":       0,
		"a b":    0,
	}
	for k, v := range m {
		fmt.Printf("Key: %v - Value: %v \n", k, v)
	}

	if item, ok := m["itemx"]; ok {
		fmt.Println(item)
		//fmt.Printf("%v",item)
	}

	if item, ok := m["zzz"]; ok {
		fmt.Println(item)
		//fmt.Printf("%v",item)
	} else {
		fmt.Println("zzz does not exists")

	}

	// func offset(tz string) int {
	// 	if seconds, ok := timeZone[tz]; ok {
	// 		return seconds
	// 	}
	// 	log.Println("unknown time zone:", tz)
	// 	return 0
	// }

}
