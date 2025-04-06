package main

import "fmt"

func main() {

	var arr [5]int
	fmt.Println("Enter array elements")
	for i := range arr {
		_, err := fmt.Scanln(&arr[i])
		if err != nil {
			fmt.Println("Error in the array", err.Error())
			return
		}
	}
	fmt.Println("Entry in the arrays is done", arr)
}
