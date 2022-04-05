// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doublylinkedlist

import (
	"encoding/json"
	"github.com/qauzy/chocolate/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*List[int])(nil)
	var _ containers.JSONDeserializer = (*List[int])(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (list *List[T]) FromJSON(data []byte) error {
	elements := []T{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}
