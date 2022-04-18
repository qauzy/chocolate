// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import "time"

// Comparator will make type assertion (see IntComparator for example),
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//    negative , if a < b
//    zero     , if a == b
//    positive , if a > b
type Comparator[T comparable] func(a, b T) int

// StringComparator provides a fast comparison on strings
func StringComparator(a, b string) int {
	s1 := a
	s2 := b
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

// IntComparator provides a basic comparison on int
func IntComparator(a, b int) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int8Comparator provides a basic comparison on int8
func Int8Comparator(a, b int8) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int16Comparator provides a basic comparison on int16
func Int16Comparator(a, b int16) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int32Comparator provides a basic comparison on int32
func Int32Comparator(a, b int32) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int64Comparator provides a basic comparison on int64
func Int64Comparator(a, b int64) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UIntComparator provides a basic comparison on uint
func UIntComparator(a, b uint) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UInt8Comparator provides a basic comparison on uint8
func UInt8Comparator(a, b uint8) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UInt16Comparator provides a basic comparison on uint16
func UInt16Comparator(a, b uint16) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UInt32Comparator provides a basic comparison on uint32
func UInt32Comparator(a, b uint32) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UInt64Comparator provides a basic comparison on uint64
func UInt64Comparator(a, b uint64) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Float32Comparator provides a basic comparison on float32
func Float32Comparator(a, b float32) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Float64Comparator provides a basic comparison on float64
func Float64Comparator(a, b float64) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// ByteComparator provides a basic comparison on byte
func ByteComparator(a, b byte) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// RuneComparator provides a basic comparison on rune
func RuneComparator(a, b rune) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// TimeComparator provides a basic comparison on time.Time
func TimeComparator(a, b time.Time) int {
	aAsserted := a
	bAsserted := b

	switch {
	case aAsserted.After(bAsserted):
		return 1
	case aAsserted.Before(bAsserted):
		return -1
	default:
		return 0
	}
}
