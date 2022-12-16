package timeouts

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/timeouts"
)

func init() { plugin.Register("timeouts", setup) }

func setup(c *caddy.Controller) error {
	err := parseTimeouts(c)
	if err != nil {
		return plugin.Error("timeouts", err)
	}
	return nil
}

func parseTimeouts(c *caddy.Controller) error {
	config := dnsserver.GetConfig(c)

	for c.Next() {
		args := c.RemainingArgs()
		if len(args) > 0 {
			return plugin.Error("timeouts", c.ArgErr())
		}

		b := 0
		for c.NextBlock() {
			block := c.Val()
			timeoutArgs := c.RemainingArgs()
			if len(timeoutArgs) != 1 {
				return c.ArgErr()
			}

			timeout, err := timeouts.NewTimeoutFromArg(timeoutArgs[0])
			if err != nil {
				return c.Err(err.Error())
			}

			switch block {
			case "read":
				config.ReadTimeout = timeout

			case "write":
				config.WriteTimeout = timeout

			case "idle":
				config.IdleTimeout = timeout

			default:
				return c.Errf("unknown option: '%s'", block)
			}
			b++
		}

		if b == 0 {
			return plugin.Error("timeouts", c.Err("timeouts block with no timeouts specified"))
		}
	}
	return nil
}
