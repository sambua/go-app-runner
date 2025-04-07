package config

import (
	"errors"
	"strconv"
)

const (
	env                                = "APP_ENV"
	debug                              = "APP_DEBUG"
	debugLevel                         = "APP_DEBUG_LEVEL"
	errMessageWrongDebugEnvKeyProvided = "wrong value for debug env key, accepted 'false' or 'true'"
)

type stateConfig struct {
	Env        string
	Debug      bool
	DebugLevel string
}

func NewState() (*stateConfig, error) {
	env := GetEnvKey(env, "prod")
	debug, err := strconv.ParseBool(GetEnvKey(debug, "false"))
	if err != nil {
		return nil, errors.New(errMessageWrongDebugEnvKeyProvided)
	}
	debugLevel := GetEnvKey(debugLevel, "info")

	return &stateConfig{
		Env:        env,
		Debug:      debug,
		DebugLevel: debugLevel,
	}, nil
}
