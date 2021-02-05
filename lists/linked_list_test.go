package lists

import (
  "strings"
  "testing"
)

func TestNewEmptyLinkedList_Single(t *testing.T) {
  got := NewEmptyLinkedList(Single)
  wanted := linkedList{
    size: 0,
    head: nil,
    tail: nil,
  }
  if got.head != wanted.head {
    t.Errorf("Incorrectly created head, got: %v, wanted: %v", got.head, wanted.head)
  }
  if got.tail != wanted.tail {
    t.Errorf("Incorrectly created head, got: %v, wanted: %v", got.tail, wanted.tail)
  }
  if got.size != wanted.size {
    t.Errorf("Incorrectly created head, got: %v, wanted: %v", got.size, wanted.size)
  }
}

func TestLinkedList_Add_Single(t *testing.T) {
  got := NewEmptyLinkedList(Single)
  got.Add("A")
  if got.size != 1 {
    t.Errorf("Incorrectly set size, got: %v, wanted: %v", got.size, 1)
  }
  if got.head.data != "A" {
    t.Errorf("Incorrectly set data, got: %v, wanted: %v", got.head.data, "A")
  }
  if got.head.next != nil {
    t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next, nil)
  }
  got.Add("B")
  if got.size != 2 {
    t.Errorf("Incorrectly set size, got: %v, wanted: %v", got.size, 2)
  }
  if got.head.data != "A" {
    t.Errorf("Incorrectly set data, got: %v, wanted: %v", got.head.data, "A")
  }
  if got.head.next != got.tail {
    t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next, got.tail)
  }
  if got.tail.data != "B" {
    t.Errorf("Incorrectly set tail, got: %v, wanted: %v", got.tail.data, "B")
  }
  if got.tail.next != nil {
    t.Errorf("Incorrectly set tail next, got: %v, wanted: %v", got.tail.next, nil)
  }
  got.Add("C")
  if got.size != 3 {
    t.Errorf("Incorrectly set size, got: %v, wanted: %v", got.size, 3)
  }
  if got.head.data != "A" {
    t.Errorf("Incorrectly set data, got: %v, wanted: %v", got.head.data, "A")
  }
  if got.head.next.data != "B" {
    t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next.data, "B")
  }
  if got.head.next.next != got.tail {
    t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next.next, got.tail)
  }
  if got.tail.data != "C" {
    t.Errorf("Incorrectly set tail, got: %v, wanted: %v", got.tail.data, "C")
  }
  if got.tail.next != nil {
    t.Errorf("Incorrectly set tail next, got: %v, wanted: %v", got.tail.next, nil)
  }
}
func TestLinkedList_AddAt_Single(t *testing.T) {
  got := NewLinkedList(Single,"A", "C")
  got.AddAt(1, "B")
  if got.size != 3 {
    t.Errorf("Incorrectly set size, got: %v, wanted: %v", got.size, 3)
  }
  if got.head.data != "A" {
    t.Errorf("Incorrectly set data, got: %v, wanted: %v", got.head.data, "A")
  }
  if got.head.next.data != "B" {
    t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next.data, "B")
  }
  if got.head.next.next != got.tail {
    t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next.next, got.tail)
  }
  if got.tail.data != "C" {
    t.Errorf("Incorrectly set tail, got: %v, wanted: %v", got.tail.data, "C")
  }
  if got.tail.next != nil {
    t.Errorf("Incorrectly set tail next, got: %v, wanted: %v", got.tail.next, nil)
  }
}

func TestLinkedList_String_Single(t *testing.T) {
  l := NewLinkedList(Single, "A", "B", "C")
  output := l.String()
  got := output[strings.Index(output, ": ")+2:]
  wanted := "[ A -> B -> C ]"
  if got != wanted {
    t.Errorf("Incorrectly printing string, got: %v, wanted: %v", got, wanted)
  }
}

func TestNewNamedLinkedList_Single(t *testing.T) {
  l := NewNamedLinkedList(Single, "new.test.list", "A", "B", "C")
  got := l.String()
  wanted := "new.test.list: [ A -> B -> C ]"
  if got != wanted {
    t.Errorf("Incorrectly printing string, got: %v, wanted: %v", got, wanted)
  }
}

func TestLinkedList_IndexOf_Single(t *testing.T) {
  l := NewLinkedList(Single,"A", "B", "C")
  got := l.IndexOf("D")
  wanted := -1
  if got != wanted {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  got = l.IndexOf("B")
  wanted = 1
  if got != wanted {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  got = l.IndexOf("C")
  wanted = 2
  if got != wanted {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  got = l.IndexOf("A")
  wanted = 0
  if got != wanted {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  l.Add("D")
  got = l.IndexOf("D")
  wanted = 3
  if got != wanted {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  l.AddAt(2,"NEW")
  got = l.IndexOf("NEW")
  wanted = 2
  if got != wanted {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
}

func TestLinkedList_Contains_Single(t *testing.T) {
  l := NewLinkedList(Single,"A", "B", "C")
  got := l.Contains("D")
  wanted := false
  if got {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  got = l.Contains("B")
  wanted = true
  if !got {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  got = l.Contains("C")
  wanted = true
  if !got {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  got = l.Contains("A")
  wanted = true
  if !got {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  l.Add("D")
  got = l.Contains("D")
  wanted = true
  if !got {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  got = l.Contains("NEW")
  wanted = false
  if got {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
  l.AddAt(2,"NEW")
  got = l.Contains("NEW")
  wanted = true
  if !got {
    t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
  }
}
