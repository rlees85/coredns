package timeouts

import (
	"fmt"
	"strconv"
	"time"
)

func NewTimeoutsConfigFromArgs(args ...string) (map[string]time.Duration, error) {
	c := make(map[string]time.Duration)

	for i := 0; i < len(args); i++ {
		t, err := validateTimeout(args[i])

		if err != nil {
			return c, err
		}

		switch i {
		case 0:
			c["read"] = t
		case 1:
			c["write"] = t
		case 2:
			c["idle"] = t
		default:
			return c, fmt.Errorf("maximum of three arguments allowed for timeouts config, found %d", len(args))
		}
	}

	return c, nil
}

func validateTimeout(t string) (time.Duration, error) {
	i, err := strconv.Atoi(t)
	if err != nil {
		return time.Duration(0), fmt.Errorf("timeout provided '%s' does not appear to be numeric", t)
	}

	if i < 1 || i > 86400 {
		return time.Duration(0), fmt.Errorf("timeout provided '%d' needs to be between 1 and 86400 second(s)", i)
	}

	return time.Duration(i) * time.Second, nil
}
