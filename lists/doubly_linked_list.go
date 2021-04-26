package lists

import (
	"fmt"

	"gitlab.com/rshmhrj/go-text-gen/names"
)

// DoublyLinkedList represents a doubly linked list
type DoublyLinkedList struct {
	name   string
	size   int
	head   *dNode
	tail   *dNode
	reader *dReader
}

// dNode represents a node in a doubly linked list which knows about both the next node and the previous node in the list
type dNode struct {
	value interface{}
	next  *dNode
}

// dReader represents a helper for traversing the linked list
type dReader struct {
	index   int
	current *dNode
}

// NewDoublyLinkedList initializes an empty linked list of type Single or Double
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		name: names.NewGenerator().Generate(),
		size: 0,
		head: nil,
		tail: nil,
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
			l.head = n
		}
		l.tail = n
		if i == len(values)-1 {
			n.next = nil
		}
	}
	return l
}

// InitNamedDoublyLinkedList initializes a named linked list of type Single or Double with the given values
func InitNamedDoublyLinkedList(name string, value ...interface{}) *DoublyLinkedList {
	l := InitDoublyLinkedList(value...)
	l.name = name
	return l
}

// Len returns size of the linked list
func (l *DoublyLinkedList) Len() int {
	return l.size
}

// Head returns a pointer to the head dNode in the linked list
func (l *DoublyLinkedList) Head() *dNode {
	return l.head
}

// Tail returns a pointer to the tail dNode in the linked list
func (l *DoublyLinkedList) Tail() *dNode {
	return l.tail
}

// IsEmpty returns true if linked list size is 0
func (l *DoublyLinkedList) IsEmpty() bool {
	return l.size == 0
}

// Index of value in list.  If not found returns -1
func (l *DoublyLinkedList) IndexOf(value interface{}) int {
	output := -1
	l.readerRestart()
	for l.reader.HasNext() || l.reader.index == l.size-1 {
		if l.reader.current.value == value {
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

// SetName updates the name of the linked list
func (l *DoublyLinkedList) SetName(name string) {
	l.name = name
}

// Prepend adds a dNode as the new Head of the linked list
func (l *DoublyLinkedList) Prepend(value interface{}) {
	n := new(dNode)

	if l.size == 0 {
		n.next = nil
	}
	if l.size > 0 {
		n.next = l.head
	}

	n.value = value
	l.size += 1
}

// Append adds a dNode as the new Tail of the linked list
func (l *DoublyLinkedList) Append(value interface{}) *dNode {
	n := new(dNode)
	if l.size == 0 {
		l.head = n
	}
	if l.size == 1 {
		l.head.next = n
	}
	if l.size > 1 {
		l.readerRestart()
		for l.reader.HasNext() || l.reader.index == l.size-1 {
			if l.reader.index == l.size-1 {
				l.reader.current.next = n
				break
			}
			l.reader.Next()
		}
	}
	n.value = value
	n.next = nil
	l.tail = n
	l.size += 1

	return n
}

// Insert adds an dNode at the specified index in the linked list (index starting at 0)
func (l *DoublyLinkedList) Insert(index int, value interface{}) {
	n := new(dNode)
	l.readerRestart()
	prevTrav := l.head
	for l.reader.HasNext() || l.reader.index == l.size-1 {
		if l.reader.index == index {
			if l.reader.index == 0 {
				l.head = n
				break
			}
			n.value = value
			n.next = l.reader.current
			prevTrav.next = n
		}
		prevTrav = l.reader.current
		l.reader.Next()
	}
	l.size += 1
}

// RemoveHead removes the head of the linked list and RETURNS the new head
// If there is no element to remove, RETURNS error
func (l *DoublyLinkedList) RemoveHead() error {
	return nil
}

// RemoveAt removes the element at the given index in the linked list
// If there is no element to remove, or if the index is out of bounds, RETURNS error
func (l *DoublyLinkedList) RemoveAt(index int) error {
	return nil
}

// RemoveValue removes the first instance of the given value from the linked list
// If value is not found, RETURNS error
func (l *DoublyLinkedList) RemoveValue(value interface{}) error {
	return nil
}

// RemoveTail removes the tail of the linked list and RETURNS the new tail
// If there is no element to remove, RETURNS error
func (l *DoublyLinkedList) RemoveTail() error {
	return nil
}

// String prints the string representation of the linked list
func (l *DoublyLinkedList) String() string {
	output := l.name + ": [ "
	l.readerRestart()
	for l.reader.HasNext() || l.reader.index == l.size-1 {
		output += fmt.Sprint(l.reader.current.value)
		if l.reader.index < l.size-1 {
			output += " -> "
		}
		l.reader.Next()
	}
	output += " ]"
	return output
}

// Next moves reader to next dNode and returns
func (r *dReader) Next() bool {
	if r.current.next != nil {
		r.current = r.current.next
		r.index++
		return true
	}
	r.current = nil
	r.index = -1
	return false
}

// HasNext returns true if reader has next
func (r *dReader) HasNext() bool {
	if r.current == nil {
		return false
	}
	if r.current.next != nil {
		return true
	}
	return false
}

// readerRestart restarts traversal at the head of the doubly linked list
func (l *DoublyLinkedList) readerRestart() {
	l.reader.index = 0
	l.reader.current = l.head
}
