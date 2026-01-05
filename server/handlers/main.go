package handlers

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/handlers/key"
	"mini-redis/server/handlers/list"
	"mini-redis/server/handlers/server"
	str "mini-redis/server/handlers/string"
	"mini-redis/types/commands"
)

func TODO(items []resp.RESPItem) ([]byte, error) {
	fmt.Printf("Command handler not yet implemented!")
	return []byte("NOT IMPLEMENTED!!!!"), nil
}

func HandleCommand(cmd commands.Command, args []resp.RESPItem) ([]byte, error) {
	if !cmd.Valid() {
		return nil, fmt.Errorf("invalid command passed to handle command")
	}

	fmt.Printf("command: %s\nargs: %v\n", cmd.String(), args)

	return commandHandlers[cmd](args)
}

// check "command" enum for order of commands
var commandHandlers = [...]func([]resp.RESPItem) ([]byte, error){
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
	list.HandleLPop,
	list.HandleRPop,
	list.HandleLRange,
	list.HandleLGet,
	TODO,
	key.HandleFlushAll,
}
