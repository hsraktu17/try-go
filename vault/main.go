package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	var k int
	for k < 5 {
		fmt.Println("looping....")
		k++
	}

	count := 0
	for {
		fmt.Println("KL")
		count++
		if count == 1 {
			break
		}
	}
}
