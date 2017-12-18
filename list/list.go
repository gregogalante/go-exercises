package list

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
func New() *List {
	list := &List{}
	list.length = 0
	return list
}

// Define a list function to insert a new value
func (list *List) Insert(value interface{}) {
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
