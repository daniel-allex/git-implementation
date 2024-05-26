package main

import "strings"

func createGitCommit(content string) string {
	return createGitObject("commit", content)
}

func WriteCommitFromContent(content string) (string, error) {
	commit := createGitCommit(content)
	return writeGitObject(commit)
}

func commitContent(sha1 string, parents []string, message string) string {
	var sb strings.Builder
	sb.WriteString("tree ")
	sb.WriteString(sha1)
	sb.WriteString("\n")

	for _, parent := range parents {
		sb.WriteString("parent ")
		sb.WriteString(parent)
		sb.WriteString("\n")
	}

	sb.WriteString("author root <root@LAPTOP-G4RJJCES> 1716153761 -0400\n")
	sb.WriteString("committer root <root@LAPTOP-G4RJJCES> 1716153761 -0400\n")
	sb.WriteString("\n")
	sb.WriteString(message)
	sb.WriteString("\n")

	return sb.String()
}

func WriteTreeCommit(sha1 string, parents []string, message string) (string, error) {
	content := commitContent(sha1, parents, message)

	return WriteCommitFromContent(content)
}
