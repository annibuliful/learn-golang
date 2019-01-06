package main

import "fmt"

// these function is like Try catch in other languages
func main() {
  // panic function process immediatly
	panic("PANIC")
	str := recover()
	fmt.Println(str)

  // So you need defer a function before call panic function
	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	// panic("PANIC")
}
