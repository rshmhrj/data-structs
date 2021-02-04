package arrays

import (
  "fmt"
)

type intArray struct {
  len int
  cap int
  data []int
}

func MakeIntArray(array ...int) *intArray {
  return &intArray{
    len: len(array),
    cap:  len(array),
    data: array,
  }
}

func NewIntArray(size int) *intArray {
  if size < 0 {
    panic("Illegal size chosen for new array")
    return nil
  }
  return &intArray {
    len: 0,
    cap: size,
    data: make([]int, size),
  }
}

// Len returns the current size of the array
func (ia *intArray) Len() int {
  return ia.len
}

// Cap returns the current capacity of the array
func (ia *intArray) Cap() int {
  return ia.cap
}

// IsEmpty returns whether the array contains any values or not
func (ia *intArray) IsEmpty() bool {
  if ia.len == 0 {
    return true
  }
  return false
}

// Get Value at index
func (ia *intArray) Get(index int) int {
  return ia.data[index]
}

// Set Value at index
func (ia *intArray) Set(index, value int) {
  if ia.data[index] == 0 {
    ia.len += 1
  }
  ia.data[index] = value
}

// Clear Array
func (ia *intArray) Clear() {
  ia.data = nil
  ia.len = 0
}

// Add new value to the array
func (ia *intArray) Add(value int) {
  if ia.len + 1 > ia.cap {
    ia.Scale()
  }
  ia.data[ia.len+1] = value
  ia.len += 1
}

// RemoveAt removes the value at the given index
func (ia *intArray) RemoveAt(index int) {
  if index > ia.len {
    panic("Index out of bounds")
  }
  tmp := make([]int, ia.cap)
   for i,j := 0,0 ; i < ia.len; i++ {
     if i == index {
       j--
     }
     tmp[j] = ia.data[i]
     j++
  }
  ia.data = tmp
  ia.len -= 1
}

// Remove removes the value from array
func (ia *intArray) Remove(value int) {
  tmp := make([]int, ia.len)
  for i, j := 0,0; i < ia.len; i++ {
    if ia.data[i] == value {
      j--
    }
    tmp[j] = ia.data[i]
    j++
  }
  ia.data = tmp
  ia.len -= 1
}

// Index of value in array.  Returns -1 if not found.
func (ia *intArray) IndexOf(value int) int {
  for i, v := range ia.data {
    if v == value {
      return i
    }
  }
  return -1
}

// Contains returns true/false if value is in array
func (ia *intArray) Contains(value int) bool {
 if ia.IndexOf(value) != -1 {
   return true
 }
 return false
}

// Scale increases the array's capacity
func (ia *intArray) Scale() {
  var value int
  // Single digit capacity: double size (100%)
  // Double digit capacity: increase by 50%
  // Triple digit capacity: increase by 25%
  // Default increase by 10%
  switch countDigits(ia.cap) {
  case 1:
    value = ia.cap * 2
  case 2:
    value = int(float32(ia.cap) * 1.5)
  case 3:
    value = int(float32(ia.cap) * 1.25)
  default:
    value = int(float32(ia.cap) * 1.1)
  }
  tmp := make([]int, value)
  for i, v := range ia.data {
    tmp[i] = v
  }
  ia.data = tmp
  ia.cap = value
}

// String representation of array
func (ia *intArray) String() string {
  output := "[ "
  for i, v := range ia.data {
    output += fmt.Sprint(v)
    if i != ia.len - 1 {
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