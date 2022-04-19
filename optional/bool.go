// Code generated by go generate
// This file was generated by robots at 2021-05-04 14:21:52.312823 +0000 UTC

package optional

import (
	"encoding/json"
	"errors"
)

// Bool is an optional bool.
type Bool struct {
	value *bool
}

// NewBool creates an optional.Bool from a bool.
func NewBool(v bool) Bool {
	return Bool{&v}
}

// Set sets the bool value.
func (b *Bool) Set(v bool) {
	b.value = &v
}

// Get returns the bool value or an error if not present.
func (b Bool) Get() (bool, error) {
	if !b.Present() {
		var zero bool
		return zero, errors.New("value not present")
	}
	return *b.value, nil
}

// MustGet returns the bool value or panics if not present.
func (b Bool) MustGet() bool {
	if !b.Present() {
		panic("value not present")
	}
	return *b.value
}

// Present returns whether or not the value is present.
func (b Bool) Present() bool {
	return b.value != nil
}

// OrElse returns the bool value or a default value if the value is not present.
func (b Bool) OrElse(v bool) bool {
	if b.Present() {
		return *b.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func (b Bool) If(fn func(bool)) {
	if b.Present() {
		fn(*b.value)
	}
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b.Present() {
		return json.Marshal(b.value)
	}
	return json.Marshal(nil)
}

func (b *Bool) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		b.value = nil
		return nil
	}

	var value bool

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	b.value = &value
	return nil
}