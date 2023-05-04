package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v8"
)

type RunMode string

const (
	ModeDebug   RunMode = "debug"
	ModeRelease RunMode = "release"
)

type Config struct {
	RunMode                RunMode `env:"APP_RUN_MODE"`
	SlackAppLevelToken     string  `env:"SLACK_APP_LEVEL_TOKEN"`
	SlackBotUserOAuthToken string  `env:"SLACK_BOT_USER_OAUTH_TOKEN"`
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{}

	opt := env.Options{
		RequiredIfNoDef: true,
	}
	if err := env.ParseWithOptions(cfg, opt); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	switch c.RunMode {
	case ModeDebug, ModeRelease:
		break
	default:
		return fmt.Errorf("Config: Invalid RunMode: %q", c.RunMode)
	}

	const appTokenPrefix = "xapp-"
	if !strings.HasPrefix(c.SlackAppLevelToken, appTokenPrefix) {
		return fmt.Errorf(
			"Config: SlackAppToken must has prefix '%s'",
			appTokenPrefix,
		)
	}

	const botTokenPrefix = "xoxb-"
	if !strings.HasPrefix(c.SlackBotUserOAuthToken, botTokenPrefix) {
		return fmt.Errorf(
			"Config: SlackBotUserOAuthToken must have the prefix '%s'",
			botTokenPrefix,
		)
	}

	return nil
}
