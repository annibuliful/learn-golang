package main

import "fmt"

func sumArraywithLength(arr []int) (int,int){
  result := 0
  for _,value := range arr{
    result += value
  }
  return result,len(arr)
}
func main(){
  listNumber := []int{1,2,3,4,5,6,7,8,9,10}

  fmt.Println(sumArraywithLength(listNumber))
}
