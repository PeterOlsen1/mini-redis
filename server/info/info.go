package info

import (
	"fmt"
	"mini-redis/server/cfg"
	"mini-redis/types/commands"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var globalInfo Info = Info{
	&MemInfo{},
	&OpInfo{
		cmd: &CmdInfo{
			cmds: make([]int, commands.NUM_COMMANDS),
		},
	},
	&ServerInfo{},
	&DbInfo{},
}

func InitServerInfo(address string, port int) {
	globalInfo.ServerInfo.address = address
	globalInfo.ServerInfo.port = port
	globalInfo.ServerInfo.startTime = time.Now()
}

func (s *ServerInfo) getServerInfo() string {
	out := "\n# Server Info:\n"

	s.mu.Lock()
	defer s.mu.Unlock()

	s.pid = os.Getpid()
	s.os = runtime.GOOS
	s.arch = runtime.GOARCH
	s.version = runtime.Version()

	s.uptimeSeconds = int(time.Since(s.startTime).Seconds())

	out += "## Process\n"
	out += "-Address: " + s.address + "\n"
	out += "-Port: " + strconv.Itoa(s.port) + "\n"
	out += "-PID: " + strconv.Itoa(s.pid) + "\n"
	out += "-Uptime (s): " + strconv.Itoa(s.uptimeSeconds) + "\n"
	out += "-Start Time: " + s.startTime.Format(time.RFC3339) + "\n"
	out += "## Machine\n"
	out += "-OS: " + s.os + "\n"
	out += "-Arch: " + s.arch + "\n"
	out += "-Version: " + s.version + "\n"
	out += "## Connections\n"
	out += "-Serviced connections: " + strconv.Itoa(s.servicedConnections) + "\n"
	out += "-Current connections: " + strconv.Itoa(s.currentConnections) + "\n"

	return out
}

func (m *MemInfo) getMemInfo() string {
	out := "\n# Memory Info:\n"

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m.mu.Lock()
	defer m.mu.Unlock()

	m.cur = memStats.Alloc
	m.total = memStats.Sys
	m.mallocs = memStats.Mallocs
	m.frees = memStats.Frees

	out += "-Current memory (kb): " + strconv.FormatUint(m.cur, 10) + "\n"
	out += "-Total memory (kb): " + strconv.FormatUint(m.total, 10) + "\n"
	out += "-Total mallocs: " + strconv.FormatUint(m.mallocs, 10) + "\n"
	out += "-Total frees: " + strconv.FormatUint(m.frees, 10) + "\n"

	return out
}

func (o *OpInfo) getOpInfo() string {
	var out strings.Builder
	out.WriteString("\n# Operation Info:\n")

	if !cfg.Info.CollectOps {
		out.WriteString("Info collection is not enabled in configuration.\n")
	}
	o.mu.Lock()
	defer o.mu.Unlock()

	out.WriteString("-Gets: " + strconv.Itoa(o.gets) + "\n")
	out.WriteString("-Sets: " + strconv.Itoa(o.sets) + "\n")
	out.WriteString("-Deletes: " + strconv.Itoa(o.deletes) + "\n")

	// if the user wants to log all commands, show them
	if cfg.Info.Command {
		o.cmd.mu.Lock()
		fmt.Fprintf(&out, "## Commands Run (%d)\n", o.cmd.total)

		for i := range commands.NUM_COMMANDS {
			fmt.Fprintf(&out, "-%s: %d\n", commands.Command(i), o.cmd.cmds[i])
		}
		o.cmd.mu.Unlock()
	} else {
		out.WriteString("> To see all commands run, enable Info.Command in config\n")
	}

	return out.String()
}

func (d *DbInfo) getDbInfo() string {
	out := "\n# Database Info:\n"

	d.mu.Lock()
	defer d.mu.Unlock()

	out += "-Expired keys: " + strconv.Itoa(d.expiredKeys) + "\n"

	return out
}

func GetInfo() string {
	serverInfo := globalInfo.ServerInfo.getServerInfo()
	memInfo := globalInfo.MemInfo.getMemInfo()
	opInfo := globalInfo.OpInfo.getOpInfo()
	dbInfo := globalInfo.DbInfo.getDbInfo()

	return serverInfo + memInfo + opInfo + dbInfo
}
