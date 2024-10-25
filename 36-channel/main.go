package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var c = make(chan int)
	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i, <-c)
		}

		for v := range c {
			//<-c
			fmt.Println(v)

		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Exit")

}
