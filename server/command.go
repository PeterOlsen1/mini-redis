package main

import "strings"

var commandStrings = [...]string{
	"NONE",
	"PING",
	"ECHO",
	"SET",
	"GET",
	"DEL",
	"EXISTS",
	"EXPIRE",
	"TTL",
	"INCR",
	"LPUSH",
	"RPUSH",
	"LPOP",
	"RPOP",
	"INFO",
}

func (c Command) String() string {
	if int(c) < 0 || int(c) >= len(commandStrings) {
		return ""
	}
	return commandStrings[c]
}

func ParseCommand(s string) Command {
	s = strings.ToUpper(s)
	for i, cmd := range commandStrings {
		if cmd == s {
			return Command(i)
		}
	}
	return Command(0)
}

func (c Command) Valid() bool {
	return c != 0
}
