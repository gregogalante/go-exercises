package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}

func main() {
	r := rect{10, 5}
	fmt.Println(r)

	fmt.Println(r.area())
	fmt.Println(r.perim())
}