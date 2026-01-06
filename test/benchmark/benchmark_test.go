package benchmark

import (
	"fmt"
	"mini-redis/client"
	"testing"
)

func setupAndFlush() (*client.RedisClient, error) {
	c, err := client.NewClient(&client.ClientOptions{Addr: "localhost:6379"})
	if err != nil {
		return nil, err
	}

	_, err = c.FlushAll()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// With all logging enabled
// BenchmarkGetAndSet-8       17340             66852 ns/op            1056 B/op         27 allocs/op

// With not enabled
// BenchmarkGetAndSet-8       20953             52240 ns/op            1057 B/op         27 allocs/op
func BenchmarkGetAndSet(b *testing.B) {
	c, err := setupAndFlush()
	if err != nil {
		fmt.Println("Error initializing test!")
		return
	}
	defer c.Close()

	// simple write/read loop
	for b.Loop() {
		_, err = c.Set("hello", "hello")
		if err != nil {
			b.Fatal("Set call failed")
		}
		_, err := c.Get("hello")
		if err != nil {
			b.Fatal("get call failed")
		}
	}
}

// add a benchmark for any pipelined version I choose to do later
