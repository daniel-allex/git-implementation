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
		WarnIfError("Could not create directory", err)
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	err := os.WriteFile(".git/HEAD", headFileContents, 0644)
	WarnIfError("Could not write file", err)

	Log("Initialized git directory")
}

func gitPrintContent(args FlagArgs) {
	fmt.Print(gitBlobContentFromSha1(args.Arg))
}

func hashObject(args FlagArgs) {
	if args.Flags.Contains("w") {
		WriteGitBlobFromFile(args.Arg)
	}
	
	fmt.Println(Sha1FromFile(args.Arg))
}

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	Log("Logs from your program will appear here!")

	args := ParseArgs()
	switch args.Command {
	case "init":
		gitInit()
	case "hash-object":
		hashObject(args)
	case "cat-file":
		gitPrintContent(args)
	default:
		ThrowException("Unknown command " + args.Command + "\n")
	}
}
