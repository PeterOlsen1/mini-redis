package miniredis_test

import (
	"mini-redis/server/start"
	"os"
	"testing"
)

// improve upon this later. not working for now

func TestMain(m *testing.M) {
	// Start the server
	start.Start("")
	defer start.Stop()

	// Run the tests
	code := m.Run()

	// Exit with the test result code
	os.Exit(code)
}
