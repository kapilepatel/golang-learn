package main

import "fmt"

type customErr struct {
	info string
}

func main() {
	c1 := customErr{
		info: "ABCD",
	}

	foo(c1)
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Println("foo ran- ", e)
}
