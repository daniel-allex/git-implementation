package main

import (
	"fmt"
	"os"
)

var LOGGING = false

func throwException(message string) {
	fmt.Fprintf(os.Stderr, "[ERROR] "+message)
	os.Exit(1)
}

func warn(message string) {
	fmt.Fprintf(os.Stderr, "[WARN] "+message)
}

func log(message string) {
	if LOGGING {
		fmt.Println("[LOG] " + message)
	}
}

func warnIfError(message string, err error) {
	if err != nil {
		warn(message + ": " + err.Error())
	}
}

func exceptIfError(message string, err error) {
	if err != nil {
		throwException(message + ": " + err.Error())
	}
}

func readFile(path string) string {
	fileContent, err := os.ReadFile(path)
	exceptIfError("Failed to read file at path "+path, err)

	return string(fileContent)
}
