package lists

import (
	"strings"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	got := NewDoublyLinkedList()
	wanted := DoublyLinkedList{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
	if got.Head != wanted.Head {
		t.Errorf("Incorrectly created Head, got: %v, wanted: %v", got.Head, wanted.Head)
	}
	if got.Tail != wanted.Tail {
		t.Errorf("Incorrectly created Head, got: %v, wanted: %v", got.Tail, wanted.Tail)
	}
	if got.Size != wanted.Size {
		t.Errorf("Incorrectly created Head, got: %v, wanted: %v", got.Size, wanted.Size)
	}
}

func TestInitDoublyLinkedList(t *testing.T) {
	got := InitDoublyLinkedList(10, 20)
	node2 := &DNode{
		Value: 20,
		Next:  nil,
	}
	node1 := &DNode{
		Value: 10,
		Next:  node2,
	}
	reader1 := &dReader{
		index:   0,
		current: node1,
	}
	wanted := DoublyLinkedList{
		Name:   "wanted",
		Size:   2,
		Head:   node1,
		Tail:   node2,
		reader: reader1,
	}
	if got.Size != wanted.Size {
		t.Errorf("Incorrectly set Size, got: %v, wanted: %v", got.Size, wanted.Size)
	}
}

func TestDoublyLinkedList_Append(t *testing.T) {
	got := NewDoublyLinkedList()
	got.Append("A")
	if got.Size != 1 {
		t.Errorf("Incorrectly set Size, got: %v, wanted: %v", got.Size, 1)
	}
	if got.Head.Value != "A" {
		t.Errorf("Incorrectly set Value, got: %v, wanted: %v", got.Head.Value, "A")
	}
	if got.Head.Next != nil {
		t.Errorf("Incorrectly set Next, got: %v, wanted: %v", got.Head.Next, nil)
	}
	got.Append("B")
	if got.Size != 2 {
		t.Errorf("Incorrectly set Size, got: %v, wanted: %v", got.Size, 2)
	}
	if got.Head.Value != "A" {
		t.Errorf("Incorrectly set Value, got: %v, wanted: %v", got.Head.Value, "A")
	}
	if got.Head.Next != got.Tail {
		t.Errorf("Incorrectly set Next, got: %v, wanted: %v", got.Head.Next, got.Tail)
	}
	if got.Tail.Value != "B" {
		t.Errorf("Incorrectly set Tail, got: %v, wanted: %v", got.Tail.Value, "B")
	}
	if got.Tail.Next != nil {
		t.Errorf("Incorrectly set Tail Next, got: %v, wanted: %v", got.Tail.Next, nil)
	}
	got.Append("C")
	if got.Size != 3 {
		t.Errorf("Incorrectly set Size, got: %v, wanted: %v", got.Size, 3)
	}
	if got.Head.Value != "A" {
		t.Errorf("Incorrectly set Value, got: %v, wanted: %v", got.Head.Value, "A")
	}
	if got.Head.Next.Value != "B" {
		t.Errorf("Incorrectly set Next, got: %v, wanted: %v", got.Head.Next.Value, "B")
	}
	if got.Head.Next.Next != got.Tail {
		t.Errorf("Incorrectly set Next, got: %v, wanted: %v", got.Head.Next.Next, got.Tail)
	}
	if got.Tail.Value != "C" {
		t.Errorf("Incorrectly set Tail, got: %v, wanted: %v", got.Tail.Value, "C")
	}
	if got.Tail.Next != nil {
		t.Errorf("Incorrectly set Tail Next, got: %v, wanted: %v", got.Tail.Next, nil)
	}
}

func TestDoublyLinkedList_Insert(t *testing.T) {
	got := InitDoublyLinkedList("A", "C")
	got.Insert(1, "B")
	if got.Size != 3 {
		t.Errorf("Incorrectly set Size, got: %v, wanted: %v", got.Size, 3)
	}
	if got.Head.Value != "A" {
		t.Errorf("Incorrectly set Value, got: %v, wanted: %v", got.Head.Value, "A")
	}
	if got.Head.Next.Value != "B" {
		t.Errorf("Incorrectly set Next, got: %v, wanted: %v", got.Head.Next.Value, "B")
	}
	if got.Head.Next.Next != got.Tail {
		t.Errorf("Incorrectly set Next, got: %v, wanted: %v", got.Head.Next.Next, got.Tail)
	}
	if got.Tail.Value != "C" {
		t.Errorf("Incorrectly set Tail, got: %v, wanted: %v", got.Tail.Value, "C")
	}
	if got.Tail.Next != nil {
		t.Errorf("Incorrectly set Tail Next, got: %v, wanted: %v", got.Tail.Next, nil)
	}
}

func TestDoublyLinkedList_String(t *testing.T) {
	l := InitDoublyLinkedList("A", "B", "C")
	output := l.String()
	got := output[strings.Index(output, ": ")+2:]
	wanted := "[ A <-> B <-> C ]"
	if got != wanted {
		t.Errorf("Incorrectly printing string, got: %v, wanted: %v", got, wanted)
	}
}

func TestInitNamedDoublyLinkedList(t *testing.T) {
	l := InitNamedDoublyLinkedList("new.test.list", "A", "B", "C")
	got := l.String()
	wanted := "new.test.list: [ A <-> B <-> C ]"
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
