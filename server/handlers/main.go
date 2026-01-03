package handlers

import (
	"fmt"
	"mini-redis/server/handlers/key"
	"mini-redis/server/handlers/list"
	"mini-redis/server/handlers/server"
	str "mini-redis/server/handlers/string"
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
	HandleNone,
	server.HandlePing,
	server.HandleEcho,
	str.HandleSet,
	str.HandleGet,
	key.HandleDel,
	key.HandleExists,
	key.HandleExpire,
	key.HandleTTL,
	str.HandleIncr,
	str.HandleDecr,
	list.HandleLPush,
	list.HandleRPush,
	TODO,
	TODO,
	TODO,
	key.HandleFlushAll,
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
