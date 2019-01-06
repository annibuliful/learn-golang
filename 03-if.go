package main

import "fmt"

func main() {
	var number int16 = 30

	if number % 2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}
}
