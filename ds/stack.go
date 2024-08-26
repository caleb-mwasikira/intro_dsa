package ds

import "errors"

var (
	ErrStackOverflow  error = errors.New("stack overflow")
	ErrStackUnderflow error = errors.New("stack underflow")
)

type ErrInvalidStackSize struct {
	msg string
}

func (e ErrInvalidStackSize) Error() string {
	if len(e.msg) != 0 {
		return e.msg
	}
	return "invalid stack size"
}

type Stack struct {
	items []int
	size  int
	top   int
}

func NewStack(size int, items ...int) (*Stack, error) {
	if size <= 0 {
		return nil, ErrInvalidStackSize{}
	}

	if len(items) > size {
		return nil, ErrInvalidStackSize{
			msg: "number of items provided greater than defined stack size",
		}
	}

	_items := make([]int, size)
	copy(_items, items)

	return &Stack{
		items: _items,
		size:  size,
		top:   -1 + len(items),
	}, nil
}

func (s *Stack) Push(items ...int) (int, error) {
	num := 0

	for _, item := range items {
		if s.isFull() {
			return num, ErrStackOverflow
		}

		s.top += 1
		s.items[s.top] = item
		num++
	}
	return num, nil
}

func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, ErrStackUnderflow
	}

	item := s.items[s.top]
	s.top--
	return item, nil
}

func (s Stack) GetItems() []int {
	return s.items
}

func (s Stack) isEmpty() bool {
	return s.top == -1
}

func (s Stack) isFull() bool {
	return s.top == s.size-1
}
