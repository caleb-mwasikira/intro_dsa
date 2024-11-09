package ds

import (
	"errors"
)

var (
	ErrQueueIsFull  error = errors.New("cannot add further items. queue is full")
	ErrQueueIsEmpty error = errors.New("cannot remove any items. empty queue")
)

type ErrInvalidQueueSize struct {
	msg string
}

func (e ErrInvalidQueueSize) Error() string {
	if len(e.msg) != 0 {
		return e.msg
	}
	return "invalid queue size"
}

type CircularQueue[T any] struct {
	rear  int
	front int
	size  int
	items []T
}

func NewCircularQueue[T any](size int, items ...T) (*CircularQueue[T], error) {
	if size <= 0 {
		return nil, ErrInvalidQueueSize{}
	}

	if len(items) > size {
		return nil, ErrInvalidQueueSize{
			msg: "number of items provided greater than defined queue size",
		}
	}

	q := CircularQueue[T]{
		rear:  -1,
		front: -1,
		size:  size,
		items: make([]T, size),
	}
	q.AddItems(items...)

	return &q, nil
}

func (q *CircularQueue[T]) AddItems(items ...T) error {
	for _, item := range items {
		if q.IsFull() {
			return ErrQueueIsFull
		}

		// if adding first element, set FRONT = 0
		if q.IsEmpty() {
			q.front = 0
		}

		// circularly increment the REAR index by 1
		if q.rear+1 >= q.size {
			q.rear = (q.rear + 1) % q.size
		} else {
			q.rear += 1
		}

		q.items[q.rear] = item
	}
	return nil
}
func (q *CircularQueue[T]) GetFrontItem() (T, error) {
	if q.IsEmpty() {
		var none T
		return none, ErrQueueIsEmpty
	}

	item := q.items[q.front]

	// circularly increment the FRONT index
	if q.front+1 > q.size {
		q.front = (q.front + 1) % q.size
	} else {
		q.front += 1
	}

	// we've just removed last element from queue
	// reset FRONT and REAR to -1
	if q.front > q.rear {
		q.front = -1
		q.rear = -1
	}

	return item, nil
}

func (q *CircularQueue[T]) GetFrontItems(n int) ([]T, error) {
	items := []T{}

	for i := 0; i < n; i++ {
		item, err := q.GetFrontItem()
		if err != nil {
			break
		}
		items = append(items, item)
	}
	return items, nil
}

func (q CircularQueue[T]) PeekFront() (T, error) {
	if q.IsEmpty() {
		var none T
		return none, ErrQueueIsEmpty
	}

	item := q.items[q.front]
	return item, nil
}

func (q CircularQueue[T]) PeekFrontItems(n int) ([]T, error) {
	items := []T{}

	for i := 0; i < n; i++ {
		item, err := q.GetFrontItem()
		if err != nil {
			break
		}
		items = append(items, item)
	}
	return items, nil
}

func (q CircularQueue[T]) IsEmpty() bool {
	return q.front == -1 && q.rear == -1
}

/*
The check for full circular queue has a new additional case:

	Case 1: FRONT = 0 && REAR == SIZE - 1
	Case 2: FRONT = (REAR + 1)%SIZE
*/
func (q CircularQueue[T]) IsFull() bool {
	return (q.front == 0 && q.rear == q.size-1) || q.front == (q.rear+1)%q.size
}

func (q CircularQueue[T]) GetAllItems() []T {
	return q.items[:q.rear+1]
}
