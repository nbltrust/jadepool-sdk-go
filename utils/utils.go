package utils

import "github.com/mitchellh/mapstructure"

// ResultToStruct ...
func ResultToStruct(input interface{}, output interface{}) (err error) {
	err = mapstructure.Decode(input, output)
	return err
}
