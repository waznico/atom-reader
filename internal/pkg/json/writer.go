package json

import (
	"io/ioutil"
)

// WriteJSON writes json content as byte array into a json file
func WriteJSON(filePath string, content []byte) error {
	err := ioutil.WriteFile(filePath, content, 0644)
	return err
}
