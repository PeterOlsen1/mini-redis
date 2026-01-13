package commands

import (
	"strings"
)

// SEE HERE FOR COMMAND IMPLEMENTATIONS
// https://redis.io/docs/latest/commands/

// To register a new command:
// * Add it to the enum
// * Add it to commandStrings
// * Add handler method
// * Register handler method in server/handlers/main
// * Create client method

type Command int

const (
	NONE Command = iota
	PING
	ECHO
	SET
	GET
	DEL
	EXISTS
	EXPIRE
	EXPIREAT
	EXPIRETIME
	TTL
	INCR
	DECR
	LPUSH
	RPUSH
	LPOP
	RPOP
	LRANGE
	LGET
	INFO
	KEYS
	FLUSHALL
	AUTH
	LOGOUT
	WHOAMI
	SETUSER
	RMUSER
	UGET
	SETRULE
	RMRULE
	SAVE
	LOAD
)

const NUM_COMMANDS = len(commandStrings)

var commandStrings = [...]string{
	"NONE",
	"PING",
	"ECHO",
	"SET",
	"GET",
	"DEL",
	"EXISTS",
	"EXPIRE",
	"EXPIREAT",
	"EXPIRETIME",
	"TTL",
	"INCR",
	"DECR",
	"LPUSH",
	"RPUSH",
	"LPOP",
	"RPOP",
	"LRANGE",
	"LGET",
	"INFO",
	"KEYS",
	"FLUSHALL",
	"AUTH",
	"LOGOUT",
	"WHOAMI",
	"SETUSER",
	"RMUSER",
	"UGET",
	"SETRULE",
	"RMRULE",
	"SAVE",
	"LOAD",
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
