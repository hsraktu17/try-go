package main

import "fmt"

func greet(name string) string {
	return "Hi " + name
}

func main() {
	var name string = "Utkarsh"
	printer := greet(name)
	fmt.Scanln()
	fmt.Println(name)
	fmt.Println(printer)
}
