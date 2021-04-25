package arrays

import (
	"math"
	"testing"
)

func TestInitDynamicArray(t *testing.T) {
	a := NewDynamicArray(10)
	if a.len != 0 {
		t.Errorf("Incorrectly assigned length, got: %v, wanted: %v", a.len, 0)
	}
	if a.cap != 10 {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, 10)
	}
	for i, v := range a.data {
		if v != nil {
			t.Errorf("Incorrectly created default empty dynamic array, got: %v at index %v, wanted %v", v, i, 0)
		}
	}
}

func TestMakeDynamicArray(t *testing.T) {
	a := InitDynamicArray(10, 20, 30, 40, 50)
	if a.len != 5 {
		t.Errorf("Incorrectly assigned length, got: %v, wanted: %v", a.len, 5)
	}
	if a.cap != 5 {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, 5)
	}
	for i, v := range a.data {
		if v != (i+1)*10 {
			t.Errorf("Incorrectly created default empty dynamic array, got: %v at index %v, wanted %v", v, i, (i+1)*10)
		}
	}
}

func TestDynamicArray_Set(t *testing.T) {
	a := NewDynamicArray(5)
	a.Set(0, 10)
	if a.data[0] != 10 {
		t.Errorf("Incorrectly setting value in array, got: %v, wanted: %v", a.data[0], 10)
	}
	if a.len != 1 {
		t.Errorf("Incorrectly updating length after set, got: %v, wanted: %v", a.len, 1)
	}
	a.Set(1, 20)
	if a.data[1] != 20 {
		t.Errorf("Incorrectly setting value in array, got: %v, wanted: %v", a.data[1], 20)
	}
	if a.len != 2 {
		t.Errorf("Incorrectly updating length after set, got: %v, wanted: %v", a.len, 2)
	}
	a.Set(2, 30)
	if a.data[2] != 30 {
		t.Errorf("Incorrectly setting value in array, got: %v, wanted: %v", a.data[2], 30)
	}
	if a.len != 3 {
		t.Errorf("Incorrectly updating length after set, got: %v, wanted: %v", a.len, 3)
	}
	a.Set(3, 40)
	if a.data[3] != 40 {
		t.Errorf("Incorrectly setting value in array, got: %v, wanted: %v", a.data[3], 40)
	}
	if a.len != 4 {
		t.Errorf("Incorrectly updating length after set, got: %v, wanted: %v", a.len, 3)
	}
	a.Set(4, 50)
	if a.data[4] != 50 {
		t.Errorf("Incorrectly setting value in array, got: %v, wanted: %v", a.data[4], 50)
	}
	if a.len != 5 {
		t.Errorf("Incorrectly updating length after set, got: %v, wanted: %v", a.len, 5)
	}
}

func TestDynamicArray_IsEmpty(t *testing.T) {
	a := NewDynamicArray(5)
	if a.IsEmpty() != true {
		t.Errorf("Incorrectly checking whether emptiness of new blank array, got: %v, wanted %v", a.IsEmpty(), true)
	}
	a.Set(0, 10)
	if a.IsEmpty() != false {
		t.Errorf("Incorrectly checking whether emptiness of array with a value, got: %v, wanted %v", a.IsEmpty(), false)
	}
	a.Clear()
	if a.IsEmpty() != true {
		t.Errorf("Incorrectly checking whether emptiness of cleared array, got: %v, wanted %v", a.IsEmpty(), true)
	}
}

func TestDynamicArray_Clear(t *testing.T) {
	a := InitDynamicArray(10, 20, 30)
	a.Clear()
	if a.len != 0 {
		t.Errorf("Incorrectly cleared array length, got: %v, wanted: %v", a.len, 0)
	}
	if a.data != nil {
		t.Errorf("Incorrectly cleared array, got: %v, wated: %v", a.data, nil)
	}
}

func TestDynamicArray_Remove(t *testing.T) {
	a := InitDynamicArray(10, 20, 30, 40, 50)
	a.Remove(20)
	if a.data[0] != 10 && a.data[1] != 30 && a.data[2] != 40 && a.data[3] != 50 && a.data[4] != 0 {
		t.Errorf("Incorrectly removed value from array, got: %v, wanted: %v", a.data, []int{10, 30, 40, 50, 0})
	}
	if a.len != 4 {
		t.Errorf("Incorrectly updated length after removing value, got %v, wanted: %v", a.len, 4)
	}
}

func TestDynamicArray_RemoveAt(t *testing.T) {
	a := InitDynamicArray(10, 20, 30, 40, 50)
	a.RemoveAt(2)
	if a.data[0] != 10 && a.data[1] != 30 && a.data[2] != 40 && a.data[3] != 50 && a.data[4] != 0 {
		t.Errorf("Incorrectly removed value from array, got: %v, wanted: %v", a.data, []int{10, 30, 40, 50, 0})
	}
	if a.len != 4 {
		t.Errorf("Incorrectly updated length after removing value at index, got %v, wanted: %v", a.len, 4)
	}
}

func TestDynamicArray_IndexOf(t *testing.T) {
	a := InitDynamicArray(10, 20, 30, 40, 50)
	if a.IndexOf(20) != 1 {
		t.Errorf("Incorrectly found index of value in array, got: %v, wanted %v", a.IndexOf(20), 1)
	}
	if a.IndexOf(60) != -1 {
		t.Errorf("Incorrectly found index of value in array, got: %v, wanted %v", a.IndexOf(60), -1)
	}
}

func TestDynamicArray_Contains(t *testing.T) {
	a := InitDynamicArray(10, 20, 30, 40, 50)
	if a.Contains(20) != true {
		t.Errorf("Incorrectly found value in array, got: %v for value %v in array %v, wanted %v", a.Contains(20), 20, a.data, true)
	}
	if a.Contains(60) != false {
		t.Errorf("Incorrectly found value in array, got: %v for value %v in array %v, wanted %v", a.Contains(60), 60, a.data, false)
	}
}

func TestDynamicArray_Scale(t *testing.T) {
	size := 5
	a := NewDynamicArray(size)
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 10
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 15
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 22
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 33
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 49
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 73
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 109
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 136
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 170
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 212
	a.Scale()
	if a.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", a.cap, size)
	}
	size = 1000
	b := NewDynamicArray(size)
	size = 1100
	b.Scale()
	if b.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", b.cap, size)
	}

	size = 10000
	c := NewDynamicArray(size)
	size = 11000
	c.Scale()
	if c.cap != size {
		t.Errorf("Incorrectly assigned capacity, got: %v, wanted: %v", c.cap, size)
	}
}

func TestDynamicArray_String(t *testing.T) {
	a := InitDynamicArray(10, 20, 30)
	got := a.String()
	wanted := "[ 10, 20, 30 ]"
	if got != wanted {
		t.Errorf("Incorrectly formatted string representation, got: %v, wanted: %v", got, wanted)
	}
}

// If you do RemoveAt(0) it'll hit line 87 making j = -1. Then on line 89 tmp[-1] will cause an index issue
func TestDynamicArray_RemoveAtIndex0(t *testing.T) {
	a := InitDynamicArray(11, 22, 33, 44, 55)
	a.RemoveAt(0)
	got := a.String()
	wanted := "[ 22, 33, 44, 55 ]"
	if got != wanted {
		t.Errorf("Incorrectly removed index 0, got: %v, wanted: %v", got, wanted)
	}
}

// If you have something like [11, 22, 33, 44, 55] and do RemoveAt(-4) it won't hit the panic statement in line 81
func TestDynamicArray_RemoveAtNegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Incorrectly handled panic")
		}
	}()

	a := InitDynamicArray(11, 22, 33, 44, 55)
	a.RemoveAt(-4)
	a.RemoveAt(-3)
	a.RemoveAt(-2)
	a.RemoveAt(-1)
	a.RemoveAt(-0)
}

func TestDynamicArray_AddAt(t *testing.T) {
	got := InitDynamicArray(20, 30, 40, 50)
	got.AddAt(0, 10)
	wanted := InitDynamicArray(10, 20, 30, 40, 50)
	if got.String() != wanted.String() {
		t.Errorf("Incorrectly added at index 0, got: %v, wanted: %v", got.String(), wanted.String())
	}
}

func TestCreatingBigAssArrays(t *testing.T) {
	maxInt32Size := math.MaxInt32   // 2147483647
	maxUint32Size := math.MaxUint32 // 4294967295
	got := NewDynamicArray(maxInt32Size)
	wanted := maxInt32Size
	if got.cap != wanted {
		t.Errorf("Could not handle creating big array, got: %v, wanted: %v", got.cap, wanted)
	}
	got = NewDynamicArray(maxInt32Size + 1)
	wanted = maxInt32Size + 1
	if got.cap != wanted {
		t.Errorf("Could not handle creating big array, got: %v, wanted: %v", got.cap, wanted)
	}
	got = NewDynamicArray(maxUint32Size)
	wanted = maxUint32Size
	if got.cap != wanted {
		t.Errorf("Could not handle creating big array, got: %v, wanted: %v", got.cap, wanted)
	}
	got = NewDynamicArray(maxUint32Size + 1)
	wanted = maxUint32Size + 1
	if got.cap != wanted {
		t.Errorf("Could not handle creating big array, got: %v, wanted: %v", got.cap, wanted)
	}
}

func BenchmarkNewDynamicArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDynamicArray(10)
	}
}
