package main

import (
	"os"
	"strings"
)

type FlagArgs struct {
	Command  string
	ArgPairs *map[string][]string
}

func (args *FlagArgs) addFlag(key string) {
	_, ok := (*args.ArgPairs)[key]
	if !ok {
		(*args.ArgPairs)[key] = []string{}
	}
}

func (args *FlagArgs) addArg(key string, val string) {
	args.addFlag(key)
	(*args.ArgPairs)[key] = append((*args.ArgPairs)[key], val)
}

func (args *FlagArgs) getArgs(key string) ([]string, bool) {
	val, ok := (*args.ArgPairs)[key]

	return val, ok
}

func (args *FlagArgs) getFirstArg(key string) (string, bool) {
	val, ok := (*args.ArgPairs)[key]

	if ok && len(val) > 0 {
		return val[0], true
	}

	return "", false
}

func (args *FlagArgs) flagExists(key string) bool {
	_, ok := args.getArgs(key)

	return ok
}

func isFlag(key string) bool {
	return key[0] == '-'
}

func extractFlag(key string) string {
	return strings.TrimLeft(key, "-")
}

func ParseArgs() *FlagArgs {
	res := FlagArgs{Command: "", ArgPairs: &map[string][]string{}}
	res.Command = os.Args[1]
	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]
		if isFlag(arg) {
			flag := extractFlag(arg)
			if i+1 < len(os.Args) {
				res.addArg(flag, os.Args[i+1])
				i += 1
			} else {
				res.addFlag(flag)
			}
		} else {
			res.addArg("", arg)
		}
	}

	return &res
}

func (args *FlagArgs) getAny() (string, bool) {
	val, ok := args.getFirstArg("")

	if ok {
		return val, true
	}

	for key, _ := range *args.ArgPairs {
		val, ok := args.getFirstArg(key)

		if ok {
			return val, ok
		}
	}

	return "", false
}
