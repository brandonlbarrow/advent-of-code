package common

import (
	"io"
	"os"
)

func GetInput(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
