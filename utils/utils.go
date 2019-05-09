package utils

import (
	"encoding/json"
)

// ResultToStruct ...
func ResultToStruct(input interface{}, output interface{}) (err error) {
	// err = mapstructure.Decode(input, output)
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, output)
	return err
}
