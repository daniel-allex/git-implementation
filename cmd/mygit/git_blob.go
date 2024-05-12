package main

import (
	"strings"
)

func contentFromGitBlob(blob string) string {
	gitObj, rem, found := strings.Cut(blob, " ")

	if !found || gitObj != "blob" {
		throwException("Git Object is not a blob")
	}

	_, content, found := strings.Cut(rem, "\x00")

	if !found {
		throwException("Could not parse Git Blob")
	}

	return content
}

func contentFromGitBlobCompressed(compressed string) string {
	blob := decompress(compressed)
	return contentFromGitBlob(blob)
}

func contentFromGitBlobPath(path string) string {
	return contentFromGitBlobCompressed(readFile(path))
}

func contentFromGitBlobShah(shah string) string {
	dir := shah[:2]
	fileName := shah[2:]
	filePath := ".git/objects/" + dir + "/" + fileName

	return contentFromGitBlobPath(filePath)
}
