package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{1, 4, 2, 9, 56, 33}
	xs := []string{"z", "Z", "zebra", "apple", "0", "9"}

	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println(xi)
	fmt.Println(xs)

}
