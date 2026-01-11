package handlers

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/cfg"
	"mini-redis/server/handlers/key"
	"mini-redis/server/handlers/list"
	"mini-redis/server/handlers/server"
	str "mini-redis/server/handlers/string"
	"mini-redis/types"
	"mini-redis/types/commands"
)

func TODO(items []resp.RESPItem) ([]byte, error) {
	return nil, fmt.Errorf("UNIMPLEMENTED!")
}

func HandleCommand(conn types.Connection, cmd commands.Command, args []resp.RESPItem) ([]byte, error) {
	if !cmd.Valid() {
		return nil, fmt.Errorf("invalid command passed to handle command")
	}

	fmt.Printf("user: %s", conn.User.Username)

	if cfg.Log.Command {
		fmt.Printf("Command: %s\nArgs: %v\n", cmd.String(), args)
	}

	return commandHandlers[cmd](conn.User, args)
}

// check "command" enum for order of commands
// must be in order of commands in the enum type, since the map is indexed 0..n
var commandHandlers = [...]func(*auth.User, []resp.RESPItem) ([]byte, error){
	HandleNone,
	server.HandlePing,
	server.HandleEcho,
	str.HandleSet,
	str.HandleGet,
	key.HandleDel,
	key.HandleExists,
	key.HandleExpire,
	key.HandleExpireAt,
	key.HandleExpireTime,
	key.HandleTTL,
	str.HandleIncr,
	str.HandleDecr,
	list.HandleLPush,
	list.HandleRPush,
	list.HandleLPop,
	list.HandleRPop,
	list.HandleLRange,
	list.HandleLGet,
	server.HandleInfo,
	key.HandleKeys,
	key.HandleFlushAll,
	server.HandleAuth,
	server.HandleLogout,
	server.HandleWhoami,
}
