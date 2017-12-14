package main

import "fmt"

type Sayan struct {
	Name string
	Power int
	Father *Sayan
}

func main() {
	goku := Sayan{
		Name: "Goku",
		Power: 1000,
	}

	gohan := Sayan{
		Name: "Gohan",
		Power: 1000,
		Father: &goku,
	}

	fmt.Println("Gohan father is", gohan.Father.Name)
}