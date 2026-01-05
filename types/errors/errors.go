package errors

import (
	"fmt"
	"mini-redis/types/commands"
)

var WRONGTYPE = fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value")
var NOT_INTEGER = fmt.Errorf("value is not an integer or out of range")
var INVALID_ARG = fmt.Errorf("INVALID_ARG failed to parse integer argument") //make into a function to return invaid argument?
func ARG_COUNT(cmd commands.Command, count int) error {
	return fmt.Errorf("%s requires %d arguments", cmd, count)
}
