package info

import (
	"sync"
	"time"
)

type Info struct {
	*MemInfo
	*OpInfo
	*ServerInfo
	*DbInfo
}

type ServerInfo struct {
	mu                  sync.Mutex
	address             string
	port                int
	pid                 int
	os                  string
	arch                string
	version             string
	startTime           time.Time
	uptimeSeconds       int
	servicedConnections int
	currentConnections  int
}

type MemInfo struct {
	mu      sync.Mutex
	total   uint64
	cur     uint64
	mallocs uint64
	frees   uint64
}

type OpInfo struct {
	mu      sync.Mutex
	gets    int
	sets    int
	deletes int
	cmd     *CmdInfo
}

type CmdInfo struct {
	mu    sync.Mutex
	total int
	cmds  []int
}

type DbInfo struct {
	mu          sync.Mutex
	totalKeys   int
	expiredKeys int
}
