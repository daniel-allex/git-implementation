package main

import (
	"strconv"
	"strings"
)

func gitBlobContentFromBlob(blob string) string {
	gitObj, rem, found := strings.Cut(blob, " ")

	if !found || gitObj != "blob" {
		ThrowException("Git Object is not a blob")
	}

	_, content, found := strings.Cut(rem, "\x00")

	if !found {
		ThrowException("Could not parse Git Blob")
	}

	return content
}

func gitBlobContentFromZlib(compressed string) string {
	blob := ZlibDecompress(compressed)
	return gitBlobContentFromBlob(blob)
}

func gitBlobContentFromFile(path string) string {
	return gitBlobContentFromZlib(ReadFile(path))
}

func gitBlobContentFromSha1(sha1 string) string {
	filePath := gitBlobPathFromSha1(sha1)
	return gitBlobContentFromFile(filePath)
}

func gitBlobPathFromSha1(sha1 string) string {
	dir := sha1[:2]
	fileName := sha1[2:]
	return ".git/objects/" + dir + "/" + fileName
}

func gitBlobFromContent(content string) string {
	return "blob " + strconv.Itoa(len(content)) + "\000" + content
}

func Sha1FromFile(path string) string {
	return Sha1Hash(gitBlobFromContent(ReadFile(path)))
}

func WriteGitBlobFromFile(path string) {
	blob := gitBlobFromContent(ReadFile(path))
	compressedBlob := ZlibCompress(blob)
	writePath := gitBlobPathFromSha1(Sha1Hash(blob))

	WriteFile(compressedBlob, writePath)
}
