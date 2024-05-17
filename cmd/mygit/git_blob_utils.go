package main

func createGitBlob(content string) string {
	return createGitObject("blob", content)
}

func gitBlobSha1FromFile(path string) string {
	return Sha1Hash(createGitBlob(ReadFile(path)))
}

func WriteGitBlobFromFile(path string) {
	blob := createGitBlob(ReadFile(path))
	compressedBlob := ZlibCompress(blob)
	writePath := pathFromSha1(Sha1Hash(blob))

	WriteFile(compressedBlob, writePath)
}
