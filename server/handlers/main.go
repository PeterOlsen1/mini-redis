package handlers

import (
	"fmt"
	"mini-redis/resp"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/cfg"
	"mini-redis/server/handlers/key"
	"mini-redis/server/handlers/list"
	"mini-redis/server/handlers/server"
	str "mini-redis/server/handlers/string"
	"mini-redis/server/internal"
	"mini-redis/types"
	"mini-redis/types/commands"
	"mini-redis/types/errors"
)

func TODO(items resp.ArgList) ([]byte, error) {
	return nil, fmt.Errorf("UNIMPLEMENTED!")
}

func HandleCommand(conn *types.Connection, cmd commands.Command, args resp.ArgList) ([]byte, error) {
	if !cmd.Valid() {
		return nil, fmt.Errorf("invalid command passed to handle command")
	}

	if cfg.Log.Command {
		fmt.Printf("Command: %s\nArgs: %v\n", cmd.String(), args)
	}

	if conn.User.DB == nil && cmd.RequiresDB() {
		conn.User.DB = internal.GetDB(0)
		return nil, errors.BAD_DB
	}

	handler, exists := mutateHandlers[cmd]
	if exists {
		return handler(&conn.User, args)
	}

	return commandHandlers[cmd](conn.User, args)
}

// special handlers for those which mutate the connection
var mutateHandlers = map[commands.Command]func(**authtypes.User, resp.ArgList) ([]byte, error){
	commands.AUTH:   server.HandleAuth,
	commands.SELECT: server.HandleSelect,
}

// check "command" enum for order of commands
// must be in order of commands in the enum type, since the map is indexed 0..n
var commandHandlers = [...]func(*authtypes.User, resp.ArgList) ([]byte, error){
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
	nil, // mutate handler
	server.HandleLogout,
	server.HandleWhoami,
	server.HandleAddUser,
	server.HandleRMUser,
	server.HandleUGet,
	server.HandleAddRule,
	server.HandleRMRule,
	server.HandleSave,
	server.HandleLoad,
	server.HandleListSaves,
	server.HandleRMSave,
	nil, // mutate handler
	server.HandleWhichDB,
}
