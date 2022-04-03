// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import (
	"encoding/json"
)

//func assertSerializationImplementation() {
//	var _ containers.JSONSerializer = (*List)(nil)
//	var _ containers.JSONDeserializer = (*List)(nil)
//}

// ToJSON outputs the JSON representation of list's Elements.
func (list *List[T]) ToJSON() ([]byte, error) {
	return json.Marshal(list)
}

// FromJSON populates list's Elements from the input JSON representation.
func (list *List[T]) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list)

	return err
}
