package ds

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
	size int
}

func NewLinkedList(items ...int) *LinkedList {
	ll := LinkedList{}
	ll.Insert(items...)
	return &ll
}

func (ll *LinkedList) Insert(items ...int) {
	tailNode := ll.head

	for _, item := range items {
		newNode := node{
			data: item,
		}

		if ll.IsEmpty() {
			ll.head = &newNode
			tailNode = ll.head
			ll.size++
			continue
		}

		tailNode.next = &newNode
		tailNode = &newNode
		ll.size++
	}
}

/*
searches for an element in a linked list.
returns the index of the element and true if the element exists,
-1 or false otherwise.
*/
func (ll *LinkedList) Search(item int) (int, bool) {
	currentNode := ll.head
	currentIndex := 0

	for currentNode != nil {
		if currentNode.data == item {
			return currentIndex, true
		}
		currentIndex++
		currentNode = currentNode.next
	}
	return -1, false
}

/*
deletes an element from linked list
*/
func (ll *LinkedList) Delete(item int) error {
	if ll.head == nil {
		return fmt.Errorf("empty linked list")
	}

	if ll.head.data == item {
		ll.head = ll.head.next
		return nil
	}

	currentNode := ll.head
	for currentNode != nil {
		if currentNode.next != nil && currentNode.next.data == item {
			currentNode.next = currentNode.next.next
			return nil
		}
		currentNode = currentNode.next
	}

	return fmt.Errorf("item not found")
}

func (ll *LinkedList) InsertAtIndex(index, item int) error {
	if index < 0 || index >= ll.size {
		return fmt.Errorf("index out of range")
	}

	// same as insert at head
	if index == 0 {
		newNode := node{
			data: item,
		}

		if ll.IsEmpty() {
			ll.head = &newNode
			return nil
		}

		oldHead := ll.head
		newNode.next = oldHead
		ll.head = &newNode
		return nil
	}

	currentNode := ll.head
	currentIndex := 0
	for currentNode != nil {
		nextIndex := currentIndex + 1
		if nextIndex == index {
			oldNode := currentNode.next
			newNode := node{
				data: item,
				next: oldNode,
			}
			currentNode.next = &newNode
			return nil
		}

		currentNode = currentNode.next
		currentIndex++
	}

	return nil
}

func (ll LinkedList) IsEmpty() bool {
	return ll.head == nil
}

/*
counts up the number of elements in the linked list
*/
func (ll LinkedList) Size() int {
	if ll.IsEmpty() {
		return 0
	}

	numItems := 0
	currentNode := ll.head

	for currentNode != nil {
		numItems++
		currentNode = currentNode.next
	}
	return numItems
}

/*
collects all the elements of a linked list
*/
func (ll LinkedList) GetItems() []int {
	items := []int{}
	currentNode := ll.head

	for currentNode != nil {
		items = append(items, currentNode.data)
		currentNode = currentNode.next
	}
	return items
}

// TODO: implement ordered insertion to linked list 1->2->3->?
