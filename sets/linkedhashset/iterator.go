// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset

import (
	"github.com/qauzy/chocolate/containers"
	"github.com/qauzy/chocolate/lists/doublylinkedlist"
)

func assertIteratorImplementation() {
	var _ containers.ReverseIteratorWithIndex[int] = (*Iterator[int, int])(nil)
}

// Iterator holding the iterator's state
type Iterator[TP comparable] struct {
	iterator doublylinkedlist.Iterator[TP]
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (set *Set[T]) Iterator() Iterator[T] {
	return Iterator[T]{iterator: set.ordering.Iterator()}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[TP]) Next() bool {
	return iterator.iterator.Next()
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[TP]) Prev() bool {
	return iterator.iterator.Prev()
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[TP]) Value() TP {
	return iterator.iterator.Value()
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator[TP]) Index() int {
	return iterator.iterator.Index()
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator[TP]) Begin() {
	iterator.iterator.Begin()
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator[TP]) End() {
	iterator.iterator.End()
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[TP]) First() bool {
	return iterator.iterator.First()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[TP]) Last() bool {
	return iterator.iterator.Last()
}
