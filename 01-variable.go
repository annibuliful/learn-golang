package main

import "fmt"

func main(){
  const name string = "Tar";
  var age int16 = 20;
  var height float64 = 94.6;
  var man bool = true;

  fmt.Printf("name: %s, age: %d, height: %.2f, man: %t \n",name,age,height,man)
  fmt.Println("name:",name,", age:",age,", height:",height,", man:", man)
}
