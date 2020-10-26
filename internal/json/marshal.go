package json

import (
	"encoding/json"
)

func MarshalStruct(i interface{}) []byte {
	b, _ := json.MarshalIndent(i, "", "\t")
	return b
}
