package main

// This script should be used inside the go packages path.
import (
	"shopping"
	"fmt"
)

func main() {
	fmt.Println(shopping.PriceCheck(4343))
}