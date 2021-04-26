package lists

import (
	"strings"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	got := NewDoublyLinkedList()
	wanted := DoublyLinkedList{
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

func TestDoublyLinkedList_Append(t *testing.T) {
	got := NewDoublyLinkedList()
	got.Append("A")
	if got.size != 1 {
		t.Errorf("Incorrectly set size, got: %v, wanted: %v", got.size, 1)
	}
	if got.head.value != "A" {
		t.Errorf("Incorrectly set value, got: %v, wanted: %v", got.head.value, "A")
	}
	if got.head.next != nil {
		t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next, nil)
	}
	got.Append("B")
	if got.size != 2 {
		t.Errorf("Incorrectly set size, got: %v, wanted: %v", got.size, 2)
	}
	if got.head.value != "A" {
		t.Errorf("Incorrectly set value, got: %v, wanted: %v", got.head.value, "A")
	}
	if got.head.next != got.tail {
		t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next, got.tail)
	}
	if got.tail.value != "B" {
		t.Errorf("Incorrectly set tail, got: %v, wanted: %v", got.tail.value, "B")
	}
	if got.tail.next != nil {
		t.Errorf("Incorrectly set tail next, got: %v, wanted: %v", got.tail.next, nil)
	}
	got.Append("C")
	if got.size != 3 {
		t.Errorf("Incorrectly set size, got: %v, wanted: %v", got.size, 3)
	}
	if got.head.value != "A" {
		t.Errorf("Incorrectly set value, got: %v, wanted: %v", got.head.value, "A")
	}
	if got.head.next.value != "B" {
		t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next.value, "B")
	}
	if got.head.next.next != got.tail {
		t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next.next, got.tail)
	}
	if got.tail.value != "C" {
		t.Errorf("Incorrectly set tail, got: %v, wanted: %v", got.tail.value, "C")
	}
	if got.tail.next != nil {
		t.Errorf("Incorrectly set tail next, got: %v, wanted: %v", got.tail.next, nil)
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	got := InitDoublyLinkedList("A", "C")
	got.Insert(1, "B")
	if got.size != 3 {
		t.Errorf("Incorrectly set size, got: %v, wanted: %v", got.size, 3)
	}
	if got.head.value != "A" {
		t.Errorf("Incorrectly set value, got: %v, wanted: %v", got.head.value, "A")
	}
	if got.head.next.value != "B" {
		t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next.value, "B")
	}
	if got.head.next.next != got.tail {
		t.Errorf("Incorrectly set next, got: %v, wanted: %v", got.head.next.next, got.tail)
	}
	if got.tail.value != "C" {
		t.Errorf("Incorrectly set tail, got: %v, wanted: %v", got.tail.value, "C")
	}
	if got.tail.next != nil {
		t.Errorf("Incorrectly set tail next, got: %v, wanted: %v", got.tail.next, nil)
	}
}

func TestDoublyLinkedList_String(t *testing.T) {
	l := InitDoublyLinkedList("A", "B", "C")
	output := l.String()
	got := output[strings.Index(output, ": ")+2:]
	wanted := "[ A -> B -> C ]"
	if got != wanted {
		t.Errorf("Incorrectly printing string, got: %v, wanted: %v", got, wanted)
	}
}

func TestInitNamedDoublyLinkedList(t *testing.T) {
	l := InitNamedDoublyLinkedList("new.test.list", "A", "B", "C")
	got := l.String()
	wanted := "new.test.list: [ A -> B -> C ]"
	if got != wanted {
		t.Errorf("Incorrectly printing string, got: %v, wanted: %v", got, wanted)
	}
}

func TestDoublyLinkedList_IndexOf(t *testing.T) {
	l := InitDoublyLinkedList("A", "B", "C")
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
	l.Append("D")
	got = l.IndexOf("D")
	wanted = 3
	if got != wanted {
		t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
	}
	l.Insert(2, "NEW")
	got = l.IndexOf("NEW")
	wanted = 2
	if got != wanted {
		t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
	}
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	l := InitDoublyLinkedList("A", "B", "C")
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
	l.Append("D")
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
	l.Insert(2, "NEW")
	got = l.Contains("NEW")
	wanted = true
	if !got {
		t.Errorf("Incorrectly found Index, got: %v, wanted: %v", got, wanted)
	}
}
