package main

import "fmt"

type Sayan struct {
	Name string
	Power int
}

func nextLevel(s Sayan) {
	s.Power += 100
}

func nextLevelCorrect(s *Sayan) {
	s.Power += 100
}

// Fake type constructor.
func NewSayan(name string, power int) Sayan {
	return Sayan{
		Name: name,
		Power: power,
	}
}

func main() {
	goku := Sayan{"Goku", 1000}
	// step 0
	fmt.Println("The sayan", goku.Name, "has power", goku.Power)
	// step 1 (not update goku)
	nextLevel(goku)
	fmt.Println("The sayan", goku.Name, "has power", goku.Power)
	// step 2 (update goku)
	nextLevelCorrect(&goku)
	fmt.Println("The sayan", goku.Name, "has power", goku.Power)

	vegeta := NewSayan("Vegeta", 600)
	// step 0
	fmt.Println("The sayan", vegeta.Name, "has power", vegeta.Power)
	// step 1 (not update vegeta)
	nextLevel(vegeta)
	fmt.Println("The sayan", vegeta.Name, "has power", vegeta.Power)
	// step 2 (update vegeta)
	nextLevelCorrect(&vegeta)
	fmt.Println("The sayan", vegeta.Name, "has power", vegeta.Power)
}