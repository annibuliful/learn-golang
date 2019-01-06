package main

import "fmt"

func sumArray(arr []int, length *int) {
	result := 0
	for range arr {
		result += 1
	}
	*length = result
}
func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Use * keyword")
	withStar := 0
	sumArray(arr, &withStar)
	fmt.Println(withStar)

	fmt.Println("Use new Keyword")
	withNew := new(int)
	sumArray(arr, withNew)
	fmt.Println(*withNew)

	// If you scared to get confused, you can use traditional like C/C++ LOL
}
