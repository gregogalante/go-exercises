package main

import "fmt"

func main() {
	var scores1 [10]int
	scores1[0] = 442
	fmt.Println(scores1)

	scores2 := [4]int{9001, 9333, 212, 33}
	fmt.Println(scores2)

	scores := [3]int{10, 20, 30}
	fmt.Println(len(scores))

	for index, value := range scores {
		fmt.Println("At position", index, "there is", value)
	}
}