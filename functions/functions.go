package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return plus(a, b) + c
}
 
func main() {
	res1 := plus(1, 2)
	fmt.Println(res1)

	res2:= plusPlus(1, 2, 3)
	fmt.Println(res2)
}