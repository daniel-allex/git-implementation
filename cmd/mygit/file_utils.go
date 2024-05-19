package main

import (
	"os"
	"path/filepath"
)

func ReadFile(path string) (string, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil
}

func WriteFile(content string, path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
