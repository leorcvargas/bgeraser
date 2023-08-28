package config

import "github.com/leorcvargas/bgeraser/pkg/env"

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	Server struct {
		Port     string
		UseSonic bool
		Prefork  bool
	}

	Profiling struct {
		Enabled bool
		CPU     string
		Mem     string
	}

	Storage struct {
		LocalPath string
	}
}

func NewConfig() *Config {
	return &Config{
		Database: struct {
			Host     string
			Port     string
			User     string
			Password string
			Name     string
		}{
			Host:     env.GetEnvOrDie("DB_HOST"),
			Port:     env.GetEnvOrDie("DB_PORT"),
			User:     env.GetEnvOrDie("DB_USER"),
			Password: env.GetEnvOrDie("DB_PASSWORD"),
			Name:     env.GetEnvOrDie("DB_NAME"),
		},

		Server: struct {
			Port     string
			UseSonic bool
			Prefork  bool
		}{
			Port:     env.GetEnvOrDie("SERVER_PORT"),
			UseSonic: env.GetEnvOrDie("ENABLE_SONIC_JSON") == "1",
			Prefork:  env.GetEnvOrDie("ENABLE_PREFORK") == "1",
		},

		Profiling: struct {
			Enabled bool
			CPU     string
			Mem     string
		}{
			Enabled: env.GetEnvOrDie("ENABLE_PROFILING") == "1",
			CPU:     env.GetEnvOrDie("CPU_PROFILE"),
			Mem:     env.GetEnvOrDie("MEM_PROFILE"),
		},

		Storage: struct {
			LocalPath string
		}{
			LocalPath: env.GetEnvOrDie("STORAGE_LOCAL_PATH"),
		},
	}
}
