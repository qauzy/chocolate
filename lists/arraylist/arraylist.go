// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraylist implements the array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package arraylist

import (
	"fmt"
	"github.com/qauzy/chocolate/lists"
	"github.com/qauzy/chocolate/stream"
	"github.com/qauzy/chocolate/utils"
	"strings"
)

func assertListImplementation() {
	var _ lists.List[int] = (*List[int])(nil)
}

// List holds the Elements in a slice
type List[T comparable] []T

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

//New instantiates a new list and adds the passed values, if any, to the list
func New[TP comparable](values ...TP) List[TP] {
	list := List[TP]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

//
//func (t *List[T]) UnmarshalJSON(data []byte) (err error) {
//	err = json.Unmarshal(data, &t.Elements)
//	if err != nil {
//		return
//	}
//	t.size = len(t.Elements)
//	return
//}
//func (t List[T]) MarshalJSON() ([]byte, error) {
//	return json.Marshal(t.Elements)
//}

//func (t List[T]) Value() (driver.Value, error) {
//	return t.Elements, nil
//}
//
//func (t *List[T]) Scan(v interface{}) error {
//	switch vt := v.(type) {
//	case []T:
//		t = New(vt...)
//	case []byte:
//		return t.UnmarshalJSON(vt)
//	case string:
//		return t.UnmarshalJSON([]byte(vt))
//	default:
//		return errors.New(fmt.Sprintf("类型处理错误,%v", v))
//	}
//	return nil
//}

// Add appends a value at the end of the list
func (list *List[T]) Add(values ...T) {
	for _, value := range values {
		*list = append(*list, value)

	}
}

func (list List[T]) ForEach(fn func(each T)) {
	for _, x := range []T(list) {
		fn(x)
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list List[T]) Get(index int) (T, bool) {

	if !list.withinRange(index) {
		var t T
		return t, false
	}
	return list[index], true
}

// Remove removes the element at the given index from the list.
func (list *List[T]) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	//list.Elements[index] =                                    // cleanup reference
	copy((*list)[index:], (*list)[index+1:]) // shift to the left by one (slow operation, need ways to optimize this)

}

// Contains checks if Elements (one or more) are present in the set.
// All Elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list List[T]) Contains(values ...T) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range []T(list) {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values returns all Elements in the list.
func (list List[T]) Values() []T {

	return list
}

//IndexOf returns index of provided element
func (list List[T]) IndexOf(value T) int {
	if len([]T(list)) == 0 {
		return -1
	}
	for index, element := range list {
		if element == value {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any Elements.
func (list List[T]) Empty() bool {
	return len([]T(list)) == 0
}

// Size returns number of Elements within the list.
func (list List[T]) Size() int {
	return len([]T(list))
}

func (list List[T]) Stream() *stream.Stream[T] {
	return stream.Of[T](list...)
}

// Clear removes all Elements from the list.
func (list *List[T]) Clear() {
	*list = make(List[T], 0)
}

// Sort sorts values (in-place) using.
func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if len(*list) < 2 {
		return
	}
	utils.Sort(*list, comparator)
}

// Swap swaps the two values at the specified positions.
func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent Elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Insert(index int, values ...T) {

	if !list.withinRange(index) {
		// Append
		if index == len(*list) {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	copy((*list)[index+l:], (*list)[index:len(*list)-l])
	copy((*list)[index:], values)
}

// Set the value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list List[T]) Set(index int, value T) {

	if !list.withinRange(index) {
		// Append
		for index >= len([]T(list)) {
			list.Add(value)
		}

		return
	}

	list[index] = value
}

// String returns a string representation of container
func (list List[T]) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list List[T]) withinRange(index int) bool {
	return index >= 0 && index < len([]T(list))
}
