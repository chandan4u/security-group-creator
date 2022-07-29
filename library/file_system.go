package library

import (
	"io/ioutil"
)

// ReadFromFile it read file and return data in bytes
func ReadFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}
