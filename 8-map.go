package main

import "fmt"

func main() {
	mapVar := make(map[string]int)

	mapVar["age"] = 20
	mapVar["height"] = 172
	fmt.Println("\n without Key")
	fmt.Println(mapVar)
	fmt.Println("\n with Key")
	fmt.Println("age: ", mapVar["age"], ", height: ", mapVar["height"])

	fmt.Println("\n Destructoring value and check value ")
	value, notNil := mapVar["age"]
	fmt.Println(value, notNil)


	fmt.Println("\nUse delete function")
	delete(mapVar, "age")
	fmt.Println(mapVar)
}
