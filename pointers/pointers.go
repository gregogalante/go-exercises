package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPoint(point *int) {
	*point = 0
}

func main() {
	i := 1
	fmt.Println("Initial value is: ", i) // -> 1
	fmt.Println("Initial pointer value is: ", &i) // -> 0xc420018068

	zeroVal(i)
	fmt.Println("After zeroVal value is: ", i) // -> 1
	fmt.Println("After zeroVal pointer value is: ", &i) // -> 0xc420018068
	/*
	zeroVal can not update the "i" value because it receives a copy of "i" value called "val"
	in a different memeory location.
	*/

	zeroPoint(&i)
	fmt.Println("After zeroPoint value is: ", i) // -> 0
	fmt.Println("After zeroPoint pointer value is: ", &i) // -> 0xc420018068
	/*
	zeroPoint can update the "i" value because it receives a pointer to the memory location
	of the "i" value and it updates the value on this location.
	*/
}