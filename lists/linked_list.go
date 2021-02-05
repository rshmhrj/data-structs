package lists

import (
 "fmt"
 
 "gitlab.com/rshmhrj/go-text-gen/names"
)

const (
 Single = "single"
 Double = "double"
)

type list struct {
 name string
 size int
 head *node
 tail *node
}

type linkedList struct {
 name string
 size int
 listType string
 head *node
 tail *node
 reader *reader
}

type node struct {
 data interface{}
 next *node
 prev *node
}

type reader struct {
 index int
 current *node
}

// Initialize an empty linked list
func NewEmptyLinkedList(listType string) *linkedList {
 return &linkedList{
  name: names.NewGenerator().Generate(),
  size: 0,
  listType: listType,
  head: nil,
  tail: nil,
  reader: &reader{
   index: -1,
   current: nil,
  },
 }
}

// Initialize a named linked list with values
func NewNamedLinkedList(listType, name string, data ... interface{}) *linkedList {
 l := NewLinkedList(listType, data ...)
 l.listType = listType
 l.listType = listType
 l.name = name
 return l
}

// Initialize a linked list with values
func NewLinkedList(listType string, data ... interface{}) *linkedList {
  l := NewEmptyLinkedList(listType)
  for i, v := range data {
   n := l.Add(v)
   if i == 0 {
    l.head = n
   }
   l.tail = n
   if i == len(data) -1 {
    n.next = nil
   }
 }
 return l
}

// Len returns size of the linked list
func (l *linkedList) Len() int {
 return l.size
}

// GetType returns type of linked list (Single / Double)
func (l *linkedList) GetType() string {
 return l.listType
}

// Head returns a pointer to the head node in the linked list
func (l *linkedList) Head() *node {
 return l.head
}

// Tail returns a pointer to the tail node in the linked list
func (l *linkedList) Tail() *node {
 return l.tail
}

// IsEmpty returns true if linked list size is 0
func (l *linkedList) IsEmpty() bool {
 return l.size == 0
}

// Index of value in list.  If not found returns -1
func (l *linkedList) IndexOf(data interface{}) int {
 output := -1
 l.readerRestart()
 for l.reader.HasNext() || l.reader.index == l.size -1 {
  if l.reader.current.data == data {
   output = l.reader.index
   break
  }
  l.reader.Next()
 }
 return output
}

// Contains returns true/false if item is in list
func (l *linkedList) Contains(value interface{}) bool {
 if l.IndexOf(value) != -1 {
  return true
 }
 return false
}

// SetName updates the name of the linked list
func (l *linkedList) SetName(name string) {
 l.name = name
}

// Add a node to the linked list
func (l *linkedList) Add(data interface{}) *node {
 n := new(node)
 
 if l.listType == Single {
  if l.size == 0 {
   l.head = n
  }
  if l.size == 1 {
   l.head.next = n
  }
  if l.size > 1 {
   l.readerRestart()
   for l.reader.HasNext() || l.reader.index == l.size -1 {
    if l.reader.index == l.size - 1 {
     l.reader.current.next = n
     break
    }
    l.reader.Next()
   }
  }
 }
 n.data = data
 n.next = nil
 l.tail = n
 l.size += 1
 return n
}

// Add a node at the specified index in the linked list (index starting at 0)
func (l *linkedList) AddAt(index int, data interface{}) *node {
 n := new(node)
 if l.listType == Single {
  l.readerRestart()
  prevTrav := l.head
  for l.reader.HasNext() || l.reader.index == l.size -1 {
   if l.reader.index == index {
    if l.reader.index == 0 {
     l.head = n
     break
    }
    n.data = data
    n.next = l.reader.current
    prevTrav.next = n
   }
   prevTrav = l.reader.current
   l.reader.Next()
  }
 }
 l.size += 1
 return n
}
// restart traversal
func (l *linkedList) readerRestart() {
 l.reader.index = 0
 l.reader.current = l.head
}

// String representation of linked list
func (l *linkedList) String() string {
 output := l.name + ": [ "
 l.readerRestart()
 for l.reader.HasNext() || l.reader.index == l.size -1 {
  output += fmt.Sprint(l.reader.current.data)
  if l.reader.index < l.size - 1 {
   if l.listType == Single {
    output += " -> "
   }
   if l.listType == Double {
    output += " <-> "
   }
  }
  l.reader.Next()
 }
 output += " ]"
 return output
}

// Next moves reader to next node and returns
func (r *reader) Next() bool {
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
func (r reader) HasNext() bool {
 if r.current == nil {
  return false
 }
 if r.current.next != nil {
  return true
 }
 return false
}

