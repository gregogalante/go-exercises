package main

import "fmt"

// Define a list struct
type List struct {
	head *Item
	last *Item
	length int
}

// Define a item struct
type Item struct {
	value interface{}
	next *Item
	prev *Item
	list *List
}

// Define a function to initialize a new list
func newList() *List {
	list := &List{}
	list.length = 0
	return list
}

// Define a list function to insert a new value
func (list *List) insert(value interface{}) {
	// initialize a new item with next and prev as head and last
	newItem := &Item{value, list.head, list.last, list}

	if list.head == nil {
		// set item as head and last of the list
		list.head = newItem
		list.last = newItem
	} else {
		// update head prev pointer
		list.head.prev = newItem
		// set item as new head of the list
		list.head = newItem
		// update last next pointer
		list.last.next = newItem
	}

	list.length++
}

// Define a helper function to generate an array of integers
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
			a[i] = min + i
	}
	return a
}

func main() {
	// create a new list
	list := newList()
	fmt.Println("Created a new list on memory address:", &list, "and length", list.length)
	// create some items for the list
	items := makeRange(1, 10)
	for item := range items {
		list.insert(item)
	}
	// check some results
	fmt.Println("- From head to last:")
	fmt.Println(list.head.value)
	fmt.Println(list.head.next.value)
	fmt.Println(list.head.next.next.value)
	fmt.Println(list.head.next.next.next.value)
	fmt.Println(list.head.next.next.next.next.value)
	fmt.Println(list.head.next.next.next.next.next.value)
	fmt.Println(list.head.next.next.next.next.next.next.value)
	fmt.Println(list.head.next.next.next.next.next.next.next.value)
	fmt.Println(list.head.next.next.next.next.next.next.next.next.value)
	fmt.Println(list.head.next.next.next.next.next.next.next.next.next.value)
	fmt.Println("- From last to head:")
	fmt.Println(list.last.value)
	fmt.Println(list.last.prev.value)
	fmt.Println(list.last.prev.prev.value)
	fmt.Println(list.last.prev.prev.prev.value)
	fmt.Println(list.last.prev.prev.prev.prev.value)
	fmt.Println(list.last.prev.prev.prev.prev.prev.value)
	fmt.Println(list.last.prev.prev.prev.prev.prev.prev.value)
	fmt.Println(list.last.prev.prev.prev.prev.prev.prev.prev.value)
	fmt.Println(list.last.prev.prev.prev.prev.prev.prev.prev.prev.value)
	fmt.Println(list.last.prev.prev.prev.prev.prev.prev.prev.prev.prev.value)
}