package main

import "os"

type FlagArgs struct {
	Command string
	Flags   *HashSet
	Arg     string
}

func ParseArgs() FlagArgs {
	res := FlagArgs{Command: "", Flags: NewHashSet(), Arg: ""}
	for i, arg := range os.Args {
		if i == 0 {
			continue
		} else if i == 1 {
			res.Command = arg
		} else if len(arg) > 1 && arg[:2] == "--" {
			res.Flags.Add(arg[2:])
		} else if arg[0] == '-' {
			res.Flags.Add(arg[1:])
		} else {
			res.Arg = arg
		}
	}

	return res
}
