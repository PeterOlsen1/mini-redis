package miniredis_test

import (
	"mini-redis/server/start"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Start the server
	go start.Start("/")

	time.Sleep(1 * time.Second)
	// Run the tests
	code := m.Run()

	start.Stop()
	os.Exit(code)
}

// Uncomment this and use for testing when the server is already running
// func TestMain(m *testing.M) {
// 	m.Run()
// }
