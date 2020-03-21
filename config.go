package main

import (
	"log"
	"os"
)

type routerConfig struct {
	Host string
	Port string
	Name string
}

type dbConfig struct {
	Host string
	Port string
	Name string
}

type Config struct {
	Router routerConfig
	DB     dbConfig
}

func LoadConfig() *Config {
	config := Config{
		Router: routerConfig{
			Host: getEnvD("MONIT_ROUTER_HOST", "192.168.1.1"),
			Port: getEnvD("MONIT_ROUTER_PORT", "80"),
			Name: getEnv("MONIT_ROUTER_NAME"),
		},
		DB: dbConfig{
			Host: getEnv("MONIT_DB_HOST"),
			Port: getEnvD("MONIT_DB_PORT", "8086"),
			Name: getEnvD("MONIT_DB_NAME", "network"),
		},
	}

	return &config
}

func getEnvD(name string, defaultValue string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}

	return defaultValue
}

func getEnv(name string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}

	log.Fatalf("Environment variable %q is required", name)
	return ""
}
