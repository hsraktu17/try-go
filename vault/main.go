package main

import "fmt"

func main() {
	var x int
	fmt.Println("Enter the number")
	fmt.Scanln(&x)

	if x > 10 {
		fmt.Println("Dont know what to do")
	} else {
		fmt.Println("knows what to do")
	}

	if y := 1; y > 5 {
		fmt.Println("Hello")
	} else {
		fmt.Println("not hello")
	}
}
