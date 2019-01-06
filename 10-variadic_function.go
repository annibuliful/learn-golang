package main

import "fmt"
func sumArray(args ...int) int{
  result := 0
  for _,value := range(args){
    result += value
  }
  return result
}
func main(){
  returnData := sumArray(1,2,3,4,5,6,7,8,9,10)

  fmt.Println(returnData)
}
