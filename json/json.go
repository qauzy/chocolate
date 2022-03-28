package JSON

import (
	json "github.com/json-iterator/go"
)

func ToJSONString(v interface{}) string {
	res, _ := json.MarshalToString(v)
	return res
}
