package main

import (
	"strconv"
	"strings"
)

func isGitObject(name string) bool {
	return name == "blob" || name == "tree"
}

func contentFromGitObject(objectStorage string) string {
	gitObj, rem, found := strings.Cut(objectStorage, " ")

	if !found || !isGitObject(gitObj) {
		ThrowException("Git Object is not a blob")
	}

	_, content, found := strings.Cut(rem, "\x00")

	if !found {
		ThrowException("Could not parse Git Blob")
	}

	return content
}

func contentFromZlib(compressed string) string {
	blob := ZlibDecompress(compressed)
	return contentFromGitObject(blob)
}

func contentFromFile(path string) string {
	return contentFromZlib(ReadFile(path))
}

func contentFromSha1(sha1 string) string {
	filePath := pathFromSha1(sha1)
	return contentFromFile(filePath)
}

func pathFromSha1(sha1 string) string {
	dir := sha1[:2]
	fileName := sha1[2:]
	return ".git/objects/" + dir + "/" + fileName
}

func createGitObject(objectType string, content string) string {
	return objectType + " " + strconv.Itoa(len(content)) + "\000" + content
}
