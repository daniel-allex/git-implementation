package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func isGitObject(name string) bool {
	return name == "blob" || name == "tree"
}

func contentFromGitObject(objectStorage string) (string, error) {
	gitObj, rem, found := strings.Cut(objectStorage, " ")

	if !found || !isGitObject(gitObj) {
		return "", errors.New("not a git object")
	}

	_, content, found := strings.Cut(rem, "\x00")

	if !found {
		return "", errors.New("could not parse git object")
	}

	return content, nil
}

func contentFromZlib(compressed string) (string, error) {
	errorMessage := "failed to get content with zlib"
	blob, err := ZlibDecompress(compressed)
	if err != nil {
		return "", fmt.Errorf("%v: %w", errorMessage, err)
	}

	content, err := contentFromGitObject(blob)
	if err != nil {
		return "", fmt.Errorf("%v: %w", errorMessage, err)
	}

	return content, nil
}

func contentFromFile(path string) (string, error) {
	file, err := ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to get content from file: %w", err)
	}

	content, err := contentFromZlib(file)
	if err != nil {
		return "", fmt.Errorf("failed to get content from file: %w", err)
	}

	return content, nil
}

func contentFromSha1(sha1 string) (string, error) {
	filePath := pathFromSha1(sha1)
	content, err := contentFromFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get content from sha1: %w", err)
	}

	return content, nil
}

func pathFromSha1(sha1 string) string {
	dir := sha1[:2]
	fileName := sha1[2:]
	return ".git/objects/" + dir + "/" + fileName
}

func createGitObject(objectType string, content string) string {
	return objectType + " " + strconv.Itoa(len(content)) + "\000" + content
}

func writeGitObject(gitObject string) (string, error) {
	compressed, err := ZlibCompress(gitObject)
	if err != nil {
		return "", fmt.Errorf("failed to write git object: %w", err)
	}

	sha1 := Sha1Hash(gitObject)
	writePath := pathFromSha1(sha1)

	err = WriteFile(compressed, writePath)
	if err != nil {
		return "", fmt.Errorf("failed to write git object: %w", err)
	}

	return sha1, nil
}
