package main

import "fmt"

func showFirst(){
  fmt.Println("show First")
}
func showSecond(){
  fmt.Println("show Second")
}
func main(){
  fmt.Println("with Defer")

  showFirst()
  defer showFirst()
  showSecond()
  showSecond()


  /*
    |              |
    | showFirst()  |
    | showSecond() |
    |______________|
  */
}
