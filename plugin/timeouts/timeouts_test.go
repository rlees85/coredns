package timeouts

import (
	"strings"
	"testing"

	"github.com/coredns/caddy"
)

func TestTimeouts(t *testing.T) {
	tests := []struct {
		input              string
		shouldErr          bool
		expectedRoot       string // expected root, set to the controller. Empty for negative cases.
		expectedErrContent string // substring from the expected error. Empty for positive cases.
	}{
		// positive
		{"timeouts 30", false, "", ""},
		{"timeouts 30 60", false, "", ""},
		{"timeouts 30 60 300", false, "", ""},
		// negative
		{"timeouts", true, "", "Wrong argument"},
		{"timeouts 30 60 300 600", true, "", "Wrong argument"},
		{"timeouts ten", true, "", "timeout provided 'ten' does not appear to be numeric"},
		{"timeouts 0", true, "", "timeout provided '0' needs to be between 1 and 86400 second(s)"},
		{"timeouts 86401", true, "", "timeout provided '86401' needs to be between 1 and 86400 second(s)"},
	}

	for i, test := range tests {
		c := caddy.NewTestController("dns", test.input)
		err := setup(c)
		//cfg := dnsserver.GetConfig(c)

		if test.shouldErr && err == nil {
			t.Errorf("Test %d: Expected error but found %s for input %s", i, err, test.input)
		}

		if err != nil {
			if !test.shouldErr {
				t.Errorf("Test %d: Expected no error but found one for input %s. Error was: %v", i, test.input, err)
			}

			if !strings.Contains(err.Error(), test.expectedErrContent) {
				t.Errorf("Test %d: Expected error to contain: %v, found error: %v, input: %s", i, test.expectedErrContent, err, test.input)
			}
		}
	}
}
