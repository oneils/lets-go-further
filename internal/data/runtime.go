package data

import (
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntime = fmt.Errorf("invalid runtime format")

// Runtime represents the runtime of a movie in minutes.
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	json := fmt.Sprintf("%d mins", r)

	// must wrap the string in quotes to make it a valid JSON string
	quotedJSONValue := strconv.Quote(json)
	return []byte(quotedJSONValue), nil
}

// UnmarshalJSON unmarshals the JSON value into the Runtime type.
// We need pointer here in order to modify the receiver
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	// remove the quotes from the JSON string
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntime
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntime
	}

	//  parse the string containing the number into an int32
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntime
	}

	*r = Runtime(i)

	return nil
}
