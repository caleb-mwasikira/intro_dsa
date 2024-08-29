package ds

import (
	"slices"
)

type ErrStackEmpty struct{}

func (e ErrStackEmpty) Error() string {
	return "cannot pop any more items; stack is empty"
}

type ErrStackFull struct{}

func (e ErrStackFull) Error() string {
	return "cannot add any more items; stack is full"
}

type ErrInvalidStackSize struct {
	msg string
}

func (e ErrInvalidStackSize) Error() string {
	if len(e.msg) != 0 {
		return e.msg
	}
	return "invalid stack size"
}

type Number interface {
	int | rune | float32 | float64
}

type Stack[T Number] struct {
	items []T
	size  int
	top   int
}

func NewStack[T Number](size int, items ...T) (*Stack[T], error) {
	if size <= 0 {
		return nil, ErrInvalidStackSize{}
	}

	if len(items) > size {
		return nil, ErrInvalidStackSize{
			msg: "number of items provided greater than defined stack size",
		}
	}

	_items := make([]T, size)
	copy(_items, items)

	return &Stack[T]{
		items: _items,
		size:  size,
		top:   -1 + len(items),
	}, nil
}

func (s *Stack[T]) Push(items ...T) (int, error) {
	num := 0

	for _, item := range items {
		if s.IsFull() {
			return num, ErrStackFull{}
		}

		s.top += 1
		s.items[s.top] = item
		num++
	}
	return num, nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return 0, ErrStackEmpty{}
	}

	item := s.items[s.top]
	s.items[s.top] = 0
	s.top--
	return item, nil
}

func (s Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return 0, ErrStackEmpty{}
	}

	item := s.items[s.top]
	return item, nil
}

/*
returns all items in the stack from top to bottom;
in the order they would have been popped
*/
func (s Stack[T]) GetItems() []T {
	items := s.items[0 : s.top+1]
	slices.Reverse(items)
	return items
}

func (s Stack[T]) IsEmpty() bool {
	return s.top == -1
}

func (s Stack[T]) IsFull() bool {
	return s.top == s.size-1
}
