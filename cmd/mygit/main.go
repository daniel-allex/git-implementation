package main

import (
	"fmt"
	"os"
	// Uncomment this block to pass the first stage!
	// "os"
)

func gitInit() {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		err := os.MkdirAll(dir, 0755)
		warnIfError("Could not create directory", err)
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	err := os.WriteFile(".git/HEAD", headFileContents, 0644)
	warnIfError("Could not write file", err)

	log("Initialized git directory")
}

func gitPrintContent(shah string) {
	fmt.Print(contentFromGitBlobShah(shah))
}

func getArg(i int) string {
	if i >= len(os.Args) {
		throwException("usage: mygit <command> [<args>...]\n")
	}

	return os.Args[i]
}

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	log("Logs from your program will appear here!")

	command := getArg(1)
	switch command {
	case "init":
		gitInit()
	case "cat-file":
		gitPrintContent(getArg(3))
	default:
		throwException("Unknown command " + command + "\n")
	}
}
