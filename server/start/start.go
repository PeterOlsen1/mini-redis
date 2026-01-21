package start

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"mini-redis/resp"
	"mini-redis/server/auth"
	"mini-redis/server/auth/authtypes"
	"mini-redis/server/cfg"
	"mini-redis/server/handlers"
	"mini-redis/server/info"
	"mini-redis/server/internal"
	logger "mini-redis/server/log"
	"mini-redis/types"
	"mini-redis/types/commands"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

var ctx, cancel = context.WithCancel(context.Background())

func Start(configPath string) {
	if (configPath)[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to get home directory:", err)
			os.Exit(1)
		}
		configPath = filepath.Join(homeDir, (configPath)[1:])
	}

	err := cfg.LoadConfig(configPath)
	if err != nil {
		fmt.Println("Failed to read config, exiting")
		fmt.Printf("Error: %e\n", err)
		os.Exit(1)
	}
	authtypes.SetAuthRequired(cfg.Server.RequireAuth)

	err = auth.LoadACLUsers()
	if err != nil {
		fmt.Println("Failed to load users")
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	internal.InitStore(cfg.Server.Databases)

	logger.StartLogger(ctx)
	// internal.StartTTLScan(ctx)
	startServer(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		auth.UpdateACLFile() // update user file on stop server
		cancel()
	}()
}

func Stop() {
	cancel()
}

func startServer(ctx context.Context) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		return err
	}

	fmt.Println("server running on localhost:6379")

	info.InitServerInfo(cfg.Server.Address, cfg.Server.Port)

	go func() {
		<-ctx.Done()
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
				log.Println("error with connection")
				continue
			}
		}

		connWrapper := types.Connection{
			Conn: conn,
			User: &authtypes.User{
				Username: "",
				Password: "",
				Perms:    0,
				DB:       internal.GetDB(0),
			},
		}
		go func() {
			err := handleConnection(&connWrapper)
			if err != nil {
				log.Printf("error handling connection: %e", err)
			}
		}()
	}
}

func handleConnection(conn *types.Connection) error {
	if cfg.Log.Connect {
		log.Printf("Established connection: %s\n", conn.Conn.RemoteAddr())
	}

	info.Connect()
	defer info.Disconnect()
	if cfg.Log.Disconnect {
		defer log.Printf("Disconnect: %s\n", conn.Conn.RemoteAddr())
	}
	defer conn.Conn.Close()

	for {
		array, err := parseArray(conn)
		if err != nil {
			if err == io.EOF {
				return nil // connection closed
			}

			fmt.Printf("error parsing array: %e\n", err)
			return err
		}

		if cfg.Log.DataEvent {
			log.Printf("Data sent on connection: %s\n", conn.Conn.RemoteAddr())
		}

		err = processArray(conn, array)
		if err != nil {
			fmt.Printf("error processing array: %e\n", err)
			return err
		}
	}
}

func parseArray(conn *types.Connection) (resp.ArgList, error) {
	reader := bufio.NewReader(conn.Conn)
	header, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	header = strings.TrimSuffix(header, "\r\n")
	header = strings.TrimPrefix(header, "*")
	arrayLen, err := strconv.Atoi(header)
	if err != nil {
		return nil, err
	}

	array := make(resp.ArgList, 0, arrayLen)

	for range arrayLen {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimSuffix(line, "\r\n")
		line = strings.TrimPrefix(line, "$")
		len, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		line, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		line = strings.TrimSuffix(line, "\r\n")
		array = append(array, resp.RESPItem{
			Len:     len,
			Content: line,
		})
	}

	return array, nil
}

func processArray(conn *types.Connection, array resp.ArgList) error {
	i := 0
	for i < len(array) {
		item := array[i]
		cmd := commands.ParseCommand(item.Content)

		if cfg.Info.CollectOps && cfg.Info.Command {
			info.Command(cmd)
		}

		if cmd != 0 {
			args := make(resp.ArgList, 0)

			i += 1
			for i < len(array) && !commands.ParseCommand(array[i].Content).Valid() {
				args = append(args, array[i])
				i += 1
			}

			cmdResp, err := handlers.HandleCommand(conn, cmd, args)
			if err != nil {
				if _, writeErr := conn.Conn.Write(resp.BYTE_ERR(err)); writeErr != nil {
					return writeErr
				}
			} else {
				if _, writeErr := conn.Conn.Write(cmdResp); writeErr != nil {
					return writeErr
				}
			}
		} else {
			i += 1
		}
	}

	return nil
}
