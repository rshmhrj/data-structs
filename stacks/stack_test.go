package stacks

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/rshmhrj/data-structs/lists"
)

func TestNewStack(t *testing.T) {
	got := NewStack()
	wanted := lists.NewSinglyLinkedList()
	if got.Size() != wanted.Size {
		t.Errorf("Incorrectly set length of new empty stack, got: %v, wanted: %v", got.Size(), wanted.Size)
	}
	if got.stack.Head != nil {
		t.Errorf("Incorrectly set top value of new empty stack, got: %v, wanted: %v", got.stack.Head, nil)
	}
}

func TestInitStack(t *testing.T) {
	got := InitStack(10, 20, 30, 40, 50)
	wanted := lists.InitSinglyLinkedList(50, 40, 30, 20, 10)

	if got.Size() != wanted.Len() {
		t.Errorf("Incorrectly set length of new empty stack, got: %v, wanted: %v", got.Size(), wanted.Len())
	}
	if got.Peek() != wanted.Head.Value {
		t.Errorf("Incorrectly set top value of new empty stack, got: %v, wanted: %v", got.Peek(), wanted.Head.Value)
	}
	wantedString := fmt.Sprintf("{ Top: 50, Length: 5, %v: [ 50 -> 40 -> 30 -> 20 -> 10 ] }", got.stack.Name)
	if strings.Compare(got.String(), wantedString) != 0 {
		t.Errorf("Incorrectly created new empty stack, \ngot: \t\t%v (len: %v) \nwanted: \t%v (len: %v)\nCompare: %v",
			got, len(got.String()), wantedString, len(wantedString), strings.Compare(got.String(), wantedString))
	}
}

func TestStack_String(t *testing.T) {
	a := InitStack(10, 20, 30, 40, 50)
	got := a.String()
	wantedString := fmt.Sprintf("{ Top: 50, Length: 5, %v: [ 50 -> 40 -> 30 -> 20 -> 10 ] }", a.stack.Name)
	if strings.Compare(got, wantedString) != 0 {
		t.Errorf("Incorrectly created string, \ngot: \t\t%v\nwanted: \t%v", got, wantedString)
	}
}

func TestIntStack_Push(t *testing.T) {
	got := InitStack(10, 20, 30, 40)
	got.Push(50)
	wantedString := fmt.Sprintf("{ Top: 50, Length: 5, %v: [ 50 -> 40 -> 30 -> 20 -> 10 ] }", got.stack.Name)
	if got.Size() != 5 {
		t.Errorf("Incorrectly set length of stack with new push, got: %v, wanted: %v", got.Size(), 5)
	}
	if got.Peek() != 50 {
		t.Errorf("Incorrectly set top value for stack with new push, got: %v, wanted: %v", got.Peek(), 50)
	}
	if strings.Compare(got.String(), wantedString) != 0 {
		t.Errorf("Incorrectly created string, \ngot: \t\t%v\nwanted: \t%v", got.String(), wantedString)
	}

}

//func TestIntStack_Pop(t *testing.T) {
//
//}
//
//func TestIntStack_Peek(t *testing.T) {
//
//}
//
//func TestIntStack_Search(t *testing.T) {
//
//}
//
//func TestIntStack_Size(t *testing.T) {
//
//}
//
//func TestIntStack_IsEmpty(t *testing.T) {
//
//}
