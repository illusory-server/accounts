package dependency

import (
	"context"
	"github.com/illusory-server/accounts/internal/infra/config"
	"github.com/illusory-server/accounts/pkg/logger"
	"github.com/illusory-server/accounts/pkg/logger/log"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"os"
)

func parseConfigFromYaml(path string, target any) error {
	_, err := os.Stat(path)
	if err != nil {
		return errors.Wrap(err, "cannot stat config file")
	}
	err = cleanenv.ReadConfig(path, target)
	if err != nil {
		return errors.Wrap(err, "cannot read config file")
	}
	return nil
}

func getLevelFromEnv(env string) logger.Level {
	switch env {
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

func (f *Factory) initConfigAndLogger(_ context.Context) error {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		return errors.New("CONFIG_PATH environment variable not set")
	}
	f.deps.config = &config.Config{}
	err := parseConfigFromYaml(path, f.deps.config)
	if err != nil {
		return err
	}
	l := log.NewLogger(&log.Options{
		Level:  getLevelFromEnv(f.deps.config.Env),
		Pretty: f.deps.config.Log.IsPretty,
	})
	f.deps.log = l

	return nil
}
