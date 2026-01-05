package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"mini-redis/resp"
	"mini-redis/server/handlers"
	"mini-redis/types"
	"net"
	"strconv"
	"strings"
)

func StartServer(ctx context.Context) error {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		return err
	}

	fmt.Println("server running on localhost:6379")

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

		go func() {
			err := handleConnection(conn)
			if err != nil {
				log.Printf("error handling connection: %e", err)
			}
		}()
	}
}

func handleConnection(conn net.Conn) error {
	log.Println("Established connection")
	defer conn.Close()

	for {
		log.Println("Data sent on connection")
		array, err := parseArray(conn)
		if err != nil {
			if err == io.EOF {
				return nil // connection closed
			}

			fmt.Printf("error parsing array: %e\n", err)
			return err
		}

		err = processArray(conn, array)
		if err != nil {
			fmt.Printf("error processing array: %e\n", err)
			return err
		}
	}
}

func parseArray(conn net.Conn) ([]resp.RESPItem, error) {
	reader := bufio.NewReader(conn)
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

	array := make([]resp.RESPItem, 0, arrayLen)

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

func processArray(conn net.Conn, array []resp.RESPItem) error {
	i := 0
	for i < len(array) {
		item := array[i]
		cmd := types.ParseCommand(item.Content)
		if cmd != 0 {
			args := make([]resp.RESPItem, 0)

			i += 1
			for i < len(array) && !types.ParseCommand(array[i].Content).Valid() {
				args = append(args, array[i])
				i += 1
			}

			cmdResp, err := handlers.HandleCommand(cmd, args)
			if err != nil {
				if _, writeErr := conn.Write(resp.BYTE_ERR(err)); writeErr != nil {
					return writeErr
				}
			} else {
				if _, writeErr := conn.Write(cmdResp); writeErr != nil {
					return writeErr
				}
			}
		} else {
			i += 1
		}
	}

	return nil
}
