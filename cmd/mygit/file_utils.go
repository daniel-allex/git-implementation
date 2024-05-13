package main

import (
	"os"
	"path/filepath"
)

func ReadFile(path string) string {
	fileContent, err := os.ReadFile(path)
	ExceptIfError("Failed to read file at path "+path, err)

	return string(fileContent)
}

func WriteFile(content string, path string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	ExceptIfError("Failed to MkdirAll for path "+path, err)

	err = os.WriteFile(path, []byte(content), 0666)
	ExceptIfError("Failed to write file at path "+path, err)
}
