package main

import "fmt"

func main(){
  withOut := []int{1,2,3,4,5,6,7}

  with := make([]int,1)

  fmt.Println(withOut)
  fmt.Println(with)

  fmt.Println("Use append function")
  fmt.Println("Before use append function")
  fmt.Println(with)
  fmt.Println("After use append function")
  appendData := append(with,1,2,3,4,5,6,7)
  fmt.Println(appendData)

  fmt.Println("Use Low:High")
  lowHigh := appendData[0:5]
  fmt.Println(lowHigh)
}
