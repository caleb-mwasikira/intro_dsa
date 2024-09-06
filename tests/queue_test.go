package tests

import (
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/ds"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

func TestNewCircularQueue(t *testing.T) {
	items := utils.GenRandomArray(10, 0, 100)

	t.Run("invalid size", func(t *testing.T) {
		_, err := ds.NewCircularQueue[int](-1)
		if err == nil {
			t.Errorf("expected an error when creating queue with -ve size")
		}

		_, err = ds.NewCircularQueue(2, items...)
		if err == nil {
			t.Error("expected an error adding more items than provided size")
		}
	})

	queue, err := ds.NewCircularQueue(len(items), items...)
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments: %v", err)
	}

	storedItems := queue.GetItems()
	if !slices.Equal(items, storedItems) {
		t.Errorf("expected items %#v in queue but got items %#v", items, storedItems)
	}
}

func TestCircularQueuePush(t *testing.T) {
	items := utils.GenRandomArray(4, 0, 10)
	queue, err := ds.NewCircularQueue[int](len(items))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments: %v", err)
	}

	// test push()
	for index, item := range items {
		err := queue.Push(item)
		if err != nil {
			t.Errorf("unexpected error pushing items onto empty queue: %v", err)
		}

		storedItems := queue.GetItems()
		expectedItems := items[:index+1]
		if !slices.Equal(expectedItems, storedItems) {
			t.Errorf("(after push) expected items %#v in queue but got items %#v", expectedItems, storedItems)
		}
	}

	// test push() on full queue
	err = queue.Push(420)
	if err == nil {
		t.Error("expected an error pushing onto an already full queue")
	}
}

func TestCircularQueueFront(t *testing.T) {
	items := utils.GenRandomArray(4, 0, 10)
	emptyQueue, err := ds.NewCircularQueue[int](len(items))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments: %v", err)
	}

	// test front() on empty queue
	_, err = emptyQueue.Front()
	if err == nil {
		t.Errorf("expected an error calling front() on empty queue")
	}

	// test front() on queue with items
	queue := emptyQueue
	for _, item := range items {
		queue.Push(item)
	}

	expectedItem := items[0]
	item, err := queue.Front()
	if err != nil {
		t.Errorf("unexpected error calling front() on queue with items: %v", err)
	}

	if item != expectedItem {
		t.Errorf("expected item %v at front but got item %v", expectedItem, item)
	}
}

func TestCircularQueueIsEmpty(t *testing.T) {
	items := utils.GenRandomArray(4, 0, 10)

	// test isEmpty() on empty queue
	queue, err := ds.NewCircularQueue[int](len(items))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments: %v", err)
	}
	if !queue.IsEmpty() {
		t.Errorf("expected call to isEmpty() on empty queue to return true but got false")
	}

	// test isEmpty() on full queue
	for _, item := range items {
		queue.Push(item)
	}
	if queue.IsEmpty() {
		t.Error("expected call to isEmpty() on full queue to return false but got true")
	}

	// test isEmpty() on queue that was full and is now empty
	for err == nil {
		_, err = queue.Front()
	}
	if !queue.IsEmpty() {
		t.Errorf("expected call to isEmpty() on empty queue to return true but got false")
	}
}

func TestCircularQueueIsFull(t *testing.T) {
	items := utils.GenRandomArray(4, 0, 10)

	// test isFull() on empty queue
	queue, err := ds.NewCircularQueue[int](len(items))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments: %v", err)
	}
	if queue.IsFull() {
		t.Errorf("expected call to isFull() on empty queue to return false but got true")
	}

	// test isFull() on full queue
	for _, item := range items {
		queue.Push(item)
	}
	if !queue.IsFull() {
		t.Errorf("expected call to isFull() on full queue to return true but got false")
	}
}
