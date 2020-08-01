package json

import (
	"io/ioutil"
	"os"
)

// ReadJSON reads a json file from disk and returns plain content as byte array
func ReadJSON(filePath string) ([]byte, error) {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, err
}
