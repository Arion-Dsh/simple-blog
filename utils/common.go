package utils

import (
	"os"
	"path/filepath"
)

// GetExPath get file path in current folder
func GetExPath() string {

	ex, err := os.Executable()
	if err != nil {
		panic(err)

	}
	exPath := filepath.Dir(ex)

	return exPath
}
