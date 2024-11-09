package tests

import (
	"log"
	"slices"
	"testing"

	"github.com/caleb-mwasikira/intro_dsa/ds"
	"github.com/caleb-mwasikira/intro_dsa/utils"
)

var (
	queueItems []int = utils.GenRandomArray(10, 0, 100)
)

func TestNewCircularQueue(t *testing.T) {
	t.Run("invalid size", func(t *testing.T) {
		_, err := ds.NewCircularQueue[int](-1)
		if err == nil {
			t.Errorf("expected an error when creating queue with -ve size")
		}

		_, err = ds.NewCircularQueue(2, queueItems...)
		if err == nil {
			t.Error("expected an error adding more queueItems than provided size")
		}
	})

	queue, err := ds.NewCircularQueue(len(queueItems), queueItems...)
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments; %v", err)
	}

	storedItems := queue.GetAllItems()
	if !slices.Equal(storedItems, queueItems) {
		t.Errorf("expected items %#v in queue but got items %#v", queueItems, storedItems)
	}
}

func TestQueueAddItems(t *testing.T) {
	queue, err := ds.NewCircularQueue[int](len(queueItems))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments; %v", err)
	}

	// test push()
	for index, item := range queueItems {
		err := queue.AddItems(item)
		if err != nil {
			t.Errorf("unexpected error pushing items onto empty queue; %v", err)
		}

		storedItems := queue.GetAllItems()
		expectedItems := queueItems[:index+1]
		if !slices.Equal(storedItems, expectedItems) {
			t.Errorf("(after push) expected items %#v in queue but got items %#v", expectedItems, storedItems)
		}
	}

	// test push() on full queue
	err = queue.AddItems(1020)
	if err == nil {
		t.Error("expected an error pushing onto an already full queue")
	}
}

func TestQueueGetFrontItem(t *testing.T) {
	emptyQueue, err := ds.NewCircularQueue[int](len(queueItems))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments; %v", err)
	}

	// test front() on empty queue
	_, err = emptyQueue.GetFrontItem()
	if err == nil {
		t.Errorf("expected an error calling front() on empty queue")
	}

	// test front() on queue with items
	queue := emptyQueue
	queue.AddItems(queueItems...)

	expectedItem := queueItems[0]
	item, err := queue.GetFrontItem()
	if err != nil {
		t.Errorf("unexpected error calling front() on queue with items; %v", err)
	}

	if item != expectedItem {
		t.Errorf("expected item %v at front but got item %v", expectedItem, item)
	}
}

func TestQueueGetFrontItems(t *testing.T) {
	queue, err := ds.NewCircularQueue(len(queueItems), queueItems...)
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments; %v", err)
	}

	numOfItems := 4
	expectedItems := queueItems[0:numOfItems]
	gotItems, err := queue.GetFrontItems(numOfItems)
	if err != nil {
		t.Errorf("unexpected error getting items in circular queue")
	}

	if !slices.Equal(gotItems, expectedItems) {
		t.Errorf("expected items %#v but got %#v", expectedItems, gotItems)
	}
}

func TestQueuePeekFront(t *testing.T) {
	queue, err := ds.NewCircularQueue[int](len(queueItems))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments; %v", err)
	}

	// test peekFront() on empty queue returns an error
	_, err = queue.PeekFront()
	if err == nil {
		t.Errorf("expected error peeking empty queue but got none")
	}

	err = queue.AddItems(queueItems...)
	if err != nil {
		log.Fatalf("error adding items to queue; %v", err)
	}

	// test correct front item
	expectedFrontItem := queueItems[0]
	frontItem, err := queue.PeekFront()
	if err != nil {
		t.Errorf("unexpected error peeking front item; %v", err)
	}

	if frontItem != expectedFrontItem {
		t.Errorf("expected front item to be %v but got %v", expectedFrontItem, frontItem)
	}

	// test that peekFront() does not change queue items
	items := queue.GetAllItems()
	if !slices.Equal(items, queueItems) {
		t.Errorf("unexpected behaviour! peekFront() func changed underlying queue items")
	}
}

func TestQueuePeekFrontItems(t *testing.T) {
	queue, err := ds.NewCircularQueue[int](len(queueItems), queueItems...)
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments; %v", err)
	}

	numItemsToPeek := 5
	expectedFrontItems := queueItems[0:numItemsToPeek]
	frontItems, err := queue.PeekFrontItems(numItemsToPeek)
	if err != nil {
		t.Errorf("unexpected error peeking front items; %v", err)
	}

	if !slices.Equal(frontItems, expectedFrontItems) {
		t.Errorf("expected peeked front items to be %v but got %v", expectedFrontItems, frontItems)
	}

	// test that peekFrontItems() does not change queue items
	items := queue.GetAllItems()
	if !slices.Equal(items, queueItems) {
		t.Errorf("unexpected behaviour! peekFrontItems() func changed underlying queue items")
	}
}

func TestQueueIsEmpty(t *testing.T) {
	// test isEmpty() on empty queue
	queue, err := ds.NewCircularQueue[int](len(queueItems))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments; %v", err)
	}
	if !queue.IsEmpty() {
		t.Errorf("expected call to isEmpty() on empty queue to return true but got false")
	}

	// test isEmpty() on full queue
	for _, item := range queueItems {
		queue.AddItems(item)
	}
	if queue.IsEmpty() {
		t.Error("expected call to isEmpty() on full queue to return false but got true")
	}

	// test isEmpty() on queue that was full and is now empty
	for err == nil {
		_, err = queue.GetFrontItem()
	}
	if !queue.IsEmpty() {
		t.Errorf("expected call to isEmpty() on empty queue to return true but got false")
	}
}

func TestQueueIsFull(t *testing.T) {
	// test isFull() on empty queue
	queue, err := ds.NewCircularQueue[int](len(queueItems))
	if err != nil {
		t.Errorf("unexpected error creating a queue with valid arguments; %v", err)
	}
	if queue.IsFull() {
		t.Errorf("expected call to isFull() on empty queue to return false but got true")
	}

	// test isFull() on full queue
	for _, item := range queueItems {
		queue.AddItems(item)
	}
	if !queue.IsFull() {
		t.Errorf("expected call to isFull() on full queue to return true but got false")
	}
}
