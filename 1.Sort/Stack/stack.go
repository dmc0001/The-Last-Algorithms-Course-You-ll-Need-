package stack

// Interface is the behavior implemented by a stack.
type Interface interface {
	Push(v any)
	Pop() (any, bool)
	Peek() (any, bool)
	IsEmpty() bool
}

// compile-time check that *Stack satisfies Interface.
var _ Interface = (*Stack)(nil)

type node struct {
	value    any
	previous *node
}

// Stack is a LIFO stack backed by a singly linked list.
type Stack struct {
	head *node
}

func (s *Stack) Push(v any) {
	newNode := &node{value: v}
	newNode.previous = s.head
	s.head = newNode
}

func (s *Stack) Pop() (any, bool) {
	if s.head == nil {
		return nil, false
	}
	removed := s.head.value
	s.head = s.head.previous
	return removed, true
}

func (s *Stack) Peek() (any, bool) {
	if s.head == nil {
		return nil, false
	}
	return s.head.value, true
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}
