package main

import "fmt"

func main() {
	scores1 := []int{1,4,293,4,9}
	scores1 = 	append(scores1, 24)
	fmt.Println("Scores1 is", scores1)

	scores2 := make([]int, 5)
	scores2 = append(scores2, 5)
	fmt.Println("Scores2 is", scores2)

	scores3 := make([]int, 0, 5)
	scores3 = append(scores3, 1)
	scores3 = append(scores3, 2)
	scores3 = append(scores3, 3)
	scores3 = append(scores3, 4)
	scores3 = append(scores3, 5)
	capacity_scores3 := cap(scores3)
	fmt.Println("Capacity of scores3 is", capacity_scores3)
	scores3 = append(scores3, 6)
	capacity_scores3_new := cap(scores3)
	fmt.Println("New capacity of scores3 is", capacity_scores3_new)
}