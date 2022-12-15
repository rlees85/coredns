package timeouts

import (
	"testing"
	"time"
)

func TestNewTimeoutFromArg(t *testing.T) {
	var validSecondsTimeoutInput = "30s"
	var validSecondsTimeoutOutput = time.Duration(30 * time.Second)
	var validMinutesTimeoutInput = "2m"
	var validMinutesTimeoutOutput = time.Duration(2 * time.Minute)
	var validIntTimeoutInput = "30"
	var validIntTimeoutOutput = time.Duration(30 * time.Second)

	var invalidTimeoutString = "twenty seconds"
	var invalidTimeoutLow = "0"
	var invalidTimeoutHigh = "24h1s"

	// Valid timeout specified as Go duration (seconds)
	to, err := NewTimeoutFromArg(validSecondsTimeoutInput)
	if err != nil {
		t.Errorf("Failed to create timeout duration given a valid Go duration in seconds: %s", err)
	}

	if to != time.Duration(validSecondsTimeoutOutput) {
		t.Errorf("Timeout created given a valid Go duration in seconds appears to have unexpected value: %s", to)
	}

	// Valid timeout specified as Go duration (minutes)
	to, err = NewTimeoutFromArg(validMinutesTimeoutInput)
	if err != nil {
		t.Errorf("Failed to create timeout duration given a valid Go duration in minutes: %s", err)
	}

	if to != time.Duration(validMinutesTimeoutOutput) {
		t.Errorf("Timeout created given a valid Go duration in minutes appears to have unexpected value: %s", to)
	}

	// Valid timeout specified as int
	to, err = NewTimeoutFromArg(validIntTimeoutInput)
	if err != nil {
		t.Errorf("Failed to create timeout duration given a valid int: %s", err)
	}

	if to != time.Duration(validIntTimeoutOutput) {
		t.Errorf("Timeout created given a valid int appears to have unexpected value: %s", to)
	}

	// Invalid timeouts
	_, err = NewTimeoutFromArg(invalidTimeoutString)
	if err == nil {
		t.Error("Attempt to create timeout with non-numeric value was successful")
	}

	_, err = NewTimeoutFromArg(invalidTimeoutLow)
	if err == nil {
		t.Error("Attempt to create timeout of less than 1 second was successful")
	}

	_, err = NewTimeoutFromArg(invalidTimeoutHigh)
	if err == nil {
		t.Error("Attempt to create timeout of more than 1 day was successful")
	}
}
