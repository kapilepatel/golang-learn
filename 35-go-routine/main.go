package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(12)
	go print("a")
	go print("a")
	go print("a")

	go print("a")
	go print("a")
	go print("a")

	go print("b")
	go print("b")
	go print("b")

	go print("b")
	go print("b")
	go print("b")

	wg.Wait()
}

func print(s string) {
	fmt.Println(s)
	wg.Done()
}
