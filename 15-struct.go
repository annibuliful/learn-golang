package main

import "fmt"

type User struct {
	name   string
	age    int
	height float64
}

func showName(name *User) {
	fmt.Printf("name %s, age %d, height %.3f", name.name, name.age, name.height)
}
func main() {
	myself := new(User)
	myself.name = "Tar"
	myself.age = 20
	myself.height = 172.203
	showName(myself)
}
