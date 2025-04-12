package config

type (
	Config struct {
		Env      string     `yaml:"env" env-default:"tests"`
		Postgres PostgreSQL `yaml:"postgres"`
		Kafka    Kafka      `yaml:"kafka"`
		Log      Log        `yaml:"log"`
	}

	Log struct {
		IsPretty bool `yaml:"pretty" env-default:"false"`
	}

	Kafka struct {
		Topic string `yaml:"topic"`
		Addr  string `yaml:"address"`
	}

	PostgreSQL struct {
		Host               string `yaml:"host" env-default:"localhost"`
		Port               int    `yaml:"port" env-default:"5121"`
		User               string `yaml:"user" env-default:"postgres"`
		Password           string `yaml:"password" env-default:"root"`
		DbName             string `yaml:"dbname" env-default:"database"`
		ConnectPoolMin     int    `yaml:"connect_pool_min" env-default:"2"`
		ConnectPoolMax     int    `yaml:"connect_pool_max" env-default:"20"`
		ConnectTimeout     int    `yaml:"connect_timeout" env-default:"60"`
		ConnectIdle        int    `yaml:"connect_idle" env-default:"30"`
		ConnectHealthCheck int    `yaml:"connect_health_check" env-default:"20"`
	}
)
