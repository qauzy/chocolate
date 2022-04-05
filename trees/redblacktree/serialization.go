// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import (
	"encoding/json"
	"github.com/qauzy/chocolate/containers"
	"github.com/qauzy/chocolate/utils"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Tree[int, int])(nil)
	var _ containers.JSONDeserializer = (*Tree[int, int])(nil)
}

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree[K, V]) ToJSON() ([]byte, error) {
	elements := make(map[K]V)
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree[K, V]) FromJSON(data []byte) error {
	elements := make(map[K]V)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}
