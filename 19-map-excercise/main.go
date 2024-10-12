package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{`Shaken`, `Not stirred`, `Martinies`}
	m["moneypenny_jenny"] = []string{`intelligence`, `literature`, `computer science`}
	fmt.Println(m)

	fmt.Println("\n\n")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
