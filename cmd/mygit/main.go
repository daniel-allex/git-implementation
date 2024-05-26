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
	err := WriteFile("ref: refs/heads/main\n", ".git/HEAD")
	if err != nil {
		fmt.Println("failed to initialize .git: %w", err)
	}
	Log("Initialized git directory")
}

func gitPrintContent(args *FlagArgs) {
	val, ok := args.getAny()

	if !ok {
		fmt.Println("no argument when printing blob content")
		return
	}

	content, err := contentFromSha1(val)
	if err != nil {
		fmt.Println("failed to print blob content: %w", err)
		return
	}

	fmt.Print(content)
}

func gitTreeContent(args *FlagArgs) (string, error) {
	arg, ok := args.getAny()

	if !ok {
		fmt.Println("no argument when printing tree content")
	}

	if args.flagExists("name-only") {
		return outputTreeNamesFromSha1(arg)
	} else {
		return outputTreeInfoFromSha1(arg)
	}
}

func gitPrintTree(args *FlagArgs) {
	content, err := gitTreeContent(args)
	if err != nil {
		fmt.Println("failed to print tree: %w", err)
		return
	}
	fmt.Println(content)
}

func gitWriteTree(args *FlagArgs) {
	sha1, err := WriteTree(".")
	if err != nil {
		fmt.Println("failed to write tree: %w", err)
		return
	}
	fmt.Println(sha1)
}

func hashObject(args *FlagArgs) {
	key := "w"
	arg, ok := args.getAny()
	if !ok {
		fmt.Println("no argument when hashing object")
	}

	if args.flagExists(key) {
		err := WriteGitBlobFromFile(arg)
		if err != nil {
			fmt.Println("failed to hash object: %w", err)
			return
		}
	}

	sha1, err := gitBlobSha1FromFile(arg)
	if err != nil {
		fmt.Println("failed to hash object: %w", err)
		return
	}

	fmt.Println(sha1)
}

func commitTree(args *FlagArgs) {
	arg, ok := args.getAny()
	if !ok {
		fmt.Println("no argument when committing tree")
	}

	parents, _ := args.getArgs("p")
	message, _ := args.getFirstArg("m")

	sha1, err := WriteTreeCommit(arg, parents, message)
	if err != nil {
		fmt.Println("failed to commit tree")
	}

	fmt.Println(sha1)
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
	case "write-tree":
		gitWriteTree(args)
	case "commit-tree":
		commitTree(args)
	case "cat-file":
		gitPrintContent(args)
	default:
		fmt.Println("Unknown command " + args.Command + "\n")
	}
}
