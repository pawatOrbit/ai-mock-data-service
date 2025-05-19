package core_config

import "github.com/pawatOrbit/ai-mock-data-service/go/core/pgdb"


type Config struct {
	Env        string        `mapstructure:"env"`
	RestServer RestServer    `mapstructure:"restServer"`
	CORS       CORS          `mapstructure:"cors"`
	Postgres   pgdb.Postgres `mapstructure:"postgres"`
}

type CORS struct {
	AllowedMethods []string `mapstructure:"allowedMethods"`
	AllowedHeaders []string `mapstructure:"allowedHeaders"`
	AllowedOrigins []string `mapstructure:"allowedOrigins"` // Default: ["*"]
	ExposedHeaders []string `mapstructure:"exposedHeaders"`
	MaxAge         int      `mapstructure:"maxAge"` // Default: 7200 (seconds)
}

type RestServer struct {
	Port string `mapstructure:"port"`
}
