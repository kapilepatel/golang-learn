package main

import "fmt"

func main() {
	//Multi dimensional slice
	sname := []string{"kk", "rr", "ll"}
	fmt.Println(sname)
	score := []string{"A", "A", "B"}
	fmt.Println(score)

	s := [][]string{sname, score}
	fmt.Println(s)

	s2 := [][]string{
		{"A", "B", "C", "D"}, {"0", "1", "2", "3"}, {"AA", "BB"}, {"-", "--", "---", "----", "-----"}}
	fmt.Println(s2)

}
