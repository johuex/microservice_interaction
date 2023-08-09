package config

import (
	"os"
	"strconv"
)

type Config struct {
	ApiServiceUrl  string
	ApiPort        string
	RPCServiceUrl  string
	RPCPort        string
	KafkaURL       string
	KafkaTopic     string
	KafkaPort      string
	KafkaPartition int
}

func New() *Config {
	return &Config{
		ApiServiceUrl:  getEnv("API_SERVER_URL", ""),
		ApiPort:        getEnv("API_SERVER_PORT", ""),
		RPCServiceUrl:  getEnv("RPC_SERVER_URL", ""),
		RPCPort:        getEnv("RPC_SERVER_PORT", ""),
		KafkaURL:       getEnv("KAFKA_URL", ""),
		KafkaTopic:     getEnv("KAFKA_TOPIC", ""),
		KafkaPort:      getEnv("KAFKA_PORT", ""),
		KafkaPartition: getEnvInt("KAFKA_PARTITION", 0),
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
