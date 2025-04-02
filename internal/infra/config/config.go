package config

type (
	Config struct {
		Env        string     `yaml:"env" env-default:"tests"`
		SuperAdmin SuperAdmin `yaml:"super_admin"`
		Postgres   PostgreSQL `yaml:"postgres"`
		Kafka      Kafka      `yaml:"kafka"`
	}

	Kafka struct {
		Topic string `yaml:"topic"`
		Addr  string `yaml:"addr"`
	}

	SuperAdmin struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Email    string `yaml:"email"`
	}

	PostgreSQL struct {
		Host           string `yaml:"host" env-default:"localhost"`
		Port           int    `yaml:"port" env-default:"5121"`
		User           string `yaml:"user" env-default:"postgres"`
		Password       string `yaml:"password" env-default:"root"`
		DbName         string `yaml:"dbname" env-default:"database"`
		ConnectPoolMin int    `yaml:"connect_pool_min" env-default:"2"`
		ConnectPoolMax int    `yaml:"connect_pool_max" env-default:"20"`
	}
)
