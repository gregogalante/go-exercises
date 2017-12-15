/*
TODO: Continue
*/

package main

import (
	"fmt"
	"sync"
)

// Define a List struct
type List struct {
	head *Item
	last *Item
	len int
	locker sync.RWMutex
}

// Define a Item struct
type Item struct {
	Val interface{}
	next *Item
	prev *Item
	list *List
}

// Define a list function to get the first item.
func (list *List) First() *Item {
	return list.head
}

// Define a list function to get the last item.
func (list *List) Last() *Item {
	return list.last
}

// Define a item function to get the next item.
func (item *Item) Next() *Item {
	return item.next
}

// Define a item function to get the prev item.
func (item *Item) Prev() *Item {
	return item.prev
}

// Define a function to initialize a new list
func New() *List {
	list := &List{}
	list.len = 0
	return list
}

// Define a function to insert new elements on the list
func Insert(value interface{}, list *List) *List {
	newItem := &Item{value, list.head, list.last, list}
	list.locker.Lock()
	defer list.locker.Unlock()

	if list.head == nil {
		list.head = newItem
		list.last = newItem
	} else {
		list.head = newItem
		list.head.prev = newItem
		list.last.next = newItem
	}

	list.len++
	return list
}

// Define main function
func main() {
	// define list and two items
	list := New()
	item1 := "Hello"
	item2 := "World"
	// insert items on list
	Insert(item1, list)
	Insert(item2, list)

	fmt.Println(list.First())
	fmt.Println(list.Last())
}
