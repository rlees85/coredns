package timeouts

import (
	"fmt"
	"strconv"
	"time"
)

// NewTimeoutFromArg returns a time.Duration from a configuration argument
// (string) which has come from the Corefile. The argument has some basic
// validation applied before returning a time.Duration.
func NewTimeoutFromArg(arg string) (time.Duration, error) {
	_, err := strconv.Atoi(arg)
	if err == nil {
		// If no time unit is specified default to seconds rather than
		// GO's default of nanoseconds.
		arg = arg + "s"
	}

	d, err := time.ParseDuration(arg)
	if err != nil {
		return time.Duration(0), fmt.Errorf("failed to parse timeout duration '%s'", arg)
	}

	if d < (1*time.Second) || d > (24*time.Hour) {
		return time.Duration(0), fmt.Errorf("timeout provided '%s' needs to be between 1 second and 24 hours", arg)
	}

	return d, nil
}
