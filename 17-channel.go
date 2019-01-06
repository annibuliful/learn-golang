//1:  Nothing Because I don't understand this section yet

//2: I understand incorrectly

//3: I understand a little bit

//4: Channel which is like Data stream in NodeJS emmmmm Got it

//5: Channel always use with goroutine if you use with function.

//5: If you don't use with goroutine it will show error goroutine deadlock
package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "Check Input"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
  // c <- "test"
  // go pinger(c)
	// go printer(c)
}
