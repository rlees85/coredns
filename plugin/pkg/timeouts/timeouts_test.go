package timeouts

import (
	"testing"
)

func TestNewTimeoutsConfigFromArgs(t *testing.T) {
	var validIdleTimeoutArg = "300" // Args from configuration are always strings
	var validReadTimeoutArg = "30"
	var validWriteTimeoutArg = "60"

	var invalidTimeoutString = "twenty seconds"
	var invalidTimeoutLow = "0"
	var invalidTimeoutHigh = "86401"

	// No Arguments
	to, err := NewTimeoutsConfigFromArgs()
	if err != nil {
		t.Errorf("Failed to create timeouts map when no arguments specified: %s", err)
	}

	if len(to) != 0 {
		t.Error("Timeouts map with no arguments should be empty")
	}

	// Read Timeout only
	to, err = NewTimeoutsConfigFromArgs(validReadTimeoutArg)
	if err != nil {
		t.Errorf("Failed to create timeouts map given just a read timeout: %s", err)
	}

	if _, ok := to["read"]; !ok {
		t.Error("Timeouts map given just a read timeout did not return a read timeout")
	}

	if _, ok := to["write"]; ok {
		t.Error("Timeouts map given just a read timeout also returned a write timeout")
	}

	if _, ok := to["idle"]; ok {
		t.Error("Timeouts map given just a read timeout also returned an idle timeout")
	}

	// Read and Write Timeouts (no Idle)
	to, err = NewTimeoutsConfigFromArgs(validReadTimeoutArg, validWriteTimeoutArg)
	if err != nil {
		t.Errorf("Failed to create timeouts map given a read and write timeout: %s", err)
	}

	if _, ok := to["read"]; !ok {
		t.Error("Timeouts map given a read and write timeout did not return a read timeout")
	}

	if _, ok := to["write"]; !ok {
		t.Error("Timeouts map given a read and write timeout did not return a write timeout")
	}

	if _, ok := to["idle"]; ok {
		t.Error("Timeouts map given a read and write timeout also returned an idle timeout")
	}

	// All Timeouts
	to, err = NewTimeoutsConfigFromArgs(validReadTimeoutArg, validWriteTimeoutArg, validIdleTimeoutArg)
	if err != nil {
		t.Errorf("Failed to create timeouts map given all timeouts: %s", err)
	}

	if _, ok := to["read"]; !ok {
		t.Error("Timeouts map given all timeouts did not return a read timeout")
	}

	if _, ok := to["write"]; !ok {
		t.Error("Timeouts map given all timeouts did not return a write timeout")
	}

	if _, ok := to["idle"]; !ok {
		t.Error("Timeouts map given all timeouts did not return an idle idle timeout")
	}

	// Too Many Timeouts
	to, err = NewTimeoutsConfigFromArgs(validReadTimeoutArg, validWriteTimeoutArg, validIdleTimeoutArg, "100")
	if err == nil {
		t.Error("Attempt to create timeouts with too many arguments was successful")
	}

	// Timeout Validation
	to, err = NewTimeoutsConfigFromArgs(invalidTimeoutString)
	if err == nil {
		t.Error("Attempt to create timeouts with non-numeric value was successful")
	}

	to, err = NewTimeoutsConfigFromArgs(invalidTimeoutLow)
	if err == nil {
		t.Error("Attempt to create a timeout of less than 1 second was successful")
	}

	to, err = NewTimeoutsConfigFromArgs(invalidTimeoutHigh)
	if err == nil {
		t.Error("Attempt to create a timeout of more than 1 day was successful")
	}
}
