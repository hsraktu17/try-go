package main

import (
	"fmt"
	"math"
)

func main() {
	array1()
	array()
}

func area() {
	const pi = 3.14
	var radius float64
	fmt.Scanln(&radius)
	area := pi * radius * radius
	fmt.Printf("The area of the circle with radius %.2f is %.2f\n", radius, area)
}

func div(a, b float64) (float64, float64) {
	if b == 0 {
		return 0, 0
	}
	quoitent := a / b
	remainder := math.Mod(a, b)
	return quoitent, remainder
}

func array() {
	s := make([]int, 3, 5) // length = 3, capacity = 5

	fmt.Println("Slice:", s)         // [0 0 0]
	fmt.Println("Length:", len(s))   // 3
	fmt.Println("Capacity:", cap(s)) // 5

	s = append(s, 1, 2)              // still within capacity
	fmt.Println("After append:", s)  // [0 0 0 1 2]
	fmt.Println("Length:", len(s))   // 5
	fmt.Println("Capacity:", cap(s)) // 5

	s = append(s, 3)                        // exceeds capacity, Go allocates new array
	fmt.Println("After another append:", s) // [0 0 0 1 2 3]
	fmt.Println("Length:", len(s))          // 6
	fmt.Println("Capacity:", cap(s))        // 10 (Go doubles it)
}

func array1() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", s)                // [1 2 3 4 5]
	fmt.Println("Length:", len(s))          // 5
	fmt.Println("Capacity:", cap(s))        // 5
	s = append(s, 6, 7)                     // still within capacity
	fmt.Println("After append:", s)         // [1 2 3 4 5 6 7]
	fmt.Println("Length:", len(s))          // 7
	fmt.Println("Capacity:", cap(s))        // 10
	s = append(s, 8)                        // exceeds capacity, Go allocates new array
	fmt.Println("After another append:", s) // [1 2 3 4 5 6 7 8]
	fmt.Println("Length:", len(s))          // 8
}
