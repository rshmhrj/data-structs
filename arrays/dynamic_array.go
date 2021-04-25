package arrays

import (
	"fmt"
)

type DynamicArray struct {
	len  int
	cap  int
	data []interface{}
}

func MakeDynamicArray(array ...interface{}) *DynamicArray {
	return &DynamicArray{
		len:  len(array),
		cap:  len(array),
		data: array,
	}
}

func InitDynamicArray(size int) *DynamicArray {
	if size < 0 {
		panic("Illegal size chosen for new array")
		return nil
	}
	return &DynamicArray{
		len:  0,
		cap:  size,
		data: make([]interface{}, size),
	}
}

// Len returns the current size of the array
func (da *DynamicArray) Len() int {
	return da.len
}

// Cap returns the current capacity of the array
func (da *DynamicArray) Cap() int {
	return da.cap
}

// IsEmpty returns whether the array contains any values or not
func (da *DynamicArray) IsEmpty() bool {
	if da.len == 0 {
		return true
	}
	return false
}

// Get Value at index
func (da *DynamicArray) Get(index int) interface{} {
	return da.data[index]
}

// GetFirst returns a pointer to the value at index 0
func (da *DynamicArray) GetFirst() *interface{} {
	return &da.data[0]
}

// Set Value at index
func (da *DynamicArray) Set(index int, value interface{}) {
	if da.data[index] == nil {
		da.len += 1
	}
	da.data[index] = value
}

// Clear Array
func (da *DynamicArray) Clear() {
	da.data = nil
	da.len = 0
}

// Add new value to the array
func (da *DynamicArray) Add(value interface{}) {
	if da.len+1 > da.cap {
		da.Scale()
	}
	da.data[da.len+1] = value
	da.len += 1
}

// AddAt adds a new value to the array at the given index
func (da *DynamicArray) AddAt(index int, value interface{}) {
	if index > da.len || index < 0 {
		panic("Index out of bounds")
	}
	tmp := make([]interface{}, da.cap+1)
	for i, j := 0, 0; i < da.len; i++ {
		if i == index {
			tmp[j] = value
			j++
		}
		tmp[j] = da.data[i]
		j++
	}
	da.data = tmp
	da.len += 1
}

// RemoveAt removes the value at the given index
func (da *DynamicArray) RemoveAt(index int) {
	if index > da.len || index < 0 {
		panic("Index out of bounds")
	}
	tmp := make([]interface{}, da.cap-1)
	for i, j := 0, 0; i < da.len; i++ {
		if i == index && i == 0 {
			continue
		}
		if i == index {
			j--
		}
		tmp[j] = da.data[i]
		j++
	}
	da.data = tmp
	da.len -= 1
}

// Remove removes the value from array
func (da *DynamicArray) Remove(value int) {
	tmp := make([]interface{}, da.len)
	for i, j := 0, 0; i < da.len; i++ {
		if da.data[i] == value {
			j--
		}
		tmp[j] = da.data[i]
		j++
	}
	da.data = tmp
	da.len -= 1
}

// Index of value in array.  Returns -1 if not found.
func (da *DynamicArray) IndexOf(value interface{}) int {
	for i, v := range da.data {
		if v == value {
			return i
		}
	}
	return -1
}

// Contains returns true/false if value is in array
func (da *DynamicArray) Contains(value interface{}) bool {
	if da.IndexOf(value) != -1 {
		return true
	}
	return false
}

// Scale increases the array's capacity
func (da *DynamicArray) Scale() {
	var value int
	// Single digit capacity: double size (100%)
	// Double digit capacity: increase by 50%
	// Triple digit capacity: increase by 25%
	// Default increase by 10%
	switch countDigits(da.cap) {
	case 1:
		value = da.cap * 2
	case 2:
		value = int(float32(da.cap) * 1.5)
	case 3:
		value = int(float32(da.cap) * 1.25)
	default:
		value = int(float32(da.cap) * 1.1)
	}
	tmp := make([]interface{}, value)
	for i, v := range da.data {
		tmp[i] = v
	}
	da.data = tmp
	da.cap = value
}

// String representation of array
func (da *DynamicArray) String() string {
	output := "[ "
	for i, v := range da.data {
		output += fmt.Sprint(v)
		if i != da.len-1 {
			output += ","
		}
		output += " "
	}
	output += "]"

	return output
}

func countDigits(x int) int {
	count := 0
	for x != 0 {
		x /= 10
		count += 1
	}
	return count
}
