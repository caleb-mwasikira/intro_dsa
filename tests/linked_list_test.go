package tests

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/ds"
)

func TestNewLinkedList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	ll := ds.NewLinkedList(items...)
	storedItems := ll.GetItems()

	if !slices.Equal(storedItems, items) {
		t.Errorf("expected items %#v in linked list but got %#v", items, storedItems)
	}
}

func TestLinkedListInsert(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}

	ll := ds.LinkedList{}
	ll.Insert(items...)
	storedItems := ll.GetItems()

	if !slices.Equal(storedItems, items) {
		t.Errorf("expected items %#v in linked list but got %#v", items, storedItems)
	}
}

func TestLinkedListSearch(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	randIndex := rand.Intn(len(items))
	item := items[randIndex]

	ll := ds.NewLinkedList(items...)

	// search for item in linked list
	index, found := ll.Search(item)
	if !found {
		t.Errorf("expected found=true for search but got found=%v", found)
	}

	if index != randIndex {
		t.Errorf("expected index=%v for search but got index=%v", randIndex, index)
	}

	// search for item NOT in linked list
	item = 69
	expectedIndex := -1
	index, found = ll.Search(item)
	if found {
		t.Errorf("expected found=false for search but got found=%v", found)
	}

	if index != expectedIndex {
		t.Errorf("expected index=%v on search but got index=%v", expectedIndex, index)
	}
}

func TestLinkedListDelete(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	ll := ds.NewLinkedList(items...)

	// delete item that in linked list
	err := ll.Delete(items[0])
	if err != nil {
		t.Error("unexpected delete error deleting item in linked list")
	}

	expectedItems := items[1:]
	storedItems := ll.GetItems()
	if !slices.Equal(storedItems, expectedItems) {
		t.Errorf("expected items %#v after delete but got %#v", expectedItems, storedItems)
	}

	// delete item NOT in linked list
	err = ll.Delete(-420)
	if err == nil {
		t.Error("expected an error deleting item NOT in linked list but got err=nil")
	}

	storedItems = ll.GetItems()
	if !slices.Equal(expectedItems, storedItems) {
		t.Errorf("expected items %#v after 'deletion' of item NOT in linked list but got %#v", expectedItems, storedItems)
	}
}

func TestLinkedListInsertAtIndex(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	expectedItems := []int{1, 2, 3, 420, 4, 5}
	ll := ds.NewLinkedList(items...)

	// insert item at valid index
	err := ll.InsertAtIndex(3, 420)
	if err != nil {
		t.Error("unexpected error while inserting at valid index")
	}
	storedItems := ll.GetItems()

	if !slices.Equal(expectedItems, storedItems) {
		t.Errorf("expected items %#v after insert at index 3 but got %#v", expectedItems, storedItems)
	}

	// insert item at invalid index
	err = ll.InsertAtIndex(69, 420)
	if err == nil {
		t.Error("expected an error while inserting at invalid index")
	}

	storedItems = ll.GetItems()

	if !slices.Equal(expectedItems, storedItems) {
		t.Errorf("expected items %#v after insert at invalid index but got %#v", expectedItems, storedItems)
	}
}

func TestLinkedListSize(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	ll := ds.NewLinkedList(items...)
	expectedSize := 5

	actualSize := ll.Size()
	if actualSize != expectedSize {
		t.Errorf("expected size %v from linked list but got size %v", expectedSize, actualSize)
	}
}

func TestLinkedListIsEmpty(t *testing.T) {
	ll := ds.NewLinkedList()

	if !ll.IsEmpty() {
		t.Error("expected linked list with no added items to be empty but is NOT")
	}
}
