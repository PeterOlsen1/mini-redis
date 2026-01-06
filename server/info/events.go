package info

import (
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
)

type InfoEvent struct {
	Type InfoEventType
	Cmd  commands.Command
}

type InfoEventType int

const (
	CONNECT InfoEventType = iota
	DISCONNECT
	GET
	SET
	EXPIRE
	DELETE
	COMMAND
)

var eventChan = make(chan InfoEvent, 100)

func init() {
	server := globalInfo.ServerInfo
	op := globalInfo.OpInfo
	db := globalInfo.DbInfo

	// spawn goroutine to collect info from a chan here
	go func() {
		for e := range eventChan {
			switch e.Type {
			case CONNECT:
				server.mu.Lock()
				server.servicedConnections += 1
				server.currentConnections += 1
				server.mu.Unlock()
			case DISCONNECT:
				server.mu.Lock()
				server.currentConnections -= 1
				server.mu.Unlock()
			case GET:
				op.mu.Lock()
				op.gets += 1
				op.mu.Unlock()
			case SET:
				op.mu.Lock()
				op.sets += 1
				op.mu.Unlock()
			case EXPIRE:
				db.mu.Lock()
				db.expiredKeys += 1
				db.mu.Unlock()
			case DELETE:
				op.mu.Lock()
				op.deletes += 1
				op.mu.Unlock()
			case COMMAND:
				op.cmd.mu.Lock()
				op.cmd.cmds[e.Cmd] += 1
				op.cmd.total += 1
				op.cmd.mu.Unlock()
			}
		}
	}()
}

func Connect() {
	eventChan <- InfoEvent{Type: CONNECT, Cmd: commands.Command(0)}
}
func Disconnect() {
	eventChan <- InfoEvent{Type: DISCONNECT, Cmd: commands.Command(0)}
}

func GetOp() {
	if !cfg.Info.CollectOps {
		return
	}
	eventChan <- InfoEvent{Type: GET, Cmd: commands.Command(0)}
}

func SetOp() {
	if !cfg.Info.CollectOps {
		return
	}
	eventChan <- InfoEvent{Type: SET, Cmd: commands.Command(0)}
}

func Expire() {
	eventChan <- InfoEvent{Type: EXPIRE, Cmd: commands.Command(0)}
}

func Delete() {
	if !cfg.Info.CollectOps {
		return
	}
	eventChan <- InfoEvent{Type: DELETE, Cmd: commands.Command(0)}
}

// the Cmd field is only set on this command beucase type COMMAND is the only one to require it
func Command(cmd commands.Command) {
	if !cfg.Info.CollectOps {
		return
	}
	eventChan <- InfoEvent{Type: COMMAND, Cmd: cmd}
}
