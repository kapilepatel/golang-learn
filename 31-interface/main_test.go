package main

import (
	"fmt"
	"math"
	"testing"
)

func TestSquarearea(t *testing.T) {
	s1 := square{width: 3, length: 2}
	actual := s1.area()
	expected := 6.0
	fmt.Printf("TestSquarearea result expected %v got %v\n", expected, actual)
	if actual == expected {
		fmt.Println("TestSquarearea passed")

	} else {
		t.Errorf("TestSquarearea failed expected %v got %v", expected, actual)
		//fmt.Printf("TestSquarearea failed expected %v got %v", exexpected, actual)
	}
}

func TestCirclearea(t *testing.T) {
	c1 := circle{radius: 1.5}
	actual := c1.area()
	expected := 7.06
	// 1/10 = 0.1 = 1e-1
	// 1/100 = 0.01 = 1e-2

	fmt.Println(math.Abs(actual-expected), 1e-1, 1e-2, 1e-3, 1e-9)
	//if math.Abs(actual-expected) > 1e-2 {
	if math.Abs(actual-expected) > 0.01 {

		t.Errorf("TestCirclearea failed expected %v got %v\n", expected, actual)
	}
}
