package json_utility

import (
	"encoding/json"
)

// Utility function to lower case JSON pure structs, generally used on response of a request inside controllers.
func StructToLowerCaseJson(data interface{}) (interface{}, error) {

	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var r interface{}

	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
