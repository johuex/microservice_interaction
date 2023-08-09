package config

import (
	"os"
	"strconv"
)

type Config struct {
	BrokerUrl       string
	BrokerPort      string
	BrokerTopic     string
	BrokerPartition int
}

func New() *Config {
	return &Config{
		BrokerUrl:       getEnv("KAFKA_URL", ""),
		BrokerPort:      getEnv("KAFKA_PORT", ""),
		BrokerTopic:     getEnv("KAFKA_TOPIC", ""),
		BrokerPartition: getEnvInt("KAFKA_PARTITION", 0),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		intVal, _ := strconv.Atoi(value)
		return intVal
	}

	return defaultVal
}
