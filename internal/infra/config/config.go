package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/leorcvargas/bgeraser/pkg/env"
)

type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Storage struct {
		LocalPath string
		Bucket    string
		BucketURL string
	}
	AWS       struct{ Region string }
	Profiling struct {
		CPU     string
		Mem     string
		Enabled bool
	}
	Server struct {
		Port     string
		UseSonic bool
		Prefork  bool
	}
	Queues struct {
		Host     string
		Port     string
		User     string
		Password string
	}
}

func NewConfig() *Config {
	cfg := &Config{
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
			CPU     string
			Mem     string
			Enabled bool
		}{
			Enabled: env.GetEnvOrDie("ENABLE_PROFILING") == "1",
			CPU:     env.GetEnvOrDie("CPU_PROFILE"),
			Mem:     env.GetEnvOrDie("MEM_PROFILE"),
		},

		Storage: struct {
			LocalPath string
			Bucket    string
			BucketURL string
		}{
			LocalPath: env.GetEnvOrDie("STORAGE_LOCAL_PATH"),
			Bucket:    env.GetEnvOrDie("STORAGE_BUCKET"),
			BucketURL: env.GetEnvOrDie("STORAGE_BUCKET_URL"),
		},

		AWS: struct {
			Region string
		}{
			Region: env.GetEnvOrDie("AWS_REGION"),
		},

		Queues: struct {
			Host     string
			Port     string
			User     string
			Password string
		}{
			Host:     env.GetEnvOrDie("QUEUES_HOST"),
			Port:     env.GetEnvOrDie("QUEUES_PORT"),
			User:     env.GetEnvOrDie("QUEUES_USER"),
			Password: env.GetEnvOrDie("QUEUES_PASSWORD"),
		},
	}

	log.Infow("Config loaded", cfg)

	return cfg
}
