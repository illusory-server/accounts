package config

import "time"

type (
	Config struct {
		Env        string     `yaml:"env" env-default:"tests"`
		Server     HTTPServer `yaml:"http_server"`
		SuperAdmin SuperAdmin `yaml:"super_admin"`
		Postgres   PostgreSQL `yaml:"postgres"`
	}

	HTTPServer struct {
		Address     string        `yaml:"address" env-default:"localhost:5000"`
		Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	}

	SuperAdmin struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Email    string `yaml:"email"`
	}

	PostgreSQL struct {
		Host     string `yaml:"host" env-default:"localhost"`
		Port     int    `yaml:"port" env-default:"5121"`
		User     string `yaml:"user" env-default:"postgres"`
		Password string `yaml:"password" env-default:"root"`
		DbName   string `yaml:"dbname" env-default:"database"`
	}
)
