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

	if config.Timeouts != nil {
		return plugin.Error("timeouts", c.Errf("Timeouts already configured for this server instance"))
	}

	for c.Next() {
		args := c.RemainingArgs()
		if len(args) < 1 || len(args) > 3 {
			return plugin.Error("timeouts", c.ArgErr())
		}

		timeoutsConfig, err := timeouts.NewTimeoutsConfigFromArgs(args...)
		if err != nil {
			return err
		}

		config.Timeouts = timeoutsConfig
	}
	return nil
}
