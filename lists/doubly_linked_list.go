package lists

import (
	"fmt"

	"gitlab.com/rshmhrj/go-text-gen/names"
)

// DoublyLinkedList represents a linked list which can be traversed in both directions.
// Time Complexity: Access O(n), Search O(n), Insert O(1), Delete O(1).
// Space Complexity: O(n).
type DoublyLinkedList struct {
	Name   string
	Size   int
	Head   *DNode
	Tail   *DNode
	reader *dReader
}

// DNode represents a node in a doubly linked list which knows about both the Next node and the previous node in the list
type DNode struct {
	Value interface{}
	Next  *DNode
}

// dReader represents a helper for traversing the linked list
type dReader struct {
	index   int
	current *DNode
}

// NewDoublyLinkedList initializes an empty linked list of type Single or Double
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Name: names.NewGenerator().Generate(),
		Size: 0,
		Head: nil,
		Tail: nil,
		reader: &dReader{
			index:   -1,
			current: nil,
		},
	}
}

// InitDoublyLinkedList initializes a linked list of type Single or Double with the given values
func InitDoublyLinkedList(values ...interface{}) *DoublyLinkedList {
	l := NewDoublyLinkedList()
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

// InitNamedDoublyLinkedList initializes a named linked list of type Single or Double with the given values
func InitNamedDoublyLinkedList(name string, value ...interface{}) *DoublyLinkedList {
	l := InitDoublyLinkedList(value...)
	l.Name = name
	return l
}

// Len returns Size of the linked list
func (l *DoublyLinkedList) Len() int {
	return l.Size
}

// IsEmpty returns true if linked list Size is 0
func (l *DoublyLinkedList) IsEmpty() bool {
	return l.Size == 0
}

// Index of Value in list.  If not found returns -1
func (l *DoublyLinkedList) IndexOf(value interface{}) int {
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
func (l *DoublyLinkedList) Contains(value interface{}) bool {
	if l.IndexOf(value) != -1 {
		return true
	}
	return false
}

// SetName updates the Name of the linked list
func (l *DoublyLinkedList) SetName(name string) {
	l.Name = name
}

// Prepend adds a DNode as the new Head of the linked list
func (l *DoublyLinkedList) Prepend(value interface{}) *DNode {
	n := new(DNode)

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

// Append adds a DNode as the new Tail of the linked list
func (l *DoublyLinkedList) Append(value interface{}) *DNode {
	n := new(DNode)
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

// Insert adds an DNode at the specified index in the linked list (index starting at 0)
func (l *DoublyLinkedList) Insert(index int, value interface{}) {
	n := new(DNode)
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
func (l *DoublyLinkedList) RemoveHead() interface{} {
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
func (l *DoublyLinkedList) RemoveAt(index int) interface{} {
	return nil
}

// RemoveValue removes the first instance of the given Value from the linked list
// If Value is not found, RETURNS error
func (l *DoublyLinkedList) RemoveValue(value interface{}) interface{} {
	return nil
}

// RemoveTail removes the Tail of the linked list and RETURNS it's Value
// If there is no element to remove, RETURNS error
func (l *DoublyLinkedList) RemoveTail() interface{} {
	return nil
}

// String prints the string representation of the linked list
func (l *DoublyLinkedList) String() string {
	output := l.Name + ": [ "
	l.readerRestart()
	for l.reader.HasNext() || l.reader.index == l.Size-1 {
		output += fmt.Sprint(l.reader.current.Value)
		if l.reader.index < l.Size-1 {
			output += " <-> "
		}
		l.reader.Next()
	}
	output += " ]"
	return output
}

// Next moves reader to Next DNode and returns
func (r *dReader) Next() bool {
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
func (r *dReader) HasNext() bool {
	if r.current == nil {
		return false
	}
	if r.current.Next != nil {
		return true
	}
	return false
}

// readerRestart restarts traversal at the Head of the doubly linked list
func (l *DoublyLinkedList) readerRestart() {
	l.reader.index = 0
	l.reader.current = l.Head
}
