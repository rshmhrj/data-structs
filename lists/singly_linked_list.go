package lists

import (
	"fmt"

	"gitlab.com/rshmhrj/go-text-gen/names"
)

// SinglyLinkedList represents a linked list which can be traversed in only one direction.
// Time Complexity: Access O(n), Search O(n), Insert O(1), Delete O(1).
// Space Complexity: O(n).
type SinglyLinkedList struct {
	Name   string
	Size   int
	Head   *SNode
	Tail   *SNode
	reader *sReader
}

// SNode represents a node in a singly linked list which only knows about the Next node in the list
type SNode struct {
	Value interface{}
	Next  *SNode
}

// sReader represents a helper for traversing the linked list
type sReader struct {
	index   int
	current *SNode
}

// NewSinglyLinkedList initializes an empty linked list of type Single or Double
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		Name: names.NewGenerator().Generate(),
		Size: 0,
		Head: nil,
		Tail: nil,
		reader: &sReader{
			index:   -1,
			current: nil,
		},
	}
}

// InitSinglyLinkedList initializes a linked list of type Single or Double with the given values
func InitSinglyLinkedList(values ...interface{}) *SinglyLinkedList {
	l := NewSinglyLinkedList()
	for i, v := range values {
		n := l.Append(v)
		if i == 0 {
			l.Head = n
		}
		l.Tail = n
		if i == len(values)-1 {
			n.Next = nil
		}
	}
	return l
}

// InitNamedSinglyLinkedList initializes a named linked list of type Single or Double with the given values
func InitNamedSinglyLinkedList(name string, value ...interface{}) *SinglyLinkedList {
	l := InitSinglyLinkedList(value...)
	l.Name = name
	return l
}

// Len returns Size of the linked list
func (l *SinglyLinkedList) Len() int {
	return l.Size
}

// IsEmpty returns true if linked list Size is 0
func (l *SinglyLinkedList) IsEmpty() bool {
	return l.Size == 0
}

// Index of Value in list.  If not found returns -1
func (l *SinglyLinkedList) IndexOf(value interface{}) int {
	output := -1
	l.readerRestart()
	for l.reader.HasNext() || l.reader.index == l.Size-1 {
		if l.reader.current.Value == value {
			output = l.reader.index
			break
		}
		l.reader.Next()
	}
	return output
}

// Contains returns true/false if item is in list
func (l *SinglyLinkedList) Contains(value interface{}) bool {
	if l.IndexOf(value) != -1 {
		return true
	}
	return false
}

// SetName updates the Name of the linked list
func (l *SinglyLinkedList) SetName(name string) {
	l.Name = name
}

// Prepend adds a SNode as the new Head of the linked list
func (l *SinglyLinkedList) Prepend(value interface{}) *SNode {
	n := new(SNode)

	if l.Size == 0 {
		n.Next = nil
	}
	if l.Size > 0 {
		n.Next = l.Head
	}

	n.Value = value
	l.Head = n
	l.Size += 1

	return n
}

// Append adds a SNode as the new Tail of the linked list
func (l *SinglyLinkedList) Append(value interface{}) *SNode {
	n := new(SNode)
	if l.Size == 0 {
		l.Head = n
	}
	if l.Size == 1 {
		l.Head.Next = n
	}
	if l.Size > 1 {
		l.readerRestart()
		for l.reader.HasNext() || l.reader.index == l.Size-1 {
			if l.reader.index == l.Size-1 {
				l.reader.current.Next = n
				break
			}
			l.reader.Next()
		}
	}
	n.Value = value
	n.Next = nil
	l.Tail = n
	l.Size += 1

	return n
}

// Insert adds an SNode at the specified index in the linked list (index starting at 0)
func (l *SinglyLinkedList) Insert(index int, value interface{}) {
	n := new(SNode)
	l.readerRestart()
	prevTrav := l.Head
	for l.reader.HasNext() || l.reader.index == l.Size-1 {
		if l.reader.index == index {
			if l.reader.index == 0 {
				l.Head = n
				break
			}
			n.Value = value
			n.Next = l.reader.current
			prevTrav.Next = n
		}
		prevTrav = l.reader.current
		l.reader.Next()
	}
	l.Size += 1
}

// RemoveHead removes the Head of the linked list and RETURNS it's Value
// If there is no element to remove, RETURNS -1
func (l *SinglyLinkedList) RemoveHead() interface{} {
	value := l.Head.Value
	if l.Size == 0 {
		return -1
	}
	if l.Size > 0 {
		l.Head = l.Head.Next
	}

	l.Size -= 1
	return value
}

// RemoveAt removes the element at the given index in the linked list
// If there is no element to remove, or if the index is out of bounds, RETURNS -1
func (l *SinglyLinkedList) RemoveAt(index int) interface{} {
	return nil
}

// RemoveValue removes the first instance of the given Value from the linked list
// If Value is not found, RETURNS error
func (l *SinglyLinkedList) RemoveValue(value interface{}) interface{} {
	return nil
}

// RemoveTail removes the Tail of the linked list and RETURNS it's Value
// If there is no element to remove, RETURNS error
func (l *SinglyLinkedList) RemoveTail() interface{} {
	return nil
}

// String prints the string representation of the linked list
func (l *SinglyLinkedList) String() string {
	output := l.Name + ": [ "
	l.readerRestart()
	for l.reader.HasNext() || l.reader.index == l.Size-1 {
		output += fmt.Sprint(l.reader.current.Value)
		if l.reader.index < l.Size-1 {
			output += " -> "
		}
		l.reader.Next()
	}
	output += " ]"
	return output
}

// Next moves reader to Next SNode and returns
func (r *sReader) Next() bool {
	if r.current.Next != nil {
		r.current = r.current.Next
		r.index++
		return true
	}
	r.current = nil
	r.index = -1
	return false
}

// HasNext returns true if reader has Next
func (r *sReader) HasNext() bool {
	if r.current == nil {
		return false
	}
	if r.current.Next != nil {
		return true
	}
	return false
}

// readerRestart restarts traversal at the Head of the singly linked list
func (l *SinglyLinkedList) readerRestart() {
	l.reader.index = 0
	l.reader.current = l.Head
}
