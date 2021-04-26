package lists

import (
	"fmt"

	"gitlab.com/rshmhrj/go-text-gen/names"
)

// SinglyLinkedList represents a singly linked list
type SinglyLinkedList struct {
	name   string
	size   int
	head   *sNode
	tail   *sNode
	reader *sReader
}

// sNode represents a node in a singly linked list which only knows about the next node in the list
type sNode struct {
	value interface{}
	next  *sNode
}

// sReader represents a helper for traversing the linked list
type sReader struct {
	index   int
	current *sNode
}

// NewSinglyLinkedList initializes an empty linked list of type Single or Double
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		name: names.NewGenerator().Generate(),
		size: 0,
		head: nil,
		tail: nil,
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
			l.head = n
		}
		l.tail = n
		if i == len(values)-1 {
			n.next = nil
		}
	}
	return l
}

// InitNamedSinglyLinkedList initializes a named linked list of type Single or Double with the given values
func InitNamedSinglyLinkedList(name string, value ...interface{}) *SinglyLinkedList {
	l := InitSinglyLinkedList(value...)
	l.name = name
	return l
}

// Len returns size of the linked list
func (l *SinglyLinkedList) Len() int {
	return l.size
}

// Head returns a pointer to the head sNode in the linked list
func (l *SinglyLinkedList) Head() *sNode {
	return l.head
}

// Tail returns a pointer to the tail sNode in the linked list
func (l *SinglyLinkedList) Tail() *sNode {
	return l.tail
}

// IsEmpty returns true if linked list size is 0
func (l *SinglyLinkedList) IsEmpty() bool {
	return l.size == 0
}

// Index of value in list.  If not found returns -1
func (l *SinglyLinkedList) IndexOf(value interface{}) int {
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
func (l *SinglyLinkedList) Contains(value interface{}) bool {
	if l.IndexOf(value) != -1 {
		return true
	}
	return false
}

// SetName updates the name of the linked list
func (l *SinglyLinkedList) SetName(name string) {
	l.name = name
}

// Prepend adds a sNode as the new Head of the linked list
func (l *SinglyLinkedList) Prepend(value interface{}) {
	n := new(sNode)

	if l.size == 0 {
		n.next = nil
	}
	if l.size > 0 {
		n.next = l.head
	}

	n.value = value
	l.size += 1
}

// Append adds a sNode as the new Tail of the linked list
func (l *SinglyLinkedList) Append(value interface{}) *sNode {
	n := new(sNode)
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

// Insert adds an sNode at the specified index in the linked list (index starting at 0)
func (l *SinglyLinkedList) Insert(index int, value interface{}) {
	n := new(sNode)
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
func (l *SinglyLinkedList) RemoveHead() error {
	return nil
}

// RemoveAt removes the element at the given index in the linked list
// If there is no element to remove, or if the index is out of bounds, RETURNS error
func (l *SinglyLinkedList) RemoveAt(index int) error {
	return nil
}

// RemoveValue removes the first instance of the given value from the linked list
// If value is not found, RETURNS error
func (l *SinglyLinkedList) RemoveValue(value interface{}) error {
	return nil
}

// RemoveTail removes the tail of the linked list and RETURNS the new tail
// If there is no element to remove, RETURNS error
func (l *SinglyLinkedList) RemoveTail() error {
	return nil
}

// String prints the string representation of the linked list
func (l *SinglyLinkedList) String() string {
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

// Next moves reader to next sNode and returns
func (r *sReader) Next() bool {
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
func (r *sReader) HasNext() bool {
	if r.current == nil {
		return false
	}
	if r.current.next != nil {
		return true
	}
	return false
}

// readerRestart restarts traversal at the head of the singly linked list
func (l *SinglyLinkedList) readerRestart() {
	l.reader.index = 0
	l.reader.current = l.head
}
