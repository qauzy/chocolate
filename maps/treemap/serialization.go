// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import "github.com/qauzy/chocolate/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Map[int, int])(nil)
	var _ containers.JSONDeserializer = (*Map[int, int])(nil)
}

// ToJSON outputs the JSON representation of the map.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	return m.tree.ToJSON()
}

// FromJSON populates the map from the input JSON representation.
func (m *Map[K, V]) FromJSON(data []byte) error {
	return m.tree.FromJSON(data)
}
