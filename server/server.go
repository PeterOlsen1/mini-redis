package main

import (
	"context"
	"log"
	"net"
)

func StartServer(ctx context.Context) error {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		return err
	}

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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) error {
	// do MANY things here
	log.Println("got connection")
	conn.Write([]byte("hello"))
	conn.Close()
	return nil
}
