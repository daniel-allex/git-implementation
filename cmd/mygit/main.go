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
	WriteFile("ref: refs/heads/main\n", ".git/HEAD")
	Log("Initialized git directory")
}

func gitPrintContent(args FlagArgs) {
	fmt.Print(contentFromSha1(args.Arg))
}

func gitPrintTree(args FlagArgs) {
	content := ""
	if args.Flags.Contains("name-only") {
		content = outputTreeNamesFromSha1(args.Arg)
	} else {
		content = outputTreeInfoFromSha1(args.Arg)
	}

	fmt.Println(content)
}

func hashObject(args FlagArgs) {
	if args.Flags.Contains("w") {
		WriteGitBlobFromFile(args.Arg)
	}

	fmt.Println(gitBlobSha1FromFile(args.Arg))
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
	case "ls-tree":
		gitPrintTree(args)
	case "cat-file":
		gitPrintContent(args)
	default:
		ThrowException("Unknown command " + args.Command + "\n")
	}
}
