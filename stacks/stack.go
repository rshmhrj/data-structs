package stacks

import (
	"fmt"

	"gitlab.com/rshmhrj/data-structs/lists"
)

// Stack represents a Last-In-First-Out model or process.
// Time Complexity: Access O(n), Search O(n), Insert O(1), Delete O(1).
// Space Complexity: O(n).
type Stack struct {
	stack *lists.SinglyLinkedList
}

func InitStack(values ...interface{}) *Stack {
	l := lists.NewSinglyLinkedList()
	for i, v := range values {
		n := l.Prepend(v)
		if i == 0 {
			l.Tail = n
			n.Next = nil
		}
		l.Head = n
	}
	return &Stack{
		stack: l,
	}
}

func NewStack() *Stack {
	l := lists.NewSinglyLinkedList()
	return &Stack{
		stack: l,
	}
}

// Push
func (s *Stack) Push(value interface{}) {
	s.stack.Prepend(value)
}

// Pop
func (s *Stack) Pop() interface{} {
	return s.stack.RemoveHead()
}

// Peek
func (s *Stack) Peek() interface{} {
	return s.stack.Head.Value
}

// Search
func (s *Stack) Search(value interface{}) bool {
	return s.stack.Contains(value)
}

// DepthOf
func (s *Stack) DepthOf(value interface{}) int {
	return s.stack.IndexOf(value)
}

// Size
func (s *Stack) Size() int {
	return s.stack.Len()
}

// IsEmpty
func (s *Stack) IsEmpty() bool {
	return s.stack.IsEmpty()
}

// String
func (s *Stack) String() string {
	return fmt.Sprintf("{ Top: %v, Length: %v, %v }", s.stack.Head.Value, s.stack.Len(), s.stack)
}
