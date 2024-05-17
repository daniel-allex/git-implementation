package main

import (
	"fmt"
	"os"
)

var LOGGING = false

func ThrowException(message string) {
	fmt.Fprintf(os.Stderr, "[ERROR] "+message)
	os.Exit(1)
}

func ExceptIfNotOk(message string, ok bool) {
	if !ok {
		ThrowException(message)
	}
}

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

func ExceptIfError(message string, err error) {
	if err != nil {
		ThrowException(message + ": " + err.Error())
	}
}
