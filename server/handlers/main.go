package handlers

import (
	"fmt"
	"mini-redis/types"
)

func TODO(items []types.RESPItem) (string, error) {
	fmt.Printf("Command handler not yet implemented!")
	return "NOT IMPLEMENTED!!!!", nil
}

func HandleCommand(cmd types.Command, args []types.RESPItem) (string, error) {
	if !cmd.Valid() {
		return "", fmt.Errorf("invalid command passed to handle command")
	}

	fmt.Printf("command: %s\nargs: %v\n", cmd.String(), args)

	return commandHandlers[cmd](args)
}

var commandHandlers = [...]func([]types.RESPItem) (string, error){
	handleNone,
	handlePing,
	handleEcho,
	handleSet,
	handleGet,
	handleDel,
	handleExists,
	handleExpire,
	handleTTL,
	handleIncr,
	handleDecr,
	handleLPush,
	TODO,
	TODO,
	TODO,
	TODO,
	handleFlushAll,
}

/*
	NONE Command = iota
	PING
	ECHO
	SET
	GET
	DEL
	EXISTS
	EXPIRE
	TTL
	INCR
	DECR
	LPUSH
	RPUSH
	LPOP
	RPOP
	INFO
*/
