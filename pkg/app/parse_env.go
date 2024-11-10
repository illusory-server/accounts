package app

import (
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"time"
)

const (
	ConfigPathKey   = "CONFIG_PATH"
	StartTimeoutKey = "START_TIMEOUT"
	StopTimeoutKey  = "STOP_TIMEOUT"
	LogLevelKey     = "LOG_LEVEL"
	LogPrettyKey    = "LOG_PRETTY"
)

type envVar struct {
	LogLevel     string
	LogPretty    bool
	ConfigPath   string
	StartTimeout time.Duration
	StopTimeout  time.Duration
}

func parseEnv() (*envVar, error) {
	configPath := os.Getenv(ConfigPathKey)
	if configPath == "" {
		return nil, errors.Wrapf(ErrEnvNotFound, "env key %s", ConfigPathKey)
	}
	logLvl := os.Getenv(LogLevelKey)
	if logLvl == "" {
		logLvl = "info"
	}
	logPretty := os.Getenv(LogPrettyKey)
	logPrettyV := false
	if logPretty == "true" {
		logPrettyV = true
	}
	start := os.Getenv(StartTimeoutKey)
	if start == "" {
		start = "5"
	}
	startV, err := strconv.Atoi(start)
	if err != nil {
		return nil, errors.Wrapf(ErrIncorrectEnvValue, "env key %s with value %s", StartTimeoutKey, start)
	}
	stop := os.Getenv(StopTimeoutKey)
	if stop == "" {
		stop = "5"
	}
	stopV, err := strconv.Atoi(start)
	if err != nil {
		return nil, errors.Wrapf(ErrIncorrectEnvValue, "env key %s with value %s", StopTimeoutKey, start)
	}

	return &envVar{
		ConfigPath:   configPath,
		LogPretty:    logPrettyV,
		LogLevel:     logLvl,
		StartTimeout: time.Duration(startV) * time.Second,
		StopTimeout:  time.Duration(stopV) * time.Second,
	}, nil
}

func convertLogLvlToIntLvl(lvl string) int {
	switch lvl {
	case "debug":
		return logger.DebugLvl
	case "info":
		return logger.InfoLvl
	case "warn":
		return logger.WarnLvl
	case "error":
		return logger.ErrorLvl
	}
	return logger.InfoLvl
}
