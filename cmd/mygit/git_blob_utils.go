package main

import "fmt"

func createGitBlob(content string) string {
	return createGitObject("blob", content)
}

func gitBlobSha1FromFile(path string) (string, error) {
	file, err := ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to create git blob sha1 from file: %w", err)
	}

	return Sha1Hash(createGitBlob(file)), nil
}

func WriteGitBlobFromFile(path string) error {
	errorMessage := "failed to write git blob sha1 from file"
	file, err := ReadFile(path)
	if err != nil {
		return fmt.Errorf("%v: %w", errorMessage, err)
	}

	blob := createGitBlob(file)
	_, err = writeGitObject(blob)
	if err != nil {
		return fmt.Errorf("%v: %w", errorMessage, err)
	}

	return nil
}
