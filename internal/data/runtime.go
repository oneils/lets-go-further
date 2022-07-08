package data

import (
	"fmt"
	"strconv"
)

// Runtime represents the runtime of a movie in minutes.
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("%d mins", r)

	// must wrap the string in quotes to make it a valid JSON string
	quotedJSONValue := strconv.Quote(json)
	return []byte(quotedJSONValue), nil
}
