package main

import (
	"fmt"
	"os"
)

var LOGGING = false

func Warn(message string) {
	fmt.Fprintf(os.Stderr, "[WARN] "+message)
}

func Log(message string) {
	if LOGGING {
		fmt.Println("[LOG] " + message)
	}
}

func WarnIfError(message string, err error) {
	if err != nil {
		Warn(message + ": " + err.Error())
	}
}
