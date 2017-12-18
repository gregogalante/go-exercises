package cirlist

import (
	"testing"
)

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

// Test Insert function
func Test_Insert(t *testing.T) {
	list := New()
	items := makeRange(1, 10)

	for _, item := range items {
		list.Insert(item)
	}

	// check correct insert of elements
	currentItem := list.head.next
	if currentItem.prev.value != (currentItem.value.(int) + 1) {
		t.Error("[Error] Insert() doesn't work as expected'")
	}
	if currentItem.next.value != (currentItem.value.(int) - 1) {
		t.Error("[Error] Insert() doesn't work as expected'")
	}
}

// Test Remove function
func Test_Remove(t *testing.T) {
	list := New()
	items := makeRange(1, 10)

	for _, item := range items {
		list.Insert(item)
	}

	// check correct remove of last
	list.Remove(1)
	if list.last.value == 1 {
		t.Error("[Error] Remove() doesn't work as expected'")
	}
	// check correct remove of head
	list.Remove(10)
	if list.head.value == 10 {
		t.Error("[Error] Remove() doesn't work as expected'")
	}
	// check correct remove of custom value
	value := 6
	current := list.head
	list.Remove(value)
	for i := 0; i < list.length; i++ {
		if current.value == value {
			t.Error("[Error] Remove() doesn't work as expected'")
		}

		current = current.next
	}
}
