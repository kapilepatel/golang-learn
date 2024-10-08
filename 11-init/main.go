package main

import "fmt"

func main() {

	fmt.Println("This is from main")
}

func init() {

	fmt.Println("This is from init 1")

}

func init() {

	fmt.Println("This is from init 2")
}

func init() {

	fmt.Println("This is from init 3")

}
