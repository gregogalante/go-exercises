package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	bob := person{"Bob", 24}
	fmt.Println(bob)
	fmt.Println(bob.name)
	fmt.Println(bob.age)

	foo := person{name: "Foo"}
	fmt.Println(foo)
	fmt.Println(foo.name)
	fmt.Println(foo.age)
}