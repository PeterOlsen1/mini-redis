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
	FLUSHDB
	AUTH
	LOGOUT
	WHOAMI
	ADDUSER
	RMUSER
	UGET
	ADDRULE
	RMRULE
	SAVE
	LOAD
	LISTSAVES
	RMSAVE
	SELECT
	WHICHDB
	COMMAND
)

const NUM_COMMANDS = len(commandStrings)

var requireDB = map[Command]struct{}{
	SET:        {},
	GET:        {},
	DEL:        {},
	EXISTS:     {},
	EXPIRE:     {},
	EXPIREAT:   {},
	EXPIRETIME: {},
	TTL:        {},
	INCR:       {},
	DECR:       {},
	LPUSH:      {},
	RPUSH:      {},
	LPOP:       {},
	RPOP:       {},
	LRANGE:     {},
	LGET:       {},
	KEYS:       {},
	FLUSHALL:   {},
}

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
	"FLUSHDB",
	"AUTH",
	"LOGOUT",
	"WHOAMI",
	"ADDUSER",
	"RMUSER",
	"UGET",
	"ADDRULE",
	"RMRULE",
	"SAVE",
	"LOAD",
	"LISTSAVES",
	"RMSAVE",
	"SELECT",
	"WHICHDB",
	"COMMAND",
}

var commandInfos = map[Command]string{
	NONE:       "No operation.",
	PING:       "Ping the server to check if it is alive.",
	ECHO:       "Echo back the provided message.",
	SET:        "Set a key to a value.",
	GET:        "Get the value of a key.",
	DEL:        "Delete one or more keys.",
	EXISTS:     "Check if a key exists.",
	EXPIRE:     "Set a timeout on a key.",
	EXPIREAT:   "Set a timeout on a key, specified in Unix time.",
	EXPIRETIME: "Get the expiration time of a key.",
	TTL:        "Get the time-to-live of a key.",
	INCR:       "Increment the value of a key.",
	DECR:       "Decrement the value of a key.",
	LPUSH:      "Push a value to the left of a list.",
	RPUSH:      "Push a value to the right of a list.",
	LPOP:       "Pop a value from the left of a list.",
	RPOP:       "Pop a value from the right of a list.",
	LRANGE:     "Get a range of elements from a list.",
	LGET:       "Get all elements of a list.",
	INFO:       "Get server information.",
	KEYS:       "Get all keys matching a pattern.",
	FLUSHALL:   "Delete all keys in all databases.",
	FLUSHDB:    "Delete all keys in the current database.",
	AUTH:       "Authenticate a user.",
	LOGOUT:     "Log out the current user.",
	WHOAMI:     "Get the username of the current user.",
	ADDUSER:    "Add a new user.",
	RMUSER:     "Remove a user.",
	UGET:       "Get information about a user.",
	ADDRULE:    "Add a rule to a user.",
	RMRULE:     "Remove a rule from a user.",
	SAVE:       "Save the current instance state to disk.",
	LOAD:       "Load the instance state from disk.",
	LISTSAVES:  "List all saved instance states.",
	RMSAVE:     "Remove a saved instance state.",
	SELECT:     "Select a database by its number.",
	WHICHDB:    "Get the current database number.",
	COMMAND:    "Returns a list of all commands and a short summary.",
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

func (c Command) RequiresDB() bool {
	_, exists := requireDB[c]
	return exists
}

func (c Command) Info() string {
	return commandInfos[c]
}

func Len() int {
	return len(commandStrings)
}
