package main

import "fmt"

func main(){
   arr := [5]int{1,2,3,4,5}

   fmt.Println(arr)

   fmt.Println("Iterate each element without show index")
   for _,value := range arr{
     fmt.Println(value)
   }
fmt.Println("Iterate each element with show index")
   for index, value := range arr{
     fmt.Println(index, value)
   }
}
