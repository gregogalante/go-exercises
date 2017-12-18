package cirlist

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

// Define a list function to remove a value
func (list *List) Remove(value interface{}) {
	// do nothing if list is empty
	if list.head == nil {
		return
	}
	// loop to search item
	headValue := list.head.value
	lastValue := list.last.value
	first := list.head
	last := list.last
	for {
		// if item is present destroy it and return
		if first.value == value {
			// change next and prev pointers
			first.prev.next = first.next
			first.next.prev = first.prev
			// change head and last pointers
			if first.value == headValue {
				list.head = first.next
			}
			if first.value == lastValue {
				list.last = first.prev
			}
			// destroy item
			first.prev = nil
      first.next = nil
      first.value = nil
			first.list = nil
			// change counter
			list.length--
			return
		}
		// if list is finished return
		if first.value == last.value {
			return
		}
		// check next item
		first = first.next
	}
}
