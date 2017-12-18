package list

import "testing"

// Define a helper function to generate an array of integers
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
			a[i] = min + i
	}
	return a
}

// Test New function
func Test_New(t *testing.T) {
	list := New()

	// check correct initial length
	if list.length != 0 {
		t.Error("[Error] New() doesn't work as expected'")
	}
}
