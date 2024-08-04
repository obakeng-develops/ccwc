package pkg

import (
	"os"
)

func ValidateFilePath(path string) (bool, error) {
	file, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if file.Mode().IsRegular() {
		return true, err
	}

	return false, err
}
